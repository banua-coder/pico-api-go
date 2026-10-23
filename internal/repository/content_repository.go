package repository

import (
	"fmt"
	"log"
	"strings"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/database"
)

// ContentRepositoryInterface defines read-only access to site content tables.
// For all list methods limit <= 0 means "no limit" (load_all).
type ContentRepositoryInterface interface {
	GetBanners(includeInactive bool, limit, offset int) ([]models.Banner, int, error)
	GetInfographics(limit, offset int) ([]models.Infographic, int, error)
	GetDonations(limit, offset int) ([]models.Donation, int, error)
	GetTelemedicines(limit, offset int) ([]models.Telemedicine, int, error)
	GetPartners(inFooter *bool, limit, offset int) ([]models.Partner, int, error)
	GetContacts(ownerType string, ownerID int64, limit, offset int) ([]models.ContactEntry, int, error)
}

// ContentRepository handles database operations for site content.
type ContentRepository struct {
	db *database.DB
}

// NewContentRepository creates a new ContentRepository.
func NewContentRepository(db *database.DB) *ContentRepository {
	return &ContentRepository{db: db}
}

const contactSelect = `SELECT c.id, c.contactable_type, c.contactable_id, c.contact,
	ct.id, ct.name, ct.prefix, ct.icon, ct.bg_color, ct.label
	FROM contacts c JOIN contact_types ct ON c.contact_type_id = ct.id`

func limitClause(limit, offset int) (string, []interface{}) {
	if limit <= 0 {
		return "", nil
	}
	return " LIMIT ? OFFSET ?", []interface{}{limit, offset}
}

func inClause(ids []int64) (string, []interface{}) {
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	return "(" + strings.TrimSuffix(strings.Repeat("?,", len(ids)), ",") + ")", args
}

func (r *ContentRepository) count(query string, args ...interface{}) (int, error) {
	var total int
	if err := r.db.QueryRow(query, args...).Scan(&total); err != nil {
		return 0, fmt.Errorf("failed to count rows: %w", err)
	}
	return total, nil
}

// GetBanners returns banners ordered by id; only status=1 unless includeInactive.
func (r *ContentRepository) GetBanners(includeInactive bool, limit, offset int) ([]models.Banner, int, error) {
	where := ""
	if !includeInactive {
		where = " WHERE status = 1"
	}
	total, err := r.count(`SELECT COUNT(*) FROM banners` + where)
	if err != nil {
		return nil, 0, err
	}
	lim, args := limitClause(limit, offset)
	rows, err := r.db.Query(`SELECT id, title, CONCAT(url_type, url), image, status FROM banners`+where+` ORDER BY id`+lim, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query banners: %w", err)
	}
	defer closeRows(rows)

	banners := []models.Banner{}
	for rows.Next() {
		var b models.Banner
		if err := rows.Scan(&b.ID, &b.Title, &b.Link, &b.Image, &b.IsActive); err != nil {
			return nil, 0, fmt.Errorf("failed to scan banner: %w", err)
		}
		banners = append(banners, b)
	}
	return banners, total, rows.Err()
}

// GetInfographics returns infographics newest first with images ordered by queue.
func (r *ContentRepository) GetInfographics(limit, offset int) ([]models.Infographic, int, error) {
	total, err := r.count(`SELECT COUNT(*) FROM infographics`)
	if err != nil {
		return nil, 0, err
	}
	lim, args := limitClause(limit, offset)
	rows, err := r.db.Query(`SELECT id, title, source FROM infographics ORDER BY id DESC`+lim, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query infographics: %w", err)
	}
	defer closeRows(rows)

	items := []models.Infographic{}
	idx := map[int64]int{}
	ids := []int64{}
	for rows.Next() {
		var i models.Infographic
		if err := rows.Scan(&i.ID, &i.Title, &i.Source); err != nil {
			return nil, 0, fmt.Errorf("failed to scan infographic: %w", err)
		}
		i.Images = []models.InfographicImage{}
		idx[i.ID] = len(items)
		ids = append(ids, i.ID)
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	if len(ids) == 0 {
		return items, total, nil
	}

	in, inArgs := inClause(ids)
	imgRows, err := r.db.Query(`SELECT id, infographic_id, queue, image_url FROM infographic_images
		WHERE infographic_id IN `+in+` ORDER BY infographic_id, queue, id`, inArgs...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query infographic images: %w", err)
	}
	defer closeRows(imgRows)
	for imgRows.Next() {
		var img models.InfographicImage
		var parent int64
		if err := imgRows.Scan(&img.ID, &parent, &img.Queue, &img.ImageURL); err != nil {
			return nil, 0, fmt.Errorf("failed to scan infographic image: %w", err)
		}
		if p, ok := idx[parent]; ok {
			items[p].Images = append(items[p].Images, img)
		}
	}
	return items, total, imgRows.Err()
}

// GetDonations returns donations with bank accounts (and bank) and contacts.
func (r *ContentRepository) GetDonations(limit, offset int) ([]models.Donation, int, error) {
	total, err := r.count(`SELECT COUNT(*) FROM donations`)
	if err != nil {
		return nil, 0, err
	}
	lim, args := limitClause(limit, offset)
	// delivery_address holds the literal string 'NULL' on some rows; normalise to NULL.
	rows, err := r.db.Query(`SELECT id, title, initiator, url, description, image, accept_noncash,
		NULLIF(delivery_address, 'NULL') FROM donations ORDER BY id`+lim, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query donations: %w", err)
	}
	defer closeRows(rows)

	items := []models.Donation{}
	idx := map[int64]int{}
	ids := []int64{}
	for rows.Next() {
		var d models.Donation
		if err := rows.Scan(&d.ID, &d.Title, &d.Initiator, &d.URL, &d.Description, &d.Image,
			&d.AcceptNoncash, &d.DeliveryAddress); err != nil {
			return nil, 0, fmt.Errorf("failed to scan donation: %w", err)
		}
		d.BankAccounts = []models.BankAccount{}
		d.Contacts = []models.ContactEntry{}
		idx[d.ID] = len(items)
		ids = append(ids, d.ID)
		items = append(items, d)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	if len(ids) == 0 {
		return items, total, nil
	}

	in, inArgs := inClause(ids)
	accRows, err := r.db.Query(`SELECT ba.id, ba.accountable_id, ba.account_number, ba.account_name,
		b.id, b.name, b.nickname, b.logo
		FROM bank_accounts ba JOIN banks b ON ba.bank_id = b.id
		WHERE ba.accountable_type = ? AND ba.accountable_id IN `+in+` ORDER BY ba.id`,
		append([]interface{}{models.ContactOwnerClasses["donation"]}, inArgs...)...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query bank accounts: %w", err)
	}
	defer closeRows(accRows)
	for accRows.Next() {
		var a models.BankAccount
		var parent int64
		if err := accRows.Scan(&a.ID, &parent, &a.AccountNumber, &a.AccountName,
			&a.Bank.ID, &a.Bank.Name, &a.Bank.Nickname, &a.Bank.Logo); err != nil {
			return nil, 0, fmt.Errorf("failed to scan bank account: %w", err)
		}
		if p, ok := idx[parent]; ok {
			items[p].BankAccounts = append(items[p].BankAccounts, a)
		}
	}
	if err := accRows.Err(); err != nil {
		return nil, 0, err
	}

	contacts, err := r.contactsFor("donation", ids)
	if err != nil {
		return nil, 0, err
	}
	for i := range items {
		items[i].Contacts = append(items[i].Contacts, contacts[items[i].ID]...)
	}
	return items, total, nil
}

// GetTelemedicines returns telemedicine providers with schedules and contacts.
func (r *ContentRepository) GetTelemedicines(limit, offset int) ([]models.Telemedicine, int, error) {
	total, err := r.count(`SELECT COUNT(*) FROM telemedicines`)
	if err != nil {
		return nil, 0, err
	}
	lim, args := limitClause(limit, offset)
	rows, err := r.db.Query(`SELECT id, name, is_free FROM telemedicines ORDER BY id`+lim, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query telemedicines: %w", err)
	}
	defer closeRows(rows)

	items := []models.Telemedicine{}
	idx := map[int64]int{}
	ids := []int64{}
	for rows.Next() {
		var t models.Telemedicine
		if err := rows.Scan(&t.ID, &t.Name, &t.IsFree); err != nil {
			return nil, 0, fmt.Errorf("failed to scan telemedicine: %w", err)
		}
		t.Schedules = []models.TelemedicineSchedule{}
		t.Contacts = []models.ContactEntry{}
		idx[t.ID] = len(items)
		ids = append(ids, t.ID)
		items = append(items, t)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	if len(ids) == 0 {
		return items, total, nil
	}

	in, inArgs := inClause(ids)
	schRows, err := r.db.Query(`SELECT id, telemedicine_id, operational_day, operational_time, timezone
		FROM telemedicine_schedules WHERE telemedicine_id IN `+in+` ORDER BY id`, inArgs...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query telemedicine schedules: %w", err)
	}
	defer closeRows(schRows)
	for schRows.Next() {
		var s models.TelemedicineSchedule
		var parent int64
		if err := schRows.Scan(&s.ID, &parent, &s.OperationalDay, &s.OperationalTime, &s.Timezone); err != nil {
			return nil, 0, fmt.Errorf("failed to scan telemedicine schedule: %w", err)
		}
		if p, ok := idx[parent]; ok {
			items[p].Schedules = append(items[p].Schedules, s)
		}
	}
	if err := schRows.Err(); err != nil {
		return nil, 0, err
	}

	contacts, err := r.contactsFor("telemedicine", ids)
	if err != nil {
		return nil, 0, err
	}
	for i := range items {
		items[i].Contacts = append(items[i].Contacts, contacts[items[i].ID]...)
	}
	return items, total, nil
}

// GetPartners returns partners ordered by id, optionally filtered on in_footer.
func (r *ContentRepository) GetPartners(inFooter *bool, limit, offset int) ([]models.Partner, int, error) {
	where := ""
	var whereArgs []interface{}
	if inFooter != nil {
		where = " WHERE in_footer = ?"
		whereArgs = []interface{}{*inFooter}
	}
	total, err := r.count(`SELECT COUNT(*) FROM partners`+where, whereArgs...)
	if err != nil {
		return nil, 0, err
	}
	lim, limArgs := limitClause(limit, offset)
	rows, err := r.db.Query(`SELECT id, name, link, image, in_footer, show_anywhere FROM partners`+where+` ORDER BY id`+lim,
		append(whereArgs, limArgs...)...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query partners: %w", err)
	}
	defer closeRows(rows)

	items := []models.Partner{}
	for rows.Next() {
		var p models.Partner
		if err := rows.Scan(&p.ID, &p.Name, &p.Link, &p.Image, &p.InFooter, &p.ShowAnywhere); err != nil {
			return nil, 0, fmt.Errorf("failed to scan partner: %w", err)
		}
		items = append(items, p)
	}
	return items, total, rows.Err()
}

// GetContacts returns contacts across all owners, optionally filtered by owner type (short name) and id.
func (r *ContentRepository) GetContacts(ownerType string, ownerID int64, limit, offset int) ([]models.ContactEntry, int, error) {
	where := " WHERE 1=1"
	var whereArgs []interface{}
	if ownerType != "" {
		class, ok := models.ContactOwnerClasses[ownerType]
		if !ok {
			return nil, 0, fmt.Errorf("unknown owner type %q", ownerType)
		}
		where += " AND c.contactable_type = ?"
		whereArgs = append(whereArgs, class)
	}
	if ownerID > 0 {
		where += " AND c.contactable_id = ?"
		whereArgs = append(whereArgs, ownerID)
	}
	total, err := r.count(`SELECT COUNT(*) FROM contacts c JOIN contact_types ct ON c.contact_type_id = ct.id`+where, whereArgs...)
	if err != nil {
		return nil, 0, err
	}
	lim, limArgs := limitClause(limit, offset)
	rows, err := r.db.Query(contactSelect+where+` ORDER BY c.id`+lim, append(whereArgs, limArgs...)...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query contacts: %w", err)
	}
	defer closeRows(rows)

	items := []models.ContactEntry{}
	for rows.Next() {
		c, class, err := scanContact(rows)
		if err != nil {
			return nil, 0, err
		}
		c.OwnerType = ownerShortName(class)
		items = append(items, c)
	}
	return items, total, rows.Err()
}

// contactsFor loads contacts of the given owner type for many owners at once.
func (r *ContentRepository) contactsFor(ownerType string, ids []int64) (map[int64][]models.ContactEntry, error) {
	in, inArgs := inClause(ids)
	rows, err := r.db.Query(contactSelect+` WHERE c.contactable_type = ? AND c.contactable_id IN `+in+` ORDER BY c.id`,
		append([]interface{}{models.ContactOwnerClasses[ownerType]}, inArgs...)...)
	if err != nil {
		return nil, fmt.Errorf("failed to query contacts: %w", err)
	}
	defer closeRows(rows)

	out := map[int64][]models.ContactEntry{}
	for rows.Next() {
		c, _, err := scanContact(rows)
		if err != nil {
			return nil, err
		}
		owner := c.OwnerID
		c.OwnerID = 0 // owner is implied by nesting
		out[owner] = append(out[owner], c)
	}
	return out, rows.Err()
}

type rowScanner interface {
	Scan(dest ...interface{}) error
}

func scanContact(rows rowScanner) (models.ContactEntry, string, error) {
	var c models.ContactEntry
	var class string
	if err := rows.Scan(&c.ID, &class, &c.OwnerID, &c.Contact,
		&c.ContactType.ID, &c.ContactType.Name, &c.ContactType.Prefix, &c.ContactType.Icon,
		&c.ContactType.BgColor, &c.ContactType.Label); err != nil {
		return c, "", fmt.Errorf("failed to scan contact: %w", err)
	}
	return c, class, nil
}

func ownerShortName(class string) string {
	for short, c := range models.ContactOwnerClasses {
		if c == class {
			return short
		}
	}
	return "unknown" // never leak raw PHP class names
}

func closeRows(rows interface{ Close() error }) {
	if err := rows.Close(); err != nil {
		log.Printf("Error closing rows: %v", err)
	}
}

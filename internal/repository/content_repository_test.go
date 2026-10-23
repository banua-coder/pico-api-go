package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var contactRowCols = []string{"id", "contactable_type", "contactable_id", "contact", "ct_id", "name", "prefix", "icon", "bg_color", "label"}

func setupContentRepo(t *testing.T) (*ContentRepository, sqlmock.Sqlmock) {
	db, mock := setupMockDB(t)
	return NewContentRepository(db), mock
}

func countRows(n int) *sqlmock.Rows { return sqlmock.NewRows([]string{"count"}).AddRow(n) }

func TestContentRepository_GetBanners_ActiveOnly(t *testing.T) {
	repo, mock := setupContentRepo(t)
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM banners WHERE status = 1`).WillReturnRows(countRows(3))
	mock.ExpectQuery(`SELECT id, title, CONCAT\(url_type, url\), image, status FROM banners WHERE status = 1 ORDER BY id LIMIT \? OFFSET \?`).
		WithArgs(10, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "link", "image", "status"}).
			AddRow(1, "Isoman", "https://www.instagram.com/p/CR5q", "https://x/a.png", true))

	got, total, err := repo.GetBanners(false, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, 3, total)
	assert.Len(t, got, 1)
	assert.Equal(t, "https://www.instagram.com/p/CR5q", got[0].Link)
	assert.True(t, got[0].IsActive)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestContentRepository_GetBanners_IncludeInactiveLoadAll(t *testing.T) {
	repo, mock := setupContentRepo(t)
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM banners$`).WillReturnRows(countRows(7))
	mock.ExpectQuery(`FROM banners ORDER BY id$`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "link", "image", "status"}))

	got, total, err := repo.GetBanners(true, 0, 0)
	assert.NoError(t, err)
	assert.Equal(t, 7, total)
	assert.NotNil(t, got)
	assert.Empty(t, got)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestContentRepository_GetBanners_Errors(t *testing.T) {
	repo, mock := setupContentRepo(t)
	mock.ExpectQuery(`SELECT COUNT`).WillReturnError(errors.New("boom"))
	_, _, err := repo.GetBanners(false, 10, 0)
	assert.Error(t, err)

	repo, mock = setupContentRepo(t)
	mock.ExpectQuery(`SELECT COUNT`).WillReturnRows(countRows(1))
	mock.ExpectQuery(`FROM banners`).WillReturnError(errors.New("boom"))
	_, _, err = repo.GetBanners(false, 10, 0)
	assert.Error(t, err)
}

func TestContentRepository_GetInfographics_ImagesOrdered(t *testing.T) {
	repo, mock := setupContentRepo(t)
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM infographics`).WillReturnRows(countRows(146))
	mock.ExpectQuery(`FROM infographics ORDER BY id DESC LIMIT \? OFFSET \?`).
		WithArgs(2, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "source"}).
			AddRow(5, "B", "src-b").AddRow(4, "A", "src-a"))
	mock.ExpectQuery(`FROM infographic_images\s+WHERE infographic_id IN \(\?,\?\) ORDER BY infographic_id, queue, id`).
		WithArgs(int64(5), int64(4)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "infographic_id", "queue", "image_url"}).
			AddRow(10, 4, 1, "u1").AddRow(11, 4, 2, "u2").AddRow(12, 99, 1, "orphan"))

	got, total, err := repo.GetInfographics(2, 0)
	assert.NoError(t, err)
	assert.Equal(t, 146, total)
	assert.Len(t, got, 2)
	assert.Equal(t, int64(5), got[0].ID)
	assert.NotNil(t, got[0].Images)
	assert.Empty(t, got[0].Images)
	assert.Len(t, got[1].Images, 2)
	assert.Equal(t, "u2", got[1].Images[1].ImageURL)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestContentRepository_GetInfographics_Empty(t *testing.T) {
	repo, mock := setupContentRepo(t)
	mock.ExpectQuery(`SELECT COUNT`).WillReturnRows(countRows(0))
	mock.ExpectQuery(`FROM infographics`).WillReturnRows(sqlmock.NewRows([]string{"id", "title", "source"}))
	got, _, err := repo.GetInfographics(10, 0)
	assert.NoError(t, err)
	assert.Empty(t, got)
	assert.NoError(t, mock.ExpectationsWereMet()) // no image query for an empty page
}

func TestContentRepository_GetDonations(t *testing.T) {
	repo, mock := setupContentRepo(t)
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM donations`).WillReturnRows(countRows(8))
	addr := "Jl. Emy Saelan"
	mock.ExpectQuery(`NULLIF\(delivery_address, 'NULL'\) FROM donations ORDER BY id LIMIT`).
		WithArgs(10, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "initiator", "url", "description", "image", "accept_noncash", "delivery_address"}).
			AddRow(2, "T2", "@a", "u", "d", "donations/x.jpg", true, addr).
			AddRow(3, "T3", "@b", "u", "d", "donations/y.jpg", false, nil))
	mock.ExpectQuery(`FROM bank_accounts ba JOIN banks b`).
		WithArgs(`App\Models\Donation`, int64(2), int64(3)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "accountable_id", "number", "name", "bid", "bname", "nick", "logo"}).
			AddRow(1, 3, "123", "Nirmala", 1, "Bank Rakyat Indonesia, Tbk", "BRI", "logo.png"))
	mock.ExpectQuery(`FROM contacts c JOIN contact_types ct`).
		WithArgs(`App\Models\Donation`, int64(2), int64(3)).
		WillReturnRows(sqlmock.NewRows(contactRowCols).
			AddRow(7, `App\Models\Donation`, 2, "0812", 2, "Telepon", "tel:", "fas fa-phone", "bg-blue-500", "Telpon"))

	got, total, err := repo.GetDonations(10, 0)
	assert.NoError(t, err)
	assert.Equal(t, 8, total)
	assert.Len(t, got, 2)
	assert.Equal(t, &addr, got[0].DeliveryAddress)
	assert.Nil(t, got[1].DeliveryAddress)
	assert.Empty(t, got[0].BankAccounts)
	assert.Len(t, got[1].BankAccounts, 1)
	assert.Equal(t, "BRI", got[1].BankAccounts[0].Bank.Nickname)
	assert.Len(t, got[0].Contacts, 1)
	assert.Equal(t, "Telepon", got[0].Contacts[0].ContactType.Name)
	assert.Zero(t, got[0].Contacts[0].OwnerID)
	assert.NotNil(t, got[1].Contacts)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestContentRepository_GetTelemedicines(t *testing.T) {
	repo, mock := setupContentRepo(t)
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM telemedicines`).WillReturnRows(countRows(33))
	mock.ExpectQuery(`FROM telemedicines ORDER BY id`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "is_free"}).AddRow(1, "dr. A", true))
	mock.ExpectQuery(`FROM telemedicine_schedules WHERE telemedicine_id IN \(\?\) ORDER BY id`).
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "tid", "day", "time", "tz"}).
			AddRow(1, 1, "Senin-Jumat", "07:00-20:00", "WIB"))
	mock.ExpectQuery(`FROM contacts c JOIN contact_types ct`).
		WithArgs(`App\Models\Telemedicine`, int64(1)).
		WillReturnRows(sqlmock.NewRows(contactRowCols))

	got, total, err := repo.GetTelemedicines(0, 0)
	assert.NoError(t, err)
	assert.Equal(t, 33, total)
	assert.Len(t, got, 1)
	assert.True(t, got[0].IsFree)
	assert.Equal(t, "WIB", got[0].Schedules[0].Timezone)
	assert.NotNil(t, got[0].Contacts)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestContentRepository_GetPartners_Filter(t *testing.T) {
	repo, mock := setupContentRepo(t)
	yes := true
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM partners WHERE in_footer = \?`).WithArgs(true).WillReturnRows(countRows(9))
	mock.ExpectQuery(`FROM partners WHERE in_footer = \? ORDER BY id LIMIT \? OFFSET \?`).
		WithArgs(true, 10, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "link", "image", "in_footer", "show_anywhere"}).
			AddRow(1, "Detexi", "https://detexi.id", "detexi.png", true, false))

	got, total, err := repo.GetPartners(&yes, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, 9, total)
	assert.Equal(t, "detexi.png", got[0].Image)
	assert.True(t, got[0].InFooter)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestContentRepository_GetPartners_NoFilter(t *testing.T) {
	repo, mock := setupContentRepo(t)
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM partners$`).WillReturnRows(countRows(12))
	mock.ExpectQuery(`FROM partners ORDER BY id$`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "link", "image", "in_footer", "show_anywhere"}))
	_, total, err := repo.GetPartners(nil, 0, 0)
	assert.NoError(t, err)
	assert.Equal(t, 12, total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestContentRepository_GetContacts_Filtered(t *testing.T) {
	repo, mock := setupContentRepo(t)
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM contacts c JOIN contact_types ct .* c.contactable_type = \? AND c.contactable_id = \?`).
		WithArgs(`App\Models\Hospital`, int64(4)).
		WillReturnRows(countRows(1))
	mock.ExpectQuery(`c.contactable_type = \? AND c.contactable_id = \? ORDER BY c.id LIMIT \? OFFSET \?`).
		WithArgs(`App\Models\Hospital`, int64(4), 10, 0).
		WillReturnRows(sqlmock.NewRows(contactRowCols).
			AddRow(9, `App\Models\Hospital`, 4, "0451", 2, "Telepon", "tel:", "fas fa-phone", "bg-blue-500", "Telpon"))

	got, total, err := repo.GetContacts("hospital", 4, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, 1, total)
	assert.Equal(t, "hospital", got[0].OwnerType)
	assert.Equal(t, int64(4), got[0].OwnerID)
	assert.Equal(t, "tel:", got[0].ContactType.Prefix)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestContentRepository_GetContacts_UnknownOwnerType(t *testing.T) {
	repo, _ := setupContentRepo(t)
	_, _, err := repo.GetContacts("bogus", 0, 10, 0)
	assert.Error(t, err)
}

func TestOwnerShortName_Unknown(t *testing.T) {
	assert.Equal(t, "hospital", ownerShortName(`App\Models\Hospital`))
	assert.Equal(t, "unknown", ownerShortName(`App\Models\Secret`))
}

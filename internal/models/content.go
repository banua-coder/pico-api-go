package models

// ContactOwnerClasses maps the public ?owner_type value to the polymorphic
// contactable_type stored in the contacts table.
var ContactOwnerClasses = map[string]string{
	"hospital":     `App\Models\Hospital`,
	"task_force":   `App\Models\TaskForce`,
	"donation":     `App\Models\Donation`,
	"telemedicine": `App\Models\Telemedicine`,
}

// Banner is a promotional banner. Link is url_type + url as stored.
type Banner struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Link     string `json:"link"`
	Image    string `json:"image"`
	IsActive bool   `json:"is_active"`
}

// InfographicImage is one ordered image of an infographic.
type InfographicImage struct {
	ID       int64  `json:"id"`
	Queue    int    `json:"queue"`
	ImageURL string `json:"image_url"`
}

// Infographic is an infographic with its images ordered by queue.
type Infographic struct {
	ID     int64              `json:"id"`
	Title  string             `json:"title"`
	Source string             `json:"source"`
	Images []InfographicImage `json:"images"`
}

// Bank is a bank referenced by donation bank accounts.
type Bank struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Logo     string `json:"logo"`
}

// BankAccount is a donation bank account.
type BankAccount struct {
	ID            int64  `json:"id"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
	Bank          Bank   `json:"bank"`
}

// ContactEntry is a contact with its type. OwnerType/OwnerID are only set by
// the /contacts endpoint.
type ContactEntry struct {
	ID          int64       `json:"id"`
	OwnerType   string      `json:"owner_type,omitempty"`
	OwnerID     int64       `json:"owner_id,omitempty"`
	ContactType ContactType `json:"contact_type"`
	Contact     string      `json:"contact"`
}

// Donation is a donation campaign. DeliveryAddress is nil when not set.
type Donation struct {
	ID              int64          `json:"id"`
	Title           string         `json:"title"`
	Initiator       string         `json:"initiator"`
	URL             string         `json:"url"`
	Description     string         `json:"description"`
	Image           string         `json:"image"`
	AcceptNoncash   bool           `json:"accept_noncash"`
	DeliveryAddress *string        `json:"delivery_address"`
	BankAccounts    []BankAccount  `json:"bank_accounts"`
	Contacts        []ContactEntry `json:"contacts"`
}

// TelemedicineSchedule is an operational slot of a telemedicine provider.
type TelemedicineSchedule struct {
	ID              int64  `json:"id"`
	OperationalDay  string `json:"operational_day"`
	OperationalTime string `json:"operational_time"`
	Timezone        string `json:"timezone"`
}

// Telemedicine is a telemedicine provider.
type Telemedicine struct {
	ID        int64                  `json:"id"`
	Name      string                 `json:"name"`
	IsFree    bool                   `json:"is_free"`
	Schedules []TelemedicineSchedule `json:"schedules"`
	Contacts  []ContactEntry         `json:"contacts"`
}

// Partner is a partner or data source shown on the site.
type Partner struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Image        string `json:"image"`
	InFooter     bool   `json:"in_footer"`
	ShowAnywhere bool   `json:"show_anywhere"`
}

package models

// Contact represents a contact entry (phone, email, etc.)
type Contact struct {
	ID              int64  `json:"id" db:"id"`
	ContactableType string `json:"-" db:"contactable_type"`
	ContactableID   int64  `json:"-" db:"contactable_id"`
	ContactTypeID   int64  `json:"contact_type_id" db:"contact_type_id"`
	ContactTypeName string `json:"contact_type_name,omitempty"`
	ContactTypeIcon string `json:"contact_type_icon,omitempty"`
	Contact         string `json:"contact" db:"contact"`
}

// ContactType represents a type of contact
type ContactType struct {
	ID      int64  `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Prefix  string `json:"prefix" db:"prefix"`
	Icon    string `json:"icon" db:"icon"`
	BgColor string `json:"bg_color" db:"bg_color"`
	Label   string `json:"label" db:"label"`
}

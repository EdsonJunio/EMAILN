package campaign

import "time"

type Contacts struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contacts
}

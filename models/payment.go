package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Payment is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Payment struct {
    ID int `json:"id" db:"id"`
    Amount string `json:"amount" db:"amount"`
    Gateway string `json:"gateway" db:"gateway"`
    CurrencyID int `json:"currency_id" db:"currency_id"`
    Status string `json:"status" db:"status"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Payment) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Payments is not required by pop and may be deleted
type Payments []Payment

// String is not required by pop and may be deleted
func (p Payments) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Payment) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: p.ID, Name: "ID"},
		&validators.StringIsPresent{Field: p.Amount, Name: "Amount"},
		&validators.StringIsPresent{Field: p.Gateway, Name: "Gateway"},
		&validators.IntIsPresent{Field: p.CurrencyID, Name: "CurrencyID"},
		&validators.StringIsPresent{Field: p.Status, Name: "Status"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Payment) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Payment) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

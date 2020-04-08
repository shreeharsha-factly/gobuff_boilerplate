package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// Plan is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Plan struct {
	PlanName  string    `json:"plan_name" db:"plan_name"`
	PlanInfo  string    `json:"plan_info" db:"plan_info"`
	Status    string    `json:"status" db:"status"`
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Plan) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Plans is not required by pop and may be deleted
type Plans []Plan

// String is not required by pop and may be deleted
func (p Plans) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Plan) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.PlanName, Name: "PlanName"},
		&validators.StringIsPresent{Field: p.PlanInfo, Name: "PlanInfo"},
		&validators.StringIsPresent{Field: p.Status, Name: "Status"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Plan) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Plan) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

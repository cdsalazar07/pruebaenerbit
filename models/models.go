package models

import "time"

// DB scheme

// Datameds
type Datameds struct {
	ID               string    `gorm:"primaryKey" json:"id,omitempty"`
	Brand            string    `gorm:"not null" json:"brand,omitempty"`
	Address          string    `gorm:"not null,unique" json:"address,omitempty"`
	InstallationDate time.Time `gorm:"not null" json:"installationDate,omitempty"`
	RetirementDate   time.Time `json:"retirementDate,omitempty"`
	Serial           string    `gorm:"not null" json:"serial,omitempty"`
	Lines            int       `gorm:"not null" json:"lines,omitempty"`
	IsActive         bool      `gorm:"not null" json:"isActive"`
	CreatedAt        time.Time `gorm:"not null" json:"createdAt,omitempty"`
}

// CRUD models

// LastEMeterModel
type LastEMeterModel struct {
	Brand  string `json:"brand,omitempty"`
	Serial string `json:"serial,omitempty"`
}

// NewEMeterModel
type NewEMeterModel struct {
	Brand            string `json:"brand,omitempty"`
	Address          string `json:"address,omitempty"`
	InstallationDate string `json:"installationDate,omitempty"`
	Serial           string `json:"serial,omitempty"`
	Lines            int    `json:"lines,omitempty"`
}

// DeleteEMeterModel
type DeleteEMeterModel struct {
	ID string `json:"id,omitempty"`
}

// UpdateStatus
type UpdateEMeterStatusModel struct {
	ID     string `json:"id,omitempty"`
	Status bool   `json:"status,omitempty"`
}

// Response
type StandardResponse struct {
	Message string `json:"message,omitempty"`
}

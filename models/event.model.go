package models

import (
	"time"

	"github.com/google/uuid"
)

// Event represents an event in the event management system.
type Event struct {
	ID              string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title           string    `gorm:"type:varchar(255);not null"`
	Description     string    `gorm:"type:varchar(255);not null"`
	Location        string    `gorm:"type:varchar(255);not null"`
	CoverPhotoURL   string    `gorm:"coverPhotoURL"`
	Tickets         bool      `gorm:"not null"`
	TicketDetails   Ticket    `gorm:"embedded;foreignkey:EventID"`
	PrivacySettings Privacy   `gorm:"embedded;foreignkey:EventID"`
	GuestInfo       GuestInfo `gorm:"embedded;foreignkey:EventID"`
	User            uuid.UUID `gorm:"not null" json:"user,omitempty"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt       time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

// Ticket represents the ticket details for an event.
type Ticket struct {
	Name          string  `json:"name" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
	Quantity      int     `json:"quantity" binding:"required"`
	EventCapacity int     `json:"eventCapacity" binding:"required"`
}

// Privacy represents the privacy settings for an event.
type Privacy struct {
	Public          bool   `json:"public" binding:"required"`
	GuestVisibility string `json:"guestVisibility" binding:"required"`
}

// GuestInfo represents the information about the guests attending the event.
type GuestInfo struct {
	// Add relevant fields for guest information here
}

// EventInput represents the input for creating or updating an event.
type EventInput struct {
	Title           string    `json:"title" binding:"required"`
	Description     string    `json:"description" binding:"required"`
	Location        string    `json:"location" binding:"required"`
	CoverPhotoURL   string    `json:"coverPhotoURL"`
	Tickets         bool      `json:"tickets"`
	TicketDetails   Ticket    `json:"ticketDetails"`
	PrivacySettings Privacy   `json:"privacySettings" binding:"required"`
	GuestInfo       GuestInfo `json:"guestInfo"`
}

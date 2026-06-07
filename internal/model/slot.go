package model

import "time"

type Slot struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TimeSlot  string    `json:"time_slot"`
	Status    string    `gorm:"default:'available'" json:"status"`
	BookedBy  string    `json:"booked_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

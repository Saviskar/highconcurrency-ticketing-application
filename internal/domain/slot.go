package domain

import "time"

type SlotStatus string

const (
	StatusAvailable SlotStatus = "available"
	StatusBooked    SlotStatus = "booked"
)

type Slot struct {
	ID        uint
	TimeSlot  string
	Status    SlotStatus
	BookedBy  string
	UpdatedAt time.Time
}

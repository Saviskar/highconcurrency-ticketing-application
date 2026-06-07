package domain

import "errors"

var (
	ErrSlotLocked      = errors.New("slot is currently locked by another user")
	ErrSlotNotAvailable = errors.New("slot is not available")
	ErrSlotNotFound    = errors.New("slot not found")
)

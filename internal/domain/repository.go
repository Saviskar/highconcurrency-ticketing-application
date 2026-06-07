package domain

import (
	"context"
	"time"
)

type SlotRepository interface {
	FindAll(ctx context.Context) ([]Slot, error)
	FindByID(ctx context.Context, id uint) (*Slot, error)
	Save(ctx context.Context, slot *Slot) error
}

type DistributedLock interface {
	Acquire(ctx context.Context, key string, ttl time.Duration) (bool, error)
	Release(ctx context.Context, key string) error
}

type EventPublisher interface {
	Publish(event string, data interface{}) error
}

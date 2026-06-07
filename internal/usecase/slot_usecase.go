package usecase

import (
	"context"
	"fmt"
	"time"

	"ticketing-application/internal/domain"
)

type SlotUseCase struct {
	repo   domain.SlotRepository
	locker domain.DistributedLock
	events domain.EventPublisher
}

func NewSlotUseCase(repo domain.SlotRepository, locker domain.DistributedLock, events domain.EventPublisher) *SlotUseCase {
	return &SlotUseCase{
		repo:   repo,
		locker: locker,
		events: events,
	}
}

func (uc *SlotUseCase) ListSlots(ctx context.Context) ([]domain.Slot, error) {
	return uc.repo.FindAll(ctx)
}

func (uc *SlotUseCase) BookSlot(ctx context.Context, slotID uint, user string) (*domain.Slot, error) {
	lockKey := fmt.Sprintf("lock:slot:%d", slotID)

	acquired, err := uc.locker.Acquire(ctx, lockKey, 5*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire lock: %w", err)
	}
	if !acquired {
		return nil, domain.ErrSlotLocked
	}
	defer uc.locker.Release(ctx, lockKey)

	slot, err := uc.repo.FindByID(ctx, slotID)
	if err != nil {
		return nil, err
	}

	if slot.Status != domain.StatusAvailable {
		return nil, domain.ErrSlotNotAvailable
	}

	slot.Status = domain.StatusBooked
	slot.BookedBy = user
	if err := uc.repo.Save(ctx, slot); err != nil {
		return nil, fmt.Errorf("failed to save slot: %w", err)
	}

	if uc.events != nil {
		uc.events.Publish("slot_locked", map[string]interface{}{
			"slot_id": slotID,
		})
	}

	return slot, nil
}

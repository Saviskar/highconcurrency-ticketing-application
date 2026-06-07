package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"ticketing-application/internal/domain"

	"gorm.io/gorm"
)

type slotModel struct {
	ID        uint   `gorm:"primaryKey"`
	TimeSlot  string
	Status    string `gorm:"default:'available'"`
	BookedBy  string
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (slotModel) TableName() string {
	return "slots"
}

type SlotRepository struct {
	db *gorm.DB
}

func NewSlotRepository(db *gorm.DB) *SlotRepository {
	return &SlotRepository{db: db}
}

func (r *SlotRepository) FindAll(ctx context.Context) ([]domain.Slot, error) {
	var models []slotModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	slots := make([]domain.Slot, len(models))
	for i, m := range models {
		slots[i] = toDomain(m)
	}
	return slots, nil
}

func (r *SlotRepository) FindByID(ctx context.Context, id uint) (*domain.Slot, error) {
	var m slotModel
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrSlotNotFound
		}
		return nil, err
	}
	s := toDomain(m)
	return &s, nil
}

func (r *SlotRepository) Save(ctx context.Context, slot *domain.Slot) error {
	m := toModel(slot)
	return r.db.WithContext(ctx).Save(&m).Error
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&slotModel{})
}

func SeedSlots(db *gorm.DB) error {
	var count int64
	db.Model(&slotModel{}).Count(&count)
	if count > 0 {
		fmt.Println("Database already seeded, skipping.")
		return nil
	}
	slots := []slotModel{
		{TimeSlot: "08:00 AM - 09:00 AM", Status: "available"},
		{TimeSlot: "09:00 AM - 10:00 AM", Status: "available"},
		{TimeSlot: "10:00 AM - 11:00 AM", Status: "available"},
		{TimeSlot: "11:00 AM - 12:00 PM", Status: "available"},
	}
	if err := db.Create(&slots).Error; err != nil {
		return err
	}
	fmt.Println("Database seeded with initial available slots!")
	return nil
}

func toDomain(m slotModel) domain.Slot {
	return domain.Slot{
		ID:        m.ID,
		TimeSlot:  m.TimeSlot,
		Status:    domain.SlotStatus(m.Status),
		BookedBy:  m.BookedBy,
		UpdatedAt: m.UpdatedAt,
	}
}

func toModel(s *domain.Slot) slotModel {
	return slotModel{
		ID:        s.ID,
		TimeSlot:  s.TimeSlot,
		Status:    string(s.Status),
		BookedBy:  s.BookedBy,
		UpdatedAt: s.UpdatedAt,
	}
}

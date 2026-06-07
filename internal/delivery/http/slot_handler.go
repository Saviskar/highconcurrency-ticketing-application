package http

import (
	"errors"
	"net/http"

	"ticketing-application/internal/domain"
	"ticketing-application/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SlotHandler struct {
	uc *usecase.SlotUseCase
}

func NewSlotHandler(uc *usecase.SlotUseCase) *SlotHandler {
	return &SlotHandler{uc: uc}
}

func (h *SlotHandler) GetSlots(c *gin.Context) {
	slots, err := h.uc.ListSlots(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch slots"})
		return
	}
	c.JSON(http.StatusOK, slots)
}

type bookSlotRequest struct {
	SlotID uint   `json:"slot_id" binding:"required"`
	User   string `json:"user" binding:"required"`
}

func (h *SlotHandler) BookSlot(c *gin.Context) {
	var req bookSlotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slot, err := h.uc.BookSlot(c.Request.Context(), req.SlotID, req.User)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrSlotLocked):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case errors.Is(err, domain.ErrSlotNotAvailable):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case errors.Is(err, domain.ErrSlotNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Slot successfully booked!",
		"slot":    slot,
	})
}

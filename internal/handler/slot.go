package handler

import (
	"fmt"
	"net/http"
	"time"

	"ticketing-application/internal/database"
	"ticketing-application/internal/model"

	"github.com/gin-gonic/gin"
)

func GetSlots(c *gin.Context) {
	var slots []model.Slot
	result := database.DB.Find(&slots)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch slots"})
		return
	}
	c.JSON(http.StatusOK, slots)
}

type BookRequest struct {
	SlotID uint   `json:"slot_id" binding:"required"`
	User   string `json:"user" binding:"required"`
}

func BookSlot(c *gin.Context) {
	var req BookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lockKey := fmt.Sprintf("lock:slot:%d", req.SlotID)

	success, err := database.RDB.SetNX(c.Request.Context(), lockKey, req.User, 5*time.Minute).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to acquire lock"})
		return
	}
	if !success {
		c.JSON(http.StatusConflict, gin.H{"error": "Slot is currently being booked by another user"})
		return
	}
	defer database.RDB.Del(c.Request.Context(), lockKey)

	var slot model.Slot
	if err := database.DB.First(&slot, req.SlotID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Slot not found"})
		return
	}

	if slot.Status != "available" {
		c.JSON(http.StatusConflict, gin.H{"error": "Slot is already booked or held"})
		return
	}

	slot.Status = "booked"
	slot.BookedBy = req.User
	database.DB.Save(&slot)

	c.JSON(http.StatusOK, gin.H{"message": "Slot successfully booked!", "slot": slot})
}

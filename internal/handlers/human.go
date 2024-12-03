package handlers

import (
	"net/http"
	"passport/internal/models"
	"passport/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HumanHandler struct {
	humanService *service.HumanService
}

func NewHumanHandler(humanService *service.HumanService) *HumanHandler {
	return &HumanHandler{humanService: humanService}
}

func (h *HumanHandler) CreateHuman(c *gin.Context) {
	var human models.Human
	if c.ShouldBindJSON(&human) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.humanService.CreateHuman(&human)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create human"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Human created successfully"})
}

func (h *HumanHandler) GetHuman(c *gin.Context) {
	id := c.Param("id")
	humanID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	human, err := h.humanService.GetHuman(&models.Human{ID: humanID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get human"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"human": human})
}

func (h *HumanHandler) UpdateHuman(c *gin.Context) {
	id := c.Param("id")
	var human models.Human
	if c.ShouldBindJSON(&human) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.humanService.UpdateHuman(&human)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update human"})
		return
	}

	human.ID, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Human updated successfully"})
}

func (h *HumanHandler) DeleteHuman(c *gin.Context) {
	id := c.Param("id")
	humanID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.humanService.DeleteHuman(&models.Human{ID: humanID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete human"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Human deleted successfully"})
}

func (h *HumanHandler) GetHumanBySerialNumber(c *gin.Context) {
	serialNumber := c.Param("serial_number")
	human, err := h.humanService.GetHumanBySerialNumber(serialNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get human by serial number"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"human": human})
}

func (h *HumanHandler) GetAllHumans(c *gin.Context) {
	humans, err := h.humanService.GetAllHumans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all humans"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"humans": humans})
}

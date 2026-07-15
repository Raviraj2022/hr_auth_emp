package shift

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Create Shift godoc
//
// @Summary Create Shift
// @Description Create a new shift
// @Tags Shift
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body CreateShiftRequest true "Shift"
// @Success 201 {object} common.APIResponse
// @Failure 400 {object} common.APIResponse
// @Router /shifts [post]
func (h *Handler) Create(c *gin.Context) {

	var req CreateShiftRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Shift created successfully",
		"data":    response,
	})
}

// Get All Shifts godoc
//
// @Summary Get All Shifts
// @Description Get all shifts
// @Tags Shift
// @Security BearerAuth
// @Produce json
// @Success 200 {object} common.APIResponse
// @Router /shifts [get]
func (h *Handler) GetAll(c *gin.Context) {

	response, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// Get Shift By ID godoc
//
// @Summary Get Shift By ID
// @Description Get shift by id
// @Tags Shift
// @Security BearerAuth
// @Produce json
// @Param id path string true "Shift ID"
// @Success 200 {object} common.APIResponse
// @Failure 404 {object} common.APIResponse
// @Router /shifts/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid shift id",
		})
		return
	}

	response, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// Update Shift godoc
//
// @Summary Update Shift
// @Description Update an existing shift
// @Tags Shift
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Shift ID"
// @Param request body UpdateShiftRequest true "Shift"
// @Success 200 {object} common.APIResponse
// @Failure 400 {object} common.APIResponse
// @Failure 404 {object} common.APIResponse
// @Router /shifts/{id} [put]
func (h *Handler) Update(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid shift id",
		})
		return
	}

	var req UpdateShiftRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Shift updated successfully",
		"data":    response,
	})
}

// Delete Shift godoc
//
// @Summary Delete Shift
// @Description Delete shift by id
// @Tags Shift
// @Security BearerAuth
// @Produce json
// @Param id path string true "Shift ID"
// @Success 200 {object} common.APIResponse
// @Failure 400 {object} common.APIResponse
// @Router /shifts/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid shift id",
		})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Shift deleted successfully",
	})
}
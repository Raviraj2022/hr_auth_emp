package holiday

import (
	"net/http"
    "strconv"
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

// Create Holiday
//
// @Summary Create Holiday
// @Description Create a new holiday
// @Tags Holiday
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body CreateHolidayRequest true "Holiday Details"
// @Success 201 {object} common.APIResponse
// @Failure 400 {object} common.APIResponse
// @Router /holidays [post]
func (h *Handler) Create(c *gin.Context) {

	var req CreateHolidayRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	holiday, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Holiday created successfully",
		"data":    holiday,
	})
}

// Get All Holidays
//
// @Summary Get All Holidays
// @Description Retrieve all holidays
// @Tags Holiday
// @Security BearerAuth
// @Produce json
// @Success 200 {object} common.APIResponse
// @Router /holidays [get]
func (h *Handler) GetAll(c *gin.Context) {

	holidays, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    holidays,
	})
}

// Get Holiday By ID
//
// @Summary Get Holiday By ID
// @Description Retrieve holiday by ID
// @Tags Holiday
// @Security BearerAuth
// @Produce json
// @Param id path string true "Holiday ID"
// @Success 200 {object} common.APIResponse
// @Failure 404 {object} common.APIResponse
// @Router /holidays/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid holiday ID",
		})
		return
	}

	holiday, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    holiday,
	})
}

// Get Holiday By Year
//
// @Summary Get Holiday By Year
// @Description Retrieve holiday by Year
// @Tags Holiday
// @Security BearerAuth
// @Produce json
// @Param id path string true "Holiday Year"
// @Success 200 {object} common.APIResponse
// @Failure 404 {object} common.APIResponse
// @Router /holidays/{year} [get]
func (h *Handler) GetByYear(c *gin.Context) {

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid year",
		})
		return
	}

	response, err := h.service.GetByYear(year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": response,
	})
}

// Update Holiday
//
// @Summary Update Holiday
// @Description Update holiday details
// @Tags Holiday
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Holiday ID"
// @Param request body UpdateHolidayRequest true "Holiday Details"
// @Success 200 {object} common.APIResponse
// @Failure 400 {object} common.APIResponse
// @Router /holidays/{id} [put]
func (h *Handler) Update(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid holiday ID",
		})
		return
	}

	var req UpdateHolidayRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	holiday, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Holiday updated successfully",
		"data":    holiday,
	})
}

// Delete Holiday
//
// @Summary Delete Holiday
// @Description Delete a holiday
// @Tags Holiday
// @Security BearerAuth
// @Produce json
// @Param id path string true "Holiday ID"
// @Success 200 {object} common.APIResponse
// @Failure 400 {object} common.APIResponse
// @Router /holidays/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid holiday ID",
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
		"message": "Holiday deleted successfully",
	})
}
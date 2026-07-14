package attendance

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

// Check In godoc
//
//	@Summary		Check In
//	@Description	Employee check in
//	@Tags			Attendance
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CheckInRequest	true	"Check In"
//	@Success		201		{object}	common.APIResponse
//	@Failure		400		{object}	common.APIResponse
//	@Failure		401		{object}	common.APIResponse
//	@Router			/attendance/check-in [post]
func (h *Handler) CheckIn(c *gin.Context) {

	var req CheckInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response, err := h.service.CheckIn(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Checked in successfully",
		"data":    response,
	})
}

// Check Out godoc
//
//	@Summary		Check Out
//	@Description	Employee check out
//	@Tags			Attendance
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CheckOutRequest	true	"Check Out"
//	@Success		200		{object}	common.APIResponse
//	@Failure		400		{object}	common.APIResponse
//	@Failure		401		{object}	common.APIResponse
//	@Router			/attendance/check-out [put]
func (h *Handler) CheckOut(c *gin.Context) {

	var req CheckOutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response, err := h.service.CheckOut(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Checked out successfully",
		"data":    response,
	})
}

// Get All Attendance godoc
//
//	@Summary		Get All Attendance
//	@Description	Get all attendance records
//	@Tags			Attendance
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	common.APIResponse
//	@Failure		401	{object}	common.APIResponse
//	@Router			/attendance [get]
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

// Get Attendance By ID godoc
//
//	@Summary		Get Attendance By ID
//	@Description	Get attendance by ID
//	@Tags			Attendance
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Attendance ID"
//	@Success		200	{object}	common.APIResponse
//	@Failure		400	{object}	common.APIResponse
//	@Failure		404	{object}	common.APIResponse
//	@Router			/attendance/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid attendance id",
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

// Get Attendance By Employee godoc
//
//	@Summary		Get Attendance By Employee
//	@Description	Get attendance by employee
//	@Tags			Attendance
//	@Security		BearerAuth
//	@Produce		json
//	@Param			employee_id	path		string	true	"Employee ID"
//	@Success		200			{object}	common.APIResponse
//	@Failure		400			{object}	common.APIResponse
//	@Failure		404			{object}	common.APIResponse
//	@Router			/attendance/employee/{employee_id} [get]
func (h *Handler) GetByEmployee(c *gin.Context) {

	employeeID, err := uuid.Parse(c.Param("employee_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid employee id",
		})
		return
	}

	response, err := h.service.GetByEmployee(employeeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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

// Delete Attendance godoc
//
//	@Summary		Delete Attendance
//	@Description	Delete attendance record
//	@Tags			Attendance
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Attendance ID"
//	@Success		200	{object}	common.APIResponse
//	@Failure		400	{object}	common.APIResponse
//	@Failure		404	{object}	common.APIResponse
//	@Router			/attendance/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid attendance id",
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
		"message": "Attendance deleted successfully",
	})
}

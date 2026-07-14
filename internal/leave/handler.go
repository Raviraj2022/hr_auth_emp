package leave

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

// Create Leave
//
//	@Summary		Create Leave
//	@Description	Create a new leave
//	@Tags			Leave
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		leave.CreateLeaveRequest	true	"Leave"
//	@Success		201		{object}	common.APIResponse
//	@Failure		400		{object}	common.APIResponse
//	@Router			/leaves [post]
func (h *Handler) Create(c *gin.Context) {

	var req CreateLeaveRequest

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
		"message": "Leave applied successfully",
		"data":    response,
	})
}

// Get All Leaves godoc
//
//	@Summary		Get All Leaves
//	@Description	Get All Leaves
//	@Tags			Leave
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	common.APIResponse
//	@Failure		401	{object}	common.APIResponse
//	@Failure		500	{object}	common.APIResponse
//	@Router			/leaves [get]
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

// Get All Leave By ID godoc
//
//	@Summary		Get All Leave By ID
//	@Description	Get All leave by ID
//	@Tags			Leave
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Leave ID"
//	@Success		200	{object}	common.APIResponse
//	@Failure		400	{object}	common.APIResponse
//	@Failure		404	{object}	common.APIResponse
//	@Router			/leaves/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid leave id",
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

// Get Leave By Employee godoc
//
//	@Summary		Get Leaves By Employee
//	@Description	Get all leave requests for a specific employee
//	@Tags			Leave
//	@Security		BearerAuth
//	@Produce		json
//	@Param			employee_id	path		string	true	"Employee ID"
//	@Success		200			{object}	common.APIResponse
//	@Failure		400			{object}	common.APIResponse
//	@Failure		401			{object}	common.APIResponse
//	@Router			/leaves/employee/{employee_id} [get]
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

// Approve Leave godoc
//
//	@Summary		Approve Leave
//	@Description	Approve an employee leave request
//	@Tags			Leave
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"Leave ID"
//	@Param			request	body		UpdateLeaveStatusRequest	true	"Approval Request"
//	@Success		200		{object}	common.APIResponse
//	@Failure		400		{object}	common.APIResponse
//	@Failure		401		{object}	common.APIResponse
//	@Failure		404		{object}	common.APIResponse
//	@Router			/leaves/{id}/approve [put]
func (h *Handler) Approve(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid leave id",
		})
		return
	}

	var req UpdateLeaveStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	approvedBy, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "invalid user",
		})
		return
	}

	if err := h.service.Approve(id, approvedBy, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Leave approved successfully",
	})
}

// Reject Leave godoc
//
//	@Summary		Reject Leave
//	@Description	Reject an employee leave request
//	@Tags			Leave
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"Leave ID"
//	@Param			request	body		UpdateLeaveStatusRequest	true	"Reject Request"
//	@Success		200		{object}	common.APIResponse
//	@Failure		400		{object}	common.APIResponse
//	@Failure		401		{object}	common.APIResponse
//	@Failure		404		{object}	common.APIResponse
//	@Router			/leaves/{id}/reject [put]
func (h *Handler) Reject(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid leave id",
		})
		return
	}

	var req UpdateLeaveStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	rejectedBy, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "invalid user",
		})
		return
	}

	if err := h.service.Reject(id, rejectedBy, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Leave rejected successfully",
	})
}

// Delete Leave godoc
//
//	@Summary		Delete Leave
//	@Description	Delete (cancel) a leave request
//	@Tags			Leave
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Leave ID"
//	@Success		200	{object}	common.APIResponse
//	@Failure		400	{object}	common.APIResponse
//	@Failure		401	{object}	common.APIResponse
//	@Failure		404	{object}	common.APIResponse
//	@Router			/leaves/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid leave id",
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
		"message": "Leave cancelled successfully",
	})
}

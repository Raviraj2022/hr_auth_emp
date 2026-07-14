package payroll

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

// Generate Payroll godoc
//
//	@Summary		Generate Payroll
//	@Description	Generate payroll for an employee
//	@Tags			Payroll
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		payroll.GeneratePayrollRequest	true	"Generate Payroll Request"
//	@Success		201		{object}	common.APIResponse
//	@Failure		400		{object}	common.APIResponse
//	@Failure		401		{object}	common.APIResponse
//	@Router			/payroll/generate [post]
func (h *Handler) Generate(c *gin.Context) {

	var req GeneratePayrollRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response, err := h.service.Generate(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Payroll generated successfully",
		"data":    response,
	})
}

// Get All Payrolls godoc
//
//	@Summary		Get All Payrolls
//	@Description	Get all payroll records
//	@Tags			Payroll
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	common.APIResponse
//	@Failure		401	{object}	common.APIResponse
//	@Failure		500	{object}	common.APIResponse
//	@Router			/payroll [get]
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

// Get Payroll By ID godoc
//
//	@Summary		Get Payroll By ID
//	@Description	Get payroll details by payroll ID
//	@Tags			Payroll
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Payroll ID"
//	@Success		200	{object}	common.APIResponse
//	@Failure		400	{object}	common.APIResponse
//	@Failure		404	{object}	common.APIResponse
//	@Router			/payroll/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid payroll id",
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

// Get Payroll By Employee godoc
//
//	@Summary		Get Payroll By Employee
//	@Description	Get payroll history of an employee
//	@Tags			Payroll
//	@Security		BearerAuth
//	@Produce		json
//	@Param			employee_id	path		string	true	"Employee ID"
//	@Success		200			{object}	common.APIResponse
//	@Failure		400			{object}	common.APIResponse
//	@Failure		404			{object}	common.APIResponse
//	@Router			/payroll/employee/{employee_id} [get]
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

// Mark Payroll Paid godoc
//
//	@Summary		Mark Payroll as Paid
//	@Description	Mark a payroll as paid
//	@Tags			Payroll
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Payroll ID"
//	@Success		200	{object}	common.APIResponse
//	@Failure		400	{object}	common.APIResponse
//	@Failure		404	{object}	common.APIResponse
//	@Router			/payroll/{id}/pay [put]
func (h *Handler) MarkPaid(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid payroll id",
		})
		return
	}

	if err := h.service.MarkPaid(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Salary marked as paid",
	})
}

// Delete Payroll godoc
//
//	@Summary		Delete Payroll
//	@Description	Delete a payroll record
//	@Tags			Payroll
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Payroll ID"
//	@Success		200	{object}	common.APIResponse
//	@Failure		400	{object}	common.APIResponse
//	@Failure		404	{object}	common.APIResponse
//	@Router			/payroll/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid payroll id",
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
		"message": "Payroll deleted successfully",
	})
}

package employee

import (
	"net/http"
	// "fmt"

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

// Create Employee

// Create Employee
//
//	@Summary		Create Employee
//	@Description	Create a new employee
//	@Tags			Employee
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		employee.CreateEmployeeRequest	true	"Employee"
//	@Success		201		{object}	common.APIResponse
//	@Failure		400		{object}	common.APIResponse
//	@Router			/employees [post]
func (h *Handler) Create(c *gin.Context) {

	var req CreateEmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	employee, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Employee created successfully",
		"data":    employee,
	})
}

// Get All Employees godoc
//
//	@Summary		Get All Employees
//	@Description	Get All employees
//	@Tags			Employee
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	common.APIResponse
//	@Failure		401	{object}	common.APIResponse
//	@Failure		500	{object}	common.APIResponse
//	@Router			/employees [get]
func (h *Handler) GetAll(c *gin.Context) {

	employees, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employees,
	})
}

// Get All Employee By ID godoc
//
//	@Summary		Get All Employee By ID
//	@Description	Get All employee by ID
//	@Tags			Employee
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Employee ID"
//	@Success		200	{object}	common.APIResponse
//	@Failure		400	{object}	common.APIResponse
//	@Failure		404	{object}	common.APIResponse
//	@Router			/employees/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid employee ID",
		})
		return
	}

	employee, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}

// Update Employee godoc
//
//	@Summary		Update Employee
//	@Description	Update Employee Details
//	@Tags			Employee
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@param			id		path		string							true	"Employee ID"
//	@param			request	body		employee.UpdateEmployeeRequest	true	"Employee"
//	@Success		200		{object}	common.APIResponse
//	@Failure		400		{object}	common.APIResponse
//	@Failure		404		{object}	common.APIResponse
//	@Router			/employees/{id} [put]
func (h *Handler) Update(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	// fmt.Printf("Request: %+v\n", c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid employee ID",
		})
		return
	}

	var req UpdateEmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	employee, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Employee updated successfully",
		"data":    employee,
	})
}

// Delete Employee godoc
//
//	@Summary		Delete Employee
//	@Description	Delete Employee by ID
//	@Tags			Employee
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Employee ID"
//	@Success		200	{object}	common.APIResponse
//	@Failure		400	{object}	common.APIResponse
//	@Failure		404	{object}	common.APIResponse
//	@Router			/employees/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid employee ID",
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
		"message": "Employee deleted successfully",
	})
}

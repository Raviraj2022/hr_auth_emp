package department

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

// Create Department
//
//	@Summary		Create Department
//	@Description	Create a new Department
//	@Tags			Department
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		department.CreateDepartmentRequest	true	"Department"
//	@Success		201		{object}	common.APIResponse
//	@Failure		400		{object}	common.APIResponse
//	@Router			/departments [post]
func (h *Handler) Create(c *gin.Context) {

	var req CreateDepartmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	department, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Department created successfully",
		"data":    department,
	})
}

// Get All Departments godoc
//
//	@Summary		Get All Departments
//	@Description	Get All Departments
//	@Tags			Department
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	common.APIResponse
//	@Failure		401	{object}	common.APIResponse
//	@Failure		500	{object}	common.APIResponse
//	@Router			/departments [get]
func (h *Handler) GetAll(c *gin.Context) {

	departments, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    departments,
	})
}

// Get All Department By ID godoc
//
//	@Summary		Get All Department By ID
//	@Description	Get All department by ID
//	@Tags			Department
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Department ID"
//	@Success		200	{object}	common.APIResponse
//	@Failure		400	{object}	common.APIResponse
//	@Failure		404	{object}	common.APIResponse
//	@Router			/departments/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid department ID",
		})
		return
	}

	department, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    department,
	})
}

// Update Department godoc
//
//	@Summary		Update Department
//	@Description	Update Department Details
//	@Tags			Department
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string								true	"Department ID"
//	@Param			request	body		department.UpdateDepartmentRequest	true	"Department"
//	@Success		200		{object}	common.APIResponse
//	@Failure		400		{object}	common.APIResponse
//	@Failure		404		{object}	common.APIResponse
//	@Router			/departments/{id} [put]
func (h *Handler) Update(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid department ID",
		})
		return
	}

	var req UpdateDepartmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	department, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Department updated successfully",
		"data":    department,
	})
}

// Delete Department godoc
//
//	@Summary		Delete Department
//	@Description	Delete Department by ID
//	@Tags			Department
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Department ID"
//	@Success		200	{object}	common.APIResponse
//	@Failure		400	{object}	common.APIResponse
//	@Failure		404	{object}	common.APIResponse
//	@Router			/departments/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid department ID",
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
		"message": "Department deleted successfully",
	})
}

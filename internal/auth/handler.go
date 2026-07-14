package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Register godoc
//
//	@Summary		Register User
//	@Description	Register a new user
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		RegisterRequest	true	"Register Request"
//	@Success		201		{object}	Response
//	@Failure		400		{object}	Response
//	@Router			/register [post]
func (h *Handler) Register(c *gin.Context) {

	var req RegisterRequest

	// Read JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	// Call service
	if err := h.service.Register(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User registered successfully",
	})
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login user
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		LoginRequest	true	"Login Request"
//	@Success		200		{object}	Response
//	@Failure		400		{object}	Response
//	@Failure		401		{object}	Response
//	@Router			/login [post]
func (h *Handler) Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	res, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login successful",
		"data":    res,
	})
}

// Profile godoc
//
//	@Summary		Get Profile
//	@Description	Get logged-in user profile
//	@Tags			Authentication
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	Response
//	@Failure		401	{object}	Response
//	@Router			/profile [get]
func (h *Handler) Profile(c *gin.Context) {

	userID := c.GetString("user_id")

	email := c.GetString("email")

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"email":   email,
	})
}

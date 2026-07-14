package dashboard

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

// Get Dashboard godoc
//
// @Summary Dashboard
// @Description Get HRMS dashboard statistics
// @Tags Dashboard
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dashboard.DashboardResponse
// @Failure 500 {object} common.APIResponse
// @Router /dashboard [get]
func (h *Handler) GetDashboard(c *gin.Context) {

	response, err := h.service.GetDashboard()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Dashboard fetched successfully",
		"data":    response,
	})
}
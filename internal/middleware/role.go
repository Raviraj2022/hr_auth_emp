package middleware

import (
	"net/http"
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/ravirajsahu/auth_app/internal/auth"
)

func RequireRoles(roles ...string) gin.HandlerFunc {
    //   fmt.Println(roles);
	return func(c *gin.Context) {

		role := c.GetString("role")
		// fmt.Println(role);

		for _, r := range roles {
            // fmt.Println(role == r)
			if role == r {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Access denied",
		})
	}
}

func AdminOnly() gin.HandlerFunc {
	return RequireRoles(auth.RoleAdmin)
}

func HROnly() gin.HandlerFunc {
	return RequireRoles(auth.RoleHR)
}

func ManagerOnly() gin.HandlerFunc {
	return RequireRoles(auth.RoleManager)
}

func EmployeeOnly() gin.HandlerFunc {
	return RequireRoles(auth.RoleEmployee)
}

func AdminOrHR() gin.HandlerFunc {
	return RequireRoles(
		auth.RoleAdmin,
		auth.RoleHR,
	)
}

func AdminHRManager() gin.HandlerFunc {
	return RequireRoles(
		auth.RoleAdmin,
		auth.RoleHR,
		auth.RoleManager,
	)
}
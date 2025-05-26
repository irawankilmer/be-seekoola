package middleware

import (
	"be-sakoola/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userInterface, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		user, ok := userInterface.(models.User)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "User type assertion failed"})
			return
		}

		roleMatch := false
		for _, role := range user.Roles {
			for _, required := range requiredRoles {
				if role.Name == required {
					roleMatch = true
					break
				}
			}
		}

		if !roleMatch {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: Insufficient role"})
			return
		}

		c.Next()
	}
}

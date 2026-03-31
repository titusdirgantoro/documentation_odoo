package middleware

import (
	"crypto/subtle"
	"documentation_odoo/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireAuth enforces a valid HttpOnly cookie before serving protected static files.
func RequireAuth(vault string, cfg *configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieName := "vault_" + vault
		cookie, err := c.Cookie(cookieName)

		var expectedToken string
		if vault == "impact" {
			expectedToken = cfg.ImpactToken
		} else if vault == "owner" {
			expectedToken = cfg.OwnerToken
		}

		// Check if the cookie exists and strictly matches the expected unique token for this vault
		if err != nil || subtle.ConstantTimeCompare([]byte(cookie), []byte(expectedToken)) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Access denied. Vault locked.",
				"message": "Please return to the main page and authenticate to view this document.",
			})
			return
		}

		c.Next()
	}
}

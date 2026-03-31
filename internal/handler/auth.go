package handler

import (
	"crypto/subtle"
	"documentation_odoo/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Vault    string `json:"vault" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Auth returns a handler for the authentication API endpoint
func (h *Handler) Auth(cfg *configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AuthRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		var correctPassword string
		var cookieName string
		var tokenValue string

		// Map input vault to config properties
		if req.Vault == "impact" {
			correctPassword = cfg.ImpactPassword
			cookieName = "vault_impact"
			tokenValue = cfg.ImpactToken
		} else if req.Vault == "tius" || req.Vault == "owner" {
			correctPassword = cfg.OwnerPassword
			cookieName = "vault_owner"
			tokenValue = cfg.OwnerToken
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown vault"})
			return
		}

		// Validate password securely in backend using constant-time evaluation to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(req.Password), []byte(correctPassword)) != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		// Issue a secure HttpOnly cookie containing the specific vault token
		// Expires in 24 hours (86400 seconds)
		c.SetCookie(cookieName, tokenValue, 86400, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Authentication successful"})
	}
}

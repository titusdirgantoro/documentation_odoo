package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler bundles all HTTP handler methods.
type Handler struct{}

// New returns a new Handler instance.
func New() *Handler {
	return &Handler{}
}

// Index serves the main portal page.
func (h *Handler) Index(c *gin.Context) {
	c.File("./web/index.html")
}

// Health is a liveness probe endpoint.
func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "documentation-portal",
	})
}

// NotFound handles all unmatched routes.
func (h *Handler) NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Route not found",
		"path":  c.Request.URL.Path,
	})
}

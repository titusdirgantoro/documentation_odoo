package server

import (
	"documentation_odoo/configs"
	"documentation_odoo/internal/handler"
	"documentation_odoo/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server wraps the Gin engine and application config.
type Server struct {
	cfg    *configs.Config
	engine *gin.Engine
}

// New creates and returns a fully configured Server.
func New(cfg *configs.Config) *Server {
	engine := gin.Default()
	h := handler.New()
	registerRoutes(engine, h, cfg)

	return &Server{
		cfg:    cfg,
		engine: engine,
	}
}

// Run starts the HTTP server on the configured port.
func (s *Server) Run() error {
	return s.engine.Run(":" + s.cfg.Port)
}

// registerRoutes wires all routes to their handlers.
func registerRoutes(r *gin.Engine, h *handler.Handler, cfg *configs.Config) {
	// ── Health & Auth ─────────────────────────────────────────────────────────
	r.GET("/health", h.Health)
	r.POST("/api/auth", h.Auth(cfg))

	// ── Static asset directories ─────────────────────────────────────────────
	r.Static("/public", "./web/public")   // Public docs

	// Protected directories — Require Cookie Authentication
	impactGroup := r.Group("/impact")
	impactGroup.Use(middleware.RequireAuth("impact", cfg))
	impactGroup.StaticFS("/", http.Dir("./web/impact"))

	ownerGroup := r.Group("/owner")
	ownerGroup.Use(middleware.RequireAuth("owner", cfg))
	ownerGroup.StaticFS("/", http.Dir("./web/owner"))

	// ── Portal root ──────────────────────────────────────────────────────────
	r.GET("/", h.Index)

	// ── 404 fallback ─────────────────────────────────────────────────────────
	r.NoRoute(h.NotFound)
}

package main

import (
	"log"

	"documentation_odoo/configs"
	"documentation_odoo/internal/server"
)

func main() {
	cfg := configs.Load()
	srv := server.New(cfg)

	log.Printf("🚀 Documentation Portal running at http://localhost:%s", cfg.Port)

	if err := srv.Run(); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

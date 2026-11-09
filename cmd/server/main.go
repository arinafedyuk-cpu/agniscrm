package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"agnis-crm/internal/config"
	"agnis-crm/internal/handler"
	"agnis-crm/internal/storage"
	"agnis-crm/internal/usecase"
)

func main() {
	cfg := config.Load()

	// Слои снизу вверх: storage → usecase → handler.
	st := storage.New()
	uc := usecase.New(st)
	h := handler.New(uc)

	router := gin.Default()
	h.Register(router)

	log.Printf("serving HTTP on %s", cfg.Addr)
	if err := router.Run(cfg.Addr); err != nil {
		log.Fatal(err)
	}
}

package handler

import (
	"agnis-crm/internal/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct{ uc *usecase.Usecase }

func New(uc *usecase.Usecase) *Handler { return &Handler{uc: uc} }

func (h *Handler) Register(r *gin.Engine) {
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
		log.Printf("request panic: %v", recovered)
		writeError(c, 500, "internal_error", "внутренняя ошибка сервера")
	}))
	r.HandleMethodNotAllowed = true
	r.NoRoute(func(c *gin.Context) { writeError(c, 404, "route_not_found", "путь не найден") })
	r.NoMethod(func(c *gin.Context) {
		writeError(c, 405, "method_not_allowed", "метод не поддерживается")
	})
	r.GET("/health", h.health)

	contacts := r.Group("/api/v1/contacts")
	contacts.GET("", h.ListContacts)
	contacts.GET("/:id", h.GetContact)
	contacts.POST("", h.CreateContact)
	contacts.POST("/:id/update", h.UpdateContact)
	contacts.POST("/:id/delete", h.DeleteContact)
}

func (h *Handler) health(c *gin.Context) {
	if err := h.uc.Health(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

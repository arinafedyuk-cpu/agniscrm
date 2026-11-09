package handler

import (
        "net/http"

        "github.com/gin-gonic/gin"

        "agnis-crm/internal/usecase"
)

type Handler struct {
        uc *usecase.Usecase
}

func New(uc *usecase.Usecase) *Handler {
        return &Handler{uc: uc}
}

func (h *Handler) Register(r *gin.Engine) {
        r.GET("/health", h.health)
}

func (h *Handler) health(c *gin.Context) {
        if err := h.uc.Health(c.Request.Context()); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
                return
        }
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
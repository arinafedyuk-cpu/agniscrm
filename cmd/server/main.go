package main

import (
        "log"

        "github.com/gin-gonic/gin"

        "agnis-crm/internal/handler"
        "agnis-crm/internal/storage"
        "agnis-crm/internal/usecase"
)

func main() {
        // Слои снизу вверх: storage → usecase → handler.
        st := storage.New()
        uc := usecase.New(st)
        h := handler.New(uc)

        router := gin.Default()
        h.Register(router)

        log.Printf("serving HTTP on :8080")
        if err := router.Run(":8080"); err != nil {
                log.Fatal(err)
        }
}

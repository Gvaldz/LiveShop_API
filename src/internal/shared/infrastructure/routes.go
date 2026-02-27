package infrastructure

import (
    "liveshop_api/src/internal/shared/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

type WSRoutes struct {
    wsController   *controllers.WSController
    authMiddleware gin.HandlerFunc
}

func NewWSRoutes(wsController *controllers.WSController, authMiddleware gin.HandlerFunc) *WSRoutes {
    return &WSRoutes{
        wsController:   wsController,
        authMiddleware: authMiddleware,
    }
}

func (r *WSRoutes) AttachRoutes(router *gin.Engine) {
    router.GET("/ws", r.authMiddleware, r.wsController.Connect)
}
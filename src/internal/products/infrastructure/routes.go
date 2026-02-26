package infrastructure

import (
    "liveshop_api/src/internal/products/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

type ProductRoutes struct {
    controllers    *controllers.ProductControllers
    authMiddleware gin.HandlerFunc
}

func NewProductRoutes(c *controllers.ProductControllers, authMiddleware gin.HandlerFunc) *ProductRoutes {
    return &ProductRoutes{
        controllers:    c,
        authMiddleware: authMiddleware,
    }
}

func (r *ProductRoutes) AttachRoutes(router *gin.Engine) {
    productsGroup := router.Group("/products")
    productsGroup.Use(r.authMiddleware)
    
    productsGroup.POST("", r.controllers.Create)
    productsGroup.GET("", r.controllers.GetAll)
    productsGroup.GET("/:id", r.controllers.GetById)
    productsGroup.PUT("/:id", r.controllers.Update)
    productsGroup.DELETE("/:id", r.controllers.Delete)
}
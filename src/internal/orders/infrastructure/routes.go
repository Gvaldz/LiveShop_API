package infrastructure

import (
	"liveshop_api/src/internal/orders/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type OrderRoutes struct {
	controllers    *controllers.OrderControllers
	authMiddleware gin.HandlerFunc
}

func NewOrderRoutes(c *controllers.OrderControllers, authMiddleware gin.HandlerFunc) *OrderRoutes {
	return &OrderRoutes{
		controllers:    c,
		authMiddleware: authMiddleware,
	}
}

func (r *OrderRoutes) AttachRoutes(router *gin.Engine) {
	ordersGroup := router.Group("/orders")
	ordersGroup.Use(r.authMiddleware)

	ordersGroup.POST("", r.controllers.Create)
	ordersGroup.GET("", r.controllers.GetAll)
	ordersGroup.GET("/:id", r.controllers.GetById)
	ordersGroup.PUT("/:id", r.controllers.Update)
	ordersGroup.DELETE("/:id", r.controllers.Delete)
}

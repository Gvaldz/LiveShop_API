package server

import (
	loginRouters "liveshop_api/src/internal/services/auth/infrastructure"
	userRouters "liveshop_api/src/internal/users/infrastructure"
	productsRouters "liveshop_api/src/internal/products/infrastructure"
	ordersRouters "liveshop_api/src/internal/orders/infrastructure"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(
	authRoutes *loginRouters.AuthRoutes,
	userRoutes *userRouters.UserRoutes,
	productRoutes *productsRouters.ProductRoutes,
	orderRoutes *ordersRouters.OrderRoutes,
) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	authRoutes.AttachRoutes(r)
	userRoutes.AttachRoutes(r)
	productRoutes.AttachRoutes(r)
	orderRoutes.AttachRoutes(r)

	r.Run(":8080")
}

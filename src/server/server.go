package server

import (
	loginRouters "liveshop_api/src/internal/services/auth/infrastructure"
	userRouters "liveshop_api/src/internal/users/infrastructure"
	priductsRouters "liveshop_api/src/internal/products/infrastructure"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(
	authRoutes *loginRouters.AuthRoutes,
	userRoutes *userRouters.UserRoutes,
	productRoutes *priductsRouters.ProductRoutes,
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

	r.Run(":8080")
}

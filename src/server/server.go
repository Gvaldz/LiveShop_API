package server

import (
	loginRouters "liveshop_api/src/internal/services/auth/infrastructure"
	userRouters "liveshop_api/src/internal/users/infrastructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(
	patientRoutes *patientRouters.PatientRoutes,
	authRoutes *loginRouters.AuthRoutes,
	userRoutes *userRouters.UserRoutes,
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

	r.Run(":8080")
}

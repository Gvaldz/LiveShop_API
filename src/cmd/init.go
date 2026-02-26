package cmd

import (
	"liveshop_api/src/core"
	productsDeps "liveshop_api/src/internal/products/infrastructure"
	loginDeps "liveshop_api/src/internal/services/auth/infrastructure"
	usersDeps "liveshop_api/src/internal/users/infrastructure"
	"liveshop_api/src/server"
	"liveshop_api/src/server/middleware"
	"log"
)

func Init() {
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	hasher := core.NewBcryptHasher(14)
	tokenService := core.NewJWTService()

	userRepo := usersDeps.NewUserRepository(db)
	authRepo := loginDeps.NewAuthRepository(db)

	authMiddleware := middleware.AuthMiddleware(tokenService, userRepo)

	productDependencies := productsDeps.NewProductDependencies(db, authMiddleware)
	productsRoutes := productDependencies.GetRoutes()

	userDependencies := usersDeps.NewUserDependencies(
		db,
		hasher,
		tokenService,
		authRepo,
		userRepo,
	)
	userRoutes := userDependencies.GetRoutes()

	authDependencies := loginDeps.NewAuthDependencies(db, hasher, userRepo)
	authRoutes := authDependencies.GetRoutes()

	server.Run(authRoutes, userRoutes, productsRoutes)
}

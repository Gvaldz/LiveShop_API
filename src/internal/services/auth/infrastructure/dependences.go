package infrastructure

import (
	"database/sql"
	"liveshop_api/src/core"
	"liveshop_api/src/internal/services/auth/application"
	"liveshop_api/src/internal/services/auth/infrastructure/controllers"
	users_domain "liveshop_api/src/internal/users/domain"
)

type AuthDependencies struct {
	DB       *sql.DB
	Hasher   *core.BcryptHasher
	UserRepo users_domain.UserRepository
}

func NewAuthDependencies(db *sql.DB, hasher *core.BcryptHasher, userRepo users_domain.UserRepository) *AuthDependencies {
	return &AuthDependencies{
		DB:       db,
		Hasher:   hasher,
		UserRepo: userRepo,
	}
}

func (d *AuthDependencies) GetRoutes() *AuthRoutes {
	authRepo := NewAuthRepository(d.DB)
	tokenService := core.NewJWTService()

	loginUC := application.NewLogin(
		authRepo,
		d.UserRepo,
		tokenService,
		d.Hasher,
	)

	loginController := controllers.NewLoginController(loginUC)

	return NewAuthRoutes(loginController)
}

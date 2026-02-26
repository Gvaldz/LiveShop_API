package domain

import (
	user "liveshop_api/src/internal/users/domain/entities"
)

type AuthRepository interface {
	FindUserByEmail(email string) (user.User, error)
	UpdateLastLogin(userID int32) error
}

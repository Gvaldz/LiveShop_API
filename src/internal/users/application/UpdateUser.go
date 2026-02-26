package application

import (
	"liveshop_api/src/internal/users/domain"
	"liveshop_api/src/internal/users/domain/entities"
)

type UpdateUser struct {
	userRepo domain.UserRepository
}

func NewUpdateUser(userRepo domain.UserRepository) *UpdateUser {
	return &UpdateUser{userRepo: userRepo}
}

func (uc *UpdateUser) Execute(id int32, user entities.User) error {
	user.Password = ""
	return uc.userRepo.UpdateUser(id, user)
}

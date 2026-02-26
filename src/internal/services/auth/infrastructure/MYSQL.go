package infrastructure

import (
	"database/sql"
	"fmt"
	"liveshop_api/src/internal/services/auth/domain"
	user "liveshop_api/src/internal/users/domain/entities"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(DB *sql.DB) domain.AuthRepository {
	return &AuthRepository{DB: DB}
}

func (r *AuthRepository) FindUserBynumber(number string) (user.User, error) {
	var u user.User
	query := "SELECT iduser, number, password FROM users WHERE number = ?"

	err := r.DB.QueryRow(query, number).Scan(&u.IdUser, &u.Number, &u.Password)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *AuthRepository) UpdateLastLogin(userID int32) error {
	query := "UPDATE users SET ultimo_login = NOW() WHERE iduser = ?"
	_, err := r.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error updating last login: %w", err)
	}
	return nil
}

func (r *AuthRepository) FindUserByID(userID int32) (user.User, error) {
	var user user.User
	query := `SELECT iduser, number, password, FROM users WHERE iduser = ?`
	err := r.DB.QueryRow(query, userID).Scan(&user.IdUser, &user.Number, &user.Password)
	return user, err
}

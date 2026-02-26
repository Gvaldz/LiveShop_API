package infrastructure

import (
	"database/sql"
	"fmt"
	users "liveshop_api/src/internal/users/domain"
	user "liveshop_api/src/internal/users/domain/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) users.UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) CreateUser(u user.User) (user.User, error) {

	query := "INSERT INTO users (name, password, number) VALUES (?, ?, ?, ?)"

	result, err := r.DB.Exec(query, u.Name, u.Password, u.Number)
	if err != nil {
		return user.User{}, fmt.Errorf("error al crear usuario: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user.User{}, fmt.Errorf("error al obtener ID: %w", err)
	}

	return user.User{
		IdUser: int32(id),
		Name:   u.Name,
		Number: u.Number,
	}, nil
}

func (r *UserRepository) GetAllUsers() ([]user.User, error) {
	query := "SELECT iduser, name, number FROM users"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %w", err)
	}
	defer rows.Close()

	var usersList []user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.IdUser, &u.Name, &u.Number); err != nil {
			return nil, fmt.Errorf("error al escanear user: %w", err)
		}
		usersList = append(usersList, u)
	}

	return usersList, nil
}

func (r *UserRepository) GetUserByID(iduser int32) (user.User, error) {
	if r.DB == nil {
		return user.User{}, fmt.Errorf("database connection is nil")
	}

	var u user.User
	query := "SELECT iduser, name, number FROM users WHERE iduser = ?"

	err := r.DB.QueryRow(query, iduser).Scan(&u.IdUser, &u.Name, &u.Number)

	if err != nil {
		return u, fmt.Errorf("error al obtener usuario: %w", err)
	}
	return u, nil
}

func (r *UserRepository) GetUserBynumber(number string) (user.User, error) {
	var u user.User
	query := "SELECT iduser, name, number, password FROM users WHERE number = ?"

	err := r.DB.QueryRow(query, number).Scan(
		&u.IdUser,
		&u.Name,
		&u.Password,
		&u.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("usuario no encontrado")
		}
		return u, fmt.Errorf("error al obtener usuario por number: %w", err)
	}
	return u, nil
}

func (r *UserRepository) UpdateUser(id int32, u user.User) error {
	query := "UPDATE users SET name = ?, number = ? WHERE iduser = ?"
	result, err := r.DB.Exec(query, u.Name, u.Number, id)
	if err != nil {
		return fmt.Errorf("error al actualizar usuario: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar actualización: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario no encontrado")
	}

	return nil
}

func (r *UserRepository) UpdatePassword(id int32, newHashedPassword string) error {
	query := "UPDATE users SET password = ? WHERE iduser = ?"
	result, err := r.DB.Exec(query, newHashedPassword, id)
	if err != nil {
		return fmt.Errorf("error al actualizar contraseña: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar actualización de contraseña: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario no encontrado")
	}

	return nil
}

func (r *UserRepository) DeleteUser(id int32) error {
	query := "DELETE FROM users WHERE iduser = ?"
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar eliminación: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario no encontrado")
	}

	return nil
}

package user

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type CreateUserParam struct {
	Name     string
	Email    string
	Password string
}

func (r *Repository) Create(cnt context.Context, param CreateUserParam) error {
	hashedpasswrd, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("ошибка хеширования пароля: %w", err)
	}

	query := `INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)`
	_, err = r.db.ExecContext(cnt, query, param.Name, param.Email, string(hashedpasswrd))
	if err != nil {
		return fmt.Errorf("ошибка сохранения пользователя %w", err)
	}
	return nil
}

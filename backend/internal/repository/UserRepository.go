package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/yashpatil74/nimbus/internal/domain/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) EmailExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	err := ur.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %v", err)
	}

	return exists, nil
}

func (ur *UserRepository) CreateUser(user *entities.User) error {
	tx, err := ur.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	exists, err := ur.EmailExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("email already exists")
	}

	user.ID = uuid.New().String()

	query := `INSERT INTO users (id, email, password, created_at, updated_at) 
    	      VALUES (?, ?, ?, DATETIME('now'), DATETIME('now'))`

	_, err = tx.Exec(query, user.ID, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	query := `SELECT id, email, password, created_at, updated_at 
          FROM users 
          WHERE email = ?`

	err := ur.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}

	return &user, nil
}

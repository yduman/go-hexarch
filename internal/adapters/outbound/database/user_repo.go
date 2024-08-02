package database

import (
	"database/sql"

	"github.com/yduman/go-hexarch/internal/domain/entity"
	"github.com/yduman/go-hexarch/internal/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) Save(user *entity.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := repo.db.Exec(query, user.Name, user.Email)
	return err
}

func (repo *userRepository) FindByID(id string) (*entity.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := repo.db.QueryRow(query, id)

	var user entity.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

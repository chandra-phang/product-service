package repositories

import (
	"database/sql"
	"shop-api/models"

	"github.com/labstack/echo/v4"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepositoryInstance(db *sql.DB) models.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(ctx echo.Context, user models.User) error {
	sqlStatement := `
		INSERT INTO users
			(id, name, email, password, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?)
	`

	params := []interface{}{
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	}

	_, err := r.db.Exec(sqlStatement, params...)
	if err != nil {
		return err
	}

	return nil
}

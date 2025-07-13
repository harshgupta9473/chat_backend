package repositories

import (
	"context"
	"database/sql"
	"github.com/harshgupta9473/chatapp/internal/userservice/dto"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	repo := &UserRepository{}
	repo.DB = db
	return repo
}

func (r *UserRepository) CreateUser(ctx context.Context, user *dto.User) error {
	_, err := r.DB.ExecContext(ctx, `
		INSERT INTO users (name, mobile, password, created_at)
		VALUES ($1, $2, $3, NOW())`,
		user.Name, user.Mobile, user.Password)
	return err
}

func (r *UserRepository) GetUserByMobile(ctx context.Context, mobile string) (*dto.User, error) {
	row := r.DB.QueryRowContext(ctx, `SELECT id, name, mobile, password, created_at FROM users WHERE mobile = $1`, mobile)
	user := &dto.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Mobile, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

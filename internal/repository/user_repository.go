package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jordanmarcelino/terradiscover-backend/internal/entity"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
}

type userRepositoryImpl struct {
	db dbtx
}

func NewUserRepository(db dbtx) *userRepositoryImpl {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	query := `
		SELECT email, hash_password FROM users
		WHERE id = $1
	`

	user := &entity.User{ID: id}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, hash_password FROM users
		WHERE email = $1
	`

	user := &entity.User{Email: email}
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepositoryImpl) Save(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO users(email, hash_password) VALUES ($1, $2) RETURNING id
	`

	return r.db.QueryRowContext(ctx, query, user.Email, user.Password).Scan(&user.ID)
}

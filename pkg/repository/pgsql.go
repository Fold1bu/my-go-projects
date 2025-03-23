package repository

import (
	"context"
	"go-project/pkg/models"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PGRepo struct {
	mu   sync.Mutex
	pool *pgxpool.Pool
}

func New(connStr string) (*PGRepo, error) {
	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &PGRepo{mu: sync.Mutex{}, pool: pool}, nil
}

// Функция на добавление пользователя
func (repo *PGRepo) CreateUser(user *models.User) error {
	query := `INSERT INTO users (user_name, user_number) VALUES ($1, $2) RETURNING user_id`
	err := repo.pool.QueryRow(context.Background(), query, user.Name, user.Number).Scan(&user.ID)
	return err
}



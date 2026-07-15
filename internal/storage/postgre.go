package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Storage struct {
	conn *pgx.Conn
}

func NewStorage() (*Storage, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("переменная DB_URL не найдена в .env")
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	fmt.Println("Успешно подключение к PSQL")
	return &Storage{conn: conn}, nil
}

func (s *Storage) Close() {
	s.conn.Close(context.Background())
}

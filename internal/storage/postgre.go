package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type Storage struct {
	db *sql.DB
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

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия соединения: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("база недоступна: %w", err)
	}

	fmt.Println("Успешно подключились к PostgreSQL!")
	return &Storage{db: db}, nil
}

func (s *Storage) Migrate() error {
	// 1. Создаем специальный драйвер для migrate на основе нашего соединения
	driver, err := postgres.WithInstance(s.db, &postgres.Config{})
	if err != nil {
		return err
	}

	// 2. Передаем этот драйвер в migrate
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	fmt.Println("Миграции успешно применены!")
	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}
func (s *Storage) GetDB() *sql.DB {
	return s.db
}

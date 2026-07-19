package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file" 
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

func (s *Storage) Migrate() error{
	driver,err:=pgx.WithInstance(s.conn,&pgx.Config{})
	if err!=nil{
		return err
	}

	m,err:=migrate.NewWithDataBaseInstance("file://migrations","postgres",driver)
	if err!=nil{
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	
	fmt.Println("Миграции успешно применены!")
	return nil

}

func (s *Storage) Close() {
	s.conn.Close(context.Background())
}

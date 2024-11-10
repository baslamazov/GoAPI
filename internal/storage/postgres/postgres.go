package postgres

import (
	"GoAPI/internal/config"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBService struct {
	db *pgxpool.Pool
}

func NewDBService(cfg config.Database) *DBService {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	return &DBService{db: CreatePool(connectionString)}
}

func CreatePool(constr string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(constr)

	if err != nil {
		log.Fatal(err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	return pool
}

func (db *DBService) GetAllUsers() ([]string, error) {
	return []string{"пидор"}, nil
}
func (db *DBService) CheckUser(login, password string) (bool, error) {
	var isExist bool
	db.db.Ping(context.Background())
	err := db.db.QueryRow(context.Background(), `  
        SELECT EXISTS (  
            SELECT 1  
            FROM users
            WHERE login = $1 AND password = $2  
        );  
    `, login, password).Scan(&isExist)

	if err != nil {
		return false, err // Возвращаем ошибку, если она произошла
	}

	return isExist, nil // Возвращаем результат проверки
}

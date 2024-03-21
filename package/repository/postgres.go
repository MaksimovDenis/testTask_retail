package repository

import (
	"fmt"
	logger "testTask_retail/logs"

	"github.com/jmoiron/sqlx"
)

const (
	productsTable = "products"
	shelvesTable  = "shelves"
	ordersTable   = "orders"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	//Check our connection to DB
	err = db.Ping()
	if err != nil {
		logger.Log.Errorf("Failed to connect db (repository)")
		return nil, err
	}
	return db, nil
}

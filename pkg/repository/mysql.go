package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Конфиг с описанием необходимых данных для подключения к базе данных.
type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

// В этой функции мы подключаемся к базе данных и проверяем работоспособность подключения.
func NewMysqlDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

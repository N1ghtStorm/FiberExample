package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

//NewPostgresDB sas
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if db == nil {
		println("4")
	}

	if err != nil {
		println("4TOTO NE TAK: " + err.Error())
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		println("4TOTO NE TAK: " + pingErr.Error())
	}

	return db, nil
}

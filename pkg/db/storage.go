package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	cfg "gitlab.geogracom.com/skdf/skdf-abac-go/configs/app"

	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

type DB struct {
	IItemsDB
}

func ConnectDatabase(config cfg.PostgresConfig) (*DB, error) {
	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)

	db, err := sqlx.Open(driverName, connInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{
		IItemsDB: NewAccessUserDB(db),
	}, nil
}

func (db *DB) Close() error {
	if err := db.IItemsDB.CloseDB(); err != nil {
		return err
	}

	return nil
}

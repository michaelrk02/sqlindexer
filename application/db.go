package application

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/michaelrk02/sqlindexer/config"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sqlx.DB
}

func Connect(cfg *config.DB) (*DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name))
	if err != nil {
		return nil, err
	}
	return &DB{DB: db.Unsafe()}, nil
}

func (db *DB) GetTables() ([]string, error) {
	var tables []string

	err := db.Select(&tables, "SHOW TABLES")
	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (db *DB) GetTableColumns(table string) ([]string, error) {
	var cols []struct {
		Field string `db:"Field"`
	}

	err := db.Select(&cols, fmt.Sprintf("DESCRIBE `%s`", table))
	if err != nil {
		return nil, err
	}

	columns := make([]string, len(cols))
	for i, col := range cols {
		columns[i] = col.Field
	}

	return columns, nil
}

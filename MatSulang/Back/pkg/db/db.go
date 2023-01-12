package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBHandler struct {
	DB *sql.DB
}

func ConnectDB(dsn string) (*DBHandler, error) {
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("DB Connection Error")
	}

	dbHandler := &DBHandler{
		DB: DB,
	}

	return dbHandler, err
}

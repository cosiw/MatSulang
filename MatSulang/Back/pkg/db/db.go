package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

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
func (db DBHandler) SelectData(tName string, colums, condition string) (string, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", colums, tName, condition)
	fmt.Println(query)
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Printf("SELECT Error!, %v", err)
		return "", err
	}
	defer rows.Close()

	jsonData, err := GetJSONFromRows(rows)
	if err != nil {

		return "", err
	}

	return jsonData, nil
}

func GetJSONFromRows(rows *sql.Rows) (string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
func removeArrChars(arr string) string {
	arr = strings.Replace(arr, "[", "", 1)
	arr = strings.Replace(arr, "]", "", 1)
	return arr
}

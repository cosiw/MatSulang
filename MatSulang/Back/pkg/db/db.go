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
func (db DBHandler) InsertData(tName string, columns string, values string) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tName, columns, values)
	fmt.Println(query)
	res, err := db.DB.Exec(query)

	if err != nil {
		fmt.Printf("INSERT Error!, %v", err)
		return 0, err
	}
	id, err := res.LastInsertId()

	if err != nil {
		fmt.Printf("Failed to check the last inserted date (%s)", err)
		return id, err
	}
	return id, nil
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

func (db DBHandler) DeleteData(table string, cond string) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", table, cond)
	fmt.Println(query)
	res, err := db.DB.Exec(query)
	if err != nil {
		fmt.Printf("DELETE QUERY ERROR! %v", err)
	}

	rowsCnt, err := res.RowsAffected()

	if err != nil {
		fmt.Printf("RowsAffected Error! %v", err)
	}

	if rowsCnt <= 0 {
		fmt.Printf("DELETE COUNT 0!")
	}

	return rowsCnt, err

}

func (db DBHandler) UpdateData(table string, value string, cond string) (int64, error) {
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, value, cond)

	res, err := db.DB.Exec(query)
	if err != nil {
		fmt.Printf("UPDATE Error : %v", err)
		return 0, err
	}

	n, err := res.RowsAffected()

	if err != nil {
		fmt.Printf("RowsAffected Error : %v", err)
		return 0, err
	}

	return n, nil

}

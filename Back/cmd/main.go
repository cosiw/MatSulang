package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbHandler, err := sql.Open("mysql", "admin:!mgrsol123@tcp(matsulang.cwgugsk0vnht.ap-northeast-2.rds.amazonaws.com:3306)/Matsulang")

	if err != nil {
		//fmt.Println("Failed to connect DB : %v", err)
	}
	query := fmt.Sprintf("SELECT MEMBERID, USERID FROM TMEMBER")

	a, err := dbHandler.Query(query)
	if err != nil {
		fmt.Println("SELECT Error")
	}

	for a.Next() {
		MEMBERID := ""
		USERID := ""
		if err := a.Scan(&MEMBERID, &USERID); err != nil {
			fmt.Println("err : ", err)
		}
		fmt.Println(MEMBERID, USERID)
	}
	a.Close()
	dbHandler.Close()
}

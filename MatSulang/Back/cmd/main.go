package main

import (
	"Back/pkg/api"
	"Back/pkg/db"
	"Back/pkg/router"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbHandler, err := db.ConnectDB("admin:!mgrsol123@tcp(matsulang.cwgugsk0vnht.ap-northeast-2.rds.amazonaws.com:3306)/Matsulang")

	if err != nil {
		fmt.Printf("Failed to connect DB : %v", err)
	}

	apis := api.NewAPI(dbHandler)
	r := router.Router(apis)

	r.Run(":8085")
}

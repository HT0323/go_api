package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HT0323/go_api/controllers"
	"github.com/HT0323/go_api/routers"
	"github.com/HT0323/go_api/services"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	ser := services.MewMyAppService(db)
	con := controllers.NewMyAppController(ser)

	r := routers.NewRouter(con)
	fmt.Println(r.Headers())
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

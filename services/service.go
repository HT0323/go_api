package services

import (
	"database/sql"
	"fmt"
)

type MyAppService struct {
	db *sql.DB
}

func NewMyAppService(db *sql.DB) *MyAppService {
	fmt.Println(&MyAppService{db: db})
	return &MyAppService{db: db}
}

package services

import "database/sql"

type MyAppService struct {
	db *sql.DB
}

func MewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}

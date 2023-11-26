package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(host string, port string, username, password, dbname string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		username,
		password,
		dbname)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

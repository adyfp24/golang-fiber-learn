package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func InitDB() (*gorm.DB, error) {

	dsn := "root:@tcp(127.0.0.1:3306)/db_go_fiber_learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Koneksi terhubung")

	DB = db

	return db, nil
}

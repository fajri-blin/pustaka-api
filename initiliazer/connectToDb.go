package initiliazer

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error){
	dsn := "root:@tcp(127.0.0.1:3306)/pustakapi?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
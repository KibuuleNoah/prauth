package database

import (
	"context"
	"log"
	"prauth/models"
  "gorm.io/driver/mysql"

	"gorm.io/gorm"
)


func InitDB() (*gorm.DB,context.Context) {
	
	dsn := "root:noahtri@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect:", err)
	}
	
	ctx := context.Background()

	db.AutoMigrate(dbModels()...)

	return db, ctx
}

func dbModels() []any{
	return []any{
		&models.User{},
	}
}

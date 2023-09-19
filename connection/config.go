package connection

import (
	"crud/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DataMigration() {
	dsn := "root:@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	database.AutoMigrate(models.Person{})
	DB = database
}

package main

import (
	"crud/connection"
	"crud/routes"
	"fmt"
)

func main() {
	// Initialize the database connection
	// dsn := "root:@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("Failed to connect to the database:", err)
	// 	return
	// }

	// controllers.SetDB(db)

	// Initialize and start the HTTP server
	connection.DataMigration()
	router := routes.InitializeRoutes()
	port := 8000
	fmt.Printf("Server started on :%d\n", port)
	router.Run(fmt.Sprintf(":%d", port))
}

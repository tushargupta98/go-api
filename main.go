package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/tushargupta98/todoapi/database"
	"github.com/tushargupta98/todoapi/handler"
	"log"
	"net/http"
)

func main() {
	// Initialize the Chi router
	router := gin.New()
	// Create an HTTP server
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	router.GET("/health", handler.HealthHandler)
	router.GET("/todo", handler.GeAllToDoHandler)
	router.POST("/todo", handler.CreateToDoHandler)

	dsn := fmt.Sprintf("host=postgres_todo_db port=5432 user=todo_user password=todo_password dbname=todo_db sslmode=disable")

	_, err := database.InitializeDB(dsn)
	if err != nil {
		log.Println("Failed to initialize the database: %v", err)
	}
	// Close the database connection when the application exits
	err = server.ListenAndServe()
	if err != nil {
		log.Panic("Application did not start on the given address:", err)
	}
}

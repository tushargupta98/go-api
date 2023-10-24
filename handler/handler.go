package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tushargupta98/todoapi/database"
	"github.com/tushargupta98/todoapi/model"
)

func HealthHandler(c *gin.Context) {
	fmt.Println("Health success")
	c.JSON(200, "Success")
}

func CreateToDoHandler(c *gin.Context) {
	var todo model.ToDo
	c.BindJSON(&todo)
	database.DB.Create(&todo)
	fmt.Println("DB Connected")
	c.JSON(200, &todo.ID)
}

func GeAllToDoHandler(c *gin.Context) {
	var todos []model.ToDo
	database.DB.Find(&todos)
	c.JSON(200, &todos)
}

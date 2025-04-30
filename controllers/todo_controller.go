package controllers

import (
	"go-personal-boilerplate/models"
	"go-personal-boilerplate/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodoHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	result, err := services.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetTodosHandler(c *gin.Context) {
	todos, err := services.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func UpdateTodoHandler(c *gin.Context) {
	id := c.Param("id") // ID tetap sebagai string

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := services.UpdateTodo(id, &todo) // Gunakan string untuk ID
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

func UpdateTodoStatusHandler(c *gin.Context) {
	id := c.Param("id") // ID tetap sebagai string

	updatedTodo, err := services.UpdateTodoStatus(id) // Gunakan string untuk ID
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

func DeleteTodoHandler(c *gin.Context) {
	id := c.Param("id") // ID tetap sebagai string

	if err := services.DeleteTodo(id); // Gunakan string untuk ID
	err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.Status(http.StatusNoContent)
}

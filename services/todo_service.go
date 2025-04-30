package services

import (
	"go-personal-boilerplate/models"
	"go-personal-boilerplate/utils"
)

func CreateTodo(todo *models.Todo) (*models.Todo, error) {
	db := utils.InitDB()

	err := db.Create(todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func GetTodos() ([]models.Todo, error) {
	db := utils.InitDB()

	var todos []models.Todo
	err := db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func UpdateTodo(id string, todo *models.Todo) (*models.Todo, error) { // Ganti uint ke string
	db := utils.InitDB()

	var existingTodo models.Todo
	if err := db.First(&existingTodo, "id = ?", id).Error; err != nil { // Gunakan string untuk pencarian
		return nil, err
	}

	existingTodo.Title = todo.Title
	existingTodo.Status = todo.Status
	err := db.Save(&existingTodo).Error
	if err != nil {
		return nil, err
	}
	return &existingTodo, nil
}

func UpdateTodoStatus(id string) (*models.Todo, error) { // Ganti uint ke string
	db := utils.InitDB()

	var existingTodo models.Todo
	if err := db.First(&existingTodo, "id = ?", id).Error; err != nil { // Gunakan string untuk pencarian
		return nil, err
	}

	existingTodo.Status = "success"
	err := db.Save(&existingTodo).Error
	if err != nil {
		return nil, err
	}
	return &existingTodo, nil
}

func DeleteTodo(id string) error { // Ganti uint ke string
	db := utils.InitDB()

	err := db.Delete(&models.Todo{}, "id = ?", id).Error // Gunakan string untuk penghapusan
	if err != nil {
		return err
	}
	return nil
}

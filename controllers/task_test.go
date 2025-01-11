package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-api/database"
	"task-api/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	database.DB, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	database.DB.AutoMigrate(&models.Task{}) // Migrate Task model
	router := gin.Default()
	return router
}

func TestGetAllTasks(t *testing.T) {
	router := setupRouter()
	router.GET("/tasks", GetAllTasks)

	database.DB.Create(&models.Task{Title: "Task 1", Description: "Test Task 1", Status: "pending"})
	database.DB.Create(&models.Task{Title: "Task 2", Description: "Test Task 2", Status: "in-progress"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var tasks []models.Task
	err := json.Unmarshal(w.Body.Bytes(), &tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 2) // Should return 2 tasks
}

func TestGetTaskByID(t *testing.T) {
	router := setupRouter()
	router.GET("/tasks/:id", GetTaskByID)

	task := &models.Task{Title: "Test Task", Description: "Test Task Desc", Status: "pending"}
	database.DB.Create(task)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var returnedTask models.Task
	err := json.Unmarshal(w.Body.Bytes(), &returnedTask)
	assert.NoError(t, err)
	assert.Equal(t, "Test Task", returnedTask.Title)
	assert.Equal(t, "Test Task Desc", returnedTask.Description)
}

func TestCreateTask(t *testing.T) {
	router := setupRouter()
	router.POST("/tasks", CreateTask)

	task := models.Task{Title: "New Task", Description: "New task description", Status: "pending"}
	taskJSON, _ := json.Marshal(task)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewReader(taskJSON))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdTask models.Task
	err := json.Unmarshal(w.Body.Bytes(), &createdTask)
	assert.NoError(t, err)
	assert.Equal(t, "New Task", createdTask.Title)
}

func TestUpdateTask(t *testing.T) {
	router := setupRouter()
	router.PUT("/tasks/:id", UpdateTask)

	task := &models.Task{Title: "Test Task", Description: "Test Task Desc", Status: "pending"}
	database.DB.Create(task)

	updatedTask := models.Task{Title: "Updated Task", Description: "Updated description", Status: "in-progress"}
	updatedTaskJSON, _ := json.Marshal(updatedTask)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewReader(updatedTaskJSON))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var taskAfterUpdate models.Task
	err := json.Unmarshal(w.Body.Bytes(), &taskAfterUpdate)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", taskAfterUpdate.Title)
	assert.Equal(t, "Updated description", taskAfterUpdate.Description)
}

func TestDeleteTask(t *testing.T) {
	router := setupRouter()
	router.DELETE("/tasks/:id", DeleteTask)

	task := &models.Task{Title: "Test Task", Description: "Test Task Desc", Status: "pending"}
	database.DB.Create(task)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	var deletedTask models.Task
	err := database.DB.First(&deletedTask, 1).Error
	assert.Error(t, err) // Task should not be found
}

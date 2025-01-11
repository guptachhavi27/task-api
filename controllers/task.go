package controllers

import (
	"net/http"
	"strconv"
	"task-api/database"
	"task-api/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	// var tasks []models.Task
	// database.DB.Find(&tasks)
	// c.JSON(http.StatusOK, tasks)
	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) // Default to page 1
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10")) // Default to 10 items per page
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	var tasks []models.Task
	offset := (page - 1) * pageSize

	// Query the database with pagination
	result := database.DB.Limit(pageSize).Offset(offset).Find(&tasks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	// Get the total count of tasks for pagination metadata
	var total int64
	database.DB.Model(&models.Task{}).Count(&total)

	// Return paginated response
	c.JSON(http.StatusOK, gin.H{
		"tasks":       tasks,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize), // Calculate total pages
	})
}

func GetTaskByID(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&task)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest,result.Error.Error())
	}else{
		c.JSON(http.StatusCreated, task)
	}
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	database.DB.Delete(&task)
	c.JSON(http.StatusNoContent, nil)
}

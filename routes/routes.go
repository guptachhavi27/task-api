package routes

import (
	"task-api/controllers"
	"task-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	public := r.Group("/public")
	public.GET("/tasks", controllers.GetAllTasks)

	protected := r.Group("/tasks")
	protected.Use(middleware.Authenticate())
	protected.GET("/:id", controllers.GetTaskByID)
	protected.POST("/", controllers.CreateTask)
	protected.PUT("/:id", controllers.UpdateTask)
	protected.DELETE("/:id", controllers.DeleteTask)

	return r
}

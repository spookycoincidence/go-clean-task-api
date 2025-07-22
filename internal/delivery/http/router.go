package http

import (
    "github.com/gin-gonic/gin"
    "go-clean-template/internal/usecase"
)

func SetupRouter(taskUseCase *usecase.TaskUseCase) *gin.Engine {
    r := gin.Default()

    taskHandler := NewTaskHandler(taskUseCase)

    r.POST("/tasks", taskHandler.Create)
    r.GET("/tasks", taskHandler.GetAll)
    r.GET("/tasks/completed", taskHandler.GetCompleted) // Nuevo endpoint
    r.GET("/tasks/:id", taskHandler.GetByID)
    r.PUT("/tasks/:id", taskHandler.Update)
    r.DELETE("/tasks/:id", taskHandler.Delete)

    return r
}

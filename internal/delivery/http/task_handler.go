package http

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "go-clean-template/internal/entity"
    "go-clean-template/internal/usecase"
    "net/http"
)

type TaskHandler struct {
    usecase *usecase.TaskUseCase
}

func NewTaskHandler(usecase *usecase.TaskUseCase) *TaskHandler {
    return &TaskHandler{usecase: usecase}
}

func (h *TaskHandler) Create(c *gin.Context) {
    var task entity.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.ID = uuid.New()

    if err := h.usecase.Create(c.Request.Context(), &task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetByID(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    task, err := h.usecase.GetByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
        return
    }

    c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) GetAll(c *gin.Context) {
    tasks, err := h.usecase.GetAll(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) Update(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    var task entity.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.ID = id

    if err := h.usecase.Update(c.Request.Context(), &task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Delete(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    if err := h.usecase.Delete(c.Request.Context(), id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}

// Nuevo handler para tareas completadas
func (h *TaskHandler) GetCompleted(c *gin.Context) {
    tasks, err := h.usecase.GetCompletedTasks(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

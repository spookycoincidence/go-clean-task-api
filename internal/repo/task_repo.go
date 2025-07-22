package repo

import (
    "context"
    "github.com/google/uuid"
    "go-clean-template/internal/entity"
)

type TaskRepository interface {
    Create(ctx context.Context, task *entity.Task) error
    GetByID(ctx context.Context, id uuid.UUID) (*entity.Task, error)
    GetAll(ctx context.Context) ([]*entity.Task, error)
    Update(ctx context.Context, task *entity.Task) error
    Delete(ctx context.Context, id uuid.UUID) error

    // Nuevo método para filtrar completadas
    GetTasksByCompletion(ctx context.Context, completed bool) ([]*entity.Task, error)
}

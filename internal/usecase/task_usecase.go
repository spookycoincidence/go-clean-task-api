package usecase

import (
    "context"
    "github.com/google/uuid"
    "go-clean-template/internal/entity"
    "go-clean-template/internal/repository"
)

type TaskUseCase struct {
    repo repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) *TaskUseCase {
    return &TaskUseCase{repo: repo}
}

func (uc *TaskUseCase) Create(ctx context.Context, task *entity.Task) error {
    return uc.repo.Create(ctx, task)
}

func (uc *TaskUseCase) GetByID(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
    return uc.repo.GetByID(ctx, id)
}

func (uc *TaskUseCase) GetAll(ctx context.Context) ([]*entity.Task, error) {
    return uc.repo.GetAll(ctx)
}

func (uc *TaskUseCase) Update(ctx context.Context, task *entity.Task) error {
    return uc.repo.Update(ctx, task)
}

func (uc *TaskUseCase) Delete(ctx context.Context, id uuid.UUID) error {
    return uc.repo.Delete(ctx, id)
}

// Nuevo método para tareas completadas
func (uc *TaskUseCase) GetCompletedTasks(ctx context.Context) ([]*entity.Task, error) {
    return uc.repo.GetTasksByCompletion(ctx, true)
}

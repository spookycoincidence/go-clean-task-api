package repository

import (
    "context"
    "database/sql"
    "errors"

    "github.com/google/uuid"
    "go-clean-template/internal/entity"
)

type PostgresTaskRepository struct {
    db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) *PostgresTaskRepository {
    return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) Create(ctx context.Context, task *entity.Task) error {
    query := `INSERT INTO tasks (id, title, completed) VALUES ($1, $2, $3)`
    _, err := r.db.ExecContext(ctx, query, task.ID, task.Title, task.Completed)
    return err
}

func (r *PostgresTaskRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
    query := `SELECT id, title, completed FROM tasks WHERE id = $1`
    row := r.db.QueryRowContext(ctx, query, id)

    task := &entity.Task{}
    var idStr string

    err := row.Scan(&idStr, &task.Title, &task.Completed)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("task not found")
        }
        return nil, err
    }

    task.ID, err = uuid.Parse(idStr)
    if err != nil {
        return nil, err
    }

    return task, nil
}

func (r *PostgresTaskRepository) GetAll(ctx context.Context) ([]*entity.Task, error) {
    query := `SELECT id, title, completed FROM tasks`
    rows, err := r.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []*entity.Task
    for rows.Next() {
        task := &entity.Task{}
        var idStr string

        err := rows.Scan(&idStr, &task.Title, &task.Completed)
        if err != nil {
            return nil, err
        }

        task.ID, err = uuid.Parse(idStr)
        if err != nil {
            return nil, err
        }

        tasks = append(tasks, task)
    }

    return tasks, nil
}

func (r *PostgresTaskRepository) Update(ctx context.Context, task *entity.Task) error {
    query := `UPDATE tasks SET title = $1, completed = $2 WHERE id = $3`
    result, err := r.db.ExecContext(ctx, query, task.Title, task.Completed, task.ID)
    if err != nil {
        return err
    }

    affected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if affected == 0 {
        return errors.New("task not found")
    }

    return nil
}

func (r *PostgresTaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
    query := `DELETE FROM tasks WHERE id = $1`
    result, err := r.db.ExecContext(ctx, query, id)
    if err != nil {
        return err
    }

    affected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if affected == 0 {
        return errors.New("task not found")
    }

    return nil
}

// Nuevo método para filtrar tareas completadas
func (r *PostgresTaskRepository) GetTasksByCompletion(ctx context.Context, completed bool) ([]*entity.Task, error) {
    query := `SELECT id, title, completed FROM tasks WHERE completed = $1`
    rows, err := r.db.QueryContext(ctx, query, completed)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []*entity.Task
    for rows.Next() {
        task := &entity.Task{}
        var idStr string

        err := rows.Scan(&idStr, &task.Title, &task.Completed)
        if err != nil {
            return nil, err
        }

        task.ID, err = uuid.Parse(idStr)
        if err != nil {
            return nil, err
        }

        tasks = append(tasks, task)
    }

    return tasks, nil
}

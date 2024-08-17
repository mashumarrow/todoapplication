package graph

import (
    "context"
     models"github.com/mashumarrow/todoapplication/backend/models"
    "gorm.io/gorm"
)

type Resolver struct {
    DB *gorm.DB
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
    todo := &models.Todo{
        Title:     input.Title,
        Completed: false,
        SubjectID: input.SubjectID,
    }
    
    if err := r.DB.Create(todo).Error; err != nil {
        return nil, err
    }
    
    return todo, nil
}


func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
    var todos []*models.Todo
    
    if err := r.DB.Preload("Subject").Find(&todos).Error; err != nil {
        return nil, err
    }
    
    return todos, nil
}
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }



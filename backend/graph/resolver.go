package graph

import (
    // "context"
    //  models"github.com/mashumarrow/todoapplication/backend/models"
    "gorm.io/gorm"
)

type Resolver struct {
    DB *gorm.DB
}
// func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string) (*models.User, error) {
//     user := &models.User{
//         Name:  name,
//         Email: email,
//     }

//     if err := r.DB.Create(user).Error; err != nil {
//         return nil, err
//     }

//     return user, nil
// }

// func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
//     todo := &models.Todo{
//         Title:     input.Title,
//         Completed: false,
//         SubjectID: input.SubjectID,
//     }
    
//     if err := r.DB.Create(todo).Error; err != nil {
//         return nil, err
//     }
    
//     return todo, nil
// }

// func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
//     var users []*models.User
//     if err := r.DB.Find(&users).Error; err != nil {
//         return nil, err
//     }
//     return users, nil
// }

// func (r *queryResolver) User(ctx context.Context, id uint) (*models.User, error) {
//     var user models.User
//     if err := r.DB.First(&user, id).Error; err != nil {
//         return nil, err
//     }
//     return &user, nil
// }

// func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
//     var todos []*models.Todo
    
//     if err := r.DB.Preload("Subject").Find(&todos).Error; err != nil {
//         return nil, err
//     }
    
//     return todos, nil
// }
// // type mutationResolver struct{ *Resolver }
// type queryResolver struct{ *Resolver }



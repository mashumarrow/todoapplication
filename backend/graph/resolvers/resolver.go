package resolvers

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"context"
	"fmt"

	"github.com/mashumarrow/todoapplication/backend/graph"

	// 	"github.com/mashumarrow/todoapplication/backend/graph/model"
	// 	"github.com/mashumarrow/todoapplication/backend/models"
	"github.com/mashumarrow/todoapplication/backend/models"
	"gorm.io/gorm"
)
type Resolver struct{
	DB *gorm.DB
}

// func (r *Resolver) Title(ctx context.Context, obj *models.NewTodo, data *string) error {
//     if obj == nil || data == nil {
//         return nil // 必要に応じてエラーハンドリングを追加
//     }
//     obj.Title = append(obj.Title, *data)
//     return nil
// }

type NewTodoResolverImpl struct{}

func (r *NewTodoResolverImpl) Title(ctx context.Context, obj *models.NewTodo, data *string) error {
    if data == nil {
        return fmt.Errorf("data is nil")
    }
    obj.Title = *data
    return nil
}
func (r *Resolver) NewTodo() graph.NewTodoResolver {
    return &NewTodoResolverImpl{}
}
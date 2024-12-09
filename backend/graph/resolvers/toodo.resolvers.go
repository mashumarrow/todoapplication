package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mashumarrow/todoapplication/backend/graph"
	"github.com/mashumarrow/todoapplication/backend/models"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	fmt.Println(&ctx)
	fmt.Println(ctx)
	fmt.Println("createtodo")
	// 認証されているユーザーのuserIDを取得
	userIDValue := ctx.Value("userID")
	var userID uint64

	if id, ok := userIDValue.(string); ok {
		// stringからuint64に変換
		parsedID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			fmt.Println("Failed to parse user ID:", id) // 変換失敗時のログ
			return nil, fmt.Errorf("invalid user ID format")
		}
		userID = parsedID
		fmt.Println("User ID in handler:", userID) // デバッグログ
	} else {
		fmt.Println("Failed to get user ID from context:", userIDValue) // ここで何が返ってきたか確認
		return nil, fmt.Errorf("user not authenticated")
	}

	//  認証済みのユーザーIDを使用して Todo を作成
	todo := &models.Todo{
		Title:         input.Title,
		UserID:        uint(userID),
		Period:        input.Period,
		Subjectname:   input.Subjectname,
		Classroomname: input.Classroomname,
		TodoID: 	   input.TodoID,	
	}

	// 5. データベースに新しい Todo を保存
	if err := r.DB.Create(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	var todos []*models.Todo

	if err := r.DB.Preload("Subject").Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, userid string) (*models.Todo, error) {
	var todo models.Todo
	if err := r.DB.Where("UserID = ?", userid).First(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// Todosid is the resolver for the todosid field.
func (r *queryResolver) Todosid(ctx context.Context, todoid string) (*models.Todo, error) {
	var todo models.Todo // 単一のTodoを取得するための変数

    // todoid を使用してデータベースクエリを実行
    if err := r.DB.Where("todoid = ?", todoid).First(&todo).Error; err != nil {
        return nil, err // エラーの場合はnilとエラーを返す
    }

    return &todo, nil // 取得したTodoを返す
		}

// Userid is the resolver for the userid field.
func (r *todoResolver) Userid(ctx context.Context, obj *models.Todo) (string, error) {
	return fmt.Sprintf("%d", obj.UserID), nil}

// Todo returns graph.TodoResolver implementation.
func (r *Resolver) Todo() graph.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
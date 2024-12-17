package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/mashumarrow/todoapplication/backend/graph"
	"github.com/mashumarrow/todoapplication/backend/models"
	"gorm.io/gorm"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input *models.NewTodo) (*models.Todo, error) {
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
		TodoID:        input.TodoID,
		Completed:     input.Completed,
	}

	// 5. データベースに新しい Todo を保存
	if err := r.DB.Create(todo).Error; err != nil {
		fmt.Printf("Failed to create Todo in DB: %v\n", err)
		return nil, fmt.Errorf("failed to save todo to the database")
	}

	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, userid string) ([]*models.Todo, error) {
	var todos []*models.Todo
    if err := r.DB.Where("userid = ?", userid).Find(&todos).Error; err != nil {
        return nil, errors.New("Todoデータの取得に失敗しました")
    }
    return todos, nil
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, userid string) (*models.Todo, error) {
	var todo models.Todo
	if err := r.DB.Where("userid = ?", userid).First(&todo).Error; err != nil {
		log.Println("Todoデータの取得エラー:", err)
		return nil, errors.New("Todoが見つかりません")
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
	return fmt.Sprintf("%d", obj.UserID), nil
}

// Todo returns graph.TodoResolver implementation.
func (r *Resolver) Todo() graph.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UpdateTodo(ctx context.Context, title string, completed bool) (*models.Todo, error) {
	fmt.Println("updatetodo")
	fmt.Println("UpdateTodo called with title:", title)

	// `title`が空の場合はエラーを返す
	if title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}

	var todo models.Todo

	// `title`を条件に検索
	if err := r.DB.First(&todo, "title = ?", title).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Todo not found in database")
			return nil, fmt.Errorf("Todo with title '%s' not found", title)
		}
		return nil, fmt.Errorf("failed to retrieve Todo: %v", err)
	}

	// `completed`を更新
	todo.Completed = completed
	if err := r.DB.Save(&todo).Error; err != nil {
		fmt.Printf("Failed to update Todo: %v\n", err)
		return nil, fmt.Errorf("failed to update Todo: %v", err)
	}
	fmt.Println("Todo updated successfully:", todo)
	return &todo, nil
}
func (r *todoResolver) ID(ctx context.Context, obj *models.Todo) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("Todo object is nil")
	}
	return obj.TodoID, nil
}

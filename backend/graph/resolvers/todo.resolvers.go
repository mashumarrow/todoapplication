// package 

// // This file will be automatically regenerated based on the schema, any resolver implementations
// // will be copied through when generating and any unknown code will be moved to the end.
// // Code generated by github.com/99designs/gqlgen version v0.17.49

// import (
// 	"context"
// 	"fmt"
// 	"strconv"

// 	"github.com/mashumarrow/todoapplication/backend/graph"
// 	"github.com/mashumarrow/todoapplication/backend/graph/model"
// 	"github.com/mashumarrow/todoapplication/backend/models"
// )

// // CreateTodo is the resolver for the createTodo field.
// func (r *mutationResolver) CreateTodou(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
// 	// // 1. context から UserID を取得
// 	// userIDFromCtx := ctx.Value("userID")
// 	// // 2. UserID が存在しない場合、エラーを返す
// 	// if userIDFromCtx == nil {
// 	// 	return nil, errors.New("認証されていません。ユーザーIDが見つかりません")
// 	// }

// 	// // 3. UserID が文字列であることを確認し、uint に変換
// 	// userIDStr, ok := userIDFromCtx.(string)
// 	// if !ok {
// 	// 	return nil, errors.New("無効なユーザーIDの形式です")
// 	// }

// 	// userID, err := strconv.ParseUint(userIDStr, 10, 32) // userIDStr を uint に変換
// 	// if err != nil {
// 	// 	return nil, errors.New("ユーザーIDの変換に失敗しました")
// 	// }

// 	// 認証されているユーザーのuserIDを取得
// 	userIDValue := ctx.Value("userID")
// 	var userID uint64

// 	if id, ok := userIDValue.(string); ok {
// 		// stringからuint64に変換
// 		parsedID, err := strconv.ParseUint(id, 10, 64)
// 		if err != nil {
// 			fmt.Println("Failed to parse user ID:", id) // 変換失敗時のログ
// 			return nil, fmt.Errorf("invalid user ID format")
// 		}
// 		userID = parsedID
// 		fmt.Println("User ID in handler:", userID) // デバッグログ
// 	} else {
// 		fmt.Println("Failed to get user ID from context:", userIDValue) // ここで何が返ってきたか確認
// 		return nil, fmt.Errorf("user not authenticated")
// 	}

// 	//  認証済みのユーザーIDを使用して Todo を作成
// 	todo := &models.Todo{
// 		Title:     input.Title,
// 		UserID:    uint(userID),
// 		DayOfWeek: input.DayOfWeek,
// 		Period:    input.Period,
// 		SubjectName:   input.Subjectname,
// 		ClassroomName: input.Classroomname,
// 	}

// 	// 5. データベースに新しい Todo を保存
// 	if err := r.DB.Create(todo).Error; err != nil {
// 		return nil, err
// 	}

// 	return todo, nil
// }

// // Todos is the resolver for the todos field.
// func (r *queryResolver) Todous(ctx context.Context) ([]*models.Todo, error) {
// 	var todos []*models.Todo

// 	if err := r.DB.Preload("Subject").Find(&todos).Error; err != nil {
// 		return nil, err
// 	}

// 	return todos, nil
// }

// // Todo is the resolver for the todo field.
// func (r *queryResolver) Todou(ctx context.Context, userid string) (*models.Todo, error) {
// 	var todo models.Todo
// 	if err := r.DB.Where("user_id = ?", userid).First(&todo).Error; err != nil {
// 		return nil, err
// 	}
// 	return &todo, nil
// }

// // Todoid is the resolver for the Todoid field.
// func (r *todoResolver) Todoidu(ctx context.Context, obj *models.Todo) (string, error) {
// 	panic(fmt.Errorf("not implemented: Todoid - Todoid"))
// }

// // Userid is the resolver for the userid field.
// func (r *todoResolver) Useridu(ctx context.Context, obj *models.Todo) (string, error) {
// 	return fmt.Sprintf("%d", obj.UserID), nil
// }

// // Dayofweek is the resolver for the dayofweek field.
// func (r *todoResolver) Dayofweeku(ctx context.Context, obj *models.Todo) (model.Dayofweek, error) {
// 	return model.Dayofweek(obj.Dayofweek), nil
// }

// // Dayofweek is the resolver for the dayofweek field.
// func (r *newTodoResolver) Dayofweeku(ctx context.Context, obj *models.NewTodo, data model.Dayofweek) error {
// 	return model.Dayofweek(obj.DayOfWeek), nil
// }

// // Todo returns graph.TodoResolver implementation.
// func (r *Resolver) Todo() graph.TodoResolver { return &todoResolver{r} }

// // NewTodo returns graph.NewTodoResolver implementation.
// func (r *Resolver) NewTodo() graph.NewTodoResolver { return &newTodoResolver{r} }

// type todoResolver struct{ *Resolver }
// type newTodoResolver struct{ *Resolver }

// // !!! WARNING !!!
// // The code below was going to be deleted when updating resolvers. It has been copied here so you have
// // one last chance to move it out of harms way if you want. There are two reasons this happens:
// //   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
// //     it when you're done.
// //   - You have helper methods in this file. Move them out to keep these resolver files clean.
// func (r *newTodoResolver) Period(ctx context.Context, obj *models.NewTodo, data int) error {
// 	return obj.Period, nil
// }
// func (r *todoResolver) Period(ctx context.Context, obj *models.Todo) (int, error) {
// 	return obj.Period, nil
// }
// func (r *newTodoResolver) Userid(ctx context.Context, obj *models.NewTodo, data string) error {
// 	userID := fmt.Sprintf("%d", obj.UserID)
// 	return &userID, nil
// }
// // func (r *todoResolver) SubjSectid(ctx context.Context, obj *models.Todo) (*string, error) {
// // 	subjectID := fmt.Sprintf("%d", obj.SubjectID)
// // 	return &subjectID, nil
// // }
// // func (r *todoResolver) Subject(ctx context.Context, obj *models.Todo) (*models.Subject, error) {
// // 	var subject models.Subject

// // 	if err := r.DB.First(&subject, "SubjectID = ?", obj.SubjectID).Error; err != nil {
// // 		return nil, fmt.Errorf("failed to load subject: %w", err)
// // 	}

// // 	return &subject, nil
// // }

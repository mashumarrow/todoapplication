package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/mashumarrow/todoapplication/backend/graph"
	"github.com/mashumarrow/todoapplication/backend/graph/model"
	"github.com/mashumarrow/todoapplication/backend/models"
)

// Createschedule is the resolver for the createschedule field.
func (r *mutationResolver) Createschedule(ctx context.Context, input model.NewSchedule) (*models.Schedule, error) {
	fmt.Println(&ctx)
	fmt.Println(ctx)
	fmt.Println("createschedule")
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
	fmt.Println("createschedule:認証されているユーザーのuserIDを取得")

	// 同じスケジュールが存在しないか確認
	var existingSchedule models.Schedule
	if err := r.DB.Where("user_id = ? AND dayofweek = ? AND period = ?", userID, input.Dayofweek, input.Period).First(&existingSchedule).Error; err == nil {
		return nil, fmt.Errorf("schedule already exists for this user at this time")
	}

	// スケジュールの作成
	schedule := &models.Schedule{
		UserID:        uint(userID),
		DayOfWeek:     string(input.Dayofweek),
		Period:        input.Period,
		SubjectName:   input.Subjectname,
		ClassroomName: input.Classroomname,
	}

	// スケジュールをデータベースに保存
	if err := r.DB.Create(schedule).Error; err != nil {
		return nil, err
	}
	fmt.Println("createschedule:データベースに保存")
	return schedule, nil
}

// Schedules is the resolver for the schedules field.
func (r *queryResolver) Schedules(ctx context.Context) ([]*models.Schedule, error) {
	var schedules []*models.Schedule

	// データベースからスケジュールを取得
	if err := r.DB.Find(&schedules).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve schedules: %w", err)
	}

	log.Println("取得したスケジュール:", schedules) // デバッグ用のログ

	return schedules, nil
}

// Schedule is the resolver for the schedule field.
func (r *queryResolver) Schedule(ctx context.Context, userid string) (*models.Schedule, error) {
	var schedule models.Schedule
	if err := r.DB.Where("user_id = ?", userid).First(&schedule).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

// Scheduleid is the resolver for the scheduleid field.
func (r *scheduleResolver) Scheduleid(ctx context.Context, obj *models.Schedule) (string, error) {
	if obj.ScheduleID == 0 {
		return "", fmt.Errorf("schedule ID is not available")
	}

	return fmt.Sprintf("%d", obj.ScheduleID), nil
}

// Userid is the resolver for the userid field.
func (r *scheduleResolver) Userid(ctx context.Context, obj *models.Schedule) (*string, error) {
	userID := fmt.Sprintf("%d", obj.UserID)
	return &userID, nil
}

// Dayofweek is the resolver for the dayofweek field.
func (r *scheduleResolver) Dayofweek(ctx context.Context, obj *models.Schedule) (model.Dayofweek, error) {
	return model.Dayofweek(obj.DayOfWeek), nil
}

// Schedule returns graph.ScheduleResolver implementation.
func (r *Resolver) Schedule() graph.ScheduleResolver { return &scheduleResolver{r} }

type scheduleResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *scheduleResolver) Classname(ctx context.Context, obj *models.Schedule) (*string, error) {
	return &obj.ClassroomName, nil
}
func (r *scheduleResolver) Classroom(ctx context.Context, obj *models.Schedule) (*models.Classroom, error) {
	var classroom models.Classroom

	if err := r.DB.Where("classroom_name = ?", obj.ClassroomName).First(&classroom).Error; err != nil {
		return nil, err
	}

	return &classroom, nil
}
func (r *scheduleResolver) Subject(ctx context.Context, obj *models.Schedule) (*models.Subject, error) {
	var subject models.Subject

	if err := r.DB.Where("subject_name = ?", obj.SubjectName).First(&subject).Error; err != nil {
		return nil, err
	}
	return &subject, nil
}
func (r *scheduleResolver) Todo(ctx context.Context, obj *models.Schedule) (*models.Todo, error) {
	var todo models.Todo
	if err := r.DB.First(&todo, obj.UserID).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

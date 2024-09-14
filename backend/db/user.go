package db

import (
	"context"
	"github.com/mashumarrow/todoapplication/backend/models"

	"gorm.io/gorm"
)

// データベースインスタンス
var DB *gorm.DB

// GetUserByName は、データベースから名前に基づいてユーザーを取得
func GetUserByName(ctx context.Context, name string) (*models.User, error) {
	var user models.User

	// データベースからユーザーを取得（
	if err := DB.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
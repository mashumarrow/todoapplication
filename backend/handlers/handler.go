// DBにCRUDする関数を書く
package handler

import (
	"net/http"

	"github.com/mashumarrow/todoapplication/backend/graph"
	"github.com/mashumarrow/todoapplication/backend/graph/resolvers"

	"github.com/mashumarrow/todoapplication/backend/models"

	"encoding/json"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/gorm"
)

// NewGraphQLHandler はGraphQLサーバーのハンドラーを作成
func NewGraphQLHandler(db *gorm.DB) http.Handler {
    
    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{DB: db}}))
    return srv
}

// NewPlaygroundHandler はGraphQL Playgroundのハンドラーを作成
func NewPlaygroundHandler() http.Handler {
    return playground.Handler("GraphQL playground", "/query")
}
 

//新しいユーザー作成
func CreateUserHandler(db *gorm.DB) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request){
        var user models.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
            http.Error(w, err.Error(), http.StatusBadRequest)
			return
        }

		if err := db.Create(&user).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user) //ユーザー情報をjson形式でレスポンスとして返す
    }
}

//指定されたUseridに対応するユーザーを取得するハンドラー
func GetUserHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request){
        //パスパラメータからuserIdを取得
        userID := r.URL.Query().Get("id")

        var user models.User
		if err := db.Where("user_id = ?", userID).First(&user).Error; err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
        }
        json.NewEncoder(w).Encode(user)
    }
}

//新しいTodoを作成するハンドラー
func CreateTodoHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request){
        var newTodo models.NewTodo
        if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil{
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        todo := models.Todo{
			Title:     newTodo.Title,
			SubjectID: newTodo.SubjectID,
			UserID:    newTodo.UserID,
        }
        if err := db.Create(&todo).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
    }
        w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)
}
}

// すべてのTodoを取得するハンドラー
func GetTodosHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todos []models.Todo
		if err := db.Preload("Subject").Preload("User").Find(&todos).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(todos)
	}
}
// GetSubjectsHandler は、すべての科目を取得するためのハンドラー
func GetSubjectsHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var subjects []models.Subject
        if err := db.Find(&subjects).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(subjects)
    }
}

// CreateSubjectHandler は、新しい科目を作成するためのハンドラー
func CreateSubjectHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var subject models.Subject
        if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if err := db.Create(&subject).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(subject)
    }
}

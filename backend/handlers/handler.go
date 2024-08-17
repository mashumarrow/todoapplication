package handler

import (
    "net/http"
    "github.com/mashumarrow/todoapplication/backend/graph"
    //"github.com/mashumarrow/todoapplication/backend/graph/generated"
    "github.com/mashumarrow/todoapplication/backend/models"
	
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "gorm.io/gorm"
	"encoding/json"
)

// NewGraphQLHandler はGraphQLサーバーのハンドラーを作成
func NewGraphQLHandler(db *gorm.DB) http.Handler {
    
    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))
    return srv
}

// NewPlaygroundHandler はGraphQL Playgroundのハンドラーを作成
func NewPlaygroundHandler() http.Handler {
    return playground.Handler("GraphQL playground", "/query")
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
        json.NewEncoder(w).Encode(subject)
    }
}

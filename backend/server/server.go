package server

import (
    "log"
    "net/http"
    "github.com/mashumarrow/todoapplication/backend/handlers"
    "gorm.io/gorm"
)

type Server struct {
    DB   *gorm.DB
    Port string
}

func NewServer(db *gorm.DB, port string) *Server {
    return &Server{
        DB:   db,
        Port: port,
    }
}

func (s *Server) Start() {
    mux := http.NewServeMux()
    mux.Handle("/playground", handler.NewPlaygroundHandler())
    mux.Handle("/query", handler.NewGraphQLHandler(s.DB))
    mux.HandleFunc("/users", handler.CreateUserHandler(s.DB))
    mux.HandleFunc("/users/get", handler.GetUserHandler(s.DB))
    mux.HandleFunc("/todos", handler.CreateTodoHandler(s.DB))
    mux.HandleFunc("/todos/get", handler.GetTodosHandler(s.DB))
    mux.HandleFunc("/subjects", handler.CreateSubjectHandler(s.DB))
    mux.HandleFunc("/subjects/get", handler.GetSubjectsHandler(s.DB))

    log.Printf("サーバーが起動しました: http://localhost:%s/playground", s.Port)
    log.Fatal(http.ListenAndServe(":"+s.Port, mux))
}

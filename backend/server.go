package main

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

    log.Printf("サーバーが起動しました: http://localhost:%s/playground", s.Port)
    log.Fatal(http.ListenAndServe(":"+s.Port, mux))
}


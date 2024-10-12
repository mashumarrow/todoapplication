package server

import (
    "log"
    "net/http"
    
    "github.com/rs/cors"
    "github.com/gorilla/mux"
    "github.com/mashumarrow/todoapplication/backend/handlers"
    "gorm.io/gorm"
)

type Server struct {
    DB   *gorm.DB
    Router *mux.Router
    Port string
}

func NewServer(db *gorm.DB, port string) *Server {
    return &Server{
        DB:   db,
        Router: mux.NewRouter(),
        Port: port,
    }
}

func (s *Server) Start() {
    mux := http.NewServeMux()

     // CORS設定
     c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"}, 
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Authorization", "Content-Type"},
    })
   
    mux.Handle("/graphql", handler.NewGraphQLHandler(s.DB))
   

    //mux.Handle("/playground", handler.NewPlaygroundHandler())
    // mux.HandleFunc("/users", handler.CreateUserHandler(s.DB))
    // mux.HandleFunc("/users/get", handler.GetUserHandler(s.DB))
    // mux.HandleFunc("/todos", handler.CreateTodoHandler(s.DB))
    // mux.HandleFunc("/todos/get", handler.GetTodosHandler(s.DB))
    // mux.HandleFunc("/subjects", handler.CreateSubjectHandler(s.DB))
    // mux.HandleFunc("/subjects/get", handler.GetSubjectsHandler(s.DB))

    log.Printf("サーバーが起動しました: http://localhost:%s/graphql", s.Port)
    log.Fatal(http.ListenAndServe(":"+s.Port, c.Handler(mux))) 
}

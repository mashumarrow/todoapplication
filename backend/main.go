package main

import (
    "log"
    "net/http"
    "github.com/mashumarrow/todoapplication/backend/models"
    "github.com/mashumarrow/todoapplication/backend/handlers"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)




const defaultPort = "8080"

func main() {
    dsn := "your_username:your_password@tcp(127.0.0.1:3306)/your_database_name?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }

    db.AutoMigrate(&models.Subject{}, &models.Classroom{}, &models.Schedule{}, &models.Todo{})

	
    http.Handle("/playground", handler.NewGraphQLHandler(db))
    http.Handle("/query", handler.NewGraphQLHandler(db))

    log.Printf("connect to http://localhost:%s/playground for GraphQL playground", defaultPort)
    log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
package main

import (
    
    "github.com/mashumarrow/todoapplication/backend/db/migrations"
    "github.com/mashumarrow/todoapplication/backend/server"
    //"gorm.io/driver/mysql"
   // "gorm.io/gorm"
)





func main() {
    // データベースを初期化
    database.InitDB()

    // サーバーを作成
    srv := server.NewServer(database.DB, "8080")

    // サーバーを開始
    srv.Start()
}
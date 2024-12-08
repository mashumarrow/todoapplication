package main

import (
    "net/http"
    "strings"
    "context"
    jwt "github.com/golang-jwt/jwt/v4"
    "github.com/mashumarrow/todoapplication/backend/handlers" 
    "github.com/mashumarrow/todoapplication/backend/db/migrations"
    "github.com/mashumarrow/todoapplication/backend/server"
    "fmt"
    //"github.com/gorilla/mux"
)





func main() {

    // データベースを初期化
    database.InitDB()

    
    // サーバーを作成
    srv := server.NewServer(database.DB, "8080")

     // ミドルウェア付きのルーターを作成
     srv.Router.Use(middleware)

    srv.Router.Handle("/graphql", handler.NewGraphQLHandler(srv.DB)) 

    // サーバーを開始
    srv.Start()

    
    
}
// ミドルウェア関数
func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("middleware")
        tokenString := r.Header.Get("Authorization")
        if tokenString != "" {
            tokenString = strings.TrimPrefix(tokenString, "Bearer ")
            token, err := verifyToken(tokenString)
            if err == nil {
                if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
                    UserID := claims["userid"].(string)
                    ctx := context.WithValue(r.Context(), "userID", UserID)
                    next.ServeHTTP(w, r.WithContext(ctx))  // ここで r にコンテキストをセットして渡す
                    return 
                }
            }
        }
        next.ServeHTTP(w, r)
    })
}

// トークン検証関数
func verifyToken(tokenString string) (*jwt.Token, error) {
    secretKey := []byte("your_secret_key") // 秘密鍵を設定
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        
        return secretKey, nil
    })
}
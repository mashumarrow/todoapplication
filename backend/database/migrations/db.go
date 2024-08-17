// 環境変数から取得した情報を使ってデータベースに接続する
package database

import (
	"fmt" //フォーマットされたi/o操作の標準ライブラリ
	"log" //ログ出力の標準ライブラリ
	"os"  //環境変数やファイル操作を行う標準ライブラリ
	//"github.com/mashumarrow/todoes/models"

	"github.com/joho/godotenv" //.envファイルから環境変数を読み込むためのライブラリ
	"gorm.io/driver/mysql"     //gormのmysqlドライバを提供するパッケージ
	"gorm.io/gorm"             //ormライブラリ．データベース操作をオブジェクト指向的に扱える
)

var DB *gorm.DB//gormのデータベース接続インスタンスを保持するグローバル変数．

func InitDB() {
    err := godotenv.Load()      //.envファイルから環境変数を読み込む．os.Getenvでこのデータ取得できる
    if err != nil {
        log.Fatalf(".envファイルの読み込み失敗")
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    var errDB error
    DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if errDB != nil {
        log.Fatalf("データベース接続失敗: %v", errDB)
    }

    //DB.AutoMigrate(&models.User{}, &models.Schedule{}, &models.Todo{}) //自動マイグレーション
}

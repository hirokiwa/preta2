package infrastructure

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)


func Init() *gorm.DB {
	err := godotenv.Load(".env")
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	return db
}

func GetDB() *gorm.DB {
return db
}
 

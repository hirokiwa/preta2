package client

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "github.com/go-sql-driver/mysql"
  )

func GetDatabase() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	return db, err
  }
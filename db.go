package crawler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var DB *gorm.DB

// 模块表结构
type Module struct {
	Id    uint
	Name  string
	Alias string
}

func init() {
	ConnectDB()
}

func ConnectDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "crawler.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 初始化数据库
func InitDB() {
	if DB.HasTable(&Module{}) {
		return
	}
	DB.CreateTable(&Module{}) // 创建
}

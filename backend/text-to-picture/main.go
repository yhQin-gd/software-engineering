package main

import (
	"log"

	"text-to-picture/api/auth"       // 导入注册和登录路由
	db "text-to-picture/models/init" // 为 init 包设置别名为 db

	"github.com/gin-gonic/gin"
)

func main() {
	// 连接数据库
	if err := db.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 初始化数据库
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 设置路由
	r := gin.Default()
	r.POST("/register", auth.Register) // 注册路由
	r.POST("/login", auth.Login)       // 登录路由

	r.Run() // 启动服务器
}

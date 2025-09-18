package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hufeng-code/db"
	"hufeng-code/models"
	"hufeng-code/routes"
	"net/http"
)

// TODO:提取内部配置至config

func main() {
	// 初始化gorm数据库连接
	db.Connect2DB()
	_ = db.DB.AutoMigrate(&models.Department{}, &models.Employee{}, &models.Records{})
	// 初始化gin后端
	router := gin.Default()
	router.Use() // 使用权限控制、限流、阻止攻击等中间件
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "欢迎访问！"})
	})

	routes.BindEmployeeSetting(router)
	routes.BindEmployeeJob(router)
	routes.BindEmployeeCompany(router)

	// 进行数据库增删改查

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error on gin maingoroutinue : %v", err)
		}
	}()

	fmt.Println("服务启动中！")
	if err := router.Run("localhost:9090"); err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oldshop/controller"
	"oldshop/initliazier"
	"oldshop/utils"
)

// github.com/mattn/go-sqlite3
func main() {

	// 初始化配置文件
	initliazier.InitConfig()
	// 初始化数据库
	initliazier.InitDB()

	//初始化路由
	r := gin.Default()
	r.Use(Cors())
	r.GET("/api/banner", controller.HomePicList)
	r.POST("/api/user/login", controller.Login)
	r.GET("/api/user/info", utils.Auth(), controller.GetUserInfo)
	r.POST("/api/goods/add", utils.Auth(), controller.CreateGood)
	r.POST("/api/goods/list", controller.GetGoodList)
	r.GET("/api/goods/detail", controller.GetGoodDetail)
	r.POST("/api/user/register", controller.Register)
	r.GET("/api/category/list", controller.FindCategoryList)
	r.GET("/api/category/good", controller.GetGoodListByCategoryId)
	r.GET("/api/goods/search", controller.GetGoodListByKey)

	//启动服务
	r.Run(":8854")
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST,GET, OPTIONS, DELETE,PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

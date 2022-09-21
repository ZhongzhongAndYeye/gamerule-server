package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler"
)

func Start(){
	r := gin.Default()  // 返回默认的路由引擎

	m := r.Group("/manage") 			// 后台管理系统的接口
	m.POST("/game",handler.GetLog)  	// 新建或修改游戏以及日志（方式是存在则覆盖，不存在则创建）

	r.Run(":8080") //监听8080端口
}
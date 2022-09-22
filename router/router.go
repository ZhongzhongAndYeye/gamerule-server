package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler"
)

func Start(){
	r := gin.Default()  // 返回默认的路由引擎

	m := r.Group("/manage") 			// 后台管理系统的接口
	m.POST("/game",handler.GetLog)  	// 新建或修改游戏玩法以及日志（方式是存在则覆盖，不存在则创建）
	// m.DELETE("/game/content",handler.DelLogContent) // 删除指定的日志详情
	// m.DELETE("/game/log")			// 删除游戏玩法（包括该玩法下的所有日志详情）

	r.Run(":8080") //监听8080端口
}
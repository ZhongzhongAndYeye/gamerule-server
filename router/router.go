package router

import (
	"server/internal/handler"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default() // 返回默认的路由引擎

	m := r.Group("/manage")                          // 后台管理系统的接口
	m.GET("/game", handler.GetLog)                   // 一进管理页面获取游戏玩法信息
	m.GET("/game/content", handler.GetLogAndContent) // 编辑具体游戏玩法时显示的初始信息以及历史更新记录
	m.POST("/game", handler.CreateLog)               // 新建或修改游戏玩法以及日志（方式是存在则覆盖，不存在则创建）
	m.DELETE("/game", handler.DelLog)                // 删除游戏玩法（包括该玩法下的所有日志详情）
	m.DELETE("/game/content", handler.DelLogContent) // 删除指定的日志详情

	r.Run(":8080") //监听8080端口
}

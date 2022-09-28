package router

import (
	"server/internal/handler"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default() // 返回默认的路由引擎

	m := r.Group("/manage")                     // 后台管理系统的接口
	m.POST("/log", handler.CreateLog)           // 新建游戏玩法以及日志
	m.PUT("/log", handler.EditLog)              // 编辑游戏玩法以及日志
	m.GET("/log", handler.GetLog)               // 获取log基本数据列表（包括页面初始化和模糊查询的参数）
	m.GET("/content", handler.GetContent)       // 点击编辑时获取历史更新记录详情
	m.DELETE("/log", handler.DeleteLog)         // 删除游戏玩法
	m.DELETE("/content", handler.DeleteContent) // 删除游戏玩法的指定历史记录（单条）

	r.Run(":8080") //监听8080端口
}

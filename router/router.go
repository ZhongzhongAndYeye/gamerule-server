package router

import (
	"server/internal/handler"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default() // 返回默认的路由引擎

	m := r.Group("/manage")                     // 后台管理系统的接口
	m.POST("/log", handler.CreateLog)           // 新建游戏玩法以及历史更新记录详情
	m.PUT("/log", handler.EditLog)              // 编辑游戏玩法以及历史更新记录详情
	m.GET("/log", handler.GetLog)               // 获取游戏玩法基本数据列表（页面初始化和模糊查询）
	m.GET("/content", handler.GetContent)       // 点击编辑时获取游戏玩法历史更新记录详情
	m.DELETE("/log", handler.DeleteLog)         // 删除游戏玩法以及其下的历史更新记录详情
	m.DELETE("/content", handler.DeleteContent) // 删除游戏玩法的指定历史更新记录详情（单条）

	c := r.Group("/client")                           // 客户端接口
	c.GET("/log", handler.GetAllTabCondition) // 给客户端一个所有游戏的更新情况表，让客户端比对旧表
	c.GET("/content",handler.GetTabContent)	// 给客户端指定游戏玩法的在当前时间之前的历史更新记录

	r.Run(":8080") // 监听8080端口
}

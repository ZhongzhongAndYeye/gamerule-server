package router

import (
	"server/internal/handler"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default() // 返回默认的路由引擎

	m := r.Group("/manage")            // 后台管理系统的接口
	m.POST("/game", handler.CreateLog) // 新建游戏玩法以及日志
	m.PUT("/game", handler.EditLog)    // 编辑游戏玩法以及日志

	r.Run(":8080") //监听8080端口
}

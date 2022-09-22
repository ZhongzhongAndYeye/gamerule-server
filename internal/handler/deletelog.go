package handler

import (
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func DelLog(c *gin.Context) {
	logInfo := model.LogInfo{}
	c.ShouldBind(&logInfo)
	model.DelLog(logInfo)
}
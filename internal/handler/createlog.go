package handler

import (
	"fmt"
	"net/http"
	"server/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取发来的form数据
func CreateLog(c *gin.Context) {
	log := new(model.Log) // 创建一个log对象，接收
	c.ShouldBind(&log)
	logcontents := make([]model.LogContent, log.Count) // 根据count来决定创建多大的content切片
	for i := range logcontents {                       // 遍历给content切片赋值
		content_index := fmt.Sprintf("%v%v", "content_", i+1)
		updatetime_index := fmt.Sprintf("%v%v", "update_time_", i+1)
		logcontents[i].Content = c.PostForm(content_index)
		logcontents[i].UpdateTime, _ = strconv.ParseInt(c.PostForm(updatetime_index), 10, 64)
	}

	model.CreateLog(log, logcontents) // 将获取的数据传入model和数据库进行交互
	c.JSON(http.StatusOK,gin.H{
		"status":"success",
	})
}

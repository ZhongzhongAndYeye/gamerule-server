package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func GetAllTabCondition(c *gin.Context) {
	var rsp struct {
		Msg  string                   `json:"msg"`  // 请求信息
		Data []map[string]interface{} `json:"data"` // 请求所需数据 字符串切片
	}

	rsp.Data, rsp.Msg = model.GetAllTabCondition()
	c.JSON(http.StatusOK,rsp)
}

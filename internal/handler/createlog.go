package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Status int
	Msg    string
}

// 获取发来的form数据
func CreateLog(c *gin.Context) {
	log := &model.Log{}
	req := &Request{}
	c.ShouldBindJSON(&log)
	result,msg := model.CreateLog(log) // 将获取的数据传入model和数据库进行交互
	req.Msg = msg
	if result == false {
		req.Status = http.StatusBadRequest
	}else{
		req.Status = http.StatusOK
	}
	c.JSON(req.Status, req)
}

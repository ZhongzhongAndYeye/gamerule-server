package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

// 获取发来的form数据
func CreateLog(c *gin.Context) {
	log := &model.Log{}
	req := &Request{}
	c.ShouldBindJSON(log)
	var msg string
	var err error
	if err, msg = model.CreateLog(log);err !=nil{// 将获取的数据传入model和数据库进行交互
		req.Status = http.StatusBadRequest
	}else{
		req.Status = http.StatusOK
	} 
	req.Msg = msg
	c.JSON(req.Status, req)
}

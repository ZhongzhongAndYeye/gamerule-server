package handler

import (
	_ "fmt"
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func EditLog(c *gin.Context) {
	log := &model.Log{}
	rsp := &MRsp{}
	c.ShouldBindJSON(log)
	var msg string
	var err error
	if msg,err = model.EditLog(log);err != nil{
		rsp.Status = http.StatusBadRequest
	}else{
		rsp.Status = http.StatusOK
	}
	rsp.Msg = msg
	c.JSON(rsp.Status, rsp)
}

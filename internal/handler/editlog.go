package handler

import (
	_ "fmt"
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func EditLog(c *gin.Context) {
	log := &model.Log{}
	req := &Request{}
	c.ShouldBindJSON(log)
	var msg string
	var err error
	if msg,err = model.EditLog(log);err != nil{
		req.Status = http.StatusBadRequest
	}else{
		req.Status = http.StatusOK
	}
	req.Msg = msg
	c.JSON(req.Status, req)
}

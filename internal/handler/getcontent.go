package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func GetContent(c *gin.Context) {
	idInfo := &IdInfo{}
	c.ShouldBindJSON(idInfo)
	var logcontents []model.LogContent
	var msg string
	var err error
	var status int
	if logcontents, msg, err = model.GetContent(idInfo.Id); err != nil {
		status = http.StatusBadRequest
	} else {
		status = http.StatusOK  
	}
	c.JSON(status, gin.H{
		"msg":         msg,
		"logcontents": logcontents,
	})
}

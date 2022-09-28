package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func DeleteLog(c *gin.Context){
	idInfo := &IdInfo{}
	req := &Request{}
	c.ShouldBindJSON(idInfo)
	err,msg := model.DeleteLog(idInfo.Id)
	req.Msg = msg
	if err != nil {
		req.Status = http.StatusBadRequest
	}else{
		req.Status = http.StatusOK
	}
	c.JSON(req.Status,req)
}
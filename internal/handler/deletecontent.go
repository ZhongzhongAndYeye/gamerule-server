package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func DeleteContent(c *gin.Context) {
	idInfo := &IdInfo{}
	rsp := &MRsp{}
	c.ShouldBindJSON(idInfo)
	err, msg := model.DeleteContent(idInfo.Id)
	rsp.Msg = msg
	if err != nil {
		rsp.Status = http.StatusBadRequest
	} else {
		rsp.Status = http.StatusOK
	}
	c.JSON(rsp.Status, rsp)
}

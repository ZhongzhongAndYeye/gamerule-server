package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func DeleteContent(c *gin.Context) {
	idInfo := &IdInfo{}
	req := &Request{}
	c.ShouldBindJSON(idInfo)
	err, msg := model.DeleteContent(idInfo.Id)
	req.Msg = msg
	if err != nil {
		req.Status = http.StatusBadRequest
	} else {
		req.Status = http.StatusOK
	}
	c.JSON(req.Status, req)
}

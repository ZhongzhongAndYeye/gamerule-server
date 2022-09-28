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
	result, msg := model.EditLog(log)
	req.Msg = msg
	if result == false {
		req.Status = http.StatusBadRequest
	} else {
		req.Status = http.StatusOK
	}
	c.JSON(req.Status, req)
}

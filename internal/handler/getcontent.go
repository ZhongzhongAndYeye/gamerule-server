package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func GetContent(c *gin.Context) {
	idInfo := &IdInfo{}
	c.ShouldBindJSON(idInfo)
	logcontents := model.GetContent(idInfo.Id)
	c.JSON(http.StatusOK, logcontents)
}

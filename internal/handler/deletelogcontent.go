package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)


func DelLogContent(c *gin.Context){ 
	contentInfo := new(model.ContentInfo)
	c.ShouldBind(&contentInfo)
	model.DelLogContent(*contentInfo)
	c.JSON(http.StatusOK,gin.H{
		"status":"success",
	})
}
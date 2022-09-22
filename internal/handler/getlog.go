package handler

import (
	_ "fmt"
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func GetLog(c *gin.Context){ 
	needlogs := model.GetLog()
	// fmt.Println(needlogs)
	c.JSON(http.StatusOK,needlogs)
}
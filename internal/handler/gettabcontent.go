package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func GetTabContent(c *gin.Context) {
	var req struct {
		GameId int    `json:"game_id"`
		Tab    string `json:"tab"`
	}
	var rsp struct{
		Msg string 		`json:"msg"`
		Data []map[string]interface{}	`json:"data"`	
	}

	c.ShouldBindJSON(&req)
	rsp.Data,rsp.Msg = model.GetTabContent(req.GameId,req.Tab)
	c.JSON(http.StatusOK,rsp)
}

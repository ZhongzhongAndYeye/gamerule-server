package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func GetLogAndContent(c *gin.Context) {
	gameid := c.PostForm("game_id")
	tab := c.PostForm("tab")
	logAndContent := model.GetLogAndContent(gameid,tab)
	c.JSON(http.StatusOK,logAndContent)
}
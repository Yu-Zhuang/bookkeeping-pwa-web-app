package controller

import (
	"bookkeeping/config"
	"bookkeeping/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetChartData(c *gin.Context) {
	userID := c.MustGet(config.AuthMidUserNameKey).(string)
	pieRet, ok := logic.GetPieData(userID)
	if ok == false {
		c.JSON(http.StatusBadRequest, nil)
	}
	lineRet, ok := logic.GetLineChartData(userID)
	if ok == false {
		c.JSON(http.StatusBadRequest, nil)
	}
	c.JSON(http.StatusOK, gin.H{
		"pie":  pieRet,
		"line": lineRet,
	})
}

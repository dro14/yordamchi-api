package handler

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/dro14/yordamchi-api/utils/info"
	"github.com/gin-gonic/gin"
)

func failure(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func sendSSEEvent(c *gin.Context, name string, data any) {
	c.SSEvent(name, data)
	c.Writer.Flush()
}

func notifyOnPanic(c *gin.Context, err any) {
	log.Printf("%s\n%s", err, debug.Stack())
	info.SendDocument("my.log")
	info.SendDocument("gin.log")
	c.AbortWithStatus(http.StatusInternalServerError)
}

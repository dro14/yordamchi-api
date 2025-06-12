package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/dro14/yordamchi-api/utils/info"
	"github.com/gin-gonic/gin"
)

func failure(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func sendSSEEvent(ctx *gin.Context, name string, data any) {
	ctx.SSEvent(name, data)
	ctx.Writer.Flush()
}

func notifyOnPanic(ctx *gin.Context, err any) {
	log.Printf("%s\n%s", err, debug.Stack())
	info.SendDocument("my.log")
	info.SendDocument("gin.log")
	ctx.AbortWithStatus(http.StatusInternalServerError)
}

func jsonEncode(data any) string {
	json, _ := json.Marshal(data)
	return string(json)
}

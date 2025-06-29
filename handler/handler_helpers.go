package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/dro14/yordamchi-api/models"
	"github.com/dro14/yordamchi-api/utils/info"
	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

const (
	latestVersion       = "1.0.3"
	minimumVersion      = "1.0.3"
	maxErrors           = 10
	cachedTokenPrice    = 75
	nonCachedTokenPrice = 300
	responseTokenPrice  = 2500
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

func recordUsage(request *models.Request, usageMetadata *genai.GenerateContentResponseUsageMetadata) {
	request.CachedTokens = int64(usageMetadata.CachedContentTokenCount)
	request.NonCachedTokens = int64(usageMetadata.PromptTokenCount) - request.CachedTokens
	request.ToolPromptTokens = int64(usageMetadata.ToolUsePromptTokenCount)
	request.ThoughtTokens = int64(usageMetadata.ThoughtsTokenCount)
	request.ResponseTokens = int64(usageMetadata.CandidatesTokenCount) - request.ThoughtTokens

	promptPrice := request.CachedTokens*cachedTokenPrice + request.NonCachedTokens*nonCachedTokenPrice
	responsePrice := (request.ThoughtTokens + request.ResponseTokens) * responseTokenPrice
	request.Price = float64(promptPrice+responsePrice) / (1e3 * 1e6)
}

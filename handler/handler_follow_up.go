package handler

import (
	"log"
	"net/http"

	"github.com/dro14/yordamchi-api/models"
	"github.com/dro14/yordamchi-api/utils/e"
	"github.com/dro14/yordamchi-api/utils/f"
	"github.com/gin-gonic/gin"
)

func (h *Handler) followUp(ctx *gin.Context) {
	startedAt := f.Now()
	request := &models.Request{}
	err := ctx.ShouldBindJSON(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, failure(err))
		return
	}
	request.StartedAt = startedAt

	if request.UserId == 0 ||
		request.ChatId == 0 ||
		request.Language == "" ||
		len(request.Contents) == 0 {
		ctx.JSON(http.StatusBadRequest, failure(e.ErrEmpty))
		return
	}

	response, err := h.provider.FollowUp(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, failure(err))
		return
	}

	message := &models.Message{
		UserId:           request.UserId,
		ChatId:           request.ChatId,
		Role:             "model",
		CreatedAt:        f.Now(),
		StructuredOutput: response.Text(),
	}
	err = h.data.CreateMessage(ctx, message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, failure(err))
		return
	}

	promptPrice := request.PromptTokens * promptTokenPrice
	responsePrice := request.ResponseTokens * responseTokenPrice

	request.FinishedAt = f.Now()
	request.Latency = request.FinishedAt - request.StartedAt
	request.Chunks = 1
	request.Attempts = 1
	request.Response = message
	request.FinishReason = string(response.Candidates[0].FinishReason)
	request.PromptTokens = int64(response.UsageMetadata.PromptTokenCount)
	request.ResponseTokens = int64(response.UsageMetadata.CandidatesTokenCount)
	request.Price = float64(promptPrice+responsePrice) / (1e2 * 1e6)
	err = h.data.CreateRequest(ctx, request)
	if err != nil {
		log.Print("can't create request: ", err)
	}

	ctx.JSON(http.StatusOK, message)
}

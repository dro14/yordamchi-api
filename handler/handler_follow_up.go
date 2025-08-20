package handler

import (
	"encoding/json"
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

Retry:
	response, err := h.provider.FollowUp(request)
	if err != nil {
		request.Errors++
		log.Print("can't generate follow-ups: ", err)
		if request.Errors < int64(len(h.provider.Clients)) {
			goto Retry
		}
		log.Printf("follow-up failed after %d attempts: %s", request.Errors, err)
		ctx.JSON(http.StatusInternalServerError, failure(err))
		return
	}

	var followUps []string
	err = json.Unmarshal([]byte(response.Text()), &followUps)
	if err != nil {
		request.Errors++
		log.Printf("can't unmarshal follow-ups: %s\n%s", err, response.Text())
		if request.Errors < int64(len(h.provider.Clients)) {
			goto Retry
		}
		log.Printf("follow-up failed after %d attempts: %s", request.Errors, err)
		ctx.JSON(http.StatusInternalServerError, failure(err))
		return
	}

	message := &models.Message{
		UserId:    request.UserId,
		ChatId:    request.ChatId,
		Role:      "model",
		CreatedAt: f.Now(),
		FollowUps: followUps,
	}
	err = h.data.CreateMessage(ctx, message)
	if err != nil {
		log.Print("can't create message: ", err)
		ctx.JSON(http.StatusInternalServerError, failure(err))
		return
	}

	request.FinishedAt = f.Now()
	request.Latency = request.FinishedAt - request.StartedAt
	request.Response = message
	request.FinishReason = string(response.Candidates[0].FinishReason)
	recordUsage(request, response.UsageMetadata)
	err = h.data.CreateRequest(ctx, request)
	if err != nil {
		log.Print("can't create request: ", err)
	}

	ctx.JSON(http.StatusOK, message)
}

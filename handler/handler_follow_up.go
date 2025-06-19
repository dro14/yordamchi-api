package handler

import (
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

	ctx.String(http.StatusOK, response.Text())
	ctx.Header("Content-Type", "application/json")
}

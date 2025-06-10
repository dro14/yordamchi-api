package handler

import (
	"net/http"

	"github.com/dro14/yordamchi-api/utils/f"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(ctx *gin.Context) {
	id, err := h.data.CreateUser(ctx, f.Now())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, failure(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

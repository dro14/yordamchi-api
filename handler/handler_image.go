package handler

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createImage(ctx *gin.Context) {
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, failure(err))
		return
	}

	randomStr := make([]byte, 8)
	rand.Read(randomStr)
	filename := fmt.Sprintf("%d_%x.jpeg", time.Now().UnixMilli(), randomStr)

	err = os.WriteFile("rasmlar/"+filename, body, 0644)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, failure(err))
		return
	}

	ctx.Header("X-Filename", filename)
}

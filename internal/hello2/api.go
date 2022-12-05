package hello2

import (
	"go-be-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello2(context *gin.Context) {
	context.IndentedJSON(
		http.StatusOK,
		utils.SuccessResponse(
			"hello2",
		),
	)
}

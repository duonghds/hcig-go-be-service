package hello1

import (
	"go-be-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello1(context *gin.Context) {
	context.IndentedJSON(
		http.StatusOK,
		utils.SuccessResponse(
			"hello1",
		),
	)
}

package runtime

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RuntimeController struct {
	service *RuntimeService
}

func NewRuntimeController(s *RuntimeService) *RuntimeController {
	return &RuntimeController{
		service: s}
}

func (rc RuntimeController) GetGameContext(ctx *gin.Context) {
	var data, err = rc.service.GetGameContext()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, data)
}

func (rc RuntimeController) InvokeMethod(ctx *gin.Context) {
	method := ctx.Query("method")

	var body json.RawMessage
	var err error

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	stringContext, _ := json.Marshal(body)

	rc.service.InvokeMethod(method, string(stringContext))

	data, err := rc.service.GetGameContext()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, data)
}

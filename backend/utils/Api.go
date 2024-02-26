package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"net/http"
)

type Api struct {
	Port int
	Db   gorm.DB
	Log  zerolog.Logger
}

type ApiContext struct {
	Api
	*gin.Context
}

type ApiError struct {
	Code    int
	Message string
}

type ApiSuccess struct {
	Code    int
	Message string
	Content interface{}
}

type ApiHandler func(api *ApiContext)

// CreateHandler / Creates an api handler from provided function
func (api *Api) CreateHandler(handler ApiHandler) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Create custom api context with addition information
		var apiContext = ApiContext{
			Context: context,
			Api:     *api,
		}

		handler(&apiContext)
	}
}

func (api *ApiContext) Success(data interface{}) {
	api.JSON(http.StatusOK, ApiSuccess{
		Code:    http.StatusOK,
		Message: "Success",
		Content: data,
	})
}

func (api *ApiContext) ReturnSuccess(message string) {
	api.JSON(http.StatusOK, ApiSuccess{
		Code:    http.StatusOK,
		Message: message,
		Content: nil,
	})
}

func (api *ApiContext) Error(code int, message string) {
	api.JSON(code, ApiError{
		Code:    code,
		Message: message,
	})
}

func (api *ApiContext) NotFound(message string) {
	api.Error(http.StatusNotFound, message)
}

func (api *ApiContext) Conflict(message string) {
	api.Error(http.StatusConflict, message)
}

func (api *ApiContext) InternalServerError(message string) {
	api.Error(http.StatusInternalServerError, message)
}

func BindJSON[T any](ctx *ApiContext) (T, error) {
	var obj T
	err := ctx.BindJSON(&obj)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ApiError{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
		})

		return obj, err
	}

	return obj, nil
}

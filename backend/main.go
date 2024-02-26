package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"todo.app/controller"
	"todo.app/repository"
	"todo.app/utils"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

func main() {
	logger := utils.InitializeLogging()
	db := repository.InitializeDatabase()

	api := utils.Api{
		Db:   db,
		Log:  logger,
		Port: utils.GetPortNumber(),
	}

	router := gin.Default()

	router.Use(CORSMiddleware())
	router.POST("/api/todo/filter", api.CreateHandler(controller.GetTodos))
	router.DELETE("/api/todo/:id", api.CreateHandler(controller.DeleteTodo))
	router.POST("/api/todo", api.CreateHandler(controller.CreateTodo))
	router.PATCH("/api/todo/:id", api.CreateHandler(controller.PatchTodo))

	err := router.Run(":" + strconv.Itoa(api.Port))

	if err != nil {
		api.Log.Fatal().Msgf("Failed to start server: %s", err)
	}
}

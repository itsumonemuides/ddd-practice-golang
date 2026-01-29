package router

import (
	"go-practice/interface/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(todoHandler *handler.TodoHandler) *gin.Engine {
	router := gin.Default()

	router.Use(corsMiddleware())

	api := router.Group("/api")
	{
		todos := api.Group("/todos")
		{
			todos.POST("", todoHandler.CreateTodo)
			todos.GET("", todoHandler.ListTodos)
			todos.PUT("/:id/complete", todoHandler.CompleteTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
		}
	}

	// ヘルスチェック
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

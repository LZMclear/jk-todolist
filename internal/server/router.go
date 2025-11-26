package server

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"jk-todolist/internal/handler"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Static("/static", "web")

	r.GET("/", func(c *gin.Context) {
		c.File("web/index.html")
	})

	api := r.Group("/api")
	{
		t := api.Group("/tasks")
		{
			t.GET("/", handler.ListTasks(db))
			t.POST("/", handler.CreateTask(db))
			t.GET(":id", handler.GetTask(db))
			t.PUT(":id", handler.UpdateTask(db))
			t.DELETE(":id", handler.DeleteTask(db))
		}
	}
	return r
}

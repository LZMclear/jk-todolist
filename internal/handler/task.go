package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"jk-todolist/internal/store"
	"net/http"
	"strconv"
	"time"
)

func ListTasks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks, err := store.ListTasks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tasks)
	}
}

func CreateTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload struct {
			Title       string `json:"title" binding:"required"`
			Description string `json:"description"`
			Category    string `json:"category"`
			DueDate     string `json:"due_date"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 解析前端 datetime-local（无时区），使用本地时区解析，转为 UTC 存库
		var duePtr *time.Time
		if payload.DueDate != "" {
			// 前端发送格式如 "2025-11-25T14:30"
			if t, err := time.ParseInLocation("2006-01-02T15:04", payload.DueDate, time.Local); err == nil {
				utc := t.UTC()
				duePtr = &utc
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid due_date format"})
				return
			}
		}

		t, err := store.CreateTask(db, payload.Title, payload.Description, payload.Category, duePtr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, t)
	}
}

func GetTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		t, err := store.GetTask(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusOK, t)
	}
}

func UpdateTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		var payload struct {
			Title       string `json:"title" binding:"required"`
			Description string `json:"description"`
			Category    string `json:"category"`
			DueDate     string `json:"due_date"`
			Completed   bool   `json:"completed"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var duePtr *time.Time
		if payload.DueDate != "" {
			if t, err := time.ParseInLocation("2006-01-02T15:04", payload.DueDate, time.Local); err == nil {
				utc := t.UTC()
				duePtr = &utc
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid due_date format"})
				return
			}
		}

		t, err := store.UpdateTask(db, id, payload.Title, payload.Description, payload.Category, duePtr, payload.Completed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, t)
	}
}

func DeleteTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		if err := store.DeleteTask(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

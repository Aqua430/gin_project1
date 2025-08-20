package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func statsEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"number_of_tasks": len(tasks),
	})
}

func resetEndpoint(c *gin.Context) {
	tasks = nil
	c.JSON(http.StatusOK, gin.H{
		"message": "all tasks have been cleared",
		"tasks":   tasks,
	})
}

func tasksEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func taskEndpoint(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "title is required",
		})
		return
	}
	newTask := Task{
		ID:    len(tasks) + 1,
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, gin.H{
		"task": newTask,
	})
}

func updateTaskEndpoint(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true

			c.JSON(http.StatusOK, gin.H{
				"task": tasks[i],
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "task not found",
	})
}

func deleteTaskEndpoint(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "task deleted",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "task not found",
	})
}

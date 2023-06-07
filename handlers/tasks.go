package handlers

import (
	"fmt"
	"ginCli/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) GetAlLTasks(c *gin.Context) {
	t, err := h.Service.Storage.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	c.JSON(http.StatusOK, t)

}

func (h *Handler) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid task id",
		})
		return
	}
	task, err := h.Service.Storage.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"reason": "task not found",
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) GetTaskByIDUser(c *gin.Context) {
	idStr := c.Param("userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid task id",
		})
		return
	}
	user, err := h.Service.Storage.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"reason": "user not found",
		})
		return
	}

	tasks, err := h.Service.Storage.GetTasksByIDUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"reason": fmt.Sprintf("task not found with this ID %d", id),
		})
		return
	}

	var resp models.UserResponse
	for _, value := range tasks {
		m := models.UserResponse{
			ID:       value.User.ID,
			FullName: value.User.FullName,
		}
		value.User = nil
		resp = m
	}

	response := models.UserWithTasks{
		User:  &resp,
		Tasks: tasks,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateTask(c *gin.Context) {
	var t *models.Task
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "error while binding body",
		})
		return
	}

	now := time.Now()
	t.Deadline = now.AddDate(0, 0, 7)
	id, err := h.Service.Storage.AddTask(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid task id",
		})
		return
	}

	existTask, err := h.Service.Storage.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}
	if existTask == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"reason": "task not found",
		})
		return
	}

	var t *models.Task
	if err = c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "error while binding body",
		})
		return
	}
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	err = h.Service.Storage.UpdateTask(id, t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reason": "successfully updated",
	})
}

func (h *Handler) DeletedTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid task id",
		})
		return
	}

	existTask, err := h.Service.Storage.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}
	if existTask == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"reason": "task not found",
		})
		return
	}

	err = h.Service.Storage.DeletedTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reason": "successfully deleted",
	})
}

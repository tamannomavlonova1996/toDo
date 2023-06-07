package handlers

import (
	"ginCli/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) GetAlLUsers(c *gin.Context) {
	u, err := h.Service.Storage.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid user id",
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

	c.JSON(http.StatusOK, user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u *models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "error while binding body",
		})
		return
	}

	id, err := h.Service.Storage.CreateUser(u)
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

func (h *Handler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid user id",
		})
		return
	}

	existUser, err := h.Service.Storage.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}
	if existUser == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"reason": "user not found",
		})
		return
	}

	var u *models.User
	if err = c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "error while binding body",
		})
		return
	}
	u.UpdatedAt = time.Now()
	err = h.Service.Storage.UpdateUser(id, u)
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

func (h *Handler) DeletedUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid user id",
		})
		return
	}

	existUser, err := h.Service.Storage.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}
	if existUser == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"reason": "user not found",
		})
		return
	}

	err = h.Service.Storage.DeletedUser(id)
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

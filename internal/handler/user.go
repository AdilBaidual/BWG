package handler

import (
	"BWG/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) createUserHandler(c *gin.Context) {
	var tmpUser entity.User
	if err := c.BindJSON(&tmpUser); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error(err.Error())
		return
	}

	id, err := h.services.User.CreateUser(tmpUser)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) getUserByIdHandler(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		logrus.Error(err.Error())
		return
	}

	user, err := h.services.User.GetUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) getUsersHandler(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	options := entity.Options{
		Page:    page,
		PerPage: perPage,
	}

	users, err := h.services.GetUsers(options)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

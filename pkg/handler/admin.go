package handler

import (
	"net/http"

	server "example.com/server/structs"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getUsers(c *gin.Context) {
	var input server.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Authorization.GetUser(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if user.Role == "admin" {
		users, err := h.services.GetUsers()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"users": users,
		})
	} else {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) getUserDeals(c *gin.Context) {
	var input server.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Authorization.GetUser(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if user.Role == "admin" {
		deals, err := h.services.GetDealsByUserId(input.Id)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"deals": deals,
		})
	} else {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}

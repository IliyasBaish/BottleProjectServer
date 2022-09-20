package handler

import (
	"fmt"
	"net/http"

	"example.com/server/structs"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createDeal(c *gin.Context) {
	var input structs.UnclaimedDeal

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Authorization.GetUserById(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Deal.CreateDeal(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) createUnclaimedDeal(c *gin.Context) {
	fmt.Println(c.Request)
	var input structs.UnclaimedDeal

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Authorization.GetUserById(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Deal.CreateUnverifiedDeal(input, user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllDeals(c *gin.Context) {

}

func (h *Handler) getDealById(c *gin.Context) {

}

func (h *Handler) getDealsByUserId(c *gin.Context) {
}

func (h *Handler) updateDeal(c *gin.Context) {

}

func (h *Handler) deleteDeal(c *gin.Context) {

}

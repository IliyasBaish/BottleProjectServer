package handler

import (
	"fmt"
	"net/http"

	"example.com/server/structs"
	"github.com/gin-gonic/gin"
)

func (h *Handler) VerifyDeal(c *gin.Context) {
	var input structs.User
	fmt.Println("CHECK")
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err, id := h.services.Authorization.VerifyDeal(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	fmt.Println(id)
	err = h.services.Deal.VerifyLastDeal(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"verified": "ok",
	})
}

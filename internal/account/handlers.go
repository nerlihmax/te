package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.RouterGroup, service Service) {
	res := resource{service}

	r.GET("/accounts/:id", res.GetAccountById)
}

type resource struct {
	service Service
}

func (r resource) GetAccountById(c *gin.Context) {
	id := c.Param("id")

	account, err := r.service.GetAccountById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   account.ID,
		"name": account.Name,
	})
}

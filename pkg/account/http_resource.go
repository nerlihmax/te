package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPResource interface {
	GetAccountById(c *gin.Context)
	CreateAccount(c *gin.Context)
}

type resource struct {
	service Service
}

type CreateAccountRequestBody struct {
	Name string `json:"name" binding:"required"`
}

func RegisterHTTPHandlers(r *gin.RouterGroup, service Service) {
	var res HTTPResource = resource{service}

	r.
		GET("/:id", res.GetAccountById).
		POST("/", res.CreateAccount)
}

func (r resource) GetAccountById(c *gin.Context) {
	id := c.Param("id")

	account, err := r.service.GetAccountById(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   account.ID.String(),
		"name": account.Name,
	})
}

func (r resource) CreateAccount(c *gin.Context) {
	body := &CreateAccountRequestBody{}

	if err := c.BindJSON(body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	acc, err := r.service.CreateAccount(body.Name)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":   acc.ID.String(),
		"name": acc.Name,
	})
}

package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/nerlihmax/te/internal/account"
)

func Register(app *gin.Engine, db *pgx.Conn) {
	app.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errors": []string{"route not found"},
		})
	})

	app.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	{
		repo := account.NewRepository(db)
		service := account.NewService(repo)
		account.RegisterHandlers(app.Group("/accounts"), service)
	}

}

package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	router "github.com/nerlihmax/te/app"
	"github.com/nerlihmax/te/internal/db"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()

	db := db.NewDB()

	router.Register(app, db)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

package health

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	router "github.com/nerlihmax/te/app"
	"github.com/nerlihmax/te/pkg/db"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()

	db := db.NewDB()

	router.Register(app, db)

	fmt.Println("registered routes:")
	for _, item := range app.Routes() {
		fmt.Printf("[%s] %s\n", item.Method, item.Path)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

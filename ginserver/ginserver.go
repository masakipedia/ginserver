// main.go

package ginserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// GinServer serving the application
func GinServer() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("../static/templates/*")

	// Initialize the routes
	initializeRoutes()

	http.Handle("/", router)
}

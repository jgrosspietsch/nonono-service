package api

import (
	"github.com/gin-gonic/gin"
)

// SetupInitialRouter sets up the networking and middleware
func SetupInitialRouter(listenOn string, production bool) *gin.Engine {
	router := gin.Default()

	return router
}

package api

import (
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jgrosspietsch/nonono-service/puzzle"
)

// SetupRoutes sets up all the routes for the API app
func SetupRoutes(r *gin.Engine) {
	r.GET("/puzzle", puzzleHandler)
}

func puzzleHandler(c *gin.Context) {
	heightString := c.DefaultQuery("height", "10")
	widthString := c.DefaultQuery("width", "10")

	height, hErr := strconv.ParseUint(heightString, 10, 8)
	width, wErr := strconv.ParseUint(widthString, 10, 8)

	if hErr == nil && wErr == nil {
		p, err := puzzle.GeneratePuzzle(uint8(height), uint8(width))

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, p.FormatAsSerializable())
		}
	} else if hErr != nil {
		c.AbortWithError(http.StatusBadRequest, hErr)
	} else if wErr != nil {
		c.AbortWithError(http.StatusBadRequest, wErr)
	}
}

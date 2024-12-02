package handlers

import (
	"github.com/gin-gonic/gin"
)

type AreaRequest struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Radius int `json:"radius"`
}

func CheckIfInArea(c *gin.Context) {
	var request AreaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	x := request.X
	y := request.Y
	unit := 100 * request.Radius

	if x > 400 && y < 400 {
		c.JSON(200, gin.H{"result": false})
	} else if x <= 400 && y >= 400 {
		c.JSON(200, gin.H{"result": 2*(y-400) <= (x-400)+unit})
	} else if x < 400 && y < 400 {
		c.JSON(200, gin.H{"result": x > 400-unit && y > 400-unit})
	} else {
		c.JSON(200, gin.H{"result": (400-x)*(400-x)+(400-y)*(400-y) <= unit*unit/4})
	}
}

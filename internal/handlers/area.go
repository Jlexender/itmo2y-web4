package handlers

import (
	"github.com/gin-gonic/gin"
)

type AreaRequest struct {
	X              int `json:"x"`
	Y              int `json:"y"`
	Radius         int `json:"radius"`
	UnitMultiplier int `json:"unit_multiplier"`
	CanvasLimit    int `json:"canvas_limit"`
}

func CheckIfInArea(c *gin.Context) {
	var request AreaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	x := request.X
	y := request.Y
	unit := request.UnitMultiplier * request.Radius
	lim := request.CanvasLimit

	if x > lim && y < lim {
		c.JSON(200, gin.H{"result": false})
	} else if x <= lim && y >= lim {
		c.JSON(200, gin.H{"result": 2*(y-lim) <= (x-lim)+unit})
	} else if x < lim && y < lim {
		c.JSON(200, gin.H{"result": x > lim-unit && y > lim-unit})
	} else {
		c.JSON(200, gin.H{"result": (lim-x)*(lim-x)+(lim-y)*(lim-y) <= unit*unit/4})
	}
}

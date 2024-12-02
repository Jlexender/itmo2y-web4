package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"web4/internal/models"
	"web4/internal/services"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	x := request.X
	y := request.Y
	unit := request.UnitMultiplier * request.Radius
	lim := request.CanvasLimit

	username, _ := c.Get("username")

	var result bool

	if x > lim && y < lim {
		result = 2*(x-lim) <= (lim-y)+unit
	} else if x <= lim && y >= lim {
		result = 2*(y-lim) <= (x-lim)+unit
	} else if x < lim && y < lim {
		result = x > lim-unit && y > lim-unit
	} else {
		result = (lim-x)*(lim-x)+(lim-y)*(lim-y) <= unit*unit/4
	}

	areaCheck := models.AreaCheck{
		X:        x,
		Y:        y,
		Radius:   unit,
		Result:   result,
		SendTime: time.Now().Format("2006-01-02 15:04:05"),
		Author:   username.(string),
	}
	log.Printf("Area check: %v\n", areaCheck)

	services.AddAreaCheck(areaCheck)

	c.JSON(http.StatusOK, areaCheck)
}

func GetAreaChecks(c *gin.Context) {
	username, _ := c.Get("username")
	areaChecks := services.FindAllFor(username.(string))
	c.JSON(http.StatusOK, areaChecks)
}

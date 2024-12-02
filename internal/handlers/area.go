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
	X              int     `json:"x"`
	Y              int     `json:"y"`
	Radius         float64 `json:"radius"`
	UnitMultiplier int     `json:"unit_multiplier"`
	CanvasLimit    int     `json:"canvas_limit"`
}

func CheckIfInArea(c *gin.Context) {
	var request AreaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	unit := float64(request.UnitMultiplier) * request.Radius
	limit := float64(request.CanvasLimit)
	var result bool
	x, y := float64(request.X), float64(request.Y)
	if x > limit && y < limit {
		result = false
	} else if x <= limit && y >= limit {
		result = 2*(y-400) <= (x-400)+unit
	} else if x < limit && y < limit {
		result = x > 400-unit && y > 400-unit
	} else {
		result = (400-x)*(400-x)+(400-y)*(400-y) < (unit/2)*(unit/2)
	}

	realX := (x - limit) / float64(request.UnitMultiplier)
	realY := (limit - y) / float64(request.UnitMultiplier)
	username, _ := c.Get("username")

	areaCheck := models.AreaCheck{
		X:        realX,
		Y:        realY,
		Radius:   request.Radius,
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

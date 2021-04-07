package controllers

import (
	"net/http"
	"phau/cinema-demo/models"
	"phau/cinema-demo/utils"

	"github.com/labstack/echo/v4"
)

func GetSeat(c echo.Context) error {
	seats := models.GetSeats()
	return c.JSON(http.StatusOK, map[string]interface{}{"seats": seats})
}

func PutSeat(c echo.Context) error {
	data := utils.GetJSONBody(c)
	seat := models.GetSeatByRowAndColumn(data["row"].(int), data["column"].(int))
	if seat.Status != "0" {
		c.JSON(http.StatusOK, map[string]interface{}{
			"result":  "error",
			"message": "seat not free",
			"detail":  "seat status is " + seat.Status,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

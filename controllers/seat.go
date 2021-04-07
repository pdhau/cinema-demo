package controllers

import (
	"net/http"
	"phau/cinema-demo/models"
	"phau/cinema-demo/utils"

	"github.com/labstack/echo/v4"
)

func GetSeat(c echo.Context) error {
	condition := map[string]string{}
	condition["group"] = c.Param("group")
	seats := models.GetSeats(condition)
	return c.JSON(http.StatusOK, map[string]interface{}{"seats": seats})
}

func PutSeat(c echo.Context) error {
	data := utils.GetJSONBody(c)
	seat, err := models.GetSeatByRowAndColumn(data["row"].(int), data["column"].(int))
	if err != nil {
		panic(err)
	}
	if seat.Status == "1" {
		c.JSON(http.StatusOK, map[string]interface{}{
			"result":  "error",
			"message": "seat is not free",
			"detail":  "seat status is " + seat.Status,
		})
	}
	if seat.Status == "X" && seat.Group != data["group"].(string) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"result":  "error",
			"message": "seat is near with other group",
			"detail":  "seat group is " + seat.Group + " and " + data["group"].(string),
		})
	}
	models.UpdateSeatStatusByRowAndColumn(seat.Row, seat.Column, "1")
	for _, s := range models.GetSeats(map[string]string{}) {
		if !s.Equal(seat) {
			models.UpdateSeatStatusByRowAndColumn(seat.Row, seat.Column, "X")
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

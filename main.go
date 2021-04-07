package main

import (
	"fmt"
	"phau/cinema-demo/configs"
	"phau/cinema-demo/controllers"
	"phau/cinema-demo/utils"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/v1/seat", controllers.GetSeat)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

func init() {
	db := utils.GetDBConnect()
	_, err := db.Exec("DELETE FROM seat")
	if err != nil {
		panic(err)
	}

	for col := 0; col < configs.Column; col++ {
		for row := 0; row < configs.Column; row++ {
			_, err := db.Exec("INSERT INTO seat VALUES (?,?,?,?)", row, col, "0", "")
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("Init successfully")
}

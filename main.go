package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", hello)
	e.GET("/v1/seat", getSeat)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getSeat(c echo.Context) error {
	seats := []Seat{}
	db, err := sql.Open("sqlite3", "./db/cinema-demo.db")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * from seat")
	if err != nil {
		panic(err)
	}

	for rows.Next(){
		seat := Seat{}
		err := rows.Scan(&seat.Row, &seat.Column, &seat.Status)
		if err != nil {
			panic(err)
		}
		seats = append(seats, seat)
	}
	return c.JSON(http.StatusOK, seats)
}

type Seat struct {
	Row    int
	Column int
	Status string
}

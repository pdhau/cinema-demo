package models

import (
	"phau/cinema-demo/utils"
)

type Seat struct {
	Row    int
	Column int
	Status string
}

func GetSeats() []Seat {
	seats := []Seat{}
	db := utils.GetDBConnect()
	rows, err := db.Query("SELECT * from seat")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		seat := Seat{}
		err := rows.Scan(&seat.Row, &seat.Column, &seat.Status)
		if err != nil {
			panic(err)
		}
		seats = append(seats, seat)
	}

	return seats
}

func GetSeatByRowAndColumn(row int, column int) Seat {
	db := utils.GetDBConnect()
	rows, err := db.Query("SELECT * FROM seat WHERE row = ? AND column = ?", row, column)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	seat := Seat{}
	for rows.Next() {
		err := rows.Scan(&seat.Row, &seat.Column, &seat.Status)
		if err != nil {
			panic(err)
		}
	}

	return seat
}

func UpdateSeatStatusByRowAndColumn(row int, column int, status string) {
	db := utils.GetDBConnect()
	_, err := db.Exec("UPDATE seat SET status = ? WHERE row = ? AND column = ?", status, row, column)
	if err != nil {
		panic(err)
	}
}

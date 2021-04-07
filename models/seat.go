package models

import (
	"fmt"
	"math"
	"phau/cinema-demo/utils"
	"strings"
)

type Seat struct {
	Row    int    `json:"row"`
	Column int    `json:"column"`
	Status string `json:"status"`
	Group  string `json:"group"`
}

func (s Seat) IsSameGroup(seat Seat) bool {
	return s.Group == seat.Group
}

func (s Seat) Equal(seat Seat) bool {
	return s.Row == seat.Row && s.Column == seat.Column
}

func (s Seat) GetDistance(seat Seat) int {
	return int(math.Abs(float64(s.Row-seat.Row)) + math.Abs(float64(s.Column-seat.Column)))
}

func GetSeats(condition map[string]string) []Seat {
	whereClause := []string{"TRUE"}
	valueForWhereClause := []interface{}{}
	for k, v := range condition {
		switch k {
		case "group":
			whereClause = append(whereClause, fmt.Sprintf("%s = ? AND %s = \"\"", k, k))
			valueForWhereClause = append(valueForWhereClause, v)
		default:
			whereClause = append(whereClause, fmt.Sprintf("%s = ?", k))
			valueForWhereClause = append(valueForWhereClause, v)
		}

	}
	seats := []Seat{}
	db := utils.GetDBConnect()
	selectSQL := "SELECT * FROM seat WHERE " + strings.Join(whereClause, " AND ")
	rows, err := db.Query(selectSQL, valueForWhereClause...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		seat := Seat{}
		err := rows.Scan(&seat.Row, &seat.Column, &seat.Status, &seat.Group)
		if err != nil {
			panic(err)
		}
		seats = append(seats, seat)
	}

	return seats
}

func GetSeatByRowAndColumn(row int, column int) (Seat, error) {
	db := utils.GetDBConnect()
	rows, err := db.Query("SELECT * FROM seat WHERE row = ? AND column = ?", row, column)
	if err != nil {
		return Seat{}, err
	}
	defer rows.Close()

	seat := Seat{}
	for rows.Next() {
		err := rows.Scan(&seat.Row, &seat.Column, &seat.Status)
		if err != nil {
			return Seat{}, err
		}
	}

	return seat, nil
}

func UpdateSeatStatusByRowAndColumn(row int, column int, status string) error {
	db := utils.GetDBConnect()
	_, err := db.Exec("UPDATE seat SET status = ? WHERE row = ? AND column = ?", status, row, column)
	return err
}

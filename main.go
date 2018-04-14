package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Streak struct {
	ID                     int    `db:"id"`
	AccumulatorKey         string `db:"accumulator_key"`
	AccumulatorValue       string `db:"accumulator_value"`
	AccumulatorDescription string `db:"accumulator_description"`
	DateStart              string `db:"date_start"`
	DateEnd                string `db:"date_end"`
	UserID                 int    `db:"user_id"`
	GoalID                 int    `db:"goal_id"`
}

func main() {
	db, err := sqlx.Connect("mysql", "streaking:streaking@/streaking")
	if err != nil {
		log.Fatalln(err)
	}

	streaks := []Streak{}
	db.Select(&streaks, "SELECT * FROM streaks")
	fmt.Println(streaks)
}

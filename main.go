package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type streak struct {
	ID                     int    `db:"id"`
	AccumulatorKey         string `db:"accumulator_key"`
	AccumulatorValue       string `db:"accumulator_value"`
	AccumulatorDescription string `db:"accumulator_description"`
	DateStart              string `db:"date_start"`
	DateEnd                string `db:"date_end"`
	UserID                 int    `db:"user_id"`
	GoalID                 int    `db:"goal_id"`
}

var db *sqlx.DB

func initDB(dataSourceName string) {
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

// people := []Person{}
// db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
// jason, john := people[0], people[1]

func readStreaks() {
	// streaks := []streak{}
	var ids []int
	db.Select(&ids, "SELECT id FROM streaks")

	// rows, err := db.Query("SELECT * FROM streaks")
	// streaks := []streak{}

	// err := db.Select(&streaks, "SELECT * FROM streaks")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// streaks := make([]*streak, 0)
	// for rows.Next() {
	// 	newStreak := new(streak)
	// 	err := rows.Scan(
	// 		&newStreak.ID,
	// 		&newStreak.AccumulatorKey,
	// 		&newStreak.AccumulatorValue,
	// 		&newStreak.AccumulatorDescription,
	// 		&newStreak.DateStart,
	// 		&newStreak.DateEnd,
	// 		&newStreak.UserID,
	// 		&newStreak.GoalID)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	streaks = append(streaks, newStreak)
	// }

	// if err = rows.Err(); err != nil {
	// 	log.Panic(err)
	// }

	fmt.Println(ids)
}

func main() {
	initDB("streaking:streaking@/streaking")
	readStreaks()
}

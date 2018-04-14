package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
 * types - with field tags for both db and json
 */
type streak struct {
	ID                     int    `db:"id" json:"id"`
	AccumulatorKey         string `db:"accumulator_key" json:"accumulator_key"`
	AccumulatorValue       string `db:"accumulator_value" json:"accumulator_value"`
	AccumulatorDescription string `db:"accumulator_description" json:"accumulator_description"`
	DateStart              string `db:"date_start" json:"date_start"`
	DateEnd                string `db:"date_end" json:"date_end"`
	UserID                 int    `db:"user_id" json:"user_id"`
	GoalID                 int    `db:"goal_id" json:"goal_id"`
}

type user struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

type goal struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

/*
 * models, read
 */
func readUsers(db *sqlx.DB) []user {
	users := []user{}
	db.Select(&users, "SELECT * FROM users")
	return users
}

func readGoals(db *sqlx.DB) []goal {
	goals := []goal{}
	db.Select(&goals, "SELECT * FROM goals")
	return goals
}

func readStreaks(db *sqlx.DB) []streak {
	streaks := []streak{}
	db.Select(&streaks, "SELECT * FROM streaks")
	return streaks
}

func main() {
	db, err := sqlx.Connect("mysql", "streaking:streaking@/streaking")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(readStreaks(db))
}

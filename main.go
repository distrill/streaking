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

func main() {
	db, err := sqlx.Connect("mysql", "streaking:streaking@/streaking")
	if err != nil {
		log.Panic(err)
	}

	um := userModel{db}
	gm := goalModel{db}
	sm := streakModel{db}

	// read
	fmt.Println(um.read(nil))
	fmt.Println(gm.read(nil))
	fmt.Println(sm.read(nil))

	// create
	um.create(user{0, "brent 4", "brent 4 email"})
	gm.create(goal{0, "goal 4", "this is the 4th goal"})
	sm.create(streak{0, "key 4", "value 4", "description 4", "2018-04-01", "2018-04-13", 2, 2})

	// update
	um.update(1, user{1, "another updated name", "another updated email"})
	fmt.Println(um.read(nil))
	gm.update(1, goal{1, "updated goal name", "updated goal description"})
	fmt.Println(gm.read(nil))
	sm.update(1, streak{1, "updated key", "updated value", "updated accumulator", "2018-01-01", "2018-01-02", 1, 1})
	fmt.Println(sm.read(nil))

	// delete
	fmt.Println(sm.read(nil))
	sm.delete(1)
	fmt.Println(sm.read(nil))
	fmt.Println(um.read(nil))
	um.delete(1)
	fmt.Println(um.read(nil))
	fmt.Println(gm.read(nil))
	gm.delete(1)
	fmt.Println(gm.read(nil))
}

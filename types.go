package main

/*
 * data types, from db and written to json
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
	UserID      int    `db:"user_id" json:"user_id"`
}

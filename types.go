package main

/*
 * data types, from db and written to json
 */
type user struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

type goal struct {
	ID                     int    `db:"id" json:"id"`
	Name                   string `db:"name" json:"name"`
	Description            string `db:"description" json:"description"`
	UserID                 int    `db:"user_id" json:"user_id"`
	UpdateInterval         string `db:"update_interval" json:"update_interval"`
	AccumulatorKey         string `db:"accumulator_key" json:"accumulator_key"`
	AccumulatorIncrement   string `db:"accumulator_increment" json:"accumulator_increment"`
	AccumulatorDescription string `db:"accumulator_description" json:"accumulator_description"`
}

type streak struct {
	ID        int    `db:"id" json:"id"`
	DateStart string `db:"date_start" json:"date_start"`
	DateEnd   string `db:"date_end" json:"date_end"`
	GoalID    int    `db:"goal_id" json:"goal_id"`
}

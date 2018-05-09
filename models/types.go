package models

/*
 * data types, from db and written to json
 */

// User - users
type User struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Email      string `db:"email" json:"email"`
	Source     string `db:"source" json:"source"`
	ExternalID string `db:"external_id" json:"external_id"`
}

// Goal - goals
type Goal struct {
	ID                     int    `db:"id" json:"id"`
	Name                   string `db:"name" json:"name"`
	Description            string `db:"description" json:"description"`
	Color                  string `db:"color" json:"color"`
	UserID                 int    `db:"user_id" json:"user_id"`
	UpdateInterval         string `db:"update_interval" json:"update_interval"`
	AccumulatorKey         string `db:"accumulator_key" json:"accumulator_key"`
	AccumulatorIncrement   string `db:"accumulator_increment" json:"accumulator_increment"`
	AccumulatorDescription string `db:"accumulator_description" json:"accumulator_description"`
}

// Streak - streaks
type Streak struct {
	ID        int    `db:"id" json:"id"`
	DateStart string `db:"date_start" json:"date_start"`
	DateEnd   string `db:"date_end" json:"date_end"`
	GoalID    int    `db:"goal_id" json:"goal_id"`
}

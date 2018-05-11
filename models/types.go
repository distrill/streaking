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
	ExternalID string `db:"external_id" json:"externalId"`
}

// Goal - goals
type Goal struct {
	ID                     int    `db:"id" json:"id"`
	Name                   string `db:"name" json:"name"`
	Description            string `db:"description" json:"description"`
	Color                  string `db:"color" json:"color"`
	UserID                 int    `db:"user_id" json:"userId"`
	UpdateInterval         string `db:"update_interval" json:"updateInterval"`
	AccumulatorKey         string `db:"accumulator_key" json:"accumulatorKey"`
	AccumulatorIncrement   string `db:"accumulator_increment" json:"accumulatorIncrement"`
	AccumulatorDescription string `db:"accumulator_description" json:"accumulatorDescription"`
}

// Streak - streaks
type Streak struct {
	ID        int    `db:"id" json:"id"`
	DateStart string `db:"date_start" json:"dateStart"`
	DateEnd   string `db:"date_end" json:"dateEnd"`
	GoalID    int    `db:"goal_id" json:"goalId"`
}

package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type model struct{ db *sqlx.DB }

type userModel model
type goalModel model
type streakModel model
type userGoalModel model

func applySearch(qs string, search map[string]interface{}) string {
	if search == nil {
		return qs
	}

	delim := "WHERE"
	for k, v := range search {
		if _, ok := v.(string); ok {
			v = fmt.Sprintf("'%v'", v)
		}
		// NOTE: this is bad and not escaped. Should use prepared statements.
		// This means getting something like Object.values(search) and destructuring below
		qs = fmt.Sprintf("%v %v %v = %v", qs, delim, k, v)
		delim = "AND"
	}
	return qs
}

/*
 * Read
 */
func (um userModel) read(search map[string]interface{}) ([]user, error) {
	userResults := []user{}
	qs := applySearch("SELECT * FROM users", search)

	if err := um.db.Select(&userResults, qs); err != nil {
		return nil, err
	}

	return userResults, nil
}

func (gm goalModel) read(search map[string]interface{}) ([]goal, error) {
	goalResults := []goal{}

	selectString := "SELECT goals.id, name, description"
	fromString := " FROM goals"

	if search["user_id"] != nil {
		selectString += ", user_id"
		fromString += " INNER JOIN users_goals ON goals.id = users_goals.goal_id"
	}

	qs := applySearch(selectString+fromString, search)
	fmt.Println(qs)

	if err := gm.db.Select(&goalResults, qs); err != nil {
		return nil, err
	}

	return goalResults, nil
}

func (sm streakModel) read(search map[string]interface{}) ([]streak, error) {
	streakResults := []streak{}
	qs := applySearch(`
		SELECT
			streaks.id,
			accumulator_key,
			accumulator_value,
			accumulator_description,
			date_start,
			date_end,
			user_id,
			goal_id
		FROM
			streaks
	`, search)

	if err := sm.db.Select(&streakResults, qs); err != nil {
		return nil, err
	}

	return streakResults, nil
}

/*
 * Create
 */
func (um userModel) create(u user) error {
	_, err := um.db.NamedExec(`
		INSERT INTO users (name, email)
		VALUES (:name, :email)
	`, &u)
	if err != nil {
		return err
	}
	return nil
}

func (gm goalModel) create(g goal) error {
	_, err := gm.db.NamedExec(`
		INSERT INTO goals (name, description)
		VALUES (:name, :description)
	`, &g)
	if !isErrDuplicateEntry(err) {
		return err
	}
	return nil
}

func (ugm userGoalModel) create(ug userGoal) error {
	_, err := ugm.db.NamedExec(`
		INSERT INTO users_goals (user_id, goal_id)
		VALUES (:user_id, :goal_id)
	`, &ug)
	if !isErrDuplicateEntry(err) {
		return err
	}
	return nil
}

func (sm streakModel) create(s streak) error {
	_, err := sm.db.NamedExec(`
		INSERT INTO streaks (
			accumulator_key,
			accumulator_value,
			accumulator_description,
			date_start,
			date_end,
			user_id,
			goal_id
		) VALUES (
			:accumulator_key,
			:accumulator_value,
			:accumulator_description,
			:date_start,
			:date_end,
			:user_id,
			:goal_id
		)
	`, &s)
	if !isErrDuplicateEntry(err) {
		return err
	}
	return nil
}

/*
 * Update
 */
func (um userModel) update(id int, u user) error {
	u.ID = id
	_, err := um.db.NamedExec(`
		UPDATE users
		SET name = :name, email = :email
		WHERE id = :id
	`, &u)
	if err != nil {
		return err
	}
	return nil
}

func (gm goalModel) update(id int, g goal) error {
	g.ID = id
	_, err := gm.db.NamedExec(`
		UPDATE goals
		SET name = :name, description = :description
		WHERE id = :id
	`, &g)
	if err != nil {
		return err
	}
	return nil
}

func (sm streakModel) update(id int, s streak) error {
	s.ID = id
	_, err := sm.db.NamedExec(`
		UPDATE streaks
		SET
			accumulator_key = :accumulator_key,
			accumulator_value = :accumulator_value,
			accumulator_description = :accumulator_description,
			date_start = :date_start,
			date_end = :date_end,
			user_id = :user_id,
			goal_id = :goal_id
		WHERE id = :id
	`, &s)
	if err != nil {
		return err
	}
	return nil
}

/*
 * Delete
 */
func delete(db *sqlx.DB, id int, table string) {
	db.MustExec("DELETE FROM "+table+" WHERE id = ?", id)
}

func (um userModel) delete(id int) {
	delete(um.db, id, "users")
}

func (gm goalModel) delete(id int) {
	delete(gm.db, id, "goals")
}

func (sm streakModel) delete(id int) {
	delete(sm.db, id, "streaks")
}

package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type model interface {
	create(interface{}) error
	read(interface{}) (error, interface{})
	update(interface{}) error
	delete(int) error
}

type userModel struct{ db *sqlx.DB }
type goalModel struct{ db *sqlx.DB }
type streakModel struct{ db *sqlx.DB }

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
func (um userModel) read(search map[string]interface{}) []user {
	userResults := []user{}
	qs := applySearch("SELECT * FROM users", search)
	um.db.Select(&userResults, qs)
	return userResults
}

func (gm goalModel) read(search map[string]interface{}) []goal {
	goalResults := []goal{}
	qs := applySearch("SELECT * FROM goals", search)
	gm.db.Select(&goalResults, qs)
	return goalResults
}

func (sm streakModel) read(search map[string]interface{}) []streak {
	streakResults := []streak{}
	qs := applySearch("SELECT * FROM streaks", search)
	sm.db.Select(&streakResults, qs)
	return streakResults
}

/*
 * Create
 */
func (um userModel) create(u user) {
	_, err := um.db.NamedExec(`
		INSERT INTO users (name, email)
		VALUES (:name, :email)
	`, &u)
	if err != nil {
		log.Fatal(err)
	}
}

func (gm goalModel) create(g goal) {
	_, err := gm.db.NamedExec(`
		INSERT INTO goals (name, description)
		VALUES (:name, :description)
	`, &g)
	if err != nil {
		log.Fatal(err)
	}
}

func (sm streakModel) create(s streak) {
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
	if err != nil {
		log.Fatal(err)
	}
}

/*
 * Update
 */
func (um userModel) update(id int, u user) {
	u.ID = id
	_, err := um.db.NamedExec(`
		UPDATE users
		SET name = :name, email = :email
		WHERE id = :id
	`, &u)
	if err != nil {
		log.Fatal(err)
	}
}

func (gm goalModel) update(id int, g goal) {
	g.ID = id
	_, err := gm.db.NamedExec(`
		UPDATE goals
		SET name = :name, description = :description
		WHERE id = :id
	`, &g)
	if err != nil {
		log.Fatal(err)
	}
}

func (sm streakModel) update(id int, s streak) {
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
		log.Fatal(err)
	}
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

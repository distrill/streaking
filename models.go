package main

import (
	"log"

	"github.com/jmoiron/sqlx"
)

/*
 * [C]rud - Create
 */
func createUser(db *sqlx.DB, u user) {
	_, err := db.NamedExec(`
		INSERT INTO users (name, email)
		VALUES (:name, :email)
	`, &u)
	if err != nil {
		log.Fatal(err)
	}
}

func createGoal(db *sqlx.DB, g goal) {
	_, err := db.NamedExec(`
		INSERT INTO goals (name, description)
		VALUES (:name, :description)
	`, &g)
	if err != nil {
		log.Fatal(err)
	}
}

func createStreak(db *sqlx.DB, s streak) {
	_, err := db.NamedExec(`
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
 * c[R]ud - Read
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

/*
 * cr[U]d - Update
 */
func updateUser(db *sqlx.DB, id int, u user) {
	u.ID = id
	_, err := db.NamedExec(`
		UPDATE users
		SET name = :name, email = :email
		WHERE id = :id
	`, &u)
	if err != nil {
		log.Fatal(err)
	}
}

func updateGoal(db *sqlx.DB, id int, g goal) {
	g.ID = id
	_, err := db.NamedExec(`
		UPDATE goals
		SET name = :name, description = :description
		WHERE id = :id
	`, &g)
	if err != nil {
		log.Fatal(err)
	}
}

func updateStreak(db *sqlx.DB, id int, s streak) {
	s.ID = id
	_, err := db.NamedExec(`
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
 * cru[D] - Delete
 */
func delete(db *sqlx.DB, id int, table string) {
	db.MustExec("DELETE FROM "+table+" WHERE id = ?", id)
}

func deleteUser(db *sqlx.DB, id int) {
	delete(db, id, "users")
}

func deleteGoal(db *sqlx.DB, id int) {
	delete(db, id, "goals")
}

func deleteStreak(db *sqlx.DB, id int) {
	delete(db, id, "streaks")
}

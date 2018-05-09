package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type model struct{ db *sqlx.DB }

type userModel model
type goalModel model
type streakModel model

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
func (um userModel) read(search map[string]interface{}) ([]User, error) {
	userResults := []User{}

	qs := applySearch("SELECT * FROM users", search)
	fmt.Println(formatQuery(qs))

	if err := um.db.Select(&userResults, qs); err != nil {
		return nil, err
	}

	return userResults, nil
}

func (gm goalModel) read(search map[string]interface{}) ([]Goal, error) {
	gs := []Goal{}

	selectString := `
		SELECT
			goals.id,
			goals.name,
			goals.description,
			goals.color,
			goals.user_id,
			goals.update_interval,
			goals.accumulator_key,
			goals.accumulator_increment,
			goals.accumulator_description
	`
	fromString := " FROM goals"

	if search["user_id"] != nil {
		selectString += ", users.id AS user_id"
		fromString += " INNER JOIN users ON users.id = goals.user_id"
	}

	qs := applySearch(selectString+fromString, search)
	fmt.Println(qs)

	if err := gm.db.Select(&gs, qs); err != nil {
		return nil, err
	}

	return gs, nil
}

func (sm streakModel) read(search map[string]interface{}) ([]Streak, error) {
	streakResults := []Streak{}

	selectString := "SELECT streaks.*"
	fromString := " FROM streaks"

	if search["user_id"] != nil {
		fromString += `
			INNER JOIN goals ON goals.id = streaks.goal_id
			INNER JOIN users ON users.id = goals.user_id
		`
	}

	qs := applySearch(selectString+fromString, search)
	fmt.Println(qs)

	if err := sm.db.Select(&streakResults, qs); err != nil {
		return nil, err
	}

	return streakResults, nil
}

/*
 * Create
 */
func (um userModel) create(u User) error {
	qs := `
		INSERT INTO users (name, email)
		VALUES (:name, :email)
	`
	fmt.Println(formatQuery(qs))

	if _, err := um.db.NamedExec(qs, &u); !isErrDuplicateEntry(err) {
		return err
	}

	return nil
}

func (gm goalModel) create(g Goal) error {
	fmt.Println(g)
	qs := `
		INSERT INTO goals (
			name,
			description,
			color,
			user_id,
			update_interval,
			accumulator_key,
			accumulator_increment,
			accumulator_description
		)
		VALUES (
			:name,
			:description,
			:color,
			:user_id,
			:update_interval,
			:accumulator_key,
			:accumulator_increment,
			:accumulator_description
		)
	`
	fmt.Println(formatQuery(qs))

	if _, err := gm.db.NamedExec(qs, &g); !isErrDuplicateEntry(err) {
		return err
	}

	return nil
}

func (sm streakModel) create(s Streak) error {
	qs := `
		INSERT INTO streaks (
			date_start,
			date_end,
			goal_id
		) VALUES (
			:date_start,
			:date_end,
			:goal_id
		)
	`
	fmt.Println(formatQuery(qs))

	if _, err := sm.db.NamedExec(qs, &s); !isErrDuplicateEntry(err) {
		return err
	}

	return nil
}

/*
 * Update
 */
func (um userModel) update(id int, u User) error {
	u.ID = id

	qs := `
        UPDATE users
        SET name = :name, email = :email
        WHERE id = :id
    `
	fmt.Println(formatQuery(qs))

	if _, err := um.db.NamedExec(qs, &u); err != nil {
		return err
	}
	return nil
}

func (gm goalModel) update(id int, g Goal) error {
	g.ID = id

	qs := `
		UPDATE goals
		SET
			name = :name,
			description = :description,
			color = :color,
			user_id = :user_id,
			update_interval = :update_interval,
			accumulator_key = :accumulator_key,
			accumulator_increment = :accumulator_increment,
			accumulator_description = :accumulator_description
		WHERE id = :id
	`
	fmt.Println(formatQuery(qs))

	if _, err := gm.db.NamedExec(qs, &g); err != nil {
		return err
	}
	return nil
}

func (sm streakModel) update(id int, s Streak) error {
	s.ID = id

	qs := `
		UPDATE streaks
		SET
			date_start = :date_start,
			date_end = :date_end,
			goal_id = :goal_id
		WHERE id = :id
	`
	fmt.Println(formatQuery(qs))

	if _, err := sm.db.NamedExec(qs, &s); err != nil {
		fmt.Println(err)
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

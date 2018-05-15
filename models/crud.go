package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type model struct{ DB *sqlx.DB }

// Users - users model
type Users model

// Goals - goals model
type Goals model

// Streaks - streaks model
type Streaks model

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
func (um Users) Read(search map[string]interface{}) ([]User, error) {
	userResults := []User{}

	qs := applySearch("SELECT * FROM users", search)
	fmt.Println(FormatQuery(qs))

	if err := um.DB.Select(&userResults, qs); err != nil {
		return nil, err
	}

	return userResults, nil
}

func (gm Goals) Read(search map[string]interface{}) ([]Goal, error) {
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

	if err := gm.DB.Select(&gs, qs); err != nil {
		return nil, err
	}

	return gs, nil
}

func (sm Streaks) Read(search map[string]interface{}) ([]Streak, error) {
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

	if err := sm.DB.Select(&streakResults, qs); err != nil {
		return nil, err
	}

	return streakResults, nil
}

// Create - create given user
func (um Users) Create(u User) error {
	qs := `
		INSERT INTO users (name, email, source, external_id)
		VALUES (:name, :email, :source, :external_id)
	`
	fmt.Println(FormatQuery(qs))

	if _, err := um.DB.NamedExec(qs, &u); !IsErrDuplicateEntry(err) {
		return err
	}

	return nil
}

// Create - creative given goal
func (gm Goals) Create(g Goal) error {
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
	fmt.Println(FormatQuery(qs))

	if _, err := gm.DB.NamedExec(qs, &g); !IsErrDuplicateEntry(err) {
		return err
	}

	return nil
}

// Create - create given streak
func (sm Streaks) Create(s Streak) error {
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
	fmt.Println(FormatQuery(qs))

	if _, err := sm.DB.NamedExec(qs, &s); !IsErrDuplicateEntry(err) {
		return err
	}

	return nil
}

// Update - update given user
func (um Users) Update(id int, u User) error {
	u.ID = id

	qs := `
        UPDATE users
				SET
					name = :name,
					email = :email,
					source = :source,
					external_id = :external_id
        WHERE id = :id
    `
	fmt.Println(FormatQuery(qs))

	if _, err := um.DB.NamedExec(qs, &u); err != nil {
		return err
	}
	return nil
}

// Update - update given goal
func (gm Goals) Update(id int, g Goal) error {
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
	fmt.Println(FormatQuery(qs))

	if _, err := gm.DB.NamedExec(qs, &g); err != nil {
		return err
	}
	return nil
}

// Update - update given streak
func (sm Streaks) Update(id int, s Streak) error {
	s.ID = id

	qs := `
		UPDATE streaks
		SET
			date_start = :date_start,
			date_end = :date_end,
			goal_id = :goal_id
		WHERE id = :id
	`
	fmt.Println(FormatQuery(qs))

	if _, err := sm.DB.NamedExec(qs, &s); err != nil {
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

// Delete - delete given user
func (um Users) Delete(id int) {
	delete(um.DB, id, "users")
}

// Delete - delete given goal
func (gm Goals) Delete(id int) {
	delete(gm.DB, id, "goals")
}

// Delete - delete given streak
func (sm Streaks) Delete(id int) {
	delete(sm.DB, id, "streaks")
}

package main

import (
	"strconv"
	"testing"

	"bh/streaking/models"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func TestFormatQuery(t *testing.T) {
	expected := "SELECT * FROM table"
	actual := models.FormatQuery(`
		SELECT *
		FROM table
		`)
	if expected != actual {
		t.Error("formatQuery: expected '" + actual + "' to equal '" + expected + "'")
	}

	expected = "SELECT * FROM table"
	actual = formatQuery("SELECT * FROM table")

	if expected != actual {
		t.Error("formatQuery: expected '" + actual + "' to equal '" + expected + "'")
	}
}

func TestIsErrDuplicateEntry(t *testing.T) {
	db, err := sqlx.Connect("mysql", "streaking:streaking@/streaking")
	if err != nil {
		t.Error("isDuplicateEntry failed connecting to database", err)
	}

	u := user{1, "name", "email"}
	_, err = db.NamedExec("INSERT INTO users VALUES (:id, :name, :email)", &u)
	if err != nil {
		t.Error("isDuplicateEntry failed insrting initial seed data", err)
	}

	_, err = db.NamedExec("INSERT INTO users VALUES (:id, :name, :email)", &u)
	expected := true
	actual := isErrDuplicateEntry(err)
	if expected != actual {
		t.Error("isErrDuplicateEntry: expected '" + strconv.FormatBool(actual) + "' to equal '" + strconv.FormatBool(expected) + "'")
	}

	_, err = db.NamedExec("this is a nonsense query", &u)
	expected = false
	actual = isErrDuplicateEntry(err)
	if expected != actual {
		t.Error("isErrDuplicateEntry: expected '" + strconv.FormatBool(actual) + "' to equal '" + strconv.FormatBool(expected) + "'")
	}

	expected = false
	actual = isErrDuplicateEntry(nil)
	if expected != actual {
		t.Error("isErrDuplicateEntry: expected '" + strconv.FormatBool(actual) + "' to equal '" + strconv.FormatBool(expected) + "'")
	}
}

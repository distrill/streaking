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
	actual = models.FormatQuery("SELECT * FROM table")

	if expected != actual {
		t.Error("formatQuery: expected '" + actual + "' to equal '" + expected + "'")
	}
}

func TestIsErrDuplicateEntry(t *testing.T) {
	db, err := sqlx.Connect("mysql", "streaking:streaking@/streaking")
	if err != nil {
		t.Error("isDuplicateEntry failed connecting to database", err)
	}

	u := models.User{ID: 1, Name: "name", Email: "email", ExternalID: "id", Source: "source"}
	_, err = db.NamedExec("INSERT INTO users VALUES (:id, :name, :email)", &u)
	if err != nil {
		t.Error("isDuplicateEntry failed insrting initial seed data", err)
	}

	_, err = db.NamedExec("INSERT INTO users VALUES (:id, :name, :email)", &u)
	expected := true
	actual := models.IsErrDuplicateEntry(err)
	if expected != actual {
		t.Error("isErrDuplicateEntry: expected '" + strconv.FormatBool(actual) + "' to equal '" + strconv.FormatBool(expected) + "'")
	}

	_, err = db.NamedExec("this is a nonsense query", &u)
	expected = false
	actual = models.IsErrDuplicateEntry(err)
	if expected != actual {
		t.Error("isErrDuplicateEntry: expected '" + strconv.FormatBool(actual) + "' to equal '" + strconv.FormatBool(expected) + "'")
	}

	expected = false
	actual = models.IsErrDuplicateEntry(nil)
	if expected != actual {
		t.Error("isErrDuplicateEntry: expected '" + strconv.FormatBool(actual) + "' to equal '" + strconv.FormatBool(expected) + "'")
	}
}

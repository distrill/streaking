package main

import (
	"strings"

	"github.com/go-sql-driver/mysql"
)

func formatQuery(qs string) string {
	noTabs := strings.Replace(qs, "\t", "", -1)
	noTabsOrSpaces := strings.Replace(noTabs, "\n", " ", -1)
	return strings.Trim(noTabsOrSpaces, " ")
}

func isErrDuplicateEntry(err error) bool {
	me, ok := err.(*mysql.MySQLError)
	return err != nil && ok && me.Number == 1062
}

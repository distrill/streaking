package models

import (
	"strings"

	"github.com/go-sql-driver/mysql"
)

// FormatQuery - remove additional whitespace for printing
func FormatQuery(qs string) string {
	noTabs := strings.Replace(qs, "\t", "", -1)
	noTabsOrSpaces := strings.Replace(noTabs, "\n", " ", -1)
	return strings.Trim(noTabsOrSpaces, " ")
}

func isErrDuplicateEntry(err error) bool {
	me, ok := err.(*mysql.MySQLError)
	return err != nil && ok && me.Number == 1062
}

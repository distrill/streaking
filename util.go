package main

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func printQuery(qs string) {
	noTabs := strings.Replace(qs, "\t", "", -1)
	noTabsOrSpaces := strings.Replace(noTabs, "\n", " ", -1)
	fmt.Println(noTabsOrSpaces)
}

func isErrDuplicateEntry(err error) bool {
	me, ok := err.(*mysql.MySQLError)
	return err != nil && ok && me.Number == 1062
}

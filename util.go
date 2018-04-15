package main

import (
	"fmt"
	"strings"
)

func printQuery(qs string) {
	noTabs := strings.Replace(qs, "\t", "", -1)
	noTabsOrSpaces := strings.Replace(noTabs, "\n", " ", -1)
	fmt.Println(noTabsOrSpaces)
}

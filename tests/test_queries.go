package main

import (
	"MySQLHelper/pkg"
	"fmt"
	"time"
)

func main() {
	db := pkg.MySQLConnection(true)
	defer db.Close()

	rows, err := db.Query("SELECT NOW() dt")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type row struct {
		now time.Time
	}

	for rows.Next() {
		var r row
		rows.Scan(&r.now)

		var dt string
		if r.now.IsZero() {
			dt = "0"
		} else {
			dt = r.now.Format(time.RFC3339)
		}

		fmt.Printf("datetime: %s\n", dt)
	}
}

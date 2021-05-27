package main

import (
	"MySQLHelper/pkg"
	"fmt"
	"time"
	"os"
)


func main() {
	dbConf := pkg.MySQLConfig {
		DbHost: os.Getenv("MYSQL_HOSTNAME"),
		DbPass: os.Getenv("MYSQL_PASSWORD"),
		DbUser: os.Getenv("MYSQL_USERNAME"),
		DbPort: os.Getenv("MYSQL_PORT"),
		DbName: os.Getenv("MYSQL_DATABASE"),
		UseSSH: true,
		SshKeyPath: os.Getenv("SSH_KEYPATH"),
		SshHost: os.Getenv("SSH_HOST"),
		SshUser: os.Getenv("SSH_USER"),
		SshPort: os.Getenv("SSH_PORT"),
	}

	db, err := dbConf.Connect()
	if err != nil {
		panic(err)
	}
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

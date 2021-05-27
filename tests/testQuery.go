package main

import (
	"fmt"
	"github.com/zimza/go-mysql-ssh-helper/pkg"
	"os"
)

func main() {
	db := pkg.MySQLConfig{
		DbHost:     os.Getenv("MYSQL_HOSTNAME"),
		DbPass:     os.Getenv("MYSQL_PASSWORD"),
		DbUser:     os.Getenv("MYSQL_USERNAME"),
		DbPort:     os.Getenv("MYSQL_PORT"),
		DbName:     os.Getenv("MYSQL_DATABASE"),
		UseSSH:     true,
		SshKeyPath: os.Getenv("SSH_KEYPATH"),
		SshHost:    os.Getenv("SSH_HOST"),
		SshUser:    os.Getenv("SSH_USER"),
		SshPort:    os.Getenv("SSH_PORT"),
	}

	cnx, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer cnx.Close()

	rows, err := cnx.Query("SELECT 'Hello World';")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type row struct {
		hello string
	}

	for rows.Next() {
		var r row
		rows.Scan(&r.hello)
		fmt.Printf("%s\n", r.hello)
	}
}

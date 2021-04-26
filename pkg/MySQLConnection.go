package pkg

import (
	"MySQLHelper/internal"
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
)

type viaSSHDialer struct {
	client *ssh.Client
}

func (sshDialer *viaSSHDialer) Dial(ctx context.Context, addr string) (net.Conn, error) {
	return sshDialer.client.Dial("tcp", addr)
}

func MySQLConnection(useSSH bool) (db *sql.DB) {
	var dbHost string = os.Getenv("MYSQL_HOSTNAME")
	var dbPass string = os.Getenv("MYSQL_PASSWORD")
	var dbUser string = os.Getenv("MYSQL_USERNAME")
	var dbPort string = os.Getenv("MYSQL_PORT")
	var dbName string = os.Getenv("MYSQL_DATABASE")
	var sshKeyPath string
	var sshHost string
	var sshUser string
	var sshPort string
	var dialContext string = "tcp"

	if useSSH {
		sshKeyPath = os.Getenv("SSH_KEYPATH")
		sshHost = os.Getenv("SSH_HOST")
		sshUser = os.Getenv("SSH_USER")
		sshPort = os.Getenv("SSH_PORT")
		sshtun := internal.SSHClient(&sshHost, &sshUser, &sshKeyPath, &sshPort)

		dialContext = "mysql+tcp"
		mysql.RegisterDialContext(dialContext, (&viaSSHDialer{sshtun}).Dial)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true&columnsWithAlias=true", dbUser, dbPass, dialContext, dbHost, dbPort, dbName))
	if err != nil {
		panic(err)
	}

	return db
}

package pkg

import (
	"MySQLHelper/internal"
	"context"
	"database/sql"
	"fmt"
	"net"
	"errors"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
)

type viaSSHDialer struct {
	client *ssh.Client
}

func (sshDialer *viaSSHDialer) Dial(ctx context.Context, addr string) (net.Conn, error) {
	return sshDialer.client.Dial("tcp", addr)
}

type MySQLConfig struct {
	DbHost string
	DbPass string
	DbUser string
	DbPort string
	DbName string
	UseSSH bool
	SshKeyPath string
	SshHost string
	SshUser string
	SshPort string
}
  
func(d *MySQLConfig) fill_defaults(){
    if d.DbPort == "" {
        d.DbPort = "3306"
    }
      
    if d.SshPort == "" {
        d.SshPort = "22"
    }
}

func (m *MySQLConfig) Connect() (*sql.DB, error) {
	m.fill_defaults()
	var dialContext string = "tcp"

	if m.DbHost == "" || m.DbHost == "" || m.DbHost == "" {
		return nil, errors.New("Database parameters missing")
	}

	if m.UseSSH {
		if m.SshKeyPath == "" || m.SshHost == "" || m.SshUser == "" {
			return nil, errors.New("SSH parameters missing")
		}

		sshtun, err := internal.SSHClient(&m.SshHost, &m.SshUser, &m.SshKeyPath, &m.SshPort)
		if err != nil {
			return nil, err
		}

		dialContext = "mysql+tcp"
		mysql.RegisterDialContext(dialContext, (&viaSSHDialer{sshtun}).Dial)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true&columnsWithAlias=true", m.DbUser, m.DbPass, dialContext, m.DbHost, m.DbPort, m.DbName))

	return db, err
}

# Go MySQL SSH Helper
Simple MySQL connection helper with optional SSH tunneling for Golang.

## Parameters
|type|variable|type|default|
|---|---|---|---|
|sql: mandatory|DbHost|string||
|sql: mandatory|DbPass|string||
|sql: mandatory|DbUser|string||
|sql: optional|DbPort|string|3306|
|sql: optional|DbName|string||
|ssh: optional|UseSSH|bool|false|
|ssh: mandatory|SshKeyPath|string||
|ssh: mandatory|SshHost|string||
|ssh: mandatory|SshUser|string||
|ssh: optional|SshPort|string|22|

## Example
```go
package main

import (
	"github.com/zimza/go-mysql-ssh-helper/pkg"
)

func main() {
	db := pkg.MySQLConfig {
		DbHost: "mydb.hostname.com",
		DbUser: "john",
		DbPass: "1234",
		DbPort: "3306",
		UseSSH: true,
		SshKeyPath: "/home/user/.ssh/id_rsa",
		SshHost: "ssh.hostname.com",
		SshUser: "john",
		SshPort: "22",
	}

	cnx, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer cnx.Close()
}
```

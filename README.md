# MySQL / SSH helper for Golang

## Usage
|required|variable|type|default|
|---|---|---|---|
|mandatory|DbHost|string||
|mandatory|DbPass|string||
|mandatory|DbUser|string||
|optional|DbPort|string|3306|
|optional|DbName|string||
|optional|UseSSH|bool|false|
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

### Using environment variables
```go
import (
	"github.com/zimza/go-mysql-ssh-helper/pkg"
	"os"
)
...
	db := pkg.MySQLConfig {
		DbHost: os.Getenv("MYSQL_HOSTNAME"),
		DbUser: os.Getenv("MYSQL_USERNAME"),
		DbPass: os.Getenv("MYSQL_PASSWORD"),
		DbPort: os.Getenv("MYSQL_PORT"),
		DbName: os.Getenv("MYSQL_DATABASE"),
		UseSSH: true,
		SshKeyPath: os.Getenv("SSH_KEYPATH"),
		SshHost: os.Getenv("SSH_HOST"),
		SshUser: os.Getenv("SSH_USER"),
		SshPort: os.Getenv("SSH_PORT"),
	}
...
```

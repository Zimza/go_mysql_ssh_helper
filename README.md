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
	"MySQLHelper/pkg"
	"os"
)


func main() {
	db := pkg.MySQLConfig {
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

	cnx, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer cnx.Close()
}
```

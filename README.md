# MySQL / SSH helper for Golang

## Usage
|required|variable|type|default|
|---|---|---|---|
|Mandatory|DbHost|string||
|Mandatory|DbPass|string||
|Mandatory|DbUser|string||
|Optional|DbPort|string|3306|
|Mandatory|UseSSH|bool|false|
|Optional|DbName|string||
|Optional (SSH)|SshKeyPath|string||
|Optional (SSH)|SshHost|string||
|Optional (SSH)|SshUser|string||
|Optional (SSH)|SshPort|string|22|

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

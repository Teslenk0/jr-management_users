package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var (
	Client *sql.DB
)

const (
	pgUsersUserName     = "pg_users_username"
	pgUsersUserPassword = "pg_users_password"
	pgUsersUserHost     = "pg_users_host"
	pgUsersUserSchema   = "pg_users_schema"
)

var (
	username = os.Getenv(pgUsersUserName)
	password = os.Getenv(pgUsersUserPassword)
	host     = os.Getenv(pgUsersUserHost)
	schema   = os.Getenv(pgUsersUserSchema)
)

func init() {

	//Call the function that reads the toml file
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=require",
		username,
		password,
		host,
		schema)

	var err error
	Client, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}
}

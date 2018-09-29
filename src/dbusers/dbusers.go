package dbusers

import (
	"database/sql"
	"fmt"
)

// User : User Info
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Auth     string `json:"auth"`
}

const (
	dbHost     = "localhost"
	dbPort     = "5432"
	dbUser     = "root"
	dbPassword = ""
	dbName     = "postgres"
)

// SQLDB - Books DB Object / User
type SQLDB struct {
	db    *sql.DB
	table string
}

// Create Table
func (user *SQLDB) createTable() (err error) {
	que := `
	CREATE TABLE IF NOT EXISTS "` + user.table + `"
	(
		"_id" serial NOT NULL,
		"username" character varying(255) NOT NULL,
		"email" character varying(255) NOT NULL,
		"password" character varying(255) NOT NULL,
		"auth" character varying(255) NOT NULL,

		CONSTRAINT userinfo_pkey PRIMARY KEY ("_id")
	)
	-- ) WITH (OIDS=FALSE); // Not work with CockroachDB`

	result, err := user.db.Exec(que)
	if err != nil {
		fmt.Println("Table Creation Error: ", result, err)
	}

	return
}

// Select - cRud
func (userHolder *SQLDB) getBook(id int) (User, error) {
	result := User{}

	rows, err := userHolder.db.Query(`SELECT * FROM "`+userHolder.table+`" where "_id"=$1`, id)
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&result.ID, &result.Username, &result.Email, &result.Auth)
			if err != nil {
				fmt.Println("Get Book Error: ", err)
			}
		}
	}

	return result, err
}

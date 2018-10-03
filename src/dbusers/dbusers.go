package dbusers

import (
	"database/sql"
	"fmt"
	"log"
	"settings"

	_ "github.com/lib/pq"
)

// User : User Info
type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Auth     string `json:"auth,omitempty"`
}

// SQLDB - Books DB Object / User
type SQLDB struct {
	DB    *sql.DB
	Table string
}

func dbConn() (db *sql.DB) {
	dbinfo := fmt.Sprintf(
		"host='%s' port='%s' user='%s' password='%s' dbname='%s' sslmode='disable'",
		settings.DbHost, settings.DbPort, settings.DbUser, settings.DbPassword, settings.DbName)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// CreateTable : Table creation
func (userHolder *SQLDB) CreateTable() (err error) {
	userHolder.DB = dbConn()
	defer userHolder.DB.Close()

	que := `
	CREATE TABLE IF NOT EXISTS "` + userHolder.Table + `"
	(
		"_id" serial NOT NULL,
		"username" character varying(255) NOT NULL,
		"email" character varying(255) NOT NULL,
		"password" character varying(255) NOT NULL,
		"auth" character varying(255) NOT NULL,

		CONSTRAINT userinfo_pkey PRIMARY KEY ("_id")
	)
	-- ) WITH (OIDS=FALSE); // Not work with CockroachDB`

	result, err := userHolder.DB.Exec(que)
	if err != nil {
		log.Fatal("Table Creation Error: ", result, err)
	}

	return
}

// GetUser - cRud
func (userHolder *SQLDB) GetUser(id int) (result User, err error) {
	userHolder.DB = dbConn()
	defer userHolder.DB.Close()

	rows, err := userHolder.DB.Query(`SELECT "_id", "username", "email" FROM "`+userHolder.Table+`" WHERE "_id"=$1`, id)
	// rows, err := userHolder.DB.Query(`SELECT * FROM "`+userHolder.Table+`" WHERE "_id"=$1`, id)
	if err == nil {
		for rows.Next() {
			// err = rows.Scan(&result.ID, &result.Username, &result.Email, &result.Password, &result.Auth)
			err = rows.Scan(&result.ID, &result.Username, &result.Email)
			if err != nil {
				log.Fatal("Error: ", err)
			}
		}
	}

	return
}

// GetUsers - cRud
func (userHolder *SQLDB) GetUsers() (results []User, err error) {
	userHolder.DB = dbConn()
	defer userHolder.DB.Close()

	rows, err := userHolder.DB.Query(`SELECT "_id", "username", "email" FROM "` + userHolder.Table + `" ORDER BY "_id" ASC`)
	if err == nil {
		for rows.Next() {
			var result User
			err = rows.Scan(&result.ID, &result.Username, &result.Email)
			if err != nil {
				log.Fatal("Error: ", err)
			}

			results = append(results, result)
		}
	}

	return
}

// GetUserLogin - cRud
func (userHolder *SQLDB) GetUserLogin(user *User) (result bool, err error) {
	userHolder.DB = dbConn()
	defer userHolder.DB.Close()

	rows, err := userHolder.DB.Query(`SELECT * FROM "`+userHolder.Table+`" WHERE "username"=$1 AND "password"=$2`, user.Username, user.Password)
	if err == nil {
		for rows.Next() {
			fmt.Println(rows)
			result = true
			if err != nil {
				fmt.Println("Get Book Error: ", err)
			}
		}
	} else {
		fmt.Println("Get Book Error: ", err)
	}

	return
}

// InsertUser - Crud
func (userHolder *SQLDB) InsertUser(user *User) (userID int, err error) {
	userHolder.DB = dbConn()
	defer userHolder.DB.Close()

	err = userHolder.DB.QueryRow(
		`INSERT INTO "`+userHolder.Table+`"("username","email","password","auth") VALUES($1,$2,$3,$4) RETURNING _id`,
		user.Username, user.Email, user.Password, "none").Scan(&userID)
	if err != nil {
		userID = 0
		log.Fatal(err)
	}

	return
}

// UpdateUser - crUd
func (userHolder *SQLDB) UpdateUser(user *User) (userID int, err error) {
	userHolder.DB = dbConn()
	defer userHolder.DB.Close()

	_, err = userHolder.DB.Exec(
		`UPDATE "`+userHolder.Table+`" SET "username"=$1,"email"=$2, "password"=$3 WHERE "_id"=$4 RETURNING _id`,
		// user.Username, user.Email, user.Password, user.ID).Scan(&userID)
		user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		userID = 0
		log.Fatal(err)
	}

	userID = user.ID

	return
}

// DeleteUser - cruD
func (userHolder *SQLDB) DeleteUser(user *User) (int, error) {
	userHolder.DB = dbConn()
	defer userHolder.DB.Close()

	res, err := userHolder.DB.Exec(`DELETE FROM "`+userHolder.Table+`" WHERE "_id"=$1`, user.ID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}

package main

import (
	"log"
	"auth"
	"dbusers"
	"net/http"
	"settings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	users = map[int]*dbusers.User{}
	seq   = 1
)

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// CreateUser : Create an user
func CreateUser(c echo.Context) error {
	u := &dbusers.User{}

	err := c.Bind(u)
	if err != nil {
		return err
	}

	userHolder := dbusers.SQLDB{Table: settings.UserTable}
	seq, _ = userHolder.InsertUser(u)
	u.ID = seq

	return c.JSON(http.StatusCreated, u)
}

// GetUser : Get user(s) information
func GetUser(c echo.Context) error {
	u := new(dbusers.User)

	err := c.Bind(u)
	if err != nil {
		return err
	}

	userHolder := dbusers.SQLDB{Table: settings.UserTable}

	var users []dbusers.User
	if u.ID > 0 {
		user, _ := userHolder.GetUser(u.ID)

		users = append(users, user)
	} else {
		users, _ = userHolder.GetUsers()
	}

	return c.JSON(http.StatusOK, users)
}

// UpdateUser : Change user information
func UpdateUser(c echo.Context) error {
	u := new(dbusers.User)

	err := c.Bind(u)
	if err != nil {
		return err
	}

	userHolder := dbusers.SQLDB{Table: settings.UserTable}
	seq, _ = userHolder.UpdateUser(u)

	return c.JSON(http.StatusOK, u)
}

// DeleteUser : Delete an user
func DeleteUser(c echo.Context) error {
	u := new(dbusers.User)

	err := c.Bind(u)
	if err != nil {
		return err
	}

	userHolder := dbusers.SQLDB{Table: settings.UserTable}
	res, err := userHolder.DeleteUser(u)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, res)
	// return c.NoContent(http.StatusNoContent)
}

func main() {
	userHolder := dbusers.SQLDB{Table: settings.UserTable}
	userHolder.CreateTable()

	e := echo.New()

	config := middleware.JWTConfig{
		Claims:        &auth.CustomClaims{},
		SigningKey:    []byte("mySecret"),
		SigningMethod: settings.JwtSigningMethod,
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes - Manage user
	e.GET("/", accessible)

	e.POST("/users", CreateUser)
	e.GET("/users", GetUser)
	e.PUT("/users", UpdateUser)
	e.DELETE("/users", DeleteUser)

	// Routes - JWT Auth
	e.POST("/login", auth.Login)

	// Routes - Restricted
	rGets := []string{"/restricted"}
	for _, rGet := range rGets {
		r := e.Group(rGet)
		r.Use(middleware.JWTWithConfig(config))
		r.GET("", auth.Restricted)
	}

	// Start server
	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
	// e.Logger.Fatal(e.Start(":1323"))
}

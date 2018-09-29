package main

import (
	"auth"
	"dbusers"
	"net/http"

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
	u := &dbusers.User{
		ID: seq,
	}

	err := c.Bind(u)
	if err != nil {
		return err
	}

	users[u.ID] = u
	seq++

	return c.JSON(http.StatusCreated, users[u.ID])
}

// GetUser : Get an user information
func GetUser(c echo.Context) error {
	u := new(dbusers.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users[u.ID])
}

// UpdateUser : Change user information
func UpdateUser(c echo.Context) error {
	u := new(dbusers.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	users[u.ID] = u

	return c.JSON(http.StatusOK, users[u.ID])
}

// DeleteUser : Delete an user
func DeleteUser(c echo.Context) error {
	u := new(dbusers.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	delete(users, u.ID)

	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	config := middleware.JWTConfig{
		Claims:        &auth.CustomClaims{},
		SigningKey:    []byte("mySecret"),
		SigningMethod: "HS512",
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

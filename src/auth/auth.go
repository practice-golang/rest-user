package auth

import (
	"net/http"
	"strconv"
	"time"

	"dbusers"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// CustomClaims : 추가 사용자 정보
type CustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// Login : Auth
func Login(c echo.Context) error {
	u := new(dbusers.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	username := u.Username
	password := u.Password

	if username == "jon" && password == "shhh!" {
		claims := &CustomClaims{
			"Jon Snow", true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

		t, err := token.SignedString([]byte("mySecret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

// Restricted : Auth
func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*CustomClaims)
	name := claims.Name

	return c.String(http.StatusOK, "Welcome "+name+"!"+" Admin: "+strconv.FormatBool(claims.Admin))
}

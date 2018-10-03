package auth

import (
	"fmt"
	"net/http"
	"settings"
	"strconv"
	"time"

	"dbusers"

	jwt "github.com/dgrijalva/jwt-go"
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
	// username := u.Username
	// password := u.Password

	userHolder := dbusers.SQLDB{
		Table: settings.UserTable,
	}

	userExist, _ := userHolder.GetUserLogin(u)

	fmt.Print(u.Username)
	fmt.Print("/")
	fmt.Println(u.Password)

	if userExist {
		claims := &CustomClaims{
			// "Jon Snow", true,
			u.Username, true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		var jwtMethod *jwt.SigningMethodHMAC
		switch settings.JwtSigningMethod {
		case "HS256":
			jwtMethod = jwt.SigningMethodHS256
		case "HS384":
			jwtMethod = jwt.SigningMethodHS384
		case "HS512":
			jwtMethod = jwt.SigningMethodHS512
		}

		token := jwt.NewWithClaims(jwtMethod, claims)

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

	return c.String(http.StatusOK, "Welcome "+name+" ! / Admin: "+strconv.FormatBool(claims.Admin))
}

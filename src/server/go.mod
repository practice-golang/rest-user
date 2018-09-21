module server

require (
	auth v0.0.0
	dbusers v0.0.0
	github.com/labstack/echo v3.2.1+incompatible
)

replace (
	auth => ../auth
	dbusers => ../dbusers
)

module "https://github.com/practice-golang/rest-user"

require (
	auth v0.0.0
	dbusers v0.0.0
	github.com/labstack/echo v3.2.1+incompatible
	settings v0.0.0
)

replace (
	auth => ../auth
	dbusers => ../dbusers
	settings => ../settings
)

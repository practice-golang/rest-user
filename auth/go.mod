module auth

require (
	dbusers v0.0.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/labstack/echo v3.2.1+incompatible
	github.com/labstack/gommon v0.2.7 // indirect
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.2.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v0.0.0-20170224212429-dcecefd839c4 // indirect
	golang.org/x/crypto v0.0.0-20181001202700-f7f546618e97 // indirect
	golang.org/x/sys v0.0.0-20180928133829-e4b3c5e90611 // indirect
	settings v0.0.0
)

replace (
	dbusers => ../dbusers
	settings => ../settings
)

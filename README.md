# Practice for REST user manage using Echo, PostgreSQL

## Directories
* auth - Login
* dbusers - DB control
* settings - Settings for db, jwt method
* server - main package
* sql - SQL queries to test

## Files
* pgdata.zip - ```pgsql 9.6``` data folder, unzip and run ```pg_ctl.exe -D .\\pgdata -l pglog start```
* request.http - http queries to test

## Build
* If you want to manage separated packages from projects, you should use vendor. Throw out $GOPATH/pkg
```sh
cd server

go mod tidy
go mod vendor

go build -mod vendor

./rest-server(.exe)
```

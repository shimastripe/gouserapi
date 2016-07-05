# gouserapi

Simple Rest API using gin(framework) & gorm(orm)

## Start using it

```
go get github.com/shimastripe/gouserapi
go run main.go
```

### example API

api endpoint list

```
http://localhost:8080/api/users
http://localhost:8080/api/profiles
http://localhost:8080/api/account_names
http://localhost:8080/api/emails
http://localhost:8080/api/nations
```

### query

?field
- You can choose the fields that you want returned with the fields query parameter.
- ?field=id,profile_id

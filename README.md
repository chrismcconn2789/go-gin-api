# Go Gin API

---

#### Description

An Go API built using the Gin library for handling a financial transactions based backend for an expense tracker application. Also interacting with a SQL database for persistent storage, in this example a local SQLite database.

#### Instructions

- git clone the repo - https://github.com/chrismcconn2789/go-gin-api
- ensure you have Go v1.22+ installed
- use `make run` to run the application
- use an API client i.e. Postman to hit the endpoints at http://localhost:8080
  - GET `/transactions` to get all transactions
  - GET `/transactions/:id` to get a specific transaction
  - POST `/transactions` to create a new transaction
  - DELETE `/transactions/:id` to a delete a transaction
- use `make build` to compile the binary into the `/output` folder

#### List of packages used

- github.com/gin-contrib/cors v1.7.2
- github.com/gin-gonic/gin v1.10.0
- gorm.io/gorm v1.25.10

#### Resources

[Gin Library](https://gin-gonic.com/)
[Gin CORS](https://pkg.go.dev/github.com/gin-contrib/cors)
[GORM](https://gorm.io/index.html)

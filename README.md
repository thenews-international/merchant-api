# Merchant Service (Demo)

## The Twelve Factors App

- I. Codebase :white_check_mark:
    - One codebase tracked in revision control, many deploys
    - `git`
- II. Dependencies :white_check_mark:
    - Explicitly declare and isolate dependencies
    - `go mod`
- III. Config :white_check_mark:
    - Store config in the environment
    - `viper`
- IV. Backing services :white_check_mark:
    - Treat backing services as attached resources
    - `mysql`
- V. Build, release, run :white_check_mark:
    - Strictly separate build and run stages
    - `github`
- VI. Processes :white_check_mark:
    - Execute the app as one or more stateless processes
- VII. Port binding :white_check_mark:
    - Export services via port binding
    - `docker`
- VIII. Concurrency :x:
    - Scale out via the process model
- IX. Disposability :white_check_mark:
    - Maximize robustness with fast startup and graceful shutdown
- X. Dev/prod parity :white_check_mark:
    - Keep development, staging, and production as similar as possible
    - `docker`
- XI. Logs :white_check_mark:
    - Treat logs as event streams
    - `uber zap`
- XII. Admin processes :x:
    - Run admin/management tasks as one-off processes

## Rule of Clean Architecture by Uncle Bob
- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

## Project Structure
```bash
.
|-- README.md
|-- api
|   |-- handler
|   |   |-- auth.go
|   |   |-- common.go
|   |   |-- merchant.go
|   |   |-- server.go
|   |   |-- suite_test.go
|   |   `-- team_member.go
|   `-- router
|       |-- middleware
|       |   |-- content_type_json.go
|       |   |-- content_type_json_test.go
|       |   |-- jwt_authentication.go
|       |   `-- jwt_authentication_test.go
|       `-- router.go
|-- cmd
|   `-- app
|       `-- main.go
|-- config
|   `-- config.go
|-- config.yaml
|-- deployment
|   `-- docker
|       |-- Dockerfile
|       `-- bin
|           |-- init.sh
|           `-- wait-for-mysql.sh
|-- docker-compose.yml
|-- docs
|   `-- swagger
|       |-- docs.go
|       |-- swagger.json
|       `-- swagger.yaml
|-- go.mod
|-- go.sum
|-- mock
|   `-- mock_repository
|       `-- mock_db.go
|-- model
|   |-- auth_form.go
|   |-- common.go
|   |-- jwt.go
|   |-- merchant.go
|   |-- merchant_form.go
|   |-- merchant_test.go
|   |-- team_member.go
|   `-- team_member_form.go
|-- mysql
|   |-- mysql.go
|   `-- mysql_test.go
|-- repository
|   |-- db.go
|   |-- merchant.go
|   |-- merchant_test.go
|   |-- suite_test.go
|   `-- team_member.go
|-- server
|   |-- driver
|   |   `-- driver.go
|   |-- health
|   |   |-- health.go
|   |   |-- health_test.go
|   |   `-- sqlhealth
|   |       |-- sqlhealth.go
|   |       `-- sqlhealth_test.go
|   |-- requestlog
|   |   |-- ncsa.go
|   |   |-- ncsa_test.go
|   |   |-- requestlog.go
|   |   `-- requestlog_test.go
|   |-- server.go
|   `-- server_test.go
`-- util
    |-- httputil
    |   `-- httputil.go
    |-- logutil
    |   |-- logutil.go
    |   `-- logutil_test.go
    `-- validator
        |-- validator.go
        `-- validator_test.go
```

## Setting up development environment
- Install [Docker Application](https://www.docker.com/products/docker-desktop) in the development environment.

- Run `docker-compose up`.
    ```
    cd path/to/project/development/docker
    docker-compose up
    ```
- To generate swagger documents locally,
  ```
  swag init -g cmd/server/main.go -o docs/swagger
  ```
- To run test locally,
  ```
  go test ./...
  ```
- To run linters locally,
  ```
  golangci-lint run ./...
  ```
- Generating mocks after repository code change
    ```
    mockgen -source=./repository/db.go -destination=./mock/mock_repository/mock_db.go
    ```

## References
- [Docker: Get started guild](https://docs.docker.com/get-started/)
- [GORM Documentation](http://gorm.io/docs/) for ORM functionality.
- [Chi README.md](https://github.com/go-chi/chi/blob/master/README.md) for router functionality.
- [Uber Zap README.md](https://github.com/uber-go/zap/blob/master/README.md) for logging functionality.
- [Validator.v10 GoDoc](https://github.com/go-playground/validator/v10) for validation functionality.
- [Swag Declarative Comments Formats](https://github.com/swaggo/swag#declarative-comments-format) for API documentation.
- [GoMock](https://github.com/golang/mock) for Mock Server
- [Viper README.md](https://github.com/spf13/viper/blob/master/README.md) for Config
- [Golangci-lint README.md](https://github.com/golangci/golangci-lint/blob/master/README.md) for linter
# Solution of Ingrid coding assignment.
This is the solution of Ingrid backend coding task. The proposed solution has been tested on Linux operating system.

## Technology stack

- Go 1.14
- [Gin framework](https://github.com/gin-gonic/gin)
- [Gin swagger](https://github.com/swaggo/gin-swagger)
- Docker

## Additional features
Some additional features have been added to the base requirements of the task.

### Limit and order query parameters
Two additional query parameters have been added:
- `limit` - integer that describes maximum number of routes in the returned JSON array with routes,
- `order` - parameter that describes the ordering manner of the result: `asc` for ascending sorting and `desc` for descending one.

### Swagger
For sake of simpler testing of the proposed solution, a Swagger endpoint has been exposed under the following URL address:
```
http://localhost:8080/swagger/index.html
```

### Unit tests
Unit tests of sorting feature have been implemented and are available in `route/routes_test.go` file.

## Local deployment
One can build the application locally by running the following command:
```
go build -o ingrid-app
```
<sub>Disclaimer: you can choose your own executable name and replace the `ingrid-app` placeholder.</sub>

To run the application:
```
./ingrid-app
```

## Building Docker image
Blah blah blah
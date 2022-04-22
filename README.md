# go-graphql-tutorial
[![CircleCI](https://circleci.com/gh/jrgriffiniii/go-graphql-tutorial/tree/main.svg?style=svg)](https://circleci.com/gh/jrgriffiniii/go-graphql-tutorial/tree/main)

A GraphQL API tutorial implemented using Go

## Getting Started

### Installing the Go dependencies
```bash
$ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2
$ go install gotest.tools/gotestsum@latest
$ go get .
```

### Starting the API server

#### Starting the Database

```bash
$ docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=go_graphql_tutorial -d mysql:latest
$ docker logs -f mysql # Please wait for the MySQL database to start...
$ $(go env GOPATH)/bin/migrate -database mysql://root:dbpass@localhost:3306/go_graphql_tutorial -path internal/pkg/db/migrations/mysql up
```

```bash
$ go run server.go
```

One may then please visit [the GraphiQL web browser interface](http://localhost:8080/) for querying and submitting mutations to the server API.

#### Queries

One example query would be:

```graphql
query {
  todos {
    id
    text,
    done,
    user {
      id,
      name
    }
  }
}
```

#### Mutations

One example mutation would be:

```graphql
mutation CreateTodo($input: NewTodo!) {
  createTodo(input: $input) {
    id,
    text,
    user {
      id,
      name
    }
  }
}
```

...with the variables being the following:

```json
{
  "input": {
    "text": "new todo",
    "userId": "user-id"
  }
}
```

### Starting the web browser client

#### Installing the NodeJS package dependencies

```bash
$ cd client/
$ yarn install
```

#### Starting the web client

```bash
$ cd client/
$ yarn run dev
```

One may then please visit [the JavaScript web client](http://localhost:8081/).

## API Development

### Lint the source code files

```bash
$ $(go env GOPATH)/bin/golangci-lint run
```

### Executing the Test Suite

```bash
$ $(go env GOPATH)/bin/gotestsum
```

## Web Client Development

### Lint the source code files

```bash
$ yarn lint
```

### Executing the Test Suite

```bash
$ yarn test
```


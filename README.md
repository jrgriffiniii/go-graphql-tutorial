# go-graphql-tutorial
[![CircleCI](https://circleci.com/gh/jrgriffiniii/go-graphql-tutorial/tree/main.svg?style=svg)](https://circleci.com/gh/jrgriffiniii/go-graphql-tutorial/tree/main)

A GraphQL API tutorial implemented using Go

## Getting Started

```bash
$ go get .
```

## Development

### Lint the source code files

```bash
$ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2
$ $(go env GOPATH)/bin/golangci-lint
```

### Executing the Test Suite

```bash
$ go install gotest.tools/gotestsum@latest
$ $(go env GOPATH)/bin/gotestsum
```

### Starting the server

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


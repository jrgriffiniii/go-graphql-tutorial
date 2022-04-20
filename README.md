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

Please visit the [GraphiQL](http://localhost:8080/) interface for querying and submitting mutations to the server API.

#### Queries

One example query would be:

```bash
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

```bash
mutation CreateTodo($input: NewTodo!) {
  createTodo(input: $input) {
    id,
    text,
    user{
      id,
      name
    }
  }
}
```

...with the variables being the following:

```bash
{
  "input": {
    "text": "new todo",
    "userId": "user-id"
  }
}
```


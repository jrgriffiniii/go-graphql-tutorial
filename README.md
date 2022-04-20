# go-graphql-tutorial
A GraphQL API tutorial implemented using Go

## Getting Started

```bash
$ go get .
```

## Development

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



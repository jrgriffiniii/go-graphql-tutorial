package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/jrgriffiniii/go-graphql-tutorial/graph"
	"github.com/jrgriffiniii/go-graphql-tutorial/graph/generated"
	database "github.com/jrgriffiniii/go-graphql-tutorial/internal/pkg/db/mysql"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.InitDB()
	database.Migrate()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/kalesecarpenter/quotes-starter/gqlgen/graph"
	"github.com/kalesecarpenter/quotes-starter/gqlgen/graph/generated"
)

const defaultPort = "8080"

func main() {
	router := chi.NewRouter()

	router.Use(Middleware())

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			xApiKey := r.Header.Get("x-api-key")
			// // put it in context
			ctx := context.WithValue(r.Context(), "x-api-key", xApiKey)
			// // and call the next with our new context
			r = r.WithContext(ctx)
			fmt.Println("this is an x-api-key header: ", r.Header.Get("x-api-key"))
			next.ServeHTTP(w, r)
		})
	}
}

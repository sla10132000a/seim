package main

import (
	"log"
	"net/http"
	"os"
	"seim/graph"
	"seim/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)
import Mlclient "seim/ml/client"


const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	//ML(python)
	mlClient, err := Mlclient.NewClient("localhost:50052")
	if err != nil {
		mlClient.Close()
		log.Fatalf("Failed to create article client: %v\n", err)
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{ MlClient:mlClient}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials: true,
	})

	//http.Handle("/query", srv)
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

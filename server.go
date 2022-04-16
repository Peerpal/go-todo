package main

import (
	"log"
	"os"
	"todo/graph"
	"todo/graph/generated"
	"todo/prisma"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	startGinServer(port)
}

func startGinServer(port string) {
	server := gin.Default()

	server.POST("/graphql", func(ctx *gin.Context) {
		prismaClient := prisma.NewClient()

		prismaClient.Connect()

		resolver := graph.NewResolver(prismaClient)

		h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

		h.ServeHTTP(ctx.Writer, ctx.Request)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	server.Run(port)
}

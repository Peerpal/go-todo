package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"todo/graph"
	"todo/graph/generated"
	"todo/prisma"
	"todo/utils"

	"github.com/99designs/gqlgen/graphql"
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

	server.POST("/graphql", func(ctext *gin.Context) {
		prismaClient := prisma.NewClient()

		prismaClient.Connect()

		resolver := graph.NewResolver(prismaClient)

		schemaConfig := generated.Config{Resolvers: resolver}

		schemaConfig.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
			// check context header for token
			token := ctext.Request.Header.Get("Authorization")

			ok, err := parseToken(token, prismaClient, ctx)
			if err != nil {
				return nil, err
			} else if !ok {
				return nil, fmt.Errorf("Unauthorized")
			}

			return next(ctx)
		}

		h := handler.NewDefaultServer(generated.NewExecutableSchema(schemaConfig))

		h.ServeHTTP(ctext.Writer, ctext.Request)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	server.Run(port)
}

func parseToken(token string, prismaClient *prisma.PrismaClient, ctx context.Context) (bool, error) {
	JWT := &utils.JWT{
		Secret: "jwtsigingkeyhere",
		Client: prismaClient,
	}

	ok, err := JWT.ParseToken(token, ctx)

	if err != nil {
		log.Fatalf("%s", err)
		return false, err
	}
	return ok, nil
}

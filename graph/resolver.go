package graph

import "todo/prisma"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *prisma.PrismaClient
}

func NewResolver(client *prisma.PrismaClient) *Resolver {
	return &Resolver{client: client}
}

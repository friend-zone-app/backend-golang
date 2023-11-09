package main

//go:generate go run hooks/bson.go

import (
	"parties-app/backend/directives"
	"parties-app/backend/graph/generated"
	graph "parties-app/backend/graph/resolvers"
	"parties-app/backend/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func graphqlHandler() gin.HandlerFunc {
	config := generated.Config{Resolvers: &graph.Resolver{}}

	config.Directives.IsAuthenticated = directives.IsAuthenticated

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	r.Use(middleware.CorsMiddleware())

	r.Run()
}

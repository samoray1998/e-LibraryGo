package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// // Define Mutation for adding new Book data

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic: %v", r)
		}
	}()
	var schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    xRootQuery,
			Mutation: mMutation,
		},
	)
	if err != nil {
		panic(err)
	}
	// GraphQL query to get all books
	query := `
	 {
		 books {
			 id
			 title
			 
		 }
	 }`

	// Execute the query
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("GraphQL query failed: %v", result.Errors)
		return
	}
	fmt.Printf("%+v\n", result.Data)

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	// http.Handle("/graphiql", handler.GraphiQL{})
	// http.Handle("/graphiql", handler.New(&handler.Config{
	// 	Schema:   &schema,
	// 	Pretty:   true,
	// 	GraphiQL: true,
	// }))
	fmt.Println("Now server is running on port 8040")
	// log.Fatal(http.ListenAndServe(":8040", nil))
	log.Fatal(http.ListenAndServe(":8040", nil))
}

// to add mutation and try to connecte this things with data base and connect with flutter aftre word

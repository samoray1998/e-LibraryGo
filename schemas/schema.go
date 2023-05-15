package schemas

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func SetUpGraph() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic: %v", r)
		}
	}()
	var schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:        xRootQuery,
			Mutation:     mMutation,
			Subscription: rootSubscription,
		},
	)
	if err != nil {
		panic(err)
	}
	//GraphQL query to get all books
	query := `
	 {
		 users {
			 id
			 

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
}

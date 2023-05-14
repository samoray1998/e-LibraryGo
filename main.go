package main

import (
	"fmt"
	"log"
	"net/http"

	s "testGoGraph/schemas"
)

// // Define Mutation for adding new Book data

func main() {

	s.SetUpGraph()
	// http.Handle("/graphiql", handler.GraphiQL{})
	// http.Handle("/graphiql", handler.New(&handler.Config{
	// 	Schema:   &schema,
	// 	Pretty:   true,
	// 	GraphiQL: true,
	// }))
	fmt.Println("Now server is running on port 8041")
	// log.Fatal(http.ListenAndServe(":8040", nil))
	log.Fatal(http.ListenAndServe(":8041", nil))
	
}

// to add mutation and try to connecte this things with data base and connect with flutter aftre word

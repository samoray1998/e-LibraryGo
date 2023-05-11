package schemas

import (
	"testGoGraph/models"

	"github.com/graphql-go/graphql"
)

var roleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var roles = []models.Role{
	{ID: 1, Name: "user"},
	{ID: 2, Name: "admin"},
}

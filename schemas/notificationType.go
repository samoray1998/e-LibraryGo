package schemas

import (
	"testGoGraph/models"
	"time"

	"github.com/graphql-go/graphql"
)

var metaDataType = graphql.NewObject(graphql.ObjectConfig{
	Name: "MetaData",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source, nil
			},
		},
	},
})

var notificationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Notification",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user": &graphql.Field{
			Type: userType,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"metaData": &graphql.Field{
			Type: graphql.NewList(graphql.NewNonNull(metaDataType)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				data, ok := p.Source.(models.Notification)
				if !ok {
					return nil, nil
				}
				var metaData []interface{}
				for k, v := range data.Metadata {
					metaData = append(metaData, map[string]interface{}{
						"name":  k,
						"value": v,
					})
				}
				return metaData, nil
			},
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var notifications = []models.Notification{{
	ID:        1,
	User:      &users[0],
	Message:   "Your order has shipped!",
	Metadata:  map[string]interface{}{"order_id": 12345, "userId": "1"},
	CreatedAt: time.Now(),
},

	{
		ID:        2,
		User:      &users[0],
		Message:   "You have a new message from Alice.",
		Metadata:  map[string]interface{}{"message_id": 67890},
		CreatedAt: time.Now(),
	}}

func getNotifications(params graphql.ResolveParams) (interface{}, error) {
	return notifications, nil
}

package schemas

import (
	"fmt"
	"testGoGraph/models"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

var MyNotifCha = make(chan *models.Notification)

var rootSubscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"NewNotif": &graphql.Field{
			Type: notificationType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				fmt.Println("this is a test 1", id)
				if !ok {
					return nil, gqlerrors.NewFormattedError("Invalid ID")
				}
				fmt.Println("this is a test 2", id)
				notifChan := make(chan *models.Notification)
				go func() {
					fmt.Println("this is a test 3", id)
					for {
						notif := <-MyNotifCha
						fmt.Println("this is a test ", notif)
						// if notif.User.ID == id {
						notifChan <- notif
						fmt.Println("this is a test gjd djf  ", notif)
						//p.Info.(*graphql.ResolveInfo).RootValue.(chan<-interface{})<-notif
						//p.Info.(*graphql.ResolveInfo).RootValue.(chan<- interface{}) <- notif

						// }
					}
				}()
				return <-notifChan, nil
			},
		},
	}})

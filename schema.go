package gateway

import (
	"log"

	"github.com/graphql-go/graphql"
)

// ConfigSchema returns the root level GraphQL schema instance
func ConfigSchema(q *QueryService, m *MutationService) (graphql.Schema, error) {
	log.Println("Setting up GraphQL schema...")
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"group":         queryOneGroupField(q),
				"groups":        queryGroupsField(q),
				"user":          queryOneUserField(q),
				"users":         queryUsersField(q),
				"events":        queryEventsField(q),
				"announcements": queryAnnouncementsField(q),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				//"login":              mutateLoginField(m),
				"createAnnouncement": mutateCreateAnnouncementField(m),
				"updateAnnouncement": mutateUpdateAnnouncementField(m),
				"createEvent":        mutateCreateEventField(m),
				"updateEvent":        mutateUpdateEventField(m),
				"createGroup":        mutateCreateGroupField(m),
				"updateGroup":        mutateUpdateGroupField(m),
				//"createUser": mutateCreateUserField(m),
				//"updateUser": mutateUpdateUserField(m),
			},
		}),
		/*Subscription: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Subscription",
			Fields: Subscriptions,
		}),*/
	})
}

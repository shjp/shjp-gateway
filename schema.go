package gateway

import (
	"log"

	"github.com/graphql-go/graphql"
)

// ConfigSchema returns the root level GraphQL schema instance
func ConfigSchema(q *QueryService, m *MutationService, a *AuthService) (graphql.Schema, error) {
	log.Println("Setting up GraphQL schema...")
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"announcement":  queryOneAnnouncementField(q),
				"announcements": queryAnnouncementsField(q),
				"event":         queryOneEventField(q),
				"events":        queryEventsField(q),
				"group":         queryOneGroupField(q),
				"groups":        queryGroupsField(q),
				"user":          queryOneUserField(q),
				"users":         queryUsersField(q),
				"me":            queryMeField(a),
				"massFiles":     queryMassFiles(q),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createAnnouncement":      mutateCreateAnnouncementField(m),
				"updateAnnouncement":      mutateUpdateAnnouncementField(m),
				"createEvent":             mutateCreateEventField(m),
				"updateEvent":             mutateUpdateEventField(m),
				"createGroup":             mutateCreateGroupField(m),
				"updateGroup":             mutateUpdateGroupField(m),
				"createUser":              mutateCreateUserField(m),
				"updateUser":              mutateUpdateUserField(m),
				"createRole":              mutateCreateRoleField(m),
				"login":                   mutateLoginField(a),
				"requestGroupJoin":        mutateRequestGroupJoin(a, m),
				"updateRsvp":              mutateUpdateRSVPField(a, m),
				"changeGroupMemberStatus": mutateUpdateGroupMemberStatus(a, m),
			},
		}),
		/*Subscription: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Subscription",
			Fields: Subscriptions,
		}),*/
	})
}

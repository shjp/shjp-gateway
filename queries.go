package gateway

import (
	"github.com/graphql-go/graphql"
)

func queryOneGroupField(s *QueryService) *graphql.Field {
	return &graphql.Field{
		Type: GroupType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: queryOneResolver(s, "groups"),
	}
}

func queryGroupsField(s *QueryService) *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(GroupType),
		Resolve: queryAllResolver(s, "groups"),
	}
}

func queryOneUserField(s *QueryService) *graphql.Field {
	return &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: queryOneResolver(s, "users"),
	}
}

func queryUsersField(s *QueryService) *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(UserType),
		Resolve: queryAllResolver(s, "users"),
	}
}

func queryEventsField(s *QueryService) *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(EventType),
		Resolve: queryAllResolver(s, "events"),
	}
}

func queryAnnouncementsField(s *QueryService) *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(AnnouncementType),
		Resolve: queryAllResolver(s, "announcements"),
	}
}

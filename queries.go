package gateway

import (
	"github.com/graphql-go/graphql"
)

func queryOneAnnouncementField(s *QueryService) *graphql.Field {
	return &graphql.Field{
		Type: AnnouncementType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: queryOneResolver(s, "announcements"),
	}
}

func queryAnnouncementsField(s *QueryService) *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(AnnouncementType),
		Resolve: queryAllResolver(s, "announcements"),
	}
}

func queryOneEventField(s *QueryService) *graphql.Field {
	return &graphql.Field{
		Type: EventType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: queryOneResolver(s, "events"),
	}
}

func queryEventsField(s *QueryService) *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(EventType),
		Resolve: queryAllResolver(s, "events"),
	}
}

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

func queryMeField(a *AuthService) *graphql.Field {
	return &graphql.Field{
		Type:    UserType,
		Resolve: me(a),
	}
}

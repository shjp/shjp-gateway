package gateway

import (
	"github.com/graphql-go/graphql"
)

func mutateLoginField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: UserSessionType,
		Args: graphql.FieldConfigArgument{
			"accountId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"clientId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"accountType": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"accountSecret": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"profileImage": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"nickname": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: login,
	}
}

func mutateCreateAnnouncementField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type:    MutationResponseType,
		Args:    transformTypeFieldsToArgument(*AnnouncementType, "name", "authorId", "content"),
		Resolve: createModelResolver(s, "announcement"),
	}
}

func mutateUpdateAnnouncementField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type:    MutationResponseType,
		Args:    transformTypeFieldsToArgument(*AnnouncementType, "id", "name", "authorId", "content"),
		Resolve: updateModelResolver(s, "announcement"),
	}
}

func mutateCreateEventField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type:    MutationResponseType,
		Args:    transformTypeFieldsToArgument(*EventType, "name", "date", "length", "authorId", "deadline", "allow_maybe", "description", "location", "location_description"),
		Resolve: createModelResolver(s, "event"),
	}
}

func mutateUpdateEventField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type:    MutationResponseType,
		Args:    transformTypeFieldsToArgument(*EventType, "id", "name", "date", "length", "authorId", "deadline", "allow_maybe", "description", "location", "location_description"),
		Resolve: updateModelResolver(s, "event"),
	}
}

func mutateCreateGroupField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type:    MutationResponseType,
		Args:    transformTypeFieldsToArgument(*GroupType, "name", "description", "image_url"),
		Resolve: createModelResolver(s, "group"),
	}
}

func mutateUpdateGroupField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type:    MutationResponseType,
		Args:    transformTypeFieldsToArgument(*GroupType, "id", "name", "description", "image_url"),
		Resolve: updateModelResolver(s, "group"),
	}
}

func mutateCreateUserField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type:    MutationResponseType,
		Args:    transformTypeFieldsToArgument(*UserType, "name", "baptismalName", "birthday", "feastDay", "googleId", "facebookId"),
		Resolve: updateModelResolver(s, "user"),
	}
}

func mutateUpdateUserField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*UserType, "id", "name", "baptismalName", "birthday", "feastDay", "googleId", "facebookId"),
	}
}

package gateway

import (
	"github.com/graphql-go/graphql"
)

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
		Args:    transformTypeFieldsToArgument(*UserType, "name", "accountType", "email", "password", "baptismalName", "birthday", "feastday"),
		Resolve: createModelResolver(s, "user"),
	}
}

func mutateUpdateUserField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type:    MutationResponseType,
		Args:    transformTypeFieldsToArgument(*UserType, "id", "name", "email", "password", "baptismalName", "birthday", "feastday"),
		Resolve: updateModelResolver(s, "user"),
	}
}

func mutateLoginField(a *AuthService) *graphql.Field {
	return &graphql.Field{
		Type: UserSessionType,
		Args: graphql.FieldConfigArgument{
			"accountType": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: login(a),
	}
}

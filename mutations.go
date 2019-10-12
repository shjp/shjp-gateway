package gateway

import (
	"github.com/graphql-go/graphql"
)

func mutateCreateAnnouncementField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*AnnouncementType,
			"name",
			"authorId",
			"content"),
		Resolve: createModelResolver(s, "announcement"),
	}
}

func mutateUpdateAnnouncementField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*AnnouncementType,
			"id",
			"name",
			"authorId",
			"content"),
		Resolve: updateModelResolver(s, "announcement"),
	}
}

func mutateCreateEventField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*EventType,
			"name",
			"start",
			"end",
			"authorId",
			"deadline",
			"allow_maybe",
			"description",
			"location",
			"location_description"),
		Resolve: createModelResolver(s, "event"),
	}
}

func mutateUpdateEventField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*EventType,
			"id",
			"name",
			"start",
			"end",
			"authorId",
			"deadline",
			"allow_maybe",
			"description",
			"location",
			"location_description"),
		Resolve: updateModelResolver(s, "event"),
	}
}

func mutateCreateGroupField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*GroupType,
			"name",
			"description",
			"image_url"),
		Resolve: createModelResolver(s, "group"),
	}
}

func mutateUpdateGroupField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*GroupType,
			"id",
			"name",
			"description",
			"image_url"),
		Resolve: updateModelResolver(s, "group"),
	}
}

func mutateCreateRoleField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformInputTypeFieldsToArgument(*RoleInputType,
			"name",
			"group_id",
			"can_read",
			"can_read_members",
			"can_read_comments",
			"can_write_comments",
			"can_write_announcements",
			"can_write_events",
			"can_admin_group",
			"can_edit_users"),
		Resolve: createModelResolver(s, "role"),
	}
}

func mutateCreateUserField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*UserType,
			"name",
			"accountType",
			"email",
			"password",
			"baptismal_name",
			"birthday",
			"feastday"),
		Resolve: createModelResolver(s, "user"),
	}
}

func mutateUpdateUserField(s *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*UserType,
			"id",
			"name",
			"email",
			"password",
			"baptismal_name",
			"birthday",
			"feastday"),
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

func mutateRequestGroupJoin(as *AuthService, ms *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*GroupMembership,
			"group_id",
			"user_id",
			"role_id",
			"status"),
		Resolve: upsertRelationshipResolver(as, ms, "group_membership"),
	}
}

func mutateUpdateRSVPField(as *AuthService, ms *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*EventRSVPType,
			"user_id",
			"event_id",
			"rsvp"),
		Resolve: upsertRelationshipResolver(as, ms, "update_rsvp"),
	}
}

func mutateUpdateGroupMemberStatus(as *AuthService, ms *MutationService) *graphql.Field {
	return &graphql.Field{
		Type: MutationResponseType,
		Args: transformTypeFieldsToArgument(*GroupMembership,
			"group_id",
			"user_id",
			"role_id",
			"status"),
		Resolve: upsertRelationshipResolver(as, ms, "group_membership"),
	}
}

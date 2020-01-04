package gateway

import (
	"github.com/graphql-go/graphql"
)

// AnnouncementType defines the GraphQL announcement type
var AnnouncementType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Announcement",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.ID},
		"name":     &graphql.Field{Type: graphql.String},
		"authorId": &graphql.Field{Type: graphql.String},
		"creator":  &graphql.Field{Type: UserType},
		"content":  &graphql.Field{Type: graphql.String},
		"created":  &graphql.Field{Type: graphql.String},
	},
})

// EventType defines the GraphQL event type
var EventType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Event",
	Fields: graphql.Fields{
		"id":                   &graphql.Field{Type: graphql.ID},
		"name":                 &graphql.Field{Type: graphql.String},
		"start":                &graphql.Field{Type: graphql.String},
		"end":                  &graphql.Field{Type: graphql.String},
		"authorId":             &graphql.Field{Type: graphql.String},
		"author":               &graphql.Field{Type: UserType},
		"deadline":             &graphql.Field{Type: graphql.String},
		"allow_maybe":          &graphql.Field{Type: graphql.Boolean},
		"description":          &graphql.Field{Type: graphql.String},
		"location":             &graphql.Field{Type: graphql.String},
		"location_description": &graphql.Field{Type: graphql.String},
		"created":              &graphql.Field{Type: graphql.String},
		"rsvps":                &graphql.Field{Type: graphql.NewList(EventRSVPUserType)},
	},
})

// EventRSVPType defines the GraphQL event RSVP type
var EventRSVPType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EventRSVP",
	Fields: graphql.Fields{
		"user_id":  &graphql.Field{Type: graphql.ID},
		"event_id": &graphql.Field{Type: graphql.ID},
		"rsvp":     &graphql.Field{Type: graphql.String},
	},
})

// EventRSVPUserType defines the GraphQL event RSVP user type
var EventRSVPUserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EventRSVPUser",
	Fields: graphql.Fields{
		"user": &graphql.Field{Type: UserType},
		"rsvp": &graphql.Field{Type: graphql.String},
	},
})

// GroupType defines the GraphQL group type
var GroupType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Group",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"name":        &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"image_url":   &graphql.Field{Type: graphql.String},
		"members":     &graphql.Field{Type: graphql.NewList(UserType)},
		"roles":       &graphql.Field{Type: graphql.NewList(RoleType)},
	},
})

// GroupMembership defines the GraphQL type for group-user linking
var GroupMembership = graphql.NewObject(graphql.ObjectConfig{
	Name: "GroupMembership",
	Fields: graphql.Fields{
		"group_id": &graphql.Field{Type: graphql.ID},
		"user_id":  &graphql.Field{Type: graphql.ID},
		"role_id":  &graphql.Field{Type: graphql.ID},
		"status":   &graphql.Field{Type: graphql.String},
	},
})

// GroupPermission defines the GraphQL type for user's permissions within a group
var GroupPermission = graphql.NewObject(graphql.ObjectConfig{
	Name: "GroupPermission",
	Fields: graphql.Fields{
		"can_read":                &graphql.Field{Type: graphql.Boolean},
		"can_read_members":        &graphql.Field{Type: graphql.Boolean},
		"can_read_comments":       &graphql.Field{Type: graphql.Boolean},
		"can_write_comments":      &graphql.Field{Type: graphql.Boolean},
		"can_write_announcements": &graphql.Field{Type: graphql.Boolean},
		"can_write_events":        &graphql.Field{Type: graphql.Boolean},
		"can_admin_group":         &graphql.Field{Type: graphql.Boolean},
		"can_edit_users":          &graphql.Field{Type: graphql.Boolean},
	},
})

// UserGroup defines the GraphQL type for user's groups
var UserGroup = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserGroup",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"name":        &graphql.Field{Type: graphql.String},
		"privilege":   &graphql.Field{Type: graphql.Int},
		"role_name":   &graphql.Field{Type: graphql.String},
		"status":      &graphql.Field{Type: graphql.String},
		"permissions": &graphql.Field{Type: GroupPermission},
	},
})

// UserType defines the GraphQL user type
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":             &graphql.Field{Type: graphql.ID},
		"name":           &graphql.Field{Type: graphql.String},
		"accountType":    &graphql.Field{Type: graphql.String},
		"email":          &graphql.Field{Type: graphql.String},
		"password":       &graphql.Field{Type: graphql.String},
		"baptismal_name": &graphql.Field{Type: graphql.String},
		"birthday":       &graphql.Field{Type: graphql.String},
		"feastday":       &graphql.Field{Type: graphql.String},
		"groups":         &graphql.Field{Type: graphql.NewList(UserGroup)},
		"created":        &graphql.Field{Type: graphql.DateTime},
		"lastActive":     &graphql.Field{Type: graphql.DateTime},
		"status":         &graphql.Field{Type: graphql.String},
		"role_name":      &graphql.Field{Type: graphql.String},
		"privilege":      &graphql.Field{Type: graphql.Int},
	},
})

// MeType defines the GraphQL type for the current user
var MeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Me",
})

// UserSessionType defines the GraphQL user session type
var UserSessionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserSession",
	Fields: graphql.Fields{
		"key":  &graphql.Field{Type: graphql.String},
		"user": &graphql.Field{Type: UserType},
	},
})

// RoleType defines the GraphQL role type
var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.String},
		"group_id":    &graphql.Field{Type: graphql.String},
		"name":        &graphql.Field{Type: graphql.String},
		"privilege":   &graphql.Field{Type: graphql.Int},
		"permissions": &graphql.Field{Type: GroupPermission},
	},
})

// MassFileType defines the GraphQL mass file type
var MassFileType = graphql.NewObject(graphql.ObjectConfig{
	Name: "MassFile",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"type": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"date": &graphql.Field{Type: graphql.String},
		"url":  &graphql.Field{Type: graphql.String},
	},
})

// MutationResponseType defines the response structures that the mutations use
var MutationResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "MutationResponse",
	Fields: graphql.Fields{
		"ref_id": &graphql.Field{Type: graphql.String},
	},
})

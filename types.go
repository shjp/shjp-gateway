package gateway

import (
	"github.com/graphql-go/graphql"
)

// AnnouncementType defines the GraphQL announcement type
var AnnouncementType = graphql.NewObject(graphql.ObjectConfig{
	Name: "announcement",
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
	Name: "event",
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
	Name: "eventRSVP",
	Fields: graphql.Fields{
		"user_id":  &graphql.Field{Type: graphql.ID},
		"event_id": &graphql.Field{Type: graphql.ID},
		"rsvp":     &graphql.Field{Type: graphql.String},
	},
})

// EventRSVPUserType defines the GraphQL event RSVP user type
var EventRSVPUserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "eventRSVPUser",
	Fields: graphql.Fields{
		"user": &graphql.Field{Type: UserType},
		"rsvp": &graphql.Field{Type: graphql.String},
	},
})

// GroupType defines the GraphQL group type
var GroupType = graphql.NewObject(graphql.ObjectConfig{
	Name: "group",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"name":        &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"image_url":   &graphql.Field{Type: graphql.String},
		"members":     &graphql.Field{Type: graphql.NewList(UserType)},
	},
})

// GroupMembership defines the GraphQL type for group-user linking
var GroupMembership = graphql.NewObject(graphql.ObjectConfig{
	Name: "groupMembership",
	Fields: graphql.Fields{
		"group_id": &graphql.Field{Type: graphql.ID},
		"user_id":  &graphql.Field{Type: graphql.ID},
		"role_id":  &graphql.Field{Type: graphql.ID},
		"status":   &graphql.Field{Type: graphql.String},
	},
})

// GroupPermission defines the GraphQL type for user's permissions within a group
var GroupPermission = graphql.NewObject(graphql.ObjectConfig{
	Name: "groupPermission",
	Fields: graphql.Fields{
		"can_read":                &graphql.Field{Type: graphql.Boolean},
		"can_read_members":        &graphql.Field{Type: graphql.Boolean},
		"can_read_comments":       &graphql.Field{Type: graphql.Boolean},
		"can_write_comments":      &graphql.Field{Type: graphql.Boolean},
		"can_write_announcements": &graphql.Field{Type: graphql.Boolean},
		"can_write_events":        &graphql.Field{Type: graphql.Boolean},
		"can_edit_users":          &graphql.Field{Type: graphql.Boolean},
	},
})

// UserGroup defines the GraphQL type for user's groups
var UserGroup = graphql.NewObject(graphql.ObjectConfig{
	Name: "userGroup",
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
	Name: "user",
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
	Name: "me",
})

// UserSessionType defines the GraphQL user session type
var UserSessionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "userSession",
	Fields: graphql.Fields{
		"key":  &graphql.Field{Type: graphql.String},
		"user": &graphql.Field{Type: UserType},
	},
})

// MutationResponseType defines the response structures that the mutations use
var MutationResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutationResponse",
	Fields: graphql.Fields{
		"ref_id": &graphql.Field{Type: graphql.String},
	},
})

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
	},
})

// EventType defines the GraphQL event type
var EventType = graphql.NewObject(graphql.ObjectConfig{
	Name: "event",
	Fields: graphql.Fields{
		"id":                   &graphql.Field{Type: graphql.ID},
		"name":                 &graphql.Field{Type: graphql.String},
		"date":                 &graphql.Field{Type: graphql.DateTime},
		"length":               &graphql.Field{Type: graphql.Int},
		"authorId":             &graphql.Field{Type: graphql.String},
		"author":               &graphql.Field{Type: UserType},
		"deadline":             &graphql.Field{Type: graphql.DateTime},
		"allow_maybe":          &graphql.Field{Type: graphql.Boolean},
		"description":          &graphql.Field{Type: graphql.String},
		"location":             &graphql.Field{Type: graphql.String},
		"location_description": &graphql.Field{Type: graphql.String},
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

// UserType defines the GraphQL user type
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "user",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.ID},
		"name":          &graphql.Field{Type: graphql.String},
		"accountType":   &graphql.Field{Type: graphql.String},
		"email":         &graphql.Field{Type: graphql.String},
		"password":      &graphql.Field{Type: graphql.String},
		"baptismalName": &graphql.Field{Type: graphql.String},
		"birthday":      &graphql.Field{Type: graphql.DateTime},
		"feastday":      &graphql.Field{Type: graphql.DateTime},
		"groups":        &graphql.Field{Type: graphql.NewList(graphql.String)},
		"created":       &graphql.Field{Type: graphql.DateTime},
		"lastActive":    &graphql.Field{Type: graphql.DateTime},
		"roleName":      &graphql.Field{Type: graphql.String},
		"privilege":     &graphql.Field{Type: graphql.Int},
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

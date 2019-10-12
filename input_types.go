package gateway

import "github.com/graphql-go/graphql"

// GroupPermissionInputType defines the GraphQL input type for group permissions
var GroupPermissionInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "GroupPermissionInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"can_read":                &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_read_members":        &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_read_comments":       &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_write_comments":      &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_write_announcements": &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_write_events":        &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_admin_group":         &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_edit_users":          &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
	},
})

// RoleInputType defines the GraphQL input type for role
var RoleInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "RoleInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"group_id":                &graphql.InputObjectFieldConfig{Type: graphql.String},
		"name":                    &graphql.InputObjectFieldConfig{Type: graphql.String},
		"can_read":                &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_read_members":        &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_read_comments":       &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_write_comments":      &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_write_announcements": &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_write_events":        &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_admin_group":         &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"can_edit_users":          &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
	},
})

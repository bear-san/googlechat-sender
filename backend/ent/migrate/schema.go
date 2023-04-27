// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GoogleAPIKeysColumns holds the columns for the "google_api_keys" table.
	GoogleAPIKeysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "access_token", Type: field.TypeString},
		{Name: "refresh_token", Type: field.TypeString},
	}
	// GoogleAPIKeysTable holds the schema information for the "google_api_keys" table.
	GoogleAPIKeysTable = &schema.Table{
		Name:       "google_api_keys",
		Columns:    GoogleAPIKeysColumns,
		PrimaryKey: []*schema.Column{GoogleAPIKeysColumns[0]},
	}
	// SystemUsersColumns holds the columns for the "system_users" table.
	SystemUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// SystemUsersTable holds the schema information for the "system_users" table.
	SystemUsersTable = &schema.Table{
		Name:       "system_users",
		Columns:    SystemUsersColumns,
		PrimaryKey: []*schema.Column{SystemUsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GoogleAPIKeysTable,
		SystemUsersTable,
	}
)

func init() {
}

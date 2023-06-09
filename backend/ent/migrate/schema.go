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
		{Name: "expiration_date", Type: field.TypeTime},
	}
	// GoogleAPIKeysTable holds the schema information for the "google_api_keys" table.
	GoogleAPIKeysTable = &schema.Table{
		Name:       "google_api_keys",
		Columns:    GoogleAPIKeysColumns,
		PrimaryKey: []*schema.Column{GoogleAPIKeysColumns[0]},
	}
	// PostSchedulesColumns holds the columns for the "post_schedules" table.
	PostSchedulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "uid", Type: field.TypeString},
		{Name: "target", Type: field.TypeString},
		{Name: "display_name", Type: field.TypeString},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "is_sent", Type: field.TypeBool},
		{Name: "send_at", Type: field.TypeTime},
	}
	// PostSchedulesTable holds the schema information for the "post_schedules" table.
	PostSchedulesTable = &schema.Table{
		Name:       "post_schedules",
		Columns:    PostSchedulesColumns,
		PrimaryKey: []*schema.Column{PostSchedulesColumns[0]},
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
		PostSchedulesTable,
		SystemUsersTable,
	}
)

func init() {
}

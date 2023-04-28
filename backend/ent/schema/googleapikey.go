package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GoogleApiKey holds the schema definition for the GoogleApiKey entity.
type GoogleApiKey struct {
	ent.Schema
}

// Fields of the GoogleApiKey.
func (GoogleApiKey) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("access_token"),
		field.String("refresh_token"),
		field.Time("expiration_date"),
	}
}

// Edges of the GoogleApiKey.
func (GoogleApiKey) Edges() []ent.Edge {
	return nil
}

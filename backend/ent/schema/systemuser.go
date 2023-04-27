package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SystemUser holds the schema definition for the SystemUser entity.
type SystemUser struct {
	ent.Schema
}

// Fields of the SystemUser.
func (SystemUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("email"),
		field.String("name"),
	}
}

// Edges of the SystemUser.
func (SystemUser) Edges() []ent.Edge {
	return nil
}

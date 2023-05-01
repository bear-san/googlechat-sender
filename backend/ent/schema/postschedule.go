package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PostSchedule holds the schema definition for the PostSchedule entity.
type PostSchedule struct {
	ent.Schema
}

// Fields of the PostSchedule.
func (PostSchedule) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("uid"),
		field.String("target"),
		field.String("displayName"),
		field.Text("text"),
		field.Bool("is_sent"),
		field.Time("send_at"),
	}
}

// Edges of the PostSchedule.
func (PostSchedule) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("name").Unique(),
		field.String("description").Optional(),
		field.Enum("status").Values("NEW", "DELETED", "ACTIVE"),
		field.String("created_by").Immutable(),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return nil
}

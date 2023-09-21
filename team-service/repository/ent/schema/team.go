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
		field.String("id").
			Unique().
			StructTag(`json:"teamId"`),
		field.String("name").
			Unique().
			StructTag(`json:"name"`),
		field.String("description").
			Optional().
			StructTag(`json:"description"`),
		field.Enum("status").
			Values("NEW", "DELETED", "ACTIVE").
			StructTag(`json:"status"`),
		field.String("created_by").
			Immutable().
			StructTag(`json:"createdBy"`),
		field.Time("created_at").
			Default(time.Now()).
			Immutable().
			StructTag(`json:"createdAt"`),
		field.Time("updated_at").
			Default(time.Now()).
			StructTag(`json:"updatedAt"`),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return nil
}

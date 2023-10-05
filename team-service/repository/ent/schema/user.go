package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			StructTag(`json:"userId"`),
		field.String("team_id").
			Optional().
			Nillable().
			StructTag(`json:"teamId"`),
		field.String("assigned_by").
			Optional().
			Nillable().
			StructTag(`json:"assignedBy"`),
		field.String("approved_by").
			Optional().
			Nillable().
			StructTag(`json:"approvedBy"`),
		field.Enum("status").
			Values("FREE", "PENDING", "IN_TEAM").
			StructTag(`json:"status"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("users"),
		edge.To("teams", Team.Type).
			Unique().
			Field("team_id"),
		edge.To("user", User.Type).
			From("assigned").
			Unique().
			Field("assigned_by"),
		edge.To("approve", User.Type).
			From("approved").
			Unique().
			Field("approved_by"),
	}
}

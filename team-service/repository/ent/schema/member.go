package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			StructTag(`json:"memberId"`),
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

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).
			Ref("members"),
		edge.To("teams", Team.Type).
			Unique().
			Field("team_id"),
		edge.To("member", Member.Type).
			From("assigned").
			Unique().
			Field("assigned_by"),
		edge.To("approve", Member.Type).
			From("approved").
			Unique().
			Field("approved_by"),
	}
}

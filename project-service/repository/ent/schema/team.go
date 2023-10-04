package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.String("project_id").
			Optional().
			Nillable().
			StructTag(`json:"projectId"`),
		field.Enum("status").
			Values("PENDING", "IN_PROGRESS", "DONE").
			StructTag(`json:"status"`),
		field.String("assigned_by").
			Optional().
			Nillable().
			StructTag(`json:"assignedBy"`),
		field.Time("assigned_at").
			Default(time.Now()).
			StructTag(`json:"assignedAt"`),
		field.Time("updated_at").
			Default(time.Now()).
			StructTag(`json:"updatedAt"`),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("team", Team.Type).
		//	Ref("members"),
		//edge.To("teams", Team.Type).
		//	Unique().
		//	Field("team_id"),
		//edge.To("member", Member.Type).
		//	From("assigned").
		//	Unique().
		//	Field("assigned_by"),
		edge.To("assignBy", Member.Type).
			Unique().
			Field("assigned_by"),
		edge.To("assignation", Project.Type).
			Unique().
			Field("project_id"),
	}
}

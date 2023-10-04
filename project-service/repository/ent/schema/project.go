package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			StructTag(`json:"projectId"`),
		field.String("name").
			StructTag(`json:"name"`),
		field.String("image").
			StructTag(`json:"image"`),
		field.String("description").
			Optional().
			StructTag(`json:"description"`),
		field.Enum("status").
			Values("PENDING", "ANALYSIS", "DESIGN", "DEVELOPMENT", "TESTING", "RELEASED", "DELETED", "BLOCKED").
			StructTag(`json:"status"`),
		field.String("created_by").
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

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("creation", Member.Type).
			Unique().
			Field("created_by").
			Required(),
		edge.From("assigned", Team.Type).
			Ref("assignation"),
		edge.From("attachment", Attachment.Type).
			Ref("assignation"),
	}
}

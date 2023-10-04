package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Attachment holds the schema definition for the Attachment entity.
type Attachment struct {
	ent.Schema
}

// Fields of the Attachment.
func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			StructTag(`json:"attachId"`),
		field.String("attachment").
			StructTag(`json:"attachment"`),
		field.String("project_id").
			StructTag(`json:"projectId"`),
		field.String("description").
			Optional().
			StructTag(`json:"description"`),
		field.String("added_by").
			StructTag(`json:"addedBy"`),
		field.Time("created_at").
			Default(time.Now()).
			Immutable().
			StructTag(`json:"createdAt"`),
		field.Time("updated_at").
			Default(time.Now()).
			StructTag(`json:"updatedAt"`),
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("insertion", Member.Type).
			Unique().
			Field("added_by").
			Required(),
		edge.To("assignation", Project.Type).
			Unique().
			Field("project_id").
			Required(),
	}
}

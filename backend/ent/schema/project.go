package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		// field.String("project_id"),
		field.String("name").NotEmpty(),
		field.String("description"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userOwner", User.Type).
			Ref("project"),
		edge.To("task", Task.Type).StorageKey(edge.Column("Project_id")),
	}
}

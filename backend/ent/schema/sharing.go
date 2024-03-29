package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Sharing holds the schema definition for the Sharing entity.
type Sharing struct {
	ent.Schema
}

// Fields of the Sharing.
func (Sharing) Fields() []ent.Field {
	return []ent.Field{
		// field.String("share_id"),
		field.Int("user_id"),
		field.String("shared_with_user_id"),
		field.String("task_id"),
		field.String("project_id"),
		field.String("permission_level"),
	}
}

// Edges of the Sharing.
func (Sharing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userOwner", User.Type).
			Ref("sharing"),
	}
}

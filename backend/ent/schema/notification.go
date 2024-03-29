package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		// field.String("notification_id"),
		field.String("task_id"),
		field.Int("user_id"),
		field.String("method"),
		field.Time("created_at"),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userOwner", User.Type).
			Ref("notification").
			Unique(),
	}
}

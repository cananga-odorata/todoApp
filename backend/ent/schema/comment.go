package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		// field.String("comment_id"),
		field.String("task_id").NotEmpty(),
		field.Int("user_id"),
		field.String("content").NotEmpty(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userOwner", User.Type).
			Ref("comment").
			Unique(),
		edge.From("taskOwner", Task.Type).
			Ref("comment").
			Unique(),
	}
}

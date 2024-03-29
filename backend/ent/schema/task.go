package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		// field.String("task_id"),
		field.String("title").NotEmpty(),
		field.String("description"),
		field.Time("due_date"),
		field.String("status"),
		field.String("piority"),
		field.String("tags"),
		field.String("project_id"),
		field.String("category_id"),
		field.String("comments"),
		field.Time("created_at").Default(time.Now),
		field.Time("update_at"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userOwner", User.Type).
			Ref("task").
			Unique(),
		edge.From("projectOwner", Project.Type).
			Ref("task").
			Unique(),
		edge.To("comment", Comment.Type).StorageKey(edge.Column("Task_id")),
		edge.From("categoryOwner", Category.Type).
			Ref("task").
			Unique(),
	}
}

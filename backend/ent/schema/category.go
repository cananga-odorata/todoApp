package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		// field.String("category_id"),
		field.String("name").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at"),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("task", Task.Type).StorageKey(edge.Column("Category_id")),
	}
}

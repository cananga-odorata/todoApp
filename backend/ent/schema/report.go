package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return []ent.Field{
		// field.String("report_id"),
		field.Int("user_id"),
		field.String("content").NotEmpty(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userOwner", User.Type).
			Ref("report").
			Unique(),
	}
}

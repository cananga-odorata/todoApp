package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return []ent.Field{
		// field.String("session_id"),
		field.Int("user_id"),
		field.Time("login_time").Default(time.Now),
		field.Time("logout_time").Default(time.Now),
	}
}

// Edges of the Settings.
func (Settings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userOwner", User.Type).
			Ref("settings").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required(),
	}
}

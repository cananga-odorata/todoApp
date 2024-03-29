package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Authentication holds the schema definition for the Authentication entity.
type Authentication struct {
	ent.Schema
}

// Fields of the Authentication.
func (Authentication) Fields() []ent.Field {
	return []ent.Field{
		field.String("session_id").NotEmpty(),
		field.Int("user_id"),
		field.Time("login_time").Default(time.Now),
		field.Time("logout_time").Default(time.Now),
	}
}

// Edges of the Authentication.
func (Authentication) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userOwner", User.Type).
			Ref("authentication").
			Unique(),
	}
}

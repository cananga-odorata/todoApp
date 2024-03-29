package schema

import (
	"errors"
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// field.Int("user_id"),
		field.String("username"),
		field.String("email").NotEmpty().MaxLen(255).
			Validate(func(s string) error {
				math, _ := regexp.MatchString("^\\w+([.-]?\\w+)*@\\w+([.-]?\\w+)*(\\.\\w{2,3})+$", s)
				if !math {
					return errors.New("invalid email format")
				}
				return nil
			}),
		field.String("password").NotEmpty().MinLen(8),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("task", Task.Type).StorageKey(edge.Column("User_id")),
		edge.To("comment", Comment.Type).StorageKey(edge.Column("User_id")),
		edge.To("report", Report.Type).StorageKey(edge.Column("User_id")),
		edge.To("settings", Settings.Type).StorageKey(edge.Column("User_id")),
		edge.To("authentication", Authentication.Type).StorageKey(edge.Column("User_id")),
		edge.To("sharing", Sharing.Type).StorageKey(edge.Columns("User_id", "Sharing_id")),
		edge.To("notification", Notification.Type).StorageKey(edge.Column("User_id")),
		edge.To("project", Project.Type).StorageKey(edge.Columns("User_id", "Project_id")),
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/settings"
	"backend/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Settings is the model entity for the Settings schema.
type Settings struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// LoginTime holds the value of the "login_time" field.
	LoginTime time.Time `json:"login_time,omitempty"`
	// LogoutTime holds the value of the "logout_time" field.
	LogoutTime time.Time `json:"logout_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SettingsQuery when eager-loading is set.
	Edges        SettingsEdges `json:"edges"`
	_User_id     *int
	selectValues sql.SelectValues
}

// SettingsEdges holds the relations/edges for other nodes in the graph.
type SettingsEdges struct {
	// UserOwner holds the value of the userOwner edge.
	UserOwner *User `json:"userOwner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOwnerOrErr returns the UserOwner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SettingsEdges) UserOwnerOrErr() (*User, error) {
	if e.UserOwner != nil {
		return e.UserOwner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "userOwner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Settings) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case settings.FieldID, settings.FieldUserID:
			values[i] = new(sql.NullInt64)
		case settings.FieldLoginTime, settings.FieldLogoutTime:
			values[i] = new(sql.NullTime)
		case settings.ForeignKeys[0]: // _User_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Settings fields.
func (s *Settings) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case settings.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case settings.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = int(value.Int64)
			}
		case settings.FieldLoginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field login_time", values[i])
			} else if value.Valid {
				s.LoginTime = value.Time
			}
		case settings.FieldLogoutTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field logout_time", values[i])
			} else if value.Valid {
				s.LogoutTime = value.Time
			}
		case settings.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field _User_id", value)
			} else if value.Valid {
				s._User_id = new(int)
				*s._User_id = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Settings.
// This includes values selected through modifiers, order, etc.
func (s *Settings) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryUserOwner queries the "userOwner" edge of the Settings entity.
func (s *Settings) QueryUserOwner() *UserQuery {
	return NewSettingsClient(s.config).QueryUserOwner(s)
}

// Update returns a builder for updating this Settings.
// Note that you need to call Settings.Unwrap() before calling this method if this Settings
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Settings) Update() *SettingsUpdateOne {
	return NewSettingsClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Settings entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Settings) Unwrap() *Settings {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Settings is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Settings) String() string {
	var builder strings.Builder
	builder.WriteString("Settings(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteString(", ")
	builder.WriteString("login_time=")
	builder.WriteString(s.LoginTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("logout_time=")
	builder.WriteString(s.LogoutTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SettingsSlice is a parsable slice of Settings.
type SettingsSlice []*Settings

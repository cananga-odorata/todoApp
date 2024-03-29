// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/authentication"
	"backend/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Authentication is the model entity for the Authentication schema.
type Authentication struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SessionID holds the value of the "session_id" field.
	SessionID string `json:"session_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// LoginTime holds the value of the "login_time" field.
	LoginTime time.Time `json:"login_time,omitempty"`
	// LogoutTime holds the value of the "logout_time" field.
	LogoutTime time.Time `json:"logout_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuthenticationQuery when eager-loading is set.
	Edges        AuthenticationEdges `json:"edges"`
	_User_id     *int
	selectValues sql.SelectValues
}

// AuthenticationEdges holds the relations/edges for other nodes in the graph.
type AuthenticationEdges struct {
	// UserOwner holds the value of the userOwner edge.
	UserOwner *User `json:"userOwner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOwnerOrErr returns the UserOwner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AuthenticationEdges) UserOwnerOrErr() (*User, error) {
	if e.UserOwner != nil {
		return e.UserOwner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "userOwner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Authentication) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case authentication.FieldID, authentication.FieldUserID:
			values[i] = new(sql.NullInt64)
		case authentication.FieldSessionID:
			values[i] = new(sql.NullString)
		case authentication.FieldLoginTime, authentication.FieldLogoutTime:
			values[i] = new(sql.NullTime)
		case authentication.ForeignKeys[0]: // _User_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Authentication fields.
func (a *Authentication) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authentication.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case authentication.FieldSessionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field session_id", values[i])
			} else if value.Valid {
				a.SessionID = value.String
			}
		case authentication.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				a.UserID = int(value.Int64)
			}
		case authentication.FieldLoginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field login_time", values[i])
			} else if value.Valid {
				a.LoginTime = value.Time
			}
		case authentication.FieldLogoutTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field logout_time", values[i])
			} else if value.Valid {
				a.LogoutTime = value.Time
			}
		case authentication.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field _User_id", value)
			} else if value.Valid {
				a._User_id = new(int)
				*a._User_id = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Authentication.
// This includes values selected through modifiers, order, etc.
func (a *Authentication) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUserOwner queries the "userOwner" edge of the Authentication entity.
func (a *Authentication) QueryUserOwner() *UserQuery {
	return NewAuthenticationClient(a.config).QueryUserOwner(a)
}

// Update returns a builder for updating this Authentication.
// Note that you need to call Authentication.Unwrap() before calling this method if this Authentication
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Authentication) Update() *AuthenticationUpdateOne {
	return NewAuthenticationClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Authentication entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Authentication) Unwrap() *Authentication {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Authentication is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Authentication) String() string {
	var builder strings.Builder
	builder.WriteString("Authentication(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("session_id=")
	builder.WriteString(a.SessionID)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", a.UserID))
	builder.WriteString(", ")
	builder.WriteString("login_time=")
	builder.WriteString(a.LoginTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("logout_time=")
	builder.WriteString(a.LogoutTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Authentications is a parsable slice of Authentication.
type Authentications []*Authentication

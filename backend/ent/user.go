// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Task holds the value of the task edge.
	Task []*Task `json:"task,omitempty"`
	// Comment holds the value of the comment edge.
	Comment []*Comment `json:"comment,omitempty"`
	// Report holds the value of the report edge.
	Report []*Report `json:"report,omitempty"`
	// Settings holds the value of the settings edge.
	Settings []*Settings `json:"settings,omitempty"`
	// Authentication holds the value of the authentication edge.
	Authentication []*Authentication `json:"authentication,omitempty"`
	// Sharing holds the value of the sharing edge.
	Sharing []*Sharing `json:"sharing,omitempty"`
	// Notification holds the value of the notification edge.
	Notification []*Notification `json:"notification,omitempty"`
	// Project holds the value of the project edge.
	Project []*Project `json:"project,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TaskOrErr() ([]*Task, error) {
	if e.loadedTypes[0] {
		return e.Task, nil
	}
	return nil, &NotLoadedError{edge: "task"}
}

// CommentOrErr returns the Comment value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CommentOrErr() ([]*Comment, error) {
	if e.loadedTypes[1] {
		return e.Comment, nil
	}
	return nil, &NotLoadedError{edge: "comment"}
}

// ReportOrErr returns the Report value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReportOrErr() ([]*Report, error) {
	if e.loadedTypes[2] {
		return e.Report, nil
	}
	return nil, &NotLoadedError{edge: "report"}
}

// SettingsOrErr returns the Settings value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SettingsOrErr() ([]*Settings, error) {
	if e.loadedTypes[3] {
		return e.Settings, nil
	}
	return nil, &NotLoadedError{edge: "settings"}
}

// AuthenticationOrErr returns the Authentication value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AuthenticationOrErr() ([]*Authentication, error) {
	if e.loadedTypes[4] {
		return e.Authentication, nil
	}
	return nil, &NotLoadedError{edge: "authentication"}
}

// SharingOrErr returns the Sharing value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SharingOrErr() ([]*Sharing, error) {
	if e.loadedTypes[5] {
		return e.Sharing, nil
	}
	return nil, &NotLoadedError{edge: "sharing"}
}

// NotificationOrErr returns the Notification value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NotificationOrErr() ([]*Notification, error) {
	if e.loadedTypes[6] {
		return e.Notification, nil
	}
	return nil, &NotLoadedError{edge: "notification"}
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ProjectOrErr() ([]*Project, error) {
	if e.loadedTypes[7] {
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldEmail, user.FieldPassword:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryTask queries the "task" edge of the User entity.
func (u *User) QueryTask() *TaskQuery {
	return NewUserClient(u.config).QueryTask(u)
}

// QueryComment queries the "comment" edge of the User entity.
func (u *User) QueryComment() *CommentQuery {
	return NewUserClient(u.config).QueryComment(u)
}

// QueryReport queries the "report" edge of the User entity.
func (u *User) QueryReport() *ReportQuery {
	return NewUserClient(u.config).QueryReport(u)
}

// QuerySettings queries the "settings" edge of the User entity.
func (u *User) QuerySettings() *SettingsQuery {
	return NewUserClient(u.config).QuerySettings(u)
}

// QueryAuthentication queries the "authentication" edge of the User entity.
func (u *User) QueryAuthentication() *AuthenticationQuery {
	return NewUserClient(u.config).QueryAuthentication(u)
}

// QuerySharing queries the "sharing" edge of the User entity.
func (u *User) QuerySharing() *SharingQuery {
	return NewUserClient(u.config).QuerySharing(u)
}

// QueryNotification queries the "notification" edge of the User entity.
func (u *User) QueryNotification() *NotificationQuery {
	return NewUserClient(u.config).QueryNotification(u)
}

// QueryProject queries the "project" edge of the User entity.
func (u *User) QueryProject() *ProjectQuery {
	return NewUserClient(u.config).QueryProject(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

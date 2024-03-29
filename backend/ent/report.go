// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/report"
	"backend/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Report is the model entity for the Report schema.
type Report struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReportQuery when eager-loading is set.
	Edges        ReportEdges `json:"edges"`
	_User_id     *int
	selectValues sql.SelectValues
}

// ReportEdges holds the relations/edges for other nodes in the graph.
type ReportEdges struct {
	// UserOwner holds the value of the userOwner edge.
	UserOwner *User `json:"userOwner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOwnerOrErr returns the UserOwner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReportEdges) UserOwnerOrErr() (*User, error) {
	if e.UserOwner != nil {
		return e.UserOwner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "userOwner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Report) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case report.FieldID, report.FieldUserID:
			values[i] = new(sql.NullInt64)
		case report.FieldContent:
			values[i] = new(sql.NullString)
		case report.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case report.ForeignKeys[0]: // _User_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Report fields.
func (r *Report) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case report.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case report.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				r.UserID = int(value.Int64)
			}
		case report.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				r.Content = value.String
			}
		case report.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case report.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field _User_id", value)
			} else if value.Valid {
				r._User_id = new(int)
				*r._User_id = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Report.
// This includes values selected through modifiers, order, etc.
func (r *Report) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryUserOwner queries the "userOwner" edge of the Report entity.
func (r *Report) QueryUserOwner() *UserQuery {
	return NewReportClient(r.config).QueryUserOwner(r)
}

// Update returns a builder for updating this Report.
// Note that you need to call Report.Unwrap() before calling this method if this Report
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Report) Update() *ReportUpdateOne {
	return NewReportClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Report entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Report) Unwrap() *Report {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Report is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Report) String() string {
	var builder strings.Builder
	builder.WriteString("Report(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", r.UserID))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(r.Content)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Reports is a parsable slice of Report.
type Reports []*Report
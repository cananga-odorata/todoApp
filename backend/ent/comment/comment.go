// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTaskID holds the string denoting the task_id field in the database.
	FieldTaskID = "task_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUserOwner holds the string denoting the userowner edge name in mutations.
	EdgeUserOwner = "userOwner"
	// EdgeTaskOwner holds the string denoting the taskowner edge name in mutations.
	EdgeTaskOwner = "taskOwner"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// UserOwnerTable is the table that holds the userOwner relation/edge.
	UserOwnerTable = "comments"
	// UserOwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserOwnerInverseTable = "users"
	// UserOwnerColumn is the table column denoting the userOwner relation/edge.
	UserOwnerColumn = "User_id"
	// TaskOwnerTable is the table that holds the taskOwner relation/edge.
	TaskOwnerTable = "comments"
	// TaskOwnerInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskOwnerInverseTable = "tasks"
	// TaskOwnerColumn is the table column denoting the taskOwner relation/edge.
	TaskOwnerColumn = "Task_id"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldTaskID,
	FieldUserID,
	FieldContent,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"Task_id",
	"User_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// TaskIDValidator is a validator for the "task_id" field. It is called by the builders before save.
	TaskIDValidator func(string) error
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Comment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTaskID orders the results by the task_id field.
func ByTaskID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaskID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUserOwnerField orders the results by userOwner field.
func ByUserOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByTaskOwnerField orders the results by taskOwner field.
func ByTaskOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTaskOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newUserOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserOwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserOwnerTable, UserOwnerColumn),
	)
}
func newTaskOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TaskOwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TaskOwnerTable, TaskOwnerColumn),
	)
}

// Code generated by ent, DO NOT EDIT.

package settings

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the settings type in the database.
	Label = "settings"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldLoginTime holds the string denoting the login_time field in the database.
	FieldLoginTime = "login_time"
	// FieldLogoutTime holds the string denoting the logout_time field in the database.
	FieldLogoutTime = "logout_time"
	// EdgeUserOwner holds the string denoting the userowner edge name in mutations.
	EdgeUserOwner = "userOwner"
	// Table holds the table name of the settings in the database.
	Table = "settings"
	// UserOwnerTable is the table that holds the userOwner relation/edge.
	UserOwnerTable = "settings"
	// UserOwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserOwnerInverseTable = "users"
	// UserOwnerColumn is the table column denoting the userOwner relation/edge.
	UserOwnerColumn = "User_id"
)

// Columns holds all SQL columns for settings fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldLoginTime,
	FieldLogoutTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "settings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
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
	// DefaultLoginTime holds the default value on creation for the "login_time" field.
	DefaultLoginTime func() time.Time
	// DefaultLogoutTime holds the default value on creation for the "logout_time" field.
	DefaultLogoutTime func() time.Time
)

// OrderOption defines the ordering options for the Settings queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByLoginTime orders the results by the login_time field.
func ByLoginTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLoginTime, opts...).ToFunc()
}

// ByLogoutTime orders the results by the logout_time field.
func ByLogoutTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogoutTime, opts...).ToFunc()
}

// ByUserOwnerField orders the results by userOwner field.
func ByUserOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newUserOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserOwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserOwnerTable, UserOwnerColumn),
	)
}

// Code generated by ent, DO NOT EDIT.

package task

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldDueDate holds the string denoting the due_date field in the database.
	FieldDueDate = "due_date"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPiority holds the string denoting the piority field in the database.
	FieldPiority = "piority"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// EdgeUserOwner holds the string denoting the userowner edge name in mutations.
	EdgeUserOwner = "userOwner"
	// EdgeProjectOwner holds the string denoting the projectowner edge name in mutations.
	EdgeProjectOwner = "projectOwner"
	// EdgeComment holds the string denoting the comment edge name in mutations.
	EdgeComment = "comment"
	// EdgeCategoryOwner holds the string denoting the categoryowner edge name in mutations.
	EdgeCategoryOwner = "categoryOwner"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// UserOwnerTable is the table that holds the userOwner relation/edge.
	UserOwnerTable = "tasks"
	// UserOwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserOwnerInverseTable = "users"
	// UserOwnerColumn is the table column denoting the userOwner relation/edge.
	UserOwnerColumn = "User_id"
	// ProjectOwnerTable is the table that holds the projectOwner relation/edge.
	ProjectOwnerTable = "tasks"
	// ProjectOwnerInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectOwnerInverseTable = "projects"
	// ProjectOwnerColumn is the table column denoting the projectOwner relation/edge.
	ProjectOwnerColumn = "Project_id"
	// CommentTable is the table that holds the comment relation/edge.
	CommentTable = "comments"
	// CommentInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentInverseTable = "comments"
	// CommentColumn is the table column denoting the comment relation/edge.
	CommentColumn = "Task_id"
	// CategoryOwnerTable is the table that holds the categoryOwner relation/edge.
	CategoryOwnerTable = "tasks"
	// CategoryOwnerInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryOwnerInverseTable = "categories"
	// CategoryOwnerColumn is the table column denoting the categoryOwner relation/edge.
	CategoryOwnerColumn = "Category_id"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldDueDate,
	FieldStatus,
	FieldPiority,
	FieldTags,
	FieldProjectID,
	FieldCategoryID,
	FieldComments,
	FieldCreatedAt,
	FieldUpdateAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"Category_id",
	"Project_id",
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Task queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByDueDate orders the results by the due_date field.
func ByDueDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDueDate, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPiority orders the results by the piority field.
func ByPiority(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPiority, opts...).ToFunc()
}

// ByTags orders the results by the tags field.
func ByTags(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTags, opts...).ToFunc()
}

// ByProjectID orders the results by the project_id field.
func ByProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectID, opts...).ToFunc()
}

// ByCategoryID orders the results by the category_id field.
func ByCategoryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryID, opts...).ToFunc()
}

// ByComments orders the results by the comments field.
func ByComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComments, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdateAt orders the results by the update_at field.
func ByUpdateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateAt, opts...).ToFunc()
}

// ByUserOwnerField orders the results by userOwner field.
func ByUserOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByProjectOwnerField orders the results by projectOwner field.
func ByProjectOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByCommentCount orders the results by comment count.
func ByCommentCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentStep(), opts...)
	}
}

// ByComment orders the results by comment terms.
func ByComment(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCategoryOwnerField orders the results by categoryOwner field.
func ByCategoryOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newUserOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserOwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserOwnerTable, UserOwnerColumn),
	)
}
func newProjectOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectOwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectOwnerTable, ProjectOwnerColumn),
	)
}
func newCommentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CommentTable, CommentColumn),
	)
}
func newCategoryOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryOwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CategoryOwnerTable, CategoryOwnerColumn),
	)
}

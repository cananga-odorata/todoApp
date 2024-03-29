// Code generated by ent, DO NOT EDIT.

package task

import (
	"backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDescription, v))
}

// DueDate applies equality check predicate on the "due_date" field. It's identical to DueDateEQ.
func DueDate(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDueDate, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatus, v))
}

// Piority applies equality check predicate on the "piority" field. It's identical to PiorityEQ.
func Piority(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldPiority, v))
}

// Tags applies equality check predicate on the "tags" field. It's identical to TagsEQ.
func Tags(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTags, v))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldProjectID, v))
}

// CategoryID applies equality check predicate on the "category_id" field. It's identical to CategoryIDEQ.
func CategoryID(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCategoryID, v))
}

// Comments applies equality check predicate on the "comments" field. It's identical to CommentsEQ.
func Comments(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldComments, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdateAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldDescription, v))
}

// DueDateEQ applies the EQ predicate on the "due_date" field.
func DueDateEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDueDate, v))
}

// DueDateNEQ applies the NEQ predicate on the "due_date" field.
func DueDateNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldDueDate, v))
}

// DueDateIn applies the In predicate on the "due_date" field.
func DueDateIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldDueDate, vs...))
}

// DueDateNotIn applies the NotIn predicate on the "due_date" field.
func DueDateNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldDueDate, vs...))
}

// DueDateGT applies the GT predicate on the "due_date" field.
func DueDateGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldDueDate, v))
}

// DueDateGTE applies the GTE predicate on the "due_date" field.
func DueDateGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldDueDate, v))
}

// DueDateLT applies the LT predicate on the "due_date" field.
func DueDateLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldDueDate, v))
}

// DueDateLTE applies the LTE predicate on the "due_date" field.
func DueDateLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldDueDate, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldStatus, v))
}

// PiorityEQ applies the EQ predicate on the "piority" field.
func PiorityEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldPiority, v))
}

// PiorityNEQ applies the NEQ predicate on the "piority" field.
func PiorityNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldPiority, v))
}

// PiorityIn applies the In predicate on the "piority" field.
func PiorityIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldPiority, vs...))
}

// PiorityNotIn applies the NotIn predicate on the "piority" field.
func PiorityNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldPiority, vs...))
}

// PiorityGT applies the GT predicate on the "piority" field.
func PiorityGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldPiority, v))
}

// PiorityGTE applies the GTE predicate on the "piority" field.
func PiorityGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldPiority, v))
}

// PiorityLT applies the LT predicate on the "piority" field.
func PiorityLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldPiority, v))
}

// PiorityLTE applies the LTE predicate on the "piority" field.
func PiorityLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldPiority, v))
}

// PiorityContains applies the Contains predicate on the "piority" field.
func PiorityContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldPiority, v))
}

// PiorityHasPrefix applies the HasPrefix predicate on the "piority" field.
func PiorityHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldPiority, v))
}

// PiorityHasSuffix applies the HasSuffix predicate on the "piority" field.
func PiorityHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldPiority, v))
}

// PiorityEqualFold applies the EqualFold predicate on the "piority" field.
func PiorityEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldPiority, v))
}

// PiorityContainsFold applies the ContainsFold predicate on the "piority" field.
func PiorityContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldPiority, v))
}

// TagsEQ applies the EQ predicate on the "tags" field.
func TagsEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTags, v))
}

// TagsNEQ applies the NEQ predicate on the "tags" field.
func TagsNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldTags, v))
}

// TagsIn applies the In predicate on the "tags" field.
func TagsIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldTags, vs...))
}

// TagsNotIn applies the NotIn predicate on the "tags" field.
func TagsNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldTags, vs...))
}

// TagsGT applies the GT predicate on the "tags" field.
func TagsGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldTags, v))
}

// TagsGTE applies the GTE predicate on the "tags" field.
func TagsGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldTags, v))
}

// TagsLT applies the LT predicate on the "tags" field.
func TagsLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldTags, v))
}

// TagsLTE applies the LTE predicate on the "tags" field.
func TagsLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldTags, v))
}

// TagsContains applies the Contains predicate on the "tags" field.
func TagsContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldTags, v))
}

// TagsHasPrefix applies the HasPrefix predicate on the "tags" field.
func TagsHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldTags, v))
}

// TagsHasSuffix applies the HasSuffix predicate on the "tags" field.
func TagsHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldTags, v))
}

// TagsEqualFold applies the EqualFold predicate on the "tags" field.
func TagsEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldTags, v))
}

// TagsContainsFold applies the ContainsFold predicate on the "tags" field.
func TagsContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldTags, v))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldProjectID, vs...))
}

// ProjectIDGT applies the GT predicate on the "project_id" field.
func ProjectIDGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldProjectID, v))
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldProjectID, v))
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldProjectID, v))
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldProjectID, v))
}

// ProjectIDContains applies the Contains predicate on the "project_id" field.
func ProjectIDContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldProjectID, v))
}

// ProjectIDHasPrefix applies the HasPrefix predicate on the "project_id" field.
func ProjectIDHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldProjectID, v))
}

// ProjectIDHasSuffix applies the HasSuffix predicate on the "project_id" field.
func ProjectIDHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldProjectID, v))
}

// ProjectIDEqualFold applies the EqualFold predicate on the "project_id" field.
func ProjectIDEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldProjectID, v))
}

// ProjectIDContainsFold applies the ContainsFold predicate on the "project_id" field.
func ProjectIDContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldProjectID, v))
}

// CategoryIDEQ applies the EQ predicate on the "category_id" field.
func CategoryIDEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCategoryID, v))
}

// CategoryIDNEQ applies the NEQ predicate on the "category_id" field.
func CategoryIDNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCategoryID, v))
}

// CategoryIDIn applies the In predicate on the "category_id" field.
func CategoryIDIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCategoryID, vs...))
}

// CategoryIDNotIn applies the NotIn predicate on the "category_id" field.
func CategoryIDNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCategoryID, vs...))
}

// CategoryIDGT applies the GT predicate on the "category_id" field.
func CategoryIDGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCategoryID, v))
}

// CategoryIDGTE applies the GTE predicate on the "category_id" field.
func CategoryIDGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCategoryID, v))
}

// CategoryIDLT applies the LT predicate on the "category_id" field.
func CategoryIDLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCategoryID, v))
}

// CategoryIDLTE applies the LTE predicate on the "category_id" field.
func CategoryIDLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCategoryID, v))
}

// CategoryIDContains applies the Contains predicate on the "category_id" field.
func CategoryIDContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldCategoryID, v))
}

// CategoryIDHasPrefix applies the HasPrefix predicate on the "category_id" field.
func CategoryIDHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldCategoryID, v))
}

// CategoryIDHasSuffix applies the HasSuffix predicate on the "category_id" field.
func CategoryIDHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldCategoryID, v))
}

// CategoryIDEqualFold applies the EqualFold predicate on the "category_id" field.
func CategoryIDEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldCategoryID, v))
}

// CategoryIDContainsFold applies the ContainsFold predicate on the "category_id" field.
func CategoryIDContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldCategoryID, v))
}

// CommentsEQ applies the EQ predicate on the "comments" field.
func CommentsEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldComments, v))
}

// CommentsNEQ applies the NEQ predicate on the "comments" field.
func CommentsNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldComments, v))
}

// CommentsIn applies the In predicate on the "comments" field.
func CommentsIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldComments, vs...))
}

// CommentsNotIn applies the NotIn predicate on the "comments" field.
func CommentsNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldComments, vs...))
}

// CommentsGT applies the GT predicate on the "comments" field.
func CommentsGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldComments, v))
}

// CommentsGTE applies the GTE predicate on the "comments" field.
func CommentsGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldComments, v))
}

// CommentsLT applies the LT predicate on the "comments" field.
func CommentsLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldComments, v))
}

// CommentsLTE applies the LTE predicate on the "comments" field.
func CommentsLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldComments, v))
}

// CommentsContains applies the Contains predicate on the "comments" field.
func CommentsContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldComments, v))
}

// CommentsHasPrefix applies the HasPrefix predicate on the "comments" field.
func CommentsHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldComments, v))
}

// CommentsHasSuffix applies the HasSuffix predicate on the "comments" field.
func CommentsHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldComments, v))
}

// CommentsEqualFold applies the EqualFold predicate on the "comments" field.
func CommentsEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldComments, v))
}

// CommentsContainsFold applies the ContainsFold predicate on the "comments" field.
func CommentsContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldComments, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldUpdateAt, v))
}

// HasUserOwner applies the HasEdge predicate on the "userOwner" edge.
func HasUserOwner() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserOwnerTable, UserOwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserOwnerWith applies the HasEdge predicate on the "userOwner" edge with a given conditions (other predicates).
func HasUserOwnerWith(preds ...predicate.User) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := newUserOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProjectOwner applies the HasEdge predicate on the "projectOwner" edge.
func HasProjectOwner() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectOwnerTable, ProjectOwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectOwnerWith applies the HasEdge predicate on the "projectOwner" edge with a given conditions (other predicates).
func HasProjectOwnerWith(preds ...predicate.Project) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := newProjectOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComment applies the HasEdge predicate on the "comment" edge.
func HasComment() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentTable, CommentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentWith applies the HasEdge predicate on the "comment" edge with a given conditions (other predicates).
func HasCommentWith(preds ...predicate.Comment) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := newCommentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategoryOwner applies the HasEdge predicate on the "categoryOwner" edge.
func HasCategoryOwner() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryOwnerTable, CategoryOwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryOwnerWith applies the HasEdge predicate on the "categoryOwner" edge with a given conditions (other predicates).
func HasCategoryOwnerWith(preds ...predicate.Category) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := newCategoryOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(sql.NotPredicates(p))
}

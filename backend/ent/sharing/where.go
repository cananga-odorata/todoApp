// Code generated by ent, DO NOT EDIT.

package sharing

import (
	"backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Sharing {
	return predicate.Sharing(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Sharing {
	return predicate.Sharing(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Sharing {
	return predicate.Sharing(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Sharing {
	return predicate.Sharing(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Sharing {
	return predicate.Sharing(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Sharing {
	return predicate.Sharing(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Sharing {
	return predicate.Sharing(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldUserID, v))
}

// SharedWithUserID applies equality check predicate on the "shared_with_user_id" field. It's identical to SharedWithUserIDEQ.
func SharedWithUserID(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldSharedWithUserID, v))
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldTaskID, v))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldProjectID, v))
}

// PermissionLevel applies equality check predicate on the "permission_level" field. It's identical to PermissionLevelEQ.
func PermissionLevel(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldPermissionLevel, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Sharing {
	return predicate.Sharing(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Sharing {
	return predicate.Sharing(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Sharing {
	return predicate.Sharing(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.Sharing {
	return predicate.Sharing(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.Sharing {
	return predicate.Sharing(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.Sharing {
	return predicate.Sharing(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.Sharing {
	return predicate.Sharing(sql.FieldLTE(FieldUserID, v))
}

// SharedWithUserIDEQ applies the EQ predicate on the "shared_with_user_id" field.
func SharedWithUserIDEQ(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldSharedWithUserID, v))
}

// SharedWithUserIDNEQ applies the NEQ predicate on the "shared_with_user_id" field.
func SharedWithUserIDNEQ(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldNEQ(FieldSharedWithUserID, v))
}

// SharedWithUserIDIn applies the In predicate on the "shared_with_user_id" field.
func SharedWithUserIDIn(vs ...string) predicate.Sharing {
	return predicate.Sharing(sql.FieldIn(FieldSharedWithUserID, vs...))
}

// SharedWithUserIDNotIn applies the NotIn predicate on the "shared_with_user_id" field.
func SharedWithUserIDNotIn(vs ...string) predicate.Sharing {
	return predicate.Sharing(sql.FieldNotIn(FieldSharedWithUserID, vs...))
}

// SharedWithUserIDGT applies the GT predicate on the "shared_with_user_id" field.
func SharedWithUserIDGT(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldGT(FieldSharedWithUserID, v))
}

// SharedWithUserIDGTE applies the GTE predicate on the "shared_with_user_id" field.
func SharedWithUserIDGTE(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldGTE(FieldSharedWithUserID, v))
}

// SharedWithUserIDLT applies the LT predicate on the "shared_with_user_id" field.
func SharedWithUserIDLT(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldLT(FieldSharedWithUserID, v))
}

// SharedWithUserIDLTE applies the LTE predicate on the "shared_with_user_id" field.
func SharedWithUserIDLTE(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldLTE(FieldSharedWithUserID, v))
}

// SharedWithUserIDContains applies the Contains predicate on the "shared_with_user_id" field.
func SharedWithUserIDContains(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldContains(FieldSharedWithUserID, v))
}

// SharedWithUserIDHasPrefix applies the HasPrefix predicate on the "shared_with_user_id" field.
func SharedWithUserIDHasPrefix(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldHasPrefix(FieldSharedWithUserID, v))
}

// SharedWithUserIDHasSuffix applies the HasSuffix predicate on the "shared_with_user_id" field.
func SharedWithUserIDHasSuffix(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldHasSuffix(FieldSharedWithUserID, v))
}

// SharedWithUserIDEqualFold applies the EqualFold predicate on the "shared_with_user_id" field.
func SharedWithUserIDEqualFold(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEqualFold(FieldSharedWithUserID, v))
}

// SharedWithUserIDContainsFold applies the ContainsFold predicate on the "shared_with_user_id" field.
func SharedWithUserIDContainsFold(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldContainsFold(FieldSharedWithUserID, v))
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldTaskID, v))
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldNEQ(FieldTaskID, v))
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...string) predicate.Sharing {
	return predicate.Sharing(sql.FieldIn(FieldTaskID, vs...))
}

// TaskIDNotIn applies the NotIn predicate on the "task_id" field.
func TaskIDNotIn(vs ...string) predicate.Sharing {
	return predicate.Sharing(sql.FieldNotIn(FieldTaskID, vs...))
}

// TaskIDGT applies the GT predicate on the "task_id" field.
func TaskIDGT(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldGT(FieldTaskID, v))
}

// TaskIDGTE applies the GTE predicate on the "task_id" field.
func TaskIDGTE(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldGTE(FieldTaskID, v))
}

// TaskIDLT applies the LT predicate on the "task_id" field.
func TaskIDLT(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldLT(FieldTaskID, v))
}

// TaskIDLTE applies the LTE predicate on the "task_id" field.
func TaskIDLTE(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldLTE(FieldTaskID, v))
}

// TaskIDContains applies the Contains predicate on the "task_id" field.
func TaskIDContains(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldContains(FieldTaskID, v))
}

// TaskIDHasPrefix applies the HasPrefix predicate on the "task_id" field.
func TaskIDHasPrefix(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldHasPrefix(FieldTaskID, v))
}

// TaskIDHasSuffix applies the HasSuffix predicate on the "task_id" field.
func TaskIDHasSuffix(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldHasSuffix(FieldTaskID, v))
}

// TaskIDEqualFold applies the EqualFold predicate on the "task_id" field.
func TaskIDEqualFold(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEqualFold(FieldTaskID, v))
}

// TaskIDContainsFold applies the ContainsFold predicate on the "task_id" field.
func TaskIDContainsFold(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldContainsFold(FieldTaskID, v))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...string) predicate.Sharing {
	return predicate.Sharing(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...string) predicate.Sharing {
	return predicate.Sharing(sql.FieldNotIn(FieldProjectID, vs...))
}

// ProjectIDGT applies the GT predicate on the "project_id" field.
func ProjectIDGT(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldGT(FieldProjectID, v))
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldGTE(FieldProjectID, v))
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldLT(FieldProjectID, v))
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldLTE(FieldProjectID, v))
}

// ProjectIDContains applies the Contains predicate on the "project_id" field.
func ProjectIDContains(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldContains(FieldProjectID, v))
}

// ProjectIDHasPrefix applies the HasPrefix predicate on the "project_id" field.
func ProjectIDHasPrefix(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldHasPrefix(FieldProjectID, v))
}

// ProjectIDHasSuffix applies the HasSuffix predicate on the "project_id" field.
func ProjectIDHasSuffix(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldHasSuffix(FieldProjectID, v))
}

// ProjectIDEqualFold applies the EqualFold predicate on the "project_id" field.
func ProjectIDEqualFold(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEqualFold(FieldProjectID, v))
}

// ProjectIDContainsFold applies the ContainsFold predicate on the "project_id" field.
func ProjectIDContainsFold(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldContainsFold(FieldProjectID, v))
}

// PermissionLevelEQ applies the EQ predicate on the "permission_level" field.
func PermissionLevelEQ(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEQ(FieldPermissionLevel, v))
}

// PermissionLevelNEQ applies the NEQ predicate on the "permission_level" field.
func PermissionLevelNEQ(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldNEQ(FieldPermissionLevel, v))
}

// PermissionLevelIn applies the In predicate on the "permission_level" field.
func PermissionLevelIn(vs ...string) predicate.Sharing {
	return predicate.Sharing(sql.FieldIn(FieldPermissionLevel, vs...))
}

// PermissionLevelNotIn applies the NotIn predicate on the "permission_level" field.
func PermissionLevelNotIn(vs ...string) predicate.Sharing {
	return predicate.Sharing(sql.FieldNotIn(FieldPermissionLevel, vs...))
}

// PermissionLevelGT applies the GT predicate on the "permission_level" field.
func PermissionLevelGT(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldGT(FieldPermissionLevel, v))
}

// PermissionLevelGTE applies the GTE predicate on the "permission_level" field.
func PermissionLevelGTE(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldGTE(FieldPermissionLevel, v))
}

// PermissionLevelLT applies the LT predicate on the "permission_level" field.
func PermissionLevelLT(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldLT(FieldPermissionLevel, v))
}

// PermissionLevelLTE applies the LTE predicate on the "permission_level" field.
func PermissionLevelLTE(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldLTE(FieldPermissionLevel, v))
}

// PermissionLevelContains applies the Contains predicate on the "permission_level" field.
func PermissionLevelContains(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldContains(FieldPermissionLevel, v))
}

// PermissionLevelHasPrefix applies the HasPrefix predicate on the "permission_level" field.
func PermissionLevelHasPrefix(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldHasPrefix(FieldPermissionLevel, v))
}

// PermissionLevelHasSuffix applies the HasSuffix predicate on the "permission_level" field.
func PermissionLevelHasSuffix(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldHasSuffix(FieldPermissionLevel, v))
}

// PermissionLevelEqualFold applies the EqualFold predicate on the "permission_level" field.
func PermissionLevelEqualFold(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldEqualFold(FieldPermissionLevel, v))
}

// PermissionLevelContainsFold applies the ContainsFold predicate on the "permission_level" field.
func PermissionLevelContainsFold(v string) predicate.Sharing {
	return predicate.Sharing(sql.FieldContainsFold(FieldPermissionLevel, v))
}

// HasUserOwner applies the HasEdge predicate on the "userOwner" edge.
func HasUserOwner() predicate.Sharing {
	return predicate.Sharing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserOwnerTable, UserOwnerPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserOwnerWith applies the HasEdge predicate on the "userOwner" edge with a given conditions (other predicates).
func HasUserOwnerWith(preds ...predicate.User) predicate.Sharing {
	return predicate.Sharing(func(s *sql.Selector) {
		step := newUserOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Sharing) predicate.Sharing {
	return predicate.Sharing(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Sharing) predicate.Sharing {
	return predicate.Sharing(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Sharing) predicate.Sharing {
	return predicate.Sharing(sql.NotPredicates(p))
}

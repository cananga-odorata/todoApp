// Code generated by ent, DO NOT EDIT.

package settings

import (
	"backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Settings {
	return predicate.Settings(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Settings {
	return predicate.Settings(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Settings {
	return predicate.Settings(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Settings {
	return predicate.Settings(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Settings {
	return predicate.Settings(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Settings {
	return predicate.Settings(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Settings {
	return predicate.Settings(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldUserID, v))
}

// LoginTime applies equality check predicate on the "login_time" field. It's identical to LoginTimeEQ.
func LoginTime(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldLoginTime, v))
}

// LogoutTime applies equality check predicate on the "logout_time" field. It's identical to LogoutTimeEQ.
func LogoutTime(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldLogoutTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Settings {
	return predicate.Settings(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Settings {
	return predicate.Settings(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Settings {
	return predicate.Settings(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.Settings {
	return predicate.Settings(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.Settings {
	return predicate.Settings(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.Settings {
	return predicate.Settings(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.Settings {
	return predicate.Settings(sql.FieldLTE(FieldUserID, v))
}

// LoginTimeEQ applies the EQ predicate on the "login_time" field.
func LoginTimeEQ(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldLoginTime, v))
}

// LoginTimeNEQ applies the NEQ predicate on the "login_time" field.
func LoginTimeNEQ(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldNEQ(FieldLoginTime, v))
}

// LoginTimeIn applies the In predicate on the "login_time" field.
func LoginTimeIn(vs ...time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldIn(FieldLoginTime, vs...))
}

// LoginTimeNotIn applies the NotIn predicate on the "login_time" field.
func LoginTimeNotIn(vs ...time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldNotIn(FieldLoginTime, vs...))
}

// LoginTimeGT applies the GT predicate on the "login_time" field.
func LoginTimeGT(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldGT(FieldLoginTime, v))
}

// LoginTimeGTE applies the GTE predicate on the "login_time" field.
func LoginTimeGTE(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldGTE(FieldLoginTime, v))
}

// LoginTimeLT applies the LT predicate on the "login_time" field.
func LoginTimeLT(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldLT(FieldLoginTime, v))
}

// LoginTimeLTE applies the LTE predicate on the "login_time" field.
func LoginTimeLTE(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldLTE(FieldLoginTime, v))
}

// LogoutTimeEQ applies the EQ predicate on the "logout_time" field.
func LogoutTimeEQ(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldEQ(FieldLogoutTime, v))
}

// LogoutTimeNEQ applies the NEQ predicate on the "logout_time" field.
func LogoutTimeNEQ(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldNEQ(FieldLogoutTime, v))
}

// LogoutTimeIn applies the In predicate on the "logout_time" field.
func LogoutTimeIn(vs ...time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldIn(FieldLogoutTime, vs...))
}

// LogoutTimeNotIn applies the NotIn predicate on the "logout_time" field.
func LogoutTimeNotIn(vs ...time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldNotIn(FieldLogoutTime, vs...))
}

// LogoutTimeGT applies the GT predicate on the "logout_time" field.
func LogoutTimeGT(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldGT(FieldLogoutTime, v))
}

// LogoutTimeGTE applies the GTE predicate on the "logout_time" field.
func LogoutTimeGTE(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldGTE(FieldLogoutTime, v))
}

// LogoutTimeLT applies the LT predicate on the "logout_time" field.
func LogoutTimeLT(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldLT(FieldLogoutTime, v))
}

// LogoutTimeLTE applies the LTE predicate on the "logout_time" field.
func LogoutTimeLTE(v time.Time) predicate.Settings {
	return predicate.Settings(sql.FieldLTE(FieldLogoutTime, v))
}

// HasUserOwner applies the HasEdge predicate on the "userOwner" edge.
func HasUserOwner() predicate.Settings {
	return predicate.Settings(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserOwnerTable, UserOwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserOwnerWith applies the HasEdge predicate on the "userOwner" edge with a given conditions (other predicates).
func HasUserOwnerWith(preds ...predicate.User) predicate.Settings {
	return predicate.Settings(func(s *sql.Selector) {
		step := newUserOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Settings) predicate.Settings {
	return predicate.Settings(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Settings) predicate.Settings {
	return predicate.Settings(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Settings) predicate.Settings {
	return predicate.Settings(sql.NotPredicates(p))
}

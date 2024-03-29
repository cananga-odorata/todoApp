// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/notification"
	"backend/ent/predicate"
	"backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationUpdate is the builder for updating Notification entities.
type NotificationUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationMutation
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nu *NotificationUpdate) Where(ps ...predicate.Notification) *NotificationUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetTaskID sets the "task_id" field.
func (nu *NotificationUpdate) SetTaskID(s string) *NotificationUpdate {
	nu.mutation.SetTaskID(s)
	return nu
}

// SetNillableTaskID sets the "task_id" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableTaskID(s *string) *NotificationUpdate {
	if s != nil {
		nu.SetTaskID(*s)
	}
	return nu
}

// SetUserID sets the "user_id" field.
func (nu *NotificationUpdate) SetUserID(i int) *NotificationUpdate {
	nu.mutation.ResetUserID()
	nu.mutation.SetUserID(i)
	return nu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableUserID(i *int) *NotificationUpdate {
	if i != nil {
		nu.SetUserID(*i)
	}
	return nu
}

// AddUserID adds i to the "user_id" field.
func (nu *NotificationUpdate) AddUserID(i int) *NotificationUpdate {
	nu.mutation.AddUserID(i)
	return nu
}

// SetMethod sets the "method" field.
func (nu *NotificationUpdate) SetMethod(s string) *NotificationUpdate {
	nu.mutation.SetMethod(s)
	return nu
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableMethod(s *string) *NotificationUpdate {
	if s != nil {
		nu.SetMethod(*s)
	}
	return nu
}

// SetCreatedAt sets the "created_at" field.
func (nu *NotificationUpdate) SetCreatedAt(t time.Time) *NotificationUpdate {
	nu.mutation.SetCreatedAt(t)
	return nu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableCreatedAt(t *time.Time) *NotificationUpdate {
	if t != nil {
		nu.SetCreatedAt(*t)
	}
	return nu
}

// SetUserOwnerID sets the "userOwner" edge to the User entity by ID.
func (nu *NotificationUpdate) SetUserOwnerID(id int) *NotificationUpdate {
	nu.mutation.SetUserOwnerID(id)
	return nu
}

// SetNillableUserOwnerID sets the "userOwner" edge to the User entity by ID if the given value is not nil.
func (nu *NotificationUpdate) SetNillableUserOwnerID(id *int) *NotificationUpdate {
	if id != nil {
		nu = nu.SetUserOwnerID(*id)
	}
	return nu
}

// SetUserOwner sets the "userOwner" edge to the User entity.
func (nu *NotificationUpdate) SetUserOwner(u *User) *NotificationUpdate {
	return nu.SetUserOwnerID(u.ID)
}

// Mutation returns the NotificationMutation object of the builder.
func (nu *NotificationUpdate) Mutation() *NotificationMutation {
	return nu.mutation
}

// ClearUserOwner clears the "userOwner" edge to the User entity.
func (nu *NotificationUpdate) ClearUserOwner() *NotificationUpdate {
	nu.mutation.ClearUserOwner()
	return nu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NotificationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NotificationUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NotificationUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NotificationUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NotificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.TaskID(); ok {
		_spec.SetField(notification.FieldTaskID, field.TypeString, value)
	}
	if value, ok := nu.mutation.UserID(); ok {
		_spec.SetField(notification.FieldUserID, field.TypeInt, value)
	}
	if value, ok := nu.mutation.AddedUserID(); ok {
		_spec.AddField(notification.FieldUserID, field.TypeInt, value)
	}
	if value, ok := nu.mutation.Method(); ok {
		_spec.SetField(notification.FieldMethod, field.TypeString, value)
	}
	if value, ok := nu.mutation.CreatedAt(); ok {
		_spec.SetField(notification.FieldCreatedAt, field.TypeTime, value)
	}
	if nu.mutation.UserOwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.UserOwnerTable,
			Columns: []string{notification.UserOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.UserOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.UserOwnerTable,
			Columns: []string{notification.UserOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NotificationUpdateOne is the builder for updating a single Notification entity.
type NotificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationMutation
}

// SetTaskID sets the "task_id" field.
func (nuo *NotificationUpdateOne) SetTaskID(s string) *NotificationUpdateOne {
	nuo.mutation.SetTaskID(s)
	return nuo
}

// SetNillableTaskID sets the "task_id" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableTaskID(s *string) *NotificationUpdateOne {
	if s != nil {
		nuo.SetTaskID(*s)
	}
	return nuo
}

// SetUserID sets the "user_id" field.
func (nuo *NotificationUpdateOne) SetUserID(i int) *NotificationUpdateOne {
	nuo.mutation.ResetUserID()
	nuo.mutation.SetUserID(i)
	return nuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableUserID(i *int) *NotificationUpdateOne {
	if i != nil {
		nuo.SetUserID(*i)
	}
	return nuo
}

// AddUserID adds i to the "user_id" field.
func (nuo *NotificationUpdateOne) AddUserID(i int) *NotificationUpdateOne {
	nuo.mutation.AddUserID(i)
	return nuo
}

// SetMethod sets the "method" field.
func (nuo *NotificationUpdateOne) SetMethod(s string) *NotificationUpdateOne {
	nuo.mutation.SetMethod(s)
	return nuo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableMethod(s *string) *NotificationUpdateOne {
	if s != nil {
		nuo.SetMethod(*s)
	}
	return nuo
}

// SetCreatedAt sets the "created_at" field.
func (nuo *NotificationUpdateOne) SetCreatedAt(t time.Time) *NotificationUpdateOne {
	nuo.mutation.SetCreatedAt(t)
	return nuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableCreatedAt(t *time.Time) *NotificationUpdateOne {
	if t != nil {
		nuo.SetCreatedAt(*t)
	}
	return nuo
}

// SetUserOwnerID sets the "userOwner" edge to the User entity by ID.
func (nuo *NotificationUpdateOne) SetUserOwnerID(id int) *NotificationUpdateOne {
	nuo.mutation.SetUserOwnerID(id)
	return nuo
}

// SetNillableUserOwnerID sets the "userOwner" edge to the User entity by ID if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableUserOwnerID(id *int) *NotificationUpdateOne {
	if id != nil {
		nuo = nuo.SetUserOwnerID(*id)
	}
	return nuo
}

// SetUserOwner sets the "userOwner" edge to the User entity.
func (nuo *NotificationUpdateOne) SetUserOwner(u *User) *NotificationUpdateOne {
	return nuo.SetUserOwnerID(u.ID)
}

// Mutation returns the NotificationMutation object of the builder.
func (nuo *NotificationUpdateOne) Mutation() *NotificationMutation {
	return nuo.mutation
}

// ClearUserOwner clears the "userOwner" edge to the User entity.
func (nuo *NotificationUpdateOne) ClearUserOwner() *NotificationUpdateOne {
	nuo.mutation.ClearUserOwner()
	return nuo
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nuo *NotificationUpdateOne) Where(ps ...predicate.Notification) *NotificationUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NotificationUpdateOne) Select(field string, fields ...string) *NotificationUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notification entity.
func (nuo *NotificationUpdateOne) Save(ctx context.Context) (*Notification, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NotificationUpdateOne) SaveX(ctx context.Context) *Notification {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NotificationUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NotificationUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NotificationUpdateOne) sqlSave(ctx context.Context) (_node *Notification, err error) {
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Notification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notification.FieldID)
		for _, f := range fields {
			if !notification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.TaskID(); ok {
		_spec.SetField(notification.FieldTaskID, field.TypeString, value)
	}
	if value, ok := nuo.mutation.UserID(); ok {
		_spec.SetField(notification.FieldUserID, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.AddedUserID(); ok {
		_spec.AddField(notification.FieldUserID, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.Method(); ok {
		_spec.SetField(notification.FieldMethod, field.TypeString, value)
	}
	if value, ok := nuo.mutation.CreatedAt(); ok {
		_spec.SetField(notification.FieldCreatedAt, field.TypeTime, value)
	}
	if nuo.mutation.UserOwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.UserOwnerTable,
			Columns: []string{notification.UserOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.UserOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notification.UserOwnerTable,
			Columns: []string{notification.UserOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Notification{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}

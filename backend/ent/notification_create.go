// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/notification"
	"backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationCreate is the builder for creating a Notification entity.
type NotificationCreate struct {
	config
	mutation *NotificationMutation
	hooks    []Hook
}

// SetTaskID sets the "task_id" field.
func (nc *NotificationCreate) SetTaskID(s string) *NotificationCreate {
	nc.mutation.SetTaskID(s)
	return nc
}

// SetUserID sets the "user_id" field.
func (nc *NotificationCreate) SetUserID(i int) *NotificationCreate {
	nc.mutation.SetUserID(i)
	return nc
}

// SetMethod sets the "method" field.
func (nc *NotificationCreate) SetMethod(s string) *NotificationCreate {
	nc.mutation.SetMethod(s)
	return nc
}

// SetCreatedAt sets the "created_at" field.
func (nc *NotificationCreate) SetCreatedAt(t time.Time) *NotificationCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetUserOwnerID sets the "userOwner" edge to the User entity by ID.
func (nc *NotificationCreate) SetUserOwnerID(id int) *NotificationCreate {
	nc.mutation.SetUserOwnerID(id)
	return nc
}

// SetNillableUserOwnerID sets the "userOwner" edge to the User entity by ID if the given value is not nil.
func (nc *NotificationCreate) SetNillableUserOwnerID(id *int) *NotificationCreate {
	if id != nil {
		nc = nc.SetUserOwnerID(*id)
	}
	return nc
}

// SetUserOwner sets the "userOwner" edge to the User entity.
func (nc *NotificationCreate) SetUserOwner(u *User) *NotificationCreate {
	return nc.SetUserOwnerID(u.ID)
}

// Mutation returns the NotificationMutation object of the builder.
func (nc *NotificationCreate) Mutation() *NotificationMutation {
	return nc.mutation
}

// Save creates the Notification in the database.
func (nc *NotificationCreate) Save(ctx context.Context) (*Notification, error) {
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NotificationCreate) SaveX(ctx context.Context) *Notification {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NotificationCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NotificationCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NotificationCreate) check() error {
	if _, ok := nc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "Notification.task_id"`)}
	}
	if _, ok := nc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Notification.user_id"`)}
	}
	if _, ok := nc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "Notification.method"`)}
	}
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Notification.created_at"`)}
	}
	return nil
}

func (nc *NotificationCreate) sqlSave(ctx context.Context) (*Notification, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NotificationCreate) createSpec() (*Notification, *sqlgraph.CreateSpec) {
	var (
		_node = &Notification{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(notification.Table, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt))
	)
	if value, ok := nc.mutation.TaskID(); ok {
		_spec.SetField(notification.FieldTaskID, field.TypeString, value)
		_node.TaskID = value
	}
	if value, ok := nc.mutation.UserID(); ok {
		_spec.SetField(notification.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := nc.mutation.Method(); ok {
		_spec.SetField(notification.FieldMethod, field.TypeString, value)
		_node.Method = value
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(notification.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := nc.mutation.UserOwnerIDs(); len(nodes) > 0 {
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
		_node._User_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NotificationCreateBulk is the builder for creating many Notification entities in bulk.
type NotificationCreateBulk struct {
	config
	err      error
	builders []*NotificationCreate
}

// Save creates the Notification entities in the database.
func (ncb *NotificationCreateBulk) Save(ctx context.Context) ([]*Notification, error) {
	if ncb.err != nil {
		return nil, ncb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Notification, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotificationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NotificationCreateBulk) SaveX(ctx context.Context) []*Notification {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NotificationCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NotificationCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}

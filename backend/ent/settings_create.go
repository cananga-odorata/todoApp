// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/settings"
	"backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SettingsCreate is the builder for creating a Settings entity.
type SettingsCreate struct {
	config
	mutation *SettingsMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (sc *SettingsCreate) SetUserID(i int) *SettingsCreate {
	sc.mutation.SetUserID(i)
	return sc
}

// SetLoginTime sets the "login_time" field.
func (sc *SettingsCreate) SetLoginTime(t time.Time) *SettingsCreate {
	sc.mutation.SetLoginTime(t)
	return sc
}

// SetNillableLoginTime sets the "login_time" field if the given value is not nil.
func (sc *SettingsCreate) SetNillableLoginTime(t *time.Time) *SettingsCreate {
	if t != nil {
		sc.SetLoginTime(*t)
	}
	return sc
}

// SetLogoutTime sets the "logout_time" field.
func (sc *SettingsCreate) SetLogoutTime(t time.Time) *SettingsCreate {
	sc.mutation.SetLogoutTime(t)
	return sc
}

// SetNillableLogoutTime sets the "logout_time" field if the given value is not nil.
func (sc *SettingsCreate) SetNillableLogoutTime(t *time.Time) *SettingsCreate {
	if t != nil {
		sc.SetLogoutTime(*t)
	}
	return sc
}

// SetUserOwnerID sets the "userOwner" edge to the User entity by ID.
func (sc *SettingsCreate) SetUserOwnerID(id int) *SettingsCreate {
	sc.mutation.SetUserOwnerID(id)
	return sc
}

// SetUserOwner sets the "userOwner" edge to the User entity.
func (sc *SettingsCreate) SetUserOwner(u *User) *SettingsCreate {
	return sc.SetUserOwnerID(u.ID)
}

// Mutation returns the SettingsMutation object of the builder.
func (sc *SettingsCreate) Mutation() *SettingsMutation {
	return sc.mutation
}

// Save creates the Settings in the database.
func (sc *SettingsCreate) Save(ctx context.Context) (*Settings, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SettingsCreate) SaveX(ctx context.Context) *Settings {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SettingsCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SettingsCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SettingsCreate) defaults() {
	if _, ok := sc.mutation.LoginTime(); !ok {
		v := settings.DefaultLoginTime()
		sc.mutation.SetLoginTime(v)
	}
	if _, ok := sc.mutation.LogoutTime(); !ok {
		v := settings.DefaultLogoutTime()
		sc.mutation.SetLogoutTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SettingsCreate) check() error {
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Settings.user_id"`)}
	}
	if _, ok := sc.mutation.LoginTime(); !ok {
		return &ValidationError{Name: "login_time", err: errors.New(`ent: missing required field "Settings.login_time"`)}
	}
	if _, ok := sc.mutation.LogoutTime(); !ok {
		return &ValidationError{Name: "logout_time", err: errors.New(`ent: missing required field "Settings.logout_time"`)}
	}
	if _, ok := sc.mutation.UserOwnerID(); !ok {
		return &ValidationError{Name: "userOwner", err: errors.New(`ent: missing required edge "Settings.userOwner"`)}
	}
	return nil
}

func (sc *SettingsCreate) sqlSave(ctx context.Context) (*Settings, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SettingsCreate) createSpec() (*Settings, *sqlgraph.CreateSpec) {
	var (
		_node = &Settings{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(settings.Table, sqlgraph.NewFieldSpec(settings.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.UserID(); ok {
		_spec.SetField(settings.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := sc.mutation.LoginTime(); ok {
		_spec.SetField(settings.FieldLoginTime, field.TypeTime, value)
		_node.LoginTime = value
	}
	if value, ok := sc.mutation.LogoutTime(); ok {
		_spec.SetField(settings.FieldLogoutTime, field.TypeTime, value)
		_node.LogoutTime = value
	}
	if nodes := sc.mutation.UserOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   settings.UserOwnerTable,
			Columns: []string{settings.UserOwnerColumn},
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

// SettingsCreateBulk is the builder for creating many Settings entities in bulk.
type SettingsCreateBulk struct {
	config
	err      error
	builders []*SettingsCreate
}

// Save creates the Settings entities in the database.
func (scb *SettingsCreateBulk) Save(ctx context.Context) ([]*Settings, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Settings, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SettingsMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SettingsCreateBulk) SaveX(ctx context.Context) []*Settings {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SettingsCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SettingsCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

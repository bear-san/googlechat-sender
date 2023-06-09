// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bear-san/googlechat-sender/backend/ent/systemuser"
)

// SystemUserCreate is the builder for creating a SystemUser entity.
type SystemUserCreate struct {
	config
	mutation *SystemUserMutation
	hooks    []Hook
}

// SetEmail sets the "email" field.
func (suc *SystemUserCreate) SetEmail(s string) *SystemUserCreate {
	suc.mutation.SetEmail(s)
	return suc
}

// SetName sets the "name" field.
func (suc *SystemUserCreate) SetName(s string) *SystemUserCreate {
	suc.mutation.SetName(s)
	return suc
}

// SetID sets the "id" field.
func (suc *SystemUserCreate) SetID(s string) *SystemUserCreate {
	suc.mutation.SetID(s)
	return suc
}

// Mutation returns the SystemUserMutation object of the builder.
func (suc *SystemUserCreate) Mutation() *SystemUserMutation {
	return suc.mutation
}

// Save creates the SystemUser in the database.
func (suc *SystemUserCreate) Save(ctx context.Context) (*SystemUser, error) {
	return withHooks[*SystemUser, SystemUserMutation](ctx, suc.sqlSave, suc.mutation, suc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (suc *SystemUserCreate) SaveX(ctx context.Context) *SystemUser {
	v, err := suc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (suc *SystemUserCreate) Exec(ctx context.Context) error {
	_, err := suc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suc *SystemUserCreate) ExecX(ctx context.Context) {
	if err := suc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suc *SystemUserCreate) check() error {
	if _, ok := suc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "SystemUser.email"`)}
	}
	if _, ok := suc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "SystemUser.name"`)}
	}
	return nil
}

func (suc *SystemUserCreate) sqlSave(ctx context.Context) (*SystemUser, error) {
	if err := suc.check(); err != nil {
		return nil, err
	}
	_node, _spec := suc.createSpec()
	if err := sqlgraph.CreateNode(ctx, suc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SystemUser.ID type: %T", _spec.ID.Value)
		}
	}
	suc.mutation.id = &_node.ID
	suc.mutation.done = true
	return _node, nil
}

func (suc *SystemUserCreate) createSpec() (*SystemUser, *sqlgraph.CreateSpec) {
	var (
		_node = &SystemUser{config: suc.config}
		_spec = sqlgraph.NewCreateSpec(systemuser.Table, sqlgraph.NewFieldSpec(systemuser.FieldID, field.TypeString))
	)
	if id, ok := suc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := suc.mutation.Email(); ok {
		_spec.SetField(systemuser.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := suc.mutation.Name(); ok {
		_spec.SetField(systemuser.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// SystemUserCreateBulk is the builder for creating many SystemUser entities in bulk.
type SystemUserCreateBulk struct {
	config
	builders []*SystemUserCreate
}

// Save creates the SystemUser entities in the database.
func (sucb *SystemUserCreateBulk) Save(ctx context.Context) ([]*SystemUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sucb.builders))
	nodes := make([]*SystemUser, len(sucb.builders))
	mutators := make([]Mutator, len(sucb.builders))
	for i := range sucb.builders {
		func(i int, root context.Context) {
			builder := sucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SystemUserMutation)
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
					_, err = mutators[i+1].Mutate(root, sucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, sucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sucb *SystemUserCreateBulk) SaveX(ctx context.Context) []*SystemUser {
	v, err := sucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sucb *SystemUserCreateBulk) Exec(ctx context.Context) error {
	_, err := sucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sucb *SystemUserCreateBulk) ExecX(ctx context.Context) {
	if err := sucb.Exec(ctx); err != nil {
		panic(err)
	}
}

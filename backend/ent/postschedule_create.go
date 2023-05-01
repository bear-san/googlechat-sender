// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bear-san/googlechat-sender/backend/ent/postschedule"
	"github.com/google/uuid"
)

// PostScheduleCreate is the builder for creating a PostSchedule entity.
type PostScheduleCreate struct {
	config
	mutation *PostScheduleMutation
	hooks    []Hook
}

// SetUID sets the "uid" field.
func (psc *PostScheduleCreate) SetUID(s string) *PostScheduleCreate {
	psc.mutation.SetUID(s)
	return psc
}

// SetTarget sets the "target" field.
func (psc *PostScheduleCreate) SetTarget(s string) *PostScheduleCreate {
	psc.mutation.SetTarget(s)
	return psc
}

// SetText sets the "text" field.
func (psc *PostScheduleCreate) SetText(s string) *PostScheduleCreate {
	psc.mutation.SetText(s)
	return psc
}

// SetIsSent sets the "is_sent" field.
func (psc *PostScheduleCreate) SetIsSent(b bool) *PostScheduleCreate {
	psc.mutation.SetIsSent(b)
	return psc
}

// SetSendAt sets the "send_at" field.
func (psc *PostScheduleCreate) SetSendAt(t time.Time) *PostScheduleCreate {
	psc.mutation.SetSendAt(t)
	return psc
}

// SetID sets the "id" field.
func (psc *PostScheduleCreate) SetID(u uuid.UUID) *PostScheduleCreate {
	psc.mutation.SetID(u)
	return psc
}

// Mutation returns the PostScheduleMutation object of the builder.
func (psc *PostScheduleCreate) Mutation() *PostScheduleMutation {
	return psc.mutation
}

// Save creates the PostSchedule in the database.
func (psc *PostScheduleCreate) Save(ctx context.Context) (*PostSchedule, error) {
	return withHooks[*PostSchedule, PostScheduleMutation](ctx, psc.sqlSave, psc.mutation, psc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psc *PostScheduleCreate) SaveX(ctx context.Context) *PostSchedule {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *PostScheduleCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *PostScheduleCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *PostScheduleCreate) check() error {
	if _, ok := psc.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`ent: missing required field "PostSchedule.uid"`)}
	}
	if _, ok := psc.mutation.Target(); !ok {
		return &ValidationError{Name: "target", err: errors.New(`ent: missing required field "PostSchedule.target"`)}
	}
	if _, ok := psc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "PostSchedule.text"`)}
	}
	if _, ok := psc.mutation.IsSent(); !ok {
		return &ValidationError{Name: "is_sent", err: errors.New(`ent: missing required field "PostSchedule.is_sent"`)}
	}
	if _, ok := psc.mutation.SendAt(); !ok {
		return &ValidationError{Name: "send_at", err: errors.New(`ent: missing required field "PostSchedule.send_at"`)}
	}
	return nil
}

func (psc *PostScheduleCreate) sqlSave(ctx context.Context) (*PostSchedule, error) {
	if err := psc.check(); err != nil {
		return nil, err
	}
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	psc.mutation.id = &_node.ID
	psc.mutation.done = true
	return _node, nil
}

func (psc *PostScheduleCreate) createSpec() (*PostSchedule, *sqlgraph.CreateSpec) {
	var (
		_node = &PostSchedule{config: psc.config}
		_spec = sqlgraph.NewCreateSpec(postschedule.Table, sqlgraph.NewFieldSpec(postschedule.FieldID, field.TypeUUID))
	)
	if id, ok := psc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := psc.mutation.UID(); ok {
		_spec.SetField(postschedule.FieldUID, field.TypeString, value)
		_node.UID = value
	}
	if value, ok := psc.mutation.Target(); ok {
		_spec.SetField(postschedule.FieldTarget, field.TypeString, value)
		_node.Target = value
	}
	if value, ok := psc.mutation.Text(); ok {
		_spec.SetField(postschedule.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := psc.mutation.IsSent(); ok {
		_spec.SetField(postschedule.FieldIsSent, field.TypeBool, value)
		_node.IsSent = value
	}
	if value, ok := psc.mutation.SendAt(); ok {
		_spec.SetField(postschedule.FieldSendAt, field.TypeTime, value)
		_node.SendAt = value
	}
	return _node, _spec
}

// PostScheduleCreateBulk is the builder for creating many PostSchedule entities in bulk.
type PostScheduleCreateBulk struct {
	config
	builders []*PostScheduleCreate
}

// Save creates the PostSchedule entities in the database.
func (pscb *PostScheduleCreateBulk) Save(ctx context.Context) ([]*PostSchedule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*PostSchedule, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostScheduleMutation)
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
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *PostScheduleCreateBulk) SaveX(ctx context.Context) []*PostSchedule {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *PostScheduleCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *PostScheduleCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}

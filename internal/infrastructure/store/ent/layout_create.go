// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/layout"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/section"
)

// LayoutCreate is the builder for creating a Layout entity.
type LayoutCreate struct {
	config
	mutation *LayoutMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (lc *LayoutCreate) SetName(s string) *LayoutCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetID sets the "id" field.
func (lc *LayoutCreate) SetID(u uuid.UUID) *LayoutCreate {
	lc.mutation.SetID(u)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LayoutCreate) SetNillableID(u *uuid.UUID) *LayoutCreate {
	if u != nil {
		lc.SetID(*u)
	}
	return lc
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (lc *LayoutCreate) AddEventIDs(ids ...uuid.UUID) *LayoutCreate {
	lc.mutation.AddEventIDs(ids...)
	return lc
}

// AddEvents adds the "events" edges to the Event entity.
func (lc *LayoutCreate) AddEvents(e ...*Event) *LayoutCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return lc.AddEventIDs(ids...)
}

// AddSectionIDs adds the "sections" edge to the Section entity by IDs.
func (lc *LayoutCreate) AddSectionIDs(ids ...uuid.UUID) *LayoutCreate {
	lc.mutation.AddSectionIDs(ids...)
	return lc
}

// AddSections adds the "sections" edges to the Section entity.
func (lc *LayoutCreate) AddSections(s ...*Section) *LayoutCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return lc.AddSectionIDs(ids...)
}

// Mutation returns the LayoutMutation object of the builder.
func (lc *LayoutCreate) Mutation() *LayoutMutation {
	return lc.mutation
}

// Save creates the Layout in the database.
func (lc *LayoutCreate) Save(ctx context.Context) (*Layout, error) {
	var (
		err  error
		node *Layout
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LayoutMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LayoutCreate) SaveX(ctx context.Context) *Layout {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LayoutCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LayoutCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LayoutCreate) defaults() {
	if _, ok := lc.mutation.ID(); !ok {
		v := layout.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LayoutCreate) check() error {
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Layout.name"`)}
	}
	return nil
}

func (lc *LayoutCreate) sqlSave(ctx context.Context) (*Layout, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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
	return _node, nil
}

func (lc *LayoutCreate) createSpec() (*Layout, *sqlgraph.CreateSpec) {
	var (
		_node = &Layout{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: layout.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: layout.FieldID,
			},
		}
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: layout.FieldName,
		})
		_node.Name = value
	}
	if nodes := lc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   layout.EventsTable,
			Columns: []string{layout.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.SectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   layout.SectionsTable,
			Columns: []string{layout.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: section.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LayoutCreateBulk is the builder for creating many Layout entities in bulk.
type LayoutCreateBulk struct {
	config
	builders []*LayoutCreate
}

// Save creates the Layout entities in the database.
func (lcb *LayoutCreateBulk) Save(ctx context.Context) ([]*Layout, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Layout, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LayoutMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LayoutCreateBulk) SaveX(ctx context.Context) []*Layout {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LayoutCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LayoutCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

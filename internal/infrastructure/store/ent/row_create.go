// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/row"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/seat"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/section"
)

// RowCreate is the builder for creating a Row entity.
type RowCreate struct {
	config
	mutation *RowMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rc *RowCreate) SetName(s string) *RowCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetOrder sets the "order" field.
func (rc *RowCreate) SetOrder(i int) *RowCreate {
	rc.mutation.SetOrder(i)
	return rc
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (rc *RowCreate) SetNillableOrder(i *int) *RowCreate {
	if i != nil {
		rc.SetOrder(*i)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RowCreate) SetID(u uuid.UUID) *RowCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *RowCreate) SetNillableID(u *uuid.UUID) *RowCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// SetSectionID sets the "section" edge to the Section entity by ID.
func (rc *RowCreate) SetSectionID(id uuid.UUID) *RowCreate {
	rc.mutation.SetSectionID(id)
	return rc
}

// SetNillableSectionID sets the "section" edge to the Section entity by ID if the given value is not nil.
func (rc *RowCreate) SetNillableSectionID(id *uuid.UUID) *RowCreate {
	if id != nil {
		rc = rc.SetSectionID(*id)
	}
	return rc
}

// SetSection sets the "section" edge to the Section entity.
func (rc *RowCreate) SetSection(s *Section) *RowCreate {
	return rc.SetSectionID(s.ID)
}

// AddSeatIDs adds the "seats" edge to the Seat entity by IDs.
func (rc *RowCreate) AddSeatIDs(ids ...uuid.UUID) *RowCreate {
	rc.mutation.AddSeatIDs(ids...)
	return rc
}

// AddSeats adds the "seats" edges to the Seat entity.
func (rc *RowCreate) AddSeats(s ...*Seat) *RowCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return rc.AddSeatIDs(ids...)
}

// Mutation returns the RowMutation object of the builder.
func (rc *RowCreate) Mutation() *RowMutation {
	return rc.mutation
}

// Save creates the Row in the database.
func (rc *RowCreate) Save(ctx context.Context) (*Row, error) {
	var (
		err  error
		node *Row
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RowCreate) SaveX(ctx context.Context) *Row {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RowCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RowCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RowCreate) defaults() {
	if _, ok := rc.mutation.Order(); !ok {
		v := row.DefaultOrder
		rc.mutation.SetOrder(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := row.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RowCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Row.name"`)}
	}
	if _, ok := rc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Row.order"`)}
	}
	return nil
}

func (rc *RowCreate) sqlSave(ctx context.Context) (*Row, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
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

func (rc *RowCreate) createSpec() (*Row, *sqlgraph.CreateSpec) {
	var (
		_node = &Row{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: row.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: row.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: row.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rc.mutation.Order(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: row.FieldOrder,
		})
		_node.Order = value
	}
	if nodes := rc.mutation.SectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   row.SectionTable,
			Columns: []string{row.SectionColumn},
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
		_node.section_rows = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.SeatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   row.SeatsTable,
			Columns: []string{row.SeatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: seat.FieldID,
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

// RowCreateBulk is the builder for creating many Row entities in bulk.
type RowCreateBulk struct {
	config
	builders []*RowCreate
}

// Save creates the Row entities in the database.
func (rcb *RowCreateBulk) Save(ctx context.Context) ([]*Row, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Row, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RowMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RowCreateBulk) SaveX(ctx context.Context) []*Row {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RowCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RowCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
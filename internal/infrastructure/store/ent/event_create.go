// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/layout"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/reservation"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/ticket"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	mutation *EventMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ec *EventCreate) SetName(s string) *EventCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetDescription sets the "description" field.
func (ec *EventCreate) SetDescription(s string) *EventCreate {
	ec.mutation.SetDescription(s)
	return ec
}

// SetSeatMap sets the "seat_map" field.
func (ec *EventCreate) SetSeatMap(s string) *EventCreate {
	ec.mutation.SetSeatMap(s)
	return ec
}

// SetStartAt sets the "start_at" field.
func (ec *EventCreate) SetStartAt(t time.Time) *EventCreate {
	ec.mutation.SetStartAt(t)
	return ec
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (ec *EventCreate) SetNillableStartAt(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetStartAt(*t)
	}
	return ec
}

// SetEndAt sets the "end_at" field.
func (ec *EventCreate) SetEndAt(t time.Time) *EventCreate {
	ec.mutation.SetEndAt(t)
	return ec
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (ec *EventCreate) SetNillableEndAt(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetEndAt(*t)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *EventCreate) SetID(u uuid.UUID) *EventCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EventCreate) SetNillableID(u *uuid.UUID) *EventCreate {
	if u != nil {
		ec.SetID(*u)
	}
	return ec
}

// AddReservationIDs adds the "reservations" edge to the Reservation entity by IDs.
func (ec *EventCreate) AddReservationIDs(ids ...uuid.UUID) *EventCreate {
	ec.mutation.AddReservationIDs(ids...)
	return ec
}

// AddReservations adds the "reservations" edges to the Reservation entity.
func (ec *EventCreate) AddReservations(r ...*Reservation) *EventCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ec.AddReservationIDs(ids...)
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (ec *EventCreate) AddTicketIDs(ids ...uuid.UUID) *EventCreate {
	ec.mutation.AddTicketIDs(ids...)
	return ec
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (ec *EventCreate) AddTickets(t ...*Ticket) *EventCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ec.AddTicketIDs(ids...)
}

// SetLayoutID sets the "layout" edge to the Layout entity by ID.
func (ec *EventCreate) SetLayoutID(id uuid.UUID) *EventCreate {
	ec.mutation.SetLayoutID(id)
	return ec
}

// SetNillableLayoutID sets the "layout" edge to the Layout entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillableLayoutID(id *uuid.UUID) *EventCreate {
	if id != nil {
		ec = ec.SetLayoutID(*id)
	}
	return ec
}

// SetLayout sets the "layout" edge to the Layout entity.
func (ec *EventCreate) SetLayout(l *Layout) *EventCreate {
	return ec.SetLayoutID(l.ID)
}

// Mutation returns the EventMutation object of the builder.
func (ec *EventCreate) Mutation() *EventMutation {
	return ec.mutation
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EventCreate) defaults() {
	if _, ok := ec.mutation.StartAt(); !ok {
		v := event.DefaultStartAt()
		ec.mutation.SetStartAt(v)
	}
	if _, ok := ec.mutation.EndAt(); !ok {
		v := event.DefaultEndAt()
		ec.mutation.SetEndAt(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		v := event.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventCreate) check() error {
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Event.name"`)}
	}
	if _, ok := ec.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Event.description"`)}
	}
	if _, ok := ec.mutation.SeatMap(); !ok {
		return &ValidationError{Name: "seat_map", err: errors.New(`ent: missing required field "Event.seat_map"`)}
	}
	if _, ok := ec.mutation.StartAt(); !ok {
		return &ValidationError{Name: "start_at", err: errors.New(`ent: missing required field "Event.start_at"`)}
	}
	if _, ok := ec.mutation.EndAt(); !ok {
		return &ValidationError{Name: "end_at", err: errors.New(`ent: missing required field "Event.end_at"`)}
	}
	return nil
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
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

func (ec *EventCreate) createSpec() (*Event, *sqlgraph.CreateSpec) {
	var (
		_node = &Event{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: event.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: event.FieldID,
			},
		}
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ec.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := ec.mutation.SeatMap(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldSeatMap,
		})
		_node.SeatMap = value
	}
	if value, ok := ec.mutation.StartAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldStartAt,
		})
		_node.StartAt = value
	}
	if value, ok := ec.mutation.EndAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldEndAt,
		})
		_node.EndAt = value
	}
	if nodes := ec.mutation.ReservationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.ReservationsTable,
			Columns: []string{event.ReservationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: reservation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.TicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.TicketsTable,
			Columns: []string{event.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticket.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.LayoutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.LayoutTable,
			Columns: []string{event.LayoutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: layout.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.layout_events = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	builders []*EventCreate
}

// Save creates the Event entities in the database.
func (ecb *EventCreateBulk) Save(ctx context.Context) ([]*Event, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Event, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventCreateBulk) SaveX(ctx context.Context) []*Event {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/layout"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/predicate"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/reservation"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/ticket"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// Where appends a list predicates to the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EventUpdate) SetName(s string) *EventUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetDescription sets the "description" field.
func (eu *EventUpdate) SetDescription(s string) *EventUpdate {
	eu.mutation.SetDescription(s)
	return eu
}

// SetSeatMap sets the "seat_map" field.
func (eu *EventUpdate) SetSeatMap(s string) *EventUpdate {
	eu.mutation.SetSeatMap(s)
	return eu
}

// SetStartAt sets the "start_at" field.
func (eu *EventUpdate) SetStartAt(t time.Time) *EventUpdate {
	eu.mutation.SetStartAt(t)
	return eu
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (eu *EventUpdate) SetNillableStartAt(t *time.Time) *EventUpdate {
	if t != nil {
		eu.SetStartAt(*t)
	}
	return eu
}

// SetEndAt sets the "end_at" field.
func (eu *EventUpdate) SetEndAt(t time.Time) *EventUpdate {
	eu.mutation.SetEndAt(t)
	return eu
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (eu *EventUpdate) SetNillableEndAt(t *time.Time) *EventUpdate {
	if t != nil {
		eu.SetEndAt(*t)
	}
	return eu
}

// AddReservationIDs adds the "reservations" edge to the Reservation entity by IDs.
func (eu *EventUpdate) AddReservationIDs(ids ...uuid.UUID) *EventUpdate {
	eu.mutation.AddReservationIDs(ids...)
	return eu
}

// AddReservations adds the "reservations" edges to the Reservation entity.
func (eu *EventUpdate) AddReservations(r ...*Reservation) *EventUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return eu.AddReservationIDs(ids...)
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (eu *EventUpdate) AddTicketIDs(ids ...uuid.UUID) *EventUpdate {
	eu.mutation.AddTicketIDs(ids...)
	return eu
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (eu *EventUpdate) AddTickets(t ...*Ticket) *EventUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return eu.AddTicketIDs(ids...)
}

// SetLayoutID sets the "layout" edge to the Layout entity by ID.
func (eu *EventUpdate) SetLayoutID(id uuid.UUID) *EventUpdate {
	eu.mutation.SetLayoutID(id)
	return eu
}

// SetNillableLayoutID sets the "layout" edge to the Layout entity by ID if the given value is not nil.
func (eu *EventUpdate) SetNillableLayoutID(id *uuid.UUID) *EventUpdate {
	if id != nil {
		eu = eu.SetLayoutID(*id)
	}
	return eu
}

// SetLayout sets the "layout" edge to the Layout entity.
func (eu *EventUpdate) SetLayout(l *Layout) *EventUpdate {
	return eu.SetLayoutID(l.ID)
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// ClearReservations clears all "reservations" edges to the Reservation entity.
func (eu *EventUpdate) ClearReservations() *EventUpdate {
	eu.mutation.ClearReservations()
	return eu
}

// RemoveReservationIDs removes the "reservations" edge to Reservation entities by IDs.
func (eu *EventUpdate) RemoveReservationIDs(ids ...uuid.UUID) *EventUpdate {
	eu.mutation.RemoveReservationIDs(ids...)
	return eu
}

// RemoveReservations removes "reservations" edges to Reservation entities.
func (eu *EventUpdate) RemoveReservations(r ...*Reservation) *EventUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return eu.RemoveReservationIDs(ids...)
}

// ClearTickets clears all "tickets" edges to the Ticket entity.
func (eu *EventUpdate) ClearTickets() *EventUpdate {
	eu.mutation.ClearTickets()
	return eu
}

// RemoveTicketIDs removes the "tickets" edge to Ticket entities by IDs.
func (eu *EventUpdate) RemoveTicketIDs(ids ...uuid.UUID) *EventUpdate {
	eu.mutation.RemoveTicketIDs(ids...)
	return eu
}

// RemoveTickets removes "tickets" edges to Ticket entities.
func (eu *EventUpdate) RemoveTickets(t ...*Ticket) *EventUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return eu.RemoveTicketIDs(ids...)
}

// ClearLayout clears the "layout" edge to the Layout entity.
func (eu *EventUpdate) ClearLayout() *EventUpdate {
	eu.mutation.ClearLayout()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: event.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldName,
		})
	}
	if value, ok := eu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldDescription,
		})
	}
	if value, ok := eu.mutation.SeatMap(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldSeatMap,
		})
	}
	if value, ok := eu.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldStartAt,
		})
	}
	if value, ok := eu.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldEndAt,
		})
	}
	if eu.mutation.ReservationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedReservationsIDs(); len(nodes) > 0 && !eu.mutation.ReservationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ReservationsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.TicketsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedTicketsIDs(); len(nodes) > 0 && !eu.mutation.TicketsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.TicketsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.LayoutCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.LayoutIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventMutation
}

// SetName sets the "name" field.
func (euo *EventUpdateOne) SetName(s string) *EventUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetDescription sets the "description" field.
func (euo *EventUpdateOne) SetDescription(s string) *EventUpdateOne {
	euo.mutation.SetDescription(s)
	return euo
}

// SetSeatMap sets the "seat_map" field.
func (euo *EventUpdateOne) SetSeatMap(s string) *EventUpdateOne {
	euo.mutation.SetSeatMap(s)
	return euo
}

// SetStartAt sets the "start_at" field.
func (euo *EventUpdateOne) SetStartAt(t time.Time) *EventUpdateOne {
	euo.mutation.SetStartAt(t)
	return euo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableStartAt(t *time.Time) *EventUpdateOne {
	if t != nil {
		euo.SetStartAt(*t)
	}
	return euo
}

// SetEndAt sets the "end_at" field.
func (euo *EventUpdateOne) SetEndAt(t time.Time) *EventUpdateOne {
	euo.mutation.SetEndAt(t)
	return euo
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableEndAt(t *time.Time) *EventUpdateOne {
	if t != nil {
		euo.SetEndAt(*t)
	}
	return euo
}

// AddReservationIDs adds the "reservations" edge to the Reservation entity by IDs.
func (euo *EventUpdateOne) AddReservationIDs(ids ...uuid.UUID) *EventUpdateOne {
	euo.mutation.AddReservationIDs(ids...)
	return euo
}

// AddReservations adds the "reservations" edges to the Reservation entity.
func (euo *EventUpdateOne) AddReservations(r ...*Reservation) *EventUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return euo.AddReservationIDs(ids...)
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (euo *EventUpdateOne) AddTicketIDs(ids ...uuid.UUID) *EventUpdateOne {
	euo.mutation.AddTicketIDs(ids...)
	return euo
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (euo *EventUpdateOne) AddTickets(t ...*Ticket) *EventUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return euo.AddTicketIDs(ids...)
}

// SetLayoutID sets the "layout" edge to the Layout entity by ID.
func (euo *EventUpdateOne) SetLayoutID(id uuid.UUID) *EventUpdateOne {
	euo.mutation.SetLayoutID(id)
	return euo
}

// SetNillableLayoutID sets the "layout" edge to the Layout entity by ID if the given value is not nil.
func (euo *EventUpdateOne) SetNillableLayoutID(id *uuid.UUID) *EventUpdateOne {
	if id != nil {
		euo = euo.SetLayoutID(*id)
	}
	return euo
}

// SetLayout sets the "layout" edge to the Layout entity.
func (euo *EventUpdateOne) SetLayout(l *Layout) *EventUpdateOne {
	return euo.SetLayoutID(l.ID)
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// ClearReservations clears all "reservations" edges to the Reservation entity.
func (euo *EventUpdateOne) ClearReservations() *EventUpdateOne {
	euo.mutation.ClearReservations()
	return euo
}

// RemoveReservationIDs removes the "reservations" edge to Reservation entities by IDs.
func (euo *EventUpdateOne) RemoveReservationIDs(ids ...uuid.UUID) *EventUpdateOne {
	euo.mutation.RemoveReservationIDs(ids...)
	return euo
}

// RemoveReservations removes "reservations" edges to Reservation entities.
func (euo *EventUpdateOne) RemoveReservations(r ...*Reservation) *EventUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return euo.RemoveReservationIDs(ids...)
}

// ClearTickets clears all "tickets" edges to the Ticket entity.
func (euo *EventUpdateOne) ClearTickets() *EventUpdateOne {
	euo.mutation.ClearTickets()
	return euo
}

// RemoveTicketIDs removes the "tickets" edge to Ticket entities by IDs.
func (euo *EventUpdateOne) RemoveTicketIDs(ids ...uuid.UUID) *EventUpdateOne {
	euo.mutation.RemoveTicketIDs(ids...)
	return euo
}

// RemoveTickets removes "tickets" edges to Ticket entities.
func (euo *EventUpdateOne) RemoveTickets(t ...*Ticket) *EventUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return euo.RemoveTicketIDs(ids...)
}

// ClearLayout clears the "layout" edge to the Layout entity.
func (euo *EventUpdateOne) ClearLayout() *EventUpdateOne {
	euo.mutation.ClearLayout()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventUpdateOne) Select(field string, fields ...string) *EventUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: event.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Event.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for _, f := range fields {
			if !event.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldName,
		})
	}
	if value, ok := euo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldDescription,
		})
	}
	if value, ok := euo.mutation.SeatMap(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldSeatMap,
		})
	}
	if value, ok := euo.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldStartAt,
		})
	}
	if value, ok := euo.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldEndAt,
		})
	}
	if euo.mutation.ReservationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedReservationsIDs(); len(nodes) > 0 && !euo.mutation.ReservationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ReservationsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.TicketsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedTicketsIDs(); len(nodes) > 0 && !euo.mutation.TicketsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.TicketsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.LayoutCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.LayoutIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

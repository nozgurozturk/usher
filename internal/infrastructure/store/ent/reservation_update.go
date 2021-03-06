// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/predicate"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/reservation"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/user"
)

// ReservationUpdate is the builder for updating Reservation entities.
type ReservationUpdate struct {
	config
	hooks    []Hook
	mutation *ReservationMutation
}

// Where appends a list predicates to the ReservationUpdate builder.
func (ru *ReservationUpdate) Where(ps ...predicate.Reservation) *ReservationUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetSize sets the "size" field.
func (ru *ReservationUpdate) SetSize(i int) *ReservationUpdate {
	ru.mutation.ResetSize()
	ru.mutation.SetSize(i)
	return ru
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillableSize(i *int) *ReservationUpdate {
	if i != nil {
		ru.SetSize(*i)
	}
	return ru
}

// AddSize adds i to the "size" field.
func (ru *ReservationUpdate) AddSize(i int) *ReservationUpdate {
	ru.mutation.AddSize(i)
	return ru
}

// SetRank sets the "rank" field.
func (ru *ReservationUpdate) SetRank(i int) *ReservationUpdate {
	ru.mutation.ResetRank()
	ru.mutation.SetRank(i)
	return ru
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillableRank(i *int) *ReservationUpdate {
	if i != nil {
		ru.SetRank(*i)
	}
	return ru
}

// AddRank adds i to the "rank" field.
func (ru *ReservationUpdate) AddRank(i int) *ReservationUpdate {
	ru.mutation.AddRank(i)
	return ru
}

// SetPreference sets the "preference" field.
func (ru *ReservationUpdate) SetPreference(i int) *ReservationUpdate {
	ru.mutation.ResetPreference()
	ru.mutation.SetPreference(i)
	return ru
}

// SetNillablePreference sets the "preference" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillablePreference(i *int) *ReservationUpdate {
	if i != nil {
		ru.SetPreference(*i)
	}
	return ru
}

// AddPreference adds i to the "preference" field.
func (ru *ReservationUpdate) AddPreference(i int) *ReservationUpdate {
	ru.mutation.AddPreference(i)
	return ru
}

// SetIsActive sets the "is_active" field.
func (ru *ReservationUpdate) SetIsActive(b bool) *ReservationUpdate {
	ru.mutation.SetIsActive(b)
	return ru
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillableIsActive(b *bool) *ReservationUpdate {
	if b != nil {
		ru.SetIsActive(*b)
	}
	return ru
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (ru *ReservationUpdate) SetEventID(id uuid.UUID) *ReservationUpdate {
	ru.mutation.SetEventID(id)
	return ru
}

// SetEvent sets the "event" edge to the Event entity.
func (ru *ReservationUpdate) SetEvent(e *Event) *ReservationUpdate {
	return ru.SetEventID(e.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ru *ReservationUpdate) SetUserID(id uuid.UUID) *ReservationUpdate {
	ru.mutation.SetUserID(id)
	return ru
}

// SetUser sets the "user" edge to the User entity.
func (ru *ReservationUpdate) SetUser(u *User) *ReservationUpdate {
	return ru.SetUserID(u.ID)
}

// Mutation returns the ReservationMutation object of the builder.
func (ru *ReservationUpdate) Mutation() *ReservationMutation {
	return ru.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (ru *ReservationUpdate) ClearEvent() *ReservationUpdate {
	ru.mutation.ClearEvent()
	return ru
}

// ClearUser clears the "user" edge to the User entity.
func (ru *ReservationUpdate) ClearUser() *ReservationUpdate {
	ru.mutation.ClearUser()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReservationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReservationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReservationUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReservationUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReservationUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReservationUpdate) check() error {
	if _, ok := ru.mutation.EventID(); ru.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reservation.event"`)
	}
	if _, ok := ru.mutation.UserID(); ru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reservation.user"`)
	}
	return nil
}

func (ru *ReservationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reservation.Table,
			Columns: reservation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: reservation.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldSize,
		})
	}
	if value, ok := ru.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldSize,
		})
	}
	if value, ok := ru.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldRank,
		})
	}
	if value, ok := ru.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldRank,
		})
	}
	if value, ok := ru.mutation.Preference(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldPreference,
		})
	}
	if value, ok := ru.mutation.AddedPreference(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldPreference,
		})
	}
	if value, ok := ru.mutation.IsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: reservation.FieldIsActive,
		})
	}
	if ru.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.EventTable,
			Columns: []string{reservation.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.EventTable,
			Columns: []string{reservation.EventColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.UserTable,
			Columns: []string{reservation.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.UserTable,
			Columns: []string{reservation.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reservation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ReservationUpdateOne is the builder for updating a single Reservation entity.
type ReservationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReservationMutation
}

// SetSize sets the "size" field.
func (ruo *ReservationUpdateOne) SetSize(i int) *ReservationUpdateOne {
	ruo.mutation.ResetSize()
	ruo.mutation.SetSize(i)
	return ruo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableSize(i *int) *ReservationUpdateOne {
	if i != nil {
		ruo.SetSize(*i)
	}
	return ruo
}

// AddSize adds i to the "size" field.
func (ruo *ReservationUpdateOne) AddSize(i int) *ReservationUpdateOne {
	ruo.mutation.AddSize(i)
	return ruo
}

// SetRank sets the "rank" field.
func (ruo *ReservationUpdateOne) SetRank(i int) *ReservationUpdateOne {
	ruo.mutation.ResetRank()
	ruo.mutation.SetRank(i)
	return ruo
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableRank(i *int) *ReservationUpdateOne {
	if i != nil {
		ruo.SetRank(*i)
	}
	return ruo
}

// AddRank adds i to the "rank" field.
func (ruo *ReservationUpdateOne) AddRank(i int) *ReservationUpdateOne {
	ruo.mutation.AddRank(i)
	return ruo
}

// SetPreference sets the "preference" field.
func (ruo *ReservationUpdateOne) SetPreference(i int) *ReservationUpdateOne {
	ruo.mutation.ResetPreference()
	ruo.mutation.SetPreference(i)
	return ruo
}

// SetNillablePreference sets the "preference" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillablePreference(i *int) *ReservationUpdateOne {
	if i != nil {
		ruo.SetPreference(*i)
	}
	return ruo
}

// AddPreference adds i to the "preference" field.
func (ruo *ReservationUpdateOne) AddPreference(i int) *ReservationUpdateOne {
	ruo.mutation.AddPreference(i)
	return ruo
}

// SetIsActive sets the "is_active" field.
func (ruo *ReservationUpdateOne) SetIsActive(b bool) *ReservationUpdateOne {
	ruo.mutation.SetIsActive(b)
	return ruo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableIsActive(b *bool) *ReservationUpdateOne {
	if b != nil {
		ruo.SetIsActive(*b)
	}
	return ruo
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (ruo *ReservationUpdateOne) SetEventID(id uuid.UUID) *ReservationUpdateOne {
	ruo.mutation.SetEventID(id)
	return ruo
}

// SetEvent sets the "event" edge to the Event entity.
func (ruo *ReservationUpdateOne) SetEvent(e *Event) *ReservationUpdateOne {
	return ruo.SetEventID(e.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ruo *ReservationUpdateOne) SetUserID(id uuid.UUID) *ReservationUpdateOne {
	ruo.mutation.SetUserID(id)
	return ruo
}

// SetUser sets the "user" edge to the User entity.
func (ruo *ReservationUpdateOne) SetUser(u *User) *ReservationUpdateOne {
	return ruo.SetUserID(u.ID)
}

// Mutation returns the ReservationMutation object of the builder.
func (ruo *ReservationUpdateOne) Mutation() *ReservationMutation {
	return ruo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (ruo *ReservationUpdateOne) ClearEvent() *ReservationUpdateOne {
	ruo.mutation.ClearEvent()
	return ruo
}

// ClearUser clears the "user" edge to the User entity.
func (ruo *ReservationUpdateOne) ClearUser() *ReservationUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReservationUpdateOne) Select(field string, fields ...string) *ReservationUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Reservation entity.
func (ruo *ReservationUpdateOne) Save(ctx context.Context) (*Reservation, error) {
	var (
		err  error
		node *Reservation
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReservationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReservationUpdateOne) SaveX(ctx context.Context) *Reservation {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReservationUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReservationUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReservationUpdateOne) check() error {
	if _, ok := ruo.mutation.EventID(); ruo.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reservation.event"`)
	}
	if _, ok := ruo.mutation.UserID(); ruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reservation.user"`)
	}
	return nil
}

func (ruo *ReservationUpdateOne) sqlSave(ctx context.Context) (_node *Reservation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reservation.Table,
			Columns: reservation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: reservation.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Reservation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reservation.FieldID)
		for _, f := range fields {
			if !reservation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != reservation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldSize,
		})
	}
	if value, ok := ruo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldSize,
		})
	}
	if value, ok := ruo.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldRank,
		})
	}
	if value, ok := ruo.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldRank,
		})
	}
	if value, ok := ruo.mutation.Preference(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldPreference,
		})
	}
	if value, ok := ruo.mutation.AddedPreference(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldPreference,
		})
	}
	if value, ok := ruo.mutation.IsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: reservation.FieldIsActive,
		})
	}
	if ruo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.EventTable,
			Columns: []string{reservation.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.EventTable,
			Columns: []string{reservation.EventColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.UserTable,
			Columns: []string{reservation.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.UserTable,
			Columns: []string{reservation.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Reservation{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reservation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

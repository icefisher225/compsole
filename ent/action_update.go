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
	"github.com/BradHacker/compsole/ent/action"
	"github.com/BradHacker/compsole/ent/predicate"
	"github.com/BradHacker/compsole/ent/user"
	"github.com/google/uuid"
)

// ActionUpdate is the builder for updating Action entities.
type ActionUpdate struct {
	config
	hooks    []Hook
	mutation *ActionMutation
}

// Where appends a list predicates to the ActionUpdate builder.
func (au *ActionUpdate) Where(ps ...predicate.Action) *ActionUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetIPAddress sets the "ip_address" field.
func (au *ActionUpdate) SetIPAddress(s string) *ActionUpdate {
	au.mutation.SetIPAddress(s)
	return au
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (au *ActionUpdate) SetNillableIPAddress(s *string) *ActionUpdate {
	if s != nil {
		au.SetIPAddress(*s)
	}
	return au
}

// SetType sets the "type" field.
func (au *ActionUpdate) SetType(a action.Type) *ActionUpdate {
	au.mutation.SetType(a)
	return au
}

// SetMessage sets the "message" field.
func (au *ActionUpdate) SetMessage(s string) *ActionUpdate {
	au.mutation.SetMessage(s)
	return au
}

// SetPerformedAt sets the "performed_at" field.
func (au *ActionUpdate) SetPerformedAt(t time.Time) *ActionUpdate {
	au.mutation.SetPerformedAt(t)
	return au
}

// SetNillablePerformedAt sets the "performed_at" field if the given value is not nil.
func (au *ActionUpdate) SetNillablePerformedAt(t *time.Time) *ActionUpdate {
	if t != nil {
		au.SetPerformedAt(*t)
	}
	return au
}

// SetActionToUserID sets the "ActionToUser" edge to the User entity by ID.
func (au *ActionUpdate) SetActionToUserID(id uuid.UUID) *ActionUpdate {
	au.mutation.SetActionToUserID(id)
	return au
}

// SetActionToUser sets the "ActionToUser" edge to the User entity.
func (au *ActionUpdate) SetActionToUser(u *User) *ActionUpdate {
	return au.SetActionToUserID(u.ID)
}

// Mutation returns the ActionMutation object of the builder.
func (au *ActionUpdate) Mutation() *ActionMutation {
	return au.mutation
}

// ClearActionToUser clears the "ActionToUser" edge to the User entity.
func (au *ActionUpdate) ClearActionToUser() *ActionUpdate {
	au.mutation.ClearActionToUser()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ActionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ActionUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ActionUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ActionUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ActionUpdate) check() error {
	if v, ok := au.mutation.GetType(); ok {
		if err := action.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Action.type": %w`, err)}
		}
	}
	if _, ok := au.mutation.ActionToUserID(); au.mutation.ActionToUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Action.ActionToUser"`)
	}
	return nil
}

func (au *ActionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   action.Table,
			Columns: action.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: action.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.IPAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldIPAddress,
		})
	}
	if value, ok := au.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: action.FieldType,
		})
	}
	if value, ok := au.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldMessage,
		})
	}
	if value, ok := au.mutation.PerformedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: action.FieldPerformedAt,
		})
	}
	if au.mutation.ActionToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   action.ActionToUserTable,
			Columns: []string{action.ActionToUserColumn},
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
	if nodes := au.mutation.ActionToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   action.ActionToUserTable,
			Columns: []string{action.ActionToUserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{action.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ActionUpdateOne is the builder for updating a single Action entity.
type ActionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActionMutation
}

// SetIPAddress sets the "ip_address" field.
func (auo *ActionUpdateOne) SetIPAddress(s string) *ActionUpdateOne {
	auo.mutation.SetIPAddress(s)
	return auo
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (auo *ActionUpdateOne) SetNillableIPAddress(s *string) *ActionUpdateOne {
	if s != nil {
		auo.SetIPAddress(*s)
	}
	return auo
}

// SetType sets the "type" field.
func (auo *ActionUpdateOne) SetType(a action.Type) *ActionUpdateOne {
	auo.mutation.SetType(a)
	return auo
}

// SetMessage sets the "message" field.
func (auo *ActionUpdateOne) SetMessage(s string) *ActionUpdateOne {
	auo.mutation.SetMessage(s)
	return auo
}

// SetPerformedAt sets the "performed_at" field.
func (auo *ActionUpdateOne) SetPerformedAt(t time.Time) *ActionUpdateOne {
	auo.mutation.SetPerformedAt(t)
	return auo
}

// SetNillablePerformedAt sets the "performed_at" field if the given value is not nil.
func (auo *ActionUpdateOne) SetNillablePerformedAt(t *time.Time) *ActionUpdateOne {
	if t != nil {
		auo.SetPerformedAt(*t)
	}
	return auo
}

// SetActionToUserID sets the "ActionToUser" edge to the User entity by ID.
func (auo *ActionUpdateOne) SetActionToUserID(id uuid.UUID) *ActionUpdateOne {
	auo.mutation.SetActionToUserID(id)
	return auo
}

// SetActionToUser sets the "ActionToUser" edge to the User entity.
func (auo *ActionUpdateOne) SetActionToUser(u *User) *ActionUpdateOne {
	return auo.SetActionToUserID(u.ID)
}

// Mutation returns the ActionMutation object of the builder.
func (auo *ActionUpdateOne) Mutation() *ActionMutation {
	return auo.mutation
}

// ClearActionToUser clears the "ActionToUser" edge to the User entity.
func (auo *ActionUpdateOne) ClearActionToUser() *ActionUpdateOne {
	auo.mutation.ClearActionToUser()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ActionUpdateOne) Select(field string, fields ...string) *ActionUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Action entity.
func (auo *ActionUpdateOne) Save(ctx context.Context) (*Action, error) {
	var (
		err  error
		node *Action
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ActionUpdateOne) SaveX(ctx context.Context) *Action {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ActionUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ActionUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ActionUpdateOne) check() error {
	if v, ok := auo.mutation.GetType(); ok {
		if err := action.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Action.type": %w`, err)}
		}
	}
	if _, ok := auo.mutation.ActionToUserID(); auo.mutation.ActionToUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Action.ActionToUser"`)
	}
	return nil
}

func (auo *ActionUpdateOne) sqlSave(ctx context.Context) (_node *Action, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   action.Table,
			Columns: action.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: action.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Action.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, action.FieldID)
		for _, f := range fields {
			if !action.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != action.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.IPAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldIPAddress,
		})
	}
	if value, ok := auo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: action.FieldType,
		})
	}
	if value, ok := auo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldMessage,
		})
	}
	if value, ok := auo.mutation.PerformedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: action.FieldPerformedAt,
		})
	}
	if auo.mutation.ActionToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   action.ActionToUserTable,
			Columns: []string{action.ActionToUserColumn},
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
	if nodes := auo.mutation.ActionToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   action.ActionToUserTable,
			Columns: []string{action.ActionToUserColumn},
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
	_node = &Action{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{action.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

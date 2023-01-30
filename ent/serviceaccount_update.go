// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/compsole/ent/action"
	"github.com/BradHacker/compsole/ent/predicate"
	"github.com/BradHacker/compsole/ent/serviceaccount"
	"github.com/BradHacker/compsole/ent/servicetoken"
	"github.com/google/uuid"
)

// ServiceAccountUpdate is the builder for updating ServiceAccount entities.
type ServiceAccountUpdate struct {
	config
	hooks    []Hook
	mutation *ServiceAccountMutation
}

// Where appends a list predicates to the ServiceAccountUpdate builder.
func (sau *ServiceAccountUpdate) Where(ps ...predicate.ServiceAccount) *ServiceAccountUpdate {
	sau.mutation.Where(ps...)
	return sau
}

// SetDisplayName sets the "display_name" field.
func (sau *ServiceAccountUpdate) SetDisplayName(s string) *ServiceAccountUpdate {
	sau.mutation.SetDisplayName(s)
	return sau
}

// SetAPIKey sets the "api_key" field.
func (sau *ServiceAccountUpdate) SetAPIKey(u uuid.UUID) *ServiceAccountUpdate {
	sau.mutation.SetAPIKey(u)
	return sau
}

// SetAPISecret sets the "api_secret" field.
func (sau *ServiceAccountUpdate) SetAPISecret(u uuid.UUID) *ServiceAccountUpdate {
	sau.mutation.SetAPISecret(u)
	return sau
}

// SetActive sets the "active" field.
func (sau *ServiceAccountUpdate) SetActive(b bool) *ServiceAccountUpdate {
	sau.mutation.SetActive(b)
	return sau
}

// AddServiceAccountToTokenIDs adds the "ServiceAccountToToken" edge to the ServiceToken entity by IDs.
func (sau *ServiceAccountUpdate) AddServiceAccountToTokenIDs(ids ...uuid.UUID) *ServiceAccountUpdate {
	sau.mutation.AddServiceAccountToTokenIDs(ids...)
	return sau
}

// AddServiceAccountToToken adds the "ServiceAccountToToken" edges to the ServiceToken entity.
func (sau *ServiceAccountUpdate) AddServiceAccountToToken(s ...*ServiceToken) *ServiceAccountUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sau.AddServiceAccountToTokenIDs(ids...)
}

// AddServiceAccountToActionIDs adds the "ServiceAccountToActions" edge to the Action entity by IDs.
func (sau *ServiceAccountUpdate) AddServiceAccountToActionIDs(ids ...uuid.UUID) *ServiceAccountUpdate {
	sau.mutation.AddServiceAccountToActionIDs(ids...)
	return sau
}

// AddServiceAccountToActions adds the "ServiceAccountToActions" edges to the Action entity.
func (sau *ServiceAccountUpdate) AddServiceAccountToActions(a ...*Action) *ServiceAccountUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sau.AddServiceAccountToActionIDs(ids...)
}

// Mutation returns the ServiceAccountMutation object of the builder.
func (sau *ServiceAccountUpdate) Mutation() *ServiceAccountMutation {
	return sau.mutation
}

// ClearServiceAccountToToken clears all "ServiceAccountToToken" edges to the ServiceToken entity.
func (sau *ServiceAccountUpdate) ClearServiceAccountToToken() *ServiceAccountUpdate {
	sau.mutation.ClearServiceAccountToToken()
	return sau
}

// RemoveServiceAccountToTokenIDs removes the "ServiceAccountToToken" edge to ServiceToken entities by IDs.
func (sau *ServiceAccountUpdate) RemoveServiceAccountToTokenIDs(ids ...uuid.UUID) *ServiceAccountUpdate {
	sau.mutation.RemoveServiceAccountToTokenIDs(ids...)
	return sau
}

// RemoveServiceAccountToToken removes "ServiceAccountToToken" edges to ServiceToken entities.
func (sau *ServiceAccountUpdate) RemoveServiceAccountToToken(s ...*ServiceToken) *ServiceAccountUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sau.RemoveServiceAccountToTokenIDs(ids...)
}

// ClearServiceAccountToActions clears all "ServiceAccountToActions" edges to the Action entity.
func (sau *ServiceAccountUpdate) ClearServiceAccountToActions() *ServiceAccountUpdate {
	sau.mutation.ClearServiceAccountToActions()
	return sau
}

// RemoveServiceAccountToActionIDs removes the "ServiceAccountToActions" edge to Action entities by IDs.
func (sau *ServiceAccountUpdate) RemoveServiceAccountToActionIDs(ids ...uuid.UUID) *ServiceAccountUpdate {
	sau.mutation.RemoveServiceAccountToActionIDs(ids...)
	return sau
}

// RemoveServiceAccountToActions removes "ServiceAccountToActions" edges to Action entities.
func (sau *ServiceAccountUpdate) RemoveServiceAccountToActions(a ...*Action) *ServiceAccountUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sau.RemoveServiceAccountToActionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *ServiceAccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sau.hooks) == 0 {
		affected, err = sau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sau.mutation = mutation
			affected, err = sau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sau.hooks) - 1; i >= 0; i-- {
			if sau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sau *ServiceAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *ServiceAccountUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *ServiceAccountUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sau *ServiceAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceaccount.Table,
			Columns: serviceaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: serviceaccount.FieldID,
			},
		},
	}
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sau.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceaccount.FieldDisplayName,
		})
	}
	if value, ok := sau.mutation.APIKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: serviceaccount.FieldAPIKey,
		})
	}
	if value, ok := sau.mutation.APISecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: serviceaccount.FieldAPISecret,
		})
	}
	if value, ok := sau.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: serviceaccount.FieldActive,
		})
	}
	if sau.mutation.ServiceAccountToTokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToTokenTable,
			Columns: []string{serviceaccount.ServiceAccountToTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: servicetoken.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.RemovedServiceAccountToTokenIDs(); len(nodes) > 0 && !sau.mutation.ServiceAccountToTokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToTokenTable,
			Columns: []string{serviceaccount.ServiceAccountToTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: servicetoken.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.ServiceAccountToTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToTokenTable,
			Columns: []string{serviceaccount.ServiceAccountToTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: servicetoken.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sau.mutation.ServiceAccountToActionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToActionsTable,
			Columns: []string{serviceaccount.ServiceAccountToActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: action.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.RemovedServiceAccountToActionsIDs(); len(nodes) > 0 && !sau.mutation.ServiceAccountToActionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToActionsTable,
			Columns: []string{serviceaccount.ServiceAccountToActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: action.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.ServiceAccountToActionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToActionsTable,
			Columns: []string{serviceaccount.ServiceAccountToActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: action.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ServiceAccountUpdateOne is the builder for updating a single ServiceAccount entity.
type ServiceAccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServiceAccountMutation
}

// SetDisplayName sets the "display_name" field.
func (sauo *ServiceAccountUpdateOne) SetDisplayName(s string) *ServiceAccountUpdateOne {
	sauo.mutation.SetDisplayName(s)
	return sauo
}

// SetAPIKey sets the "api_key" field.
func (sauo *ServiceAccountUpdateOne) SetAPIKey(u uuid.UUID) *ServiceAccountUpdateOne {
	sauo.mutation.SetAPIKey(u)
	return sauo
}

// SetAPISecret sets the "api_secret" field.
func (sauo *ServiceAccountUpdateOne) SetAPISecret(u uuid.UUID) *ServiceAccountUpdateOne {
	sauo.mutation.SetAPISecret(u)
	return sauo
}

// SetActive sets the "active" field.
func (sauo *ServiceAccountUpdateOne) SetActive(b bool) *ServiceAccountUpdateOne {
	sauo.mutation.SetActive(b)
	return sauo
}

// AddServiceAccountToTokenIDs adds the "ServiceAccountToToken" edge to the ServiceToken entity by IDs.
func (sauo *ServiceAccountUpdateOne) AddServiceAccountToTokenIDs(ids ...uuid.UUID) *ServiceAccountUpdateOne {
	sauo.mutation.AddServiceAccountToTokenIDs(ids...)
	return sauo
}

// AddServiceAccountToToken adds the "ServiceAccountToToken" edges to the ServiceToken entity.
func (sauo *ServiceAccountUpdateOne) AddServiceAccountToToken(s ...*ServiceToken) *ServiceAccountUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sauo.AddServiceAccountToTokenIDs(ids...)
}

// AddServiceAccountToActionIDs adds the "ServiceAccountToActions" edge to the Action entity by IDs.
func (sauo *ServiceAccountUpdateOne) AddServiceAccountToActionIDs(ids ...uuid.UUID) *ServiceAccountUpdateOne {
	sauo.mutation.AddServiceAccountToActionIDs(ids...)
	return sauo
}

// AddServiceAccountToActions adds the "ServiceAccountToActions" edges to the Action entity.
func (sauo *ServiceAccountUpdateOne) AddServiceAccountToActions(a ...*Action) *ServiceAccountUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sauo.AddServiceAccountToActionIDs(ids...)
}

// Mutation returns the ServiceAccountMutation object of the builder.
func (sauo *ServiceAccountUpdateOne) Mutation() *ServiceAccountMutation {
	return sauo.mutation
}

// ClearServiceAccountToToken clears all "ServiceAccountToToken" edges to the ServiceToken entity.
func (sauo *ServiceAccountUpdateOne) ClearServiceAccountToToken() *ServiceAccountUpdateOne {
	sauo.mutation.ClearServiceAccountToToken()
	return sauo
}

// RemoveServiceAccountToTokenIDs removes the "ServiceAccountToToken" edge to ServiceToken entities by IDs.
func (sauo *ServiceAccountUpdateOne) RemoveServiceAccountToTokenIDs(ids ...uuid.UUID) *ServiceAccountUpdateOne {
	sauo.mutation.RemoveServiceAccountToTokenIDs(ids...)
	return sauo
}

// RemoveServiceAccountToToken removes "ServiceAccountToToken" edges to ServiceToken entities.
func (sauo *ServiceAccountUpdateOne) RemoveServiceAccountToToken(s ...*ServiceToken) *ServiceAccountUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sauo.RemoveServiceAccountToTokenIDs(ids...)
}

// ClearServiceAccountToActions clears all "ServiceAccountToActions" edges to the Action entity.
func (sauo *ServiceAccountUpdateOne) ClearServiceAccountToActions() *ServiceAccountUpdateOne {
	sauo.mutation.ClearServiceAccountToActions()
	return sauo
}

// RemoveServiceAccountToActionIDs removes the "ServiceAccountToActions" edge to Action entities by IDs.
func (sauo *ServiceAccountUpdateOne) RemoveServiceAccountToActionIDs(ids ...uuid.UUID) *ServiceAccountUpdateOne {
	sauo.mutation.RemoveServiceAccountToActionIDs(ids...)
	return sauo
}

// RemoveServiceAccountToActions removes "ServiceAccountToActions" edges to Action entities.
func (sauo *ServiceAccountUpdateOne) RemoveServiceAccountToActions(a ...*Action) *ServiceAccountUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sauo.RemoveServiceAccountToActionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *ServiceAccountUpdateOne) Select(field string, fields ...string) *ServiceAccountUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated ServiceAccount entity.
func (sauo *ServiceAccountUpdateOne) Save(ctx context.Context) (*ServiceAccount, error) {
	var (
		err  error
		node *ServiceAccount
	)
	if len(sauo.hooks) == 0 {
		node, err = sauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sauo.mutation = mutation
			node, err = sauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sauo.hooks) - 1; i >= 0; i-- {
			if sauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *ServiceAccountUpdateOne) SaveX(ctx context.Context) *ServiceAccount {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *ServiceAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *ServiceAccountUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sauo *ServiceAccountUpdateOne) sqlSave(ctx context.Context) (_node *ServiceAccount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceaccount.Table,
			Columns: serviceaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: serviceaccount.FieldID,
			},
		},
	}
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ServiceAccount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serviceaccount.FieldID)
		for _, f := range fields {
			if !serviceaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != serviceaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sauo.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceaccount.FieldDisplayName,
		})
	}
	if value, ok := sauo.mutation.APIKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: serviceaccount.FieldAPIKey,
		})
	}
	if value, ok := sauo.mutation.APISecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: serviceaccount.FieldAPISecret,
		})
	}
	if value, ok := sauo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: serviceaccount.FieldActive,
		})
	}
	if sauo.mutation.ServiceAccountToTokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToTokenTable,
			Columns: []string{serviceaccount.ServiceAccountToTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: servicetoken.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.RemovedServiceAccountToTokenIDs(); len(nodes) > 0 && !sauo.mutation.ServiceAccountToTokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToTokenTable,
			Columns: []string{serviceaccount.ServiceAccountToTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: servicetoken.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.ServiceAccountToTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToTokenTable,
			Columns: []string{serviceaccount.ServiceAccountToTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: servicetoken.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sauo.mutation.ServiceAccountToActionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToActionsTable,
			Columns: []string{serviceaccount.ServiceAccountToActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: action.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.RemovedServiceAccountToActionsIDs(); len(nodes) > 0 && !sauo.mutation.ServiceAccountToActionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToActionsTable,
			Columns: []string{serviceaccount.ServiceAccountToActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: action.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.ServiceAccountToActionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceaccount.ServiceAccountToActionsTable,
			Columns: []string{serviceaccount.ServiceAccountToActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: action.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ServiceAccount{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

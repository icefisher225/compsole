// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/compsole/ent/competition"
	"github.com/BradHacker/compsole/ent/predicate"
	"github.com/BradHacker/compsole/ent/provider"
	"github.com/BradHacker/compsole/ent/team"
	"github.com/google/uuid"
)

// CompetitionUpdate is the builder for updating Competition entities.
type CompetitionUpdate struct {
	config
	hooks    []Hook
	mutation *CompetitionMutation
}

// Where appends a list predicates to the CompetitionUpdate builder.
func (cu *CompetitionUpdate) Where(ps ...predicate.Competition) *CompetitionUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CompetitionUpdate) SetName(s string) *CompetitionUpdate {
	cu.mutation.SetName(s)
	return cu
}

// AddCompetitionToTeamIDs adds the "CompetitionToTeams" edge to the Team entity by IDs.
func (cu *CompetitionUpdate) AddCompetitionToTeamIDs(ids ...uuid.UUID) *CompetitionUpdate {
	cu.mutation.AddCompetitionToTeamIDs(ids...)
	return cu
}

// AddCompetitionToTeams adds the "CompetitionToTeams" edges to the Team entity.
func (cu *CompetitionUpdate) AddCompetitionToTeams(t ...*Team) *CompetitionUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddCompetitionToTeamIDs(ids...)
}

// SetCompetitionToProviderID sets the "CompetitionToProvider" edge to the Provider entity by ID.
func (cu *CompetitionUpdate) SetCompetitionToProviderID(id uuid.UUID) *CompetitionUpdate {
	cu.mutation.SetCompetitionToProviderID(id)
	return cu
}

// SetCompetitionToProvider sets the "CompetitionToProvider" edge to the Provider entity.
func (cu *CompetitionUpdate) SetCompetitionToProvider(p *Provider) *CompetitionUpdate {
	return cu.SetCompetitionToProviderID(p.ID)
}

// Mutation returns the CompetitionMutation object of the builder.
func (cu *CompetitionUpdate) Mutation() *CompetitionMutation {
	return cu.mutation
}

// ClearCompetitionToTeams clears all "CompetitionToTeams" edges to the Team entity.
func (cu *CompetitionUpdate) ClearCompetitionToTeams() *CompetitionUpdate {
	cu.mutation.ClearCompetitionToTeams()
	return cu
}

// RemoveCompetitionToTeamIDs removes the "CompetitionToTeams" edge to Team entities by IDs.
func (cu *CompetitionUpdate) RemoveCompetitionToTeamIDs(ids ...uuid.UUID) *CompetitionUpdate {
	cu.mutation.RemoveCompetitionToTeamIDs(ids...)
	return cu
}

// RemoveCompetitionToTeams removes "CompetitionToTeams" edges to Team entities.
func (cu *CompetitionUpdate) RemoveCompetitionToTeams(t ...*Team) *CompetitionUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveCompetitionToTeamIDs(ids...)
}

// ClearCompetitionToProvider clears the "CompetitionToProvider" edge to the Provider entity.
func (cu *CompetitionUpdate) ClearCompetitionToProvider() *CompetitionUpdate {
	cu.mutation.ClearCompetitionToProvider()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CompetitionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompetitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CompetitionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CompetitionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CompetitionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CompetitionUpdate) check() error {
	if _, ok := cu.mutation.CompetitionToProviderID(); cu.mutation.CompetitionToProviderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Competition.CompetitionToProvider"`)
	}
	return nil
}

func (cu *CompetitionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   competition.Table,
			Columns: competition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: competition.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: competition.FieldName,
		})
	}
	if cu.mutation.CompetitionToTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTeamsTable,
			Columns: []string{competition.CompetitionToTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCompetitionToTeamsIDs(); len(nodes) > 0 && !cu.mutation.CompetitionToTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTeamsTable,
			Columns: []string{competition.CompetitionToTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CompetitionToTeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTeamsTable,
			Columns: []string{competition.CompetitionToTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CompetitionToProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   competition.CompetitionToProviderTable,
			Columns: []string{competition.CompetitionToProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provider.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CompetitionToProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   competition.CompetitionToProviderTable,
			Columns: []string{competition.CompetitionToProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provider.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{competition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CompetitionUpdateOne is the builder for updating a single Competition entity.
type CompetitionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CompetitionMutation
}

// SetName sets the "name" field.
func (cuo *CompetitionUpdateOne) SetName(s string) *CompetitionUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// AddCompetitionToTeamIDs adds the "CompetitionToTeams" edge to the Team entity by IDs.
func (cuo *CompetitionUpdateOne) AddCompetitionToTeamIDs(ids ...uuid.UUID) *CompetitionUpdateOne {
	cuo.mutation.AddCompetitionToTeamIDs(ids...)
	return cuo
}

// AddCompetitionToTeams adds the "CompetitionToTeams" edges to the Team entity.
func (cuo *CompetitionUpdateOne) AddCompetitionToTeams(t ...*Team) *CompetitionUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddCompetitionToTeamIDs(ids...)
}

// SetCompetitionToProviderID sets the "CompetitionToProvider" edge to the Provider entity by ID.
func (cuo *CompetitionUpdateOne) SetCompetitionToProviderID(id uuid.UUID) *CompetitionUpdateOne {
	cuo.mutation.SetCompetitionToProviderID(id)
	return cuo
}

// SetCompetitionToProvider sets the "CompetitionToProvider" edge to the Provider entity.
func (cuo *CompetitionUpdateOne) SetCompetitionToProvider(p *Provider) *CompetitionUpdateOne {
	return cuo.SetCompetitionToProviderID(p.ID)
}

// Mutation returns the CompetitionMutation object of the builder.
func (cuo *CompetitionUpdateOne) Mutation() *CompetitionMutation {
	return cuo.mutation
}

// ClearCompetitionToTeams clears all "CompetitionToTeams" edges to the Team entity.
func (cuo *CompetitionUpdateOne) ClearCompetitionToTeams() *CompetitionUpdateOne {
	cuo.mutation.ClearCompetitionToTeams()
	return cuo
}

// RemoveCompetitionToTeamIDs removes the "CompetitionToTeams" edge to Team entities by IDs.
func (cuo *CompetitionUpdateOne) RemoveCompetitionToTeamIDs(ids ...uuid.UUID) *CompetitionUpdateOne {
	cuo.mutation.RemoveCompetitionToTeamIDs(ids...)
	return cuo
}

// RemoveCompetitionToTeams removes "CompetitionToTeams" edges to Team entities.
func (cuo *CompetitionUpdateOne) RemoveCompetitionToTeams(t ...*Team) *CompetitionUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveCompetitionToTeamIDs(ids...)
}

// ClearCompetitionToProvider clears the "CompetitionToProvider" edge to the Provider entity.
func (cuo *CompetitionUpdateOne) ClearCompetitionToProvider() *CompetitionUpdateOne {
	cuo.mutation.ClearCompetitionToProvider()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CompetitionUpdateOne) Select(field string, fields ...string) *CompetitionUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Competition entity.
func (cuo *CompetitionUpdateOne) Save(ctx context.Context) (*Competition, error) {
	var (
		err  error
		node *Competition
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompetitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CompetitionUpdateOne) SaveX(ctx context.Context) *Competition {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CompetitionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CompetitionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CompetitionUpdateOne) check() error {
	if _, ok := cuo.mutation.CompetitionToProviderID(); cuo.mutation.CompetitionToProviderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Competition.CompetitionToProvider"`)
	}
	return nil
}

func (cuo *CompetitionUpdateOne) sqlSave(ctx context.Context) (_node *Competition, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   competition.Table,
			Columns: competition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: competition.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Competition.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, competition.FieldID)
		for _, f := range fields {
			if !competition.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != competition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: competition.FieldName,
		})
	}
	if cuo.mutation.CompetitionToTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTeamsTable,
			Columns: []string{competition.CompetitionToTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCompetitionToTeamsIDs(); len(nodes) > 0 && !cuo.mutation.CompetitionToTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTeamsTable,
			Columns: []string{competition.CompetitionToTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CompetitionToTeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTeamsTable,
			Columns: []string{competition.CompetitionToTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CompetitionToProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   competition.CompetitionToProviderTable,
			Columns: []string{competition.CompetitionToProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provider.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CompetitionToProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   competition.CompetitionToProviderTable,
			Columns: []string{competition.CompetitionToProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provider.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Competition{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{competition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/compsole/ent/predicate"
	"github.com/BradHacker/compsole/ent/team"
	"github.com/BradHacker/compsole/ent/vmobject"
	"github.com/google/uuid"
)

// VmObjectUpdate is the builder for updating VmObject entities.
type VmObjectUpdate struct {
	config
	hooks    []Hook
	mutation *VmObjectMutation
}

// Where appends a list predicates to the VmObjectUpdate builder.
func (vou *VmObjectUpdate) Where(ps ...predicate.VmObject) *VmObjectUpdate {
	vou.mutation.Where(ps...)
	return vou
}

// SetName sets the "name" field.
func (vou *VmObjectUpdate) SetName(s string) *VmObjectUpdate {
	vou.mutation.SetName(s)
	return vou
}

// SetIdentifier sets the "identifier" field.
func (vou *VmObjectUpdate) SetIdentifier(s string) *VmObjectUpdate {
	vou.mutation.SetIdentifier(s)
	return vou
}

// SetIPAddresses sets the "ip_addresses" field.
func (vou *VmObjectUpdate) SetIPAddresses(s []string) *VmObjectUpdate {
	vou.mutation.SetIPAddresses(s)
	return vou
}

// ClearIPAddresses clears the value of the "ip_addresses" field.
func (vou *VmObjectUpdate) ClearIPAddresses() *VmObjectUpdate {
	vou.mutation.ClearIPAddresses()
	return vou
}

// SetLocked sets the "locked" field.
func (vou *VmObjectUpdate) SetLocked(b bool) *VmObjectUpdate {
	vou.mutation.SetLocked(b)
	return vou
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (vou *VmObjectUpdate) SetNillableLocked(b *bool) *VmObjectUpdate {
	if b != nil {
		vou.SetLocked(*b)
	}
	return vou
}

// SetVmObjectToTeamID sets the "VmObjectToTeam" edge to the Team entity by ID.
func (vou *VmObjectUpdate) SetVmObjectToTeamID(id uuid.UUID) *VmObjectUpdate {
	vou.mutation.SetVmObjectToTeamID(id)
	return vou
}

// SetNillableVmObjectToTeamID sets the "VmObjectToTeam" edge to the Team entity by ID if the given value is not nil.
func (vou *VmObjectUpdate) SetNillableVmObjectToTeamID(id *uuid.UUID) *VmObjectUpdate {
	if id != nil {
		vou = vou.SetVmObjectToTeamID(*id)
	}
	return vou
}

// SetVmObjectToTeam sets the "VmObjectToTeam" edge to the Team entity.
func (vou *VmObjectUpdate) SetVmObjectToTeam(t *Team) *VmObjectUpdate {
	return vou.SetVmObjectToTeamID(t.ID)
}

// Mutation returns the VmObjectMutation object of the builder.
func (vou *VmObjectUpdate) Mutation() *VmObjectMutation {
	return vou.mutation
}

// ClearVmObjectToTeam clears the "VmObjectToTeam" edge to the Team entity.
func (vou *VmObjectUpdate) ClearVmObjectToTeam() *VmObjectUpdate {
	vou.mutation.ClearVmObjectToTeam()
	return vou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vou *VmObjectUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vou.hooks) == 0 {
		affected, err = vou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VmObjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vou.mutation = mutation
			affected, err = vou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vou.hooks) - 1; i >= 0; i-- {
			if vou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vou *VmObjectUpdate) SaveX(ctx context.Context) int {
	affected, err := vou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vou *VmObjectUpdate) Exec(ctx context.Context) error {
	_, err := vou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vou *VmObjectUpdate) ExecX(ctx context.Context) {
	if err := vou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vou *VmObjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vmobject.Table,
			Columns: vmobject.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vmobject.FieldID,
			},
		},
	}
	if ps := vou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vou.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vmobject.FieldName,
		})
	}
	if value, ok := vou.mutation.Identifier(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vmobject.FieldIdentifier,
		})
	}
	if value, ok := vou.mutation.IPAddresses(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: vmobject.FieldIPAddresses,
		})
	}
	if vou.mutation.IPAddressesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: vmobject.FieldIPAddresses,
		})
	}
	if value, ok := vou.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: vmobject.FieldLocked,
		})
	}
	if vou.mutation.VmObjectToTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vmobject.VmObjectToTeamTable,
			Columns: []string{vmobject.VmObjectToTeamColumn},
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
	if nodes := vou.mutation.VmObjectToTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vmobject.VmObjectToTeamTable,
			Columns: []string{vmobject.VmObjectToTeamColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, vou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vmobject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// VmObjectUpdateOne is the builder for updating a single VmObject entity.
type VmObjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VmObjectMutation
}

// SetName sets the "name" field.
func (vouo *VmObjectUpdateOne) SetName(s string) *VmObjectUpdateOne {
	vouo.mutation.SetName(s)
	return vouo
}

// SetIdentifier sets the "identifier" field.
func (vouo *VmObjectUpdateOne) SetIdentifier(s string) *VmObjectUpdateOne {
	vouo.mutation.SetIdentifier(s)
	return vouo
}

// SetIPAddresses sets the "ip_addresses" field.
func (vouo *VmObjectUpdateOne) SetIPAddresses(s []string) *VmObjectUpdateOne {
	vouo.mutation.SetIPAddresses(s)
	return vouo
}

// ClearIPAddresses clears the value of the "ip_addresses" field.
func (vouo *VmObjectUpdateOne) ClearIPAddresses() *VmObjectUpdateOne {
	vouo.mutation.ClearIPAddresses()
	return vouo
}

// SetLocked sets the "locked" field.
func (vouo *VmObjectUpdateOne) SetLocked(b bool) *VmObjectUpdateOne {
	vouo.mutation.SetLocked(b)
	return vouo
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (vouo *VmObjectUpdateOne) SetNillableLocked(b *bool) *VmObjectUpdateOne {
	if b != nil {
		vouo.SetLocked(*b)
	}
	return vouo
}

// SetVmObjectToTeamID sets the "VmObjectToTeam" edge to the Team entity by ID.
func (vouo *VmObjectUpdateOne) SetVmObjectToTeamID(id uuid.UUID) *VmObjectUpdateOne {
	vouo.mutation.SetVmObjectToTeamID(id)
	return vouo
}

// SetNillableVmObjectToTeamID sets the "VmObjectToTeam" edge to the Team entity by ID if the given value is not nil.
func (vouo *VmObjectUpdateOne) SetNillableVmObjectToTeamID(id *uuid.UUID) *VmObjectUpdateOne {
	if id != nil {
		vouo = vouo.SetVmObjectToTeamID(*id)
	}
	return vouo
}

// SetVmObjectToTeam sets the "VmObjectToTeam" edge to the Team entity.
func (vouo *VmObjectUpdateOne) SetVmObjectToTeam(t *Team) *VmObjectUpdateOne {
	return vouo.SetVmObjectToTeamID(t.ID)
}

// Mutation returns the VmObjectMutation object of the builder.
func (vouo *VmObjectUpdateOne) Mutation() *VmObjectMutation {
	return vouo.mutation
}

// ClearVmObjectToTeam clears the "VmObjectToTeam" edge to the Team entity.
func (vouo *VmObjectUpdateOne) ClearVmObjectToTeam() *VmObjectUpdateOne {
	vouo.mutation.ClearVmObjectToTeam()
	return vouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vouo *VmObjectUpdateOne) Select(field string, fields ...string) *VmObjectUpdateOne {
	vouo.fields = append([]string{field}, fields...)
	return vouo
}

// Save executes the query and returns the updated VmObject entity.
func (vouo *VmObjectUpdateOne) Save(ctx context.Context) (*VmObject, error) {
	var (
		err  error
		node *VmObject
	)
	if len(vouo.hooks) == 0 {
		node, err = vouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VmObjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vouo.mutation = mutation
			node, err = vouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vouo.hooks) - 1; i >= 0; i-- {
			if vouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vouo *VmObjectUpdateOne) SaveX(ctx context.Context) *VmObject {
	node, err := vouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vouo *VmObjectUpdateOne) Exec(ctx context.Context) error {
	_, err := vouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vouo *VmObjectUpdateOne) ExecX(ctx context.Context) {
	if err := vouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vouo *VmObjectUpdateOne) sqlSave(ctx context.Context) (_node *VmObject, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vmobject.Table,
			Columns: vmobject.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vmobject.FieldID,
			},
		},
	}
	id, ok := vouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "VmObject.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vmobject.FieldID)
		for _, f := range fields {
			if !vmobject.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vmobject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vouo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vmobject.FieldName,
		})
	}
	if value, ok := vouo.mutation.Identifier(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vmobject.FieldIdentifier,
		})
	}
	if value, ok := vouo.mutation.IPAddresses(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: vmobject.FieldIPAddresses,
		})
	}
	if vouo.mutation.IPAddressesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: vmobject.FieldIPAddresses,
		})
	}
	if value, ok := vouo.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: vmobject.FieldLocked,
		})
	}
	if vouo.mutation.VmObjectToTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vmobject.VmObjectToTeamTable,
			Columns: []string{vmobject.VmObjectToTeamColumn},
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
	if nodes := vouo.mutation.VmObjectToTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vmobject.VmObjectToTeamTable,
			Columns: []string{vmobject.VmObjectToTeamColumn},
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
	_node = &VmObject{config: vouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vmobject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

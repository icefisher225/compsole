// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/compsole/ent/predicate"
	"github.com/BradHacker/compsole/ent/vmobject"
)

// VmObjectDelete is the builder for deleting a VmObject entity.
type VmObjectDelete struct {
	config
	hooks    []Hook
	mutation *VmObjectMutation
}

// Where appends a list predicates to the VmObjectDelete builder.
func (vod *VmObjectDelete) Where(ps ...predicate.VmObject) *VmObjectDelete {
	vod.mutation.Where(ps...)
	return vod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vod *VmObjectDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vod.hooks) == 0 {
		affected, err = vod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VmObjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vod.mutation = mutation
			affected, err = vod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vod.hooks) - 1; i >= 0; i-- {
			if vod.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vod *VmObjectDelete) ExecX(ctx context.Context) int {
	n, err := vod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vod *VmObjectDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vmobject.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vmobject.FieldID,
			},
		},
	}
	if ps := vod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vod.driver, _spec)
}

// VmObjectDeleteOne is the builder for deleting a single VmObject entity.
type VmObjectDeleteOne struct {
	vod *VmObjectDelete
}

// Exec executes the deletion query.
func (vodo *VmObjectDeleteOne) Exec(ctx context.Context) error {
	n, err := vodo.vod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vmobject.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vodo *VmObjectDeleteOne) ExecX(ctx context.Context) {
	vodo.vod.ExecX(ctx)
}

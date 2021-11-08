// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"lin-cms-go/internal/data/model/linuser"
	"lin-cms-go/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinUserDelete is the builder for deleting a LinUser entity.
type LinUserDelete struct {
	config
	hooks    []Hook
	mutation *LinUserMutation
}

// Where appends a list predicates to the LinUserDelete builder.
func (lud *LinUserDelete) Where(ps ...predicate.LinUser) *LinUserDelete {
	lud.mutation.Where(ps...)
	return lud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lud *LinUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lud.hooks) == 0 {
		affected, err = lud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lud.mutation = mutation
			affected, err = lud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lud.hooks) - 1; i >= 0; i-- {
			if lud.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = lud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (lud *LinUserDelete) ExecX(ctx context.Context) int {
	n, err := lud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lud *LinUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: linuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: linuser.FieldID,
			},
		},
	}
	if ps := lud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, lud.driver, _spec)
}

// LinUserDeleteOne is the builder for deleting a single LinUser entity.
type LinUserDeleteOne struct {
	lud *LinUserDelete
}

// Exec executes the deletion query.
func (ludo *LinUserDeleteOne) Exec(ctx context.Context) error {
	n, err := ludo.lud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{linuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ludo *LinUserDeleteOne) ExecX(ctx context.Context) {
	ludo.lud.ExecX(ctx)
}

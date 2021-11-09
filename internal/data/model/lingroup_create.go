// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"lin-cms-go/internal/data/model/lingroup"
	"lin-cms-go/internal/data/model/linpermission"
	"lin-cms-go/internal/data/model/linuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinGroupCreate is the builder for creating a LinGroup entity.
type LinGroupCreate struct {
	config
	mutation *LinGroupMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (lgc *LinGroupCreate) SetCreateTime(t time.Time) *LinGroupCreate {
	lgc.mutation.SetCreateTime(t)
	return lgc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (lgc *LinGroupCreate) SetNillableCreateTime(t *time.Time) *LinGroupCreate {
	if t != nil {
		lgc.SetCreateTime(*t)
	}
	return lgc
}

// SetUpdateTime sets the "update_time" field.
func (lgc *LinGroupCreate) SetUpdateTime(t time.Time) *LinGroupCreate {
	lgc.mutation.SetUpdateTime(t)
	return lgc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (lgc *LinGroupCreate) SetNillableUpdateTime(t *time.Time) *LinGroupCreate {
	if t != nil {
		lgc.SetUpdateTime(*t)
	}
	return lgc
}

// SetDeleteTime sets the "delete_time" field.
func (lgc *LinGroupCreate) SetDeleteTime(t time.Time) *LinGroupCreate {
	lgc.mutation.SetDeleteTime(t)
	return lgc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (lgc *LinGroupCreate) SetNillableDeleteTime(t *time.Time) *LinGroupCreate {
	if t != nil {
		lgc.SetDeleteTime(*t)
	}
	return lgc
}

// SetName sets the "name" field.
func (lgc *LinGroupCreate) SetName(s string) *LinGroupCreate {
	lgc.mutation.SetName(s)
	return lgc
}

// SetInfo sets the "info" field.
func (lgc *LinGroupCreate) SetInfo(s string) *LinGroupCreate {
	lgc.mutation.SetInfo(s)
	return lgc
}

// SetLevel sets the "level" field.
func (lgc *LinGroupCreate) SetLevel(i int8) *LinGroupCreate {
	lgc.mutation.SetLevel(i)
	return lgc
}

// AddLinUserIDs adds the "lin_user" edge to the LinUser entity by IDs.
func (lgc *LinGroupCreate) AddLinUserIDs(ids ...int) *LinGroupCreate {
	lgc.mutation.AddLinUserIDs(ids...)
	return lgc
}

// AddLinUser adds the "lin_user" edges to the LinUser entity.
func (lgc *LinGroupCreate) AddLinUser(l ...*LinUser) *LinGroupCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lgc.AddLinUserIDs(ids...)
}

// AddLinPermissionIDs adds the "lin_permission" edge to the LinPermission entity by IDs.
func (lgc *LinGroupCreate) AddLinPermissionIDs(ids ...int) *LinGroupCreate {
	lgc.mutation.AddLinPermissionIDs(ids...)
	return lgc
}

// AddLinPermission adds the "lin_permission" edges to the LinPermission entity.
func (lgc *LinGroupCreate) AddLinPermission(l ...*LinPermission) *LinGroupCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lgc.AddLinPermissionIDs(ids...)
}

// Mutation returns the LinGroupMutation object of the builder.
func (lgc *LinGroupCreate) Mutation() *LinGroupMutation {
	return lgc.mutation
}

// Save creates the LinGroup in the database.
func (lgc *LinGroupCreate) Save(ctx context.Context) (*LinGroup, error) {
	var (
		err  error
		node *LinGroup
	)
	lgc.defaults()
	if len(lgc.hooks) == 0 {
		if err = lgc.check(); err != nil {
			return nil, err
		}
		node, err = lgc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lgc.check(); err != nil {
				return nil, err
			}
			lgc.mutation = mutation
			if node, err = lgc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lgc.hooks) - 1; i >= 0; i-- {
			if lgc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = lgc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lgc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lgc *LinGroupCreate) SaveX(ctx context.Context) *LinGroup {
	v, err := lgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lgc *LinGroupCreate) Exec(ctx context.Context) error {
	_, err := lgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lgc *LinGroupCreate) ExecX(ctx context.Context) {
	if err := lgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lgc *LinGroupCreate) defaults() {
	if _, ok := lgc.mutation.CreateTime(); !ok {
		v := lingroup.DefaultCreateTime()
		lgc.mutation.SetCreateTime(v)
	}
	if _, ok := lgc.mutation.UpdateTime(); !ok {
		v := lingroup.DefaultUpdateTime()
		lgc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lgc *LinGroupCreate) check() error {
	if _, ok := lgc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "create_time"`)}
	}
	if _, ok := lgc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "update_time"`)}
	}
	if _, ok := lgc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "name"`)}
	}
	if _, ok := lgc.mutation.Info(); !ok {
		return &ValidationError{Name: "info", err: errors.New(`model: missing required field "info"`)}
	}
	if _, ok := lgc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`model: missing required field "level"`)}
	}
	return nil
}

func (lgc *LinGroupCreate) sqlSave(ctx context.Context) (*LinGroup, error) {
	_node, _spec := lgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lgc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lgc *LinGroupCreate) createSpec() (*LinGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &LinGroup{config: lgc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lingroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lingroup.FieldID,
			},
		}
	)
	if value, ok := lgc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lingroup.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := lgc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lingroup.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := lgc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lingroup.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := lgc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lingroup.FieldName,
		})
		_node.Name = value
	}
	if value, ok := lgc.mutation.Info(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lingroup.FieldInfo,
		})
		_node.Info = value
	}
	if value, ok := lgc.mutation.Level(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: lingroup.FieldLevel,
		})
		_node.Level = value
	}
	if nodes := lgc.mutation.LinUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   lingroup.LinUserTable,
			Columns: lingroup.LinUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lgc.mutation.LinPermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   lingroup.LinPermissionTable,
			Columns: lingroup.LinPermissionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linpermission.FieldID,
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

// LinGroupCreateBulk is the builder for creating many LinGroup entities in bulk.
type LinGroupCreateBulk struct {
	config
	builders []*LinGroupCreate
}

// Save creates the LinGroup entities in the database.
func (lgcb *LinGroupCreateBulk) Save(ctx context.Context) ([]*LinGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lgcb.builders))
	nodes := make([]*LinGroup, len(lgcb.builders))
	mutators := make([]Mutator, len(lgcb.builders))
	for i := range lgcb.builders {
		func(i int, root context.Context) {
			builder := lgcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, lgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lgcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lgcb *LinGroupCreateBulk) SaveX(ctx context.Context) []*LinGroup {
	v, err := lgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lgcb *LinGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := lgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lgcb *LinGroupCreateBulk) ExecX(ctx context.Context) {
	if err := lgcb.Exec(ctx); err != nil {
		panic(err)
	}
}

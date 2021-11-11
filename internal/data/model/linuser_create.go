// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"lin-cms-go/internal/data/model/lingroup"
	"lin-cms-go/internal/data/model/linuser"
	"lin-cms-go/internal/data/model/linuseridentiy"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinUserCreate is the builder for creating a LinUser entity.
type LinUserCreate struct {
	config
	mutation *LinUserMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (luc *LinUserCreate) SetCreateTime(t time.Time) *LinUserCreate {
	luc.mutation.SetCreateTime(t)
	return luc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (luc *LinUserCreate) SetNillableCreateTime(t *time.Time) *LinUserCreate {
	if t != nil {
		luc.SetCreateTime(*t)
	}
	return luc
}

// SetUpdateTime sets the "update_time" field.
func (luc *LinUserCreate) SetUpdateTime(t time.Time) *LinUserCreate {
	luc.mutation.SetUpdateTime(t)
	return luc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (luc *LinUserCreate) SetNillableUpdateTime(t *time.Time) *LinUserCreate {
	if t != nil {
		luc.SetUpdateTime(*t)
	}
	return luc
}

// SetDeleteTime sets the "delete_time" field.
func (luc *LinUserCreate) SetDeleteTime(t time.Time) *LinUserCreate {
	luc.mutation.SetDeleteTime(t)
	return luc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (luc *LinUserCreate) SetNillableDeleteTime(t *time.Time) *LinUserCreate {
	if t != nil {
		luc.SetDeleteTime(*t)
	}
	return luc
}

// SetUsername sets the "username" field.
func (luc *LinUserCreate) SetUsername(s string) *LinUserCreate {
	luc.mutation.SetUsername(s)
	return luc
}

// SetNickname sets the "nickname" field.
func (luc *LinUserCreate) SetNickname(s string) *LinUserCreate {
	luc.mutation.SetNickname(s)
	return luc
}

// SetAvatar sets the "avatar" field.
func (luc *LinUserCreate) SetAvatar(s string) *LinUserCreate {
	luc.mutation.SetAvatar(s)
	return luc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (luc *LinUserCreate) SetNillableAvatar(s *string) *LinUserCreate {
	if s != nil {
		luc.SetAvatar(*s)
	}
	return luc
}

// SetEmail sets the "email" field.
func (luc *LinUserCreate) SetEmail(s string) *LinUserCreate {
	luc.mutation.SetEmail(s)
	return luc
}

// AddLinUserIdentiyIDs adds the "lin_user_identiy" edge to the LinUserIdentiy entity by IDs.
func (luc *LinUserCreate) AddLinUserIdentiyIDs(ids ...int) *LinUserCreate {
	luc.mutation.AddLinUserIdentiyIDs(ids...)
	return luc
}

// AddLinUserIdentiy adds the "lin_user_identiy" edges to the LinUserIdentiy entity.
func (luc *LinUserCreate) AddLinUserIdentiy(l ...*LinUserIdentiy) *LinUserCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luc.AddLinUserIdentiyIDs(ids...)
}

// AddLinGroupIDs adds the "lin_group" edge to the LinGroup entity by IDs.
func (luc *LinUserCreate) AddLinGroupIDs(ids ...int) *LinUserCreate {
	luc.mutation.AddLinGroupIDs(ids...)
	return luc
}

// AddLinGroup adds the "lin_group" edges to the LinGroup entity.
func (luc *LinUserCreate) AddLinGroup(l ...*LinGroup) *LinUserCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luc.AddLinGroupIDs(ids...)
}

// Mutation returns the LinUserMutation object of the builder.
func (luc *LinUserCreate) Mutation() *LinUserMutation {
	return luc.mutation
}

// Save creates the LinUser in the database.
func (luc *LinUserCreate) Save(ctx context.Context) (*LinUser, error) {
	var (
		err  error
		node *LinUser
	)
	luc.defaults()
	if len(luc.hooks) == 0 {
		if err = luc.check(); err != nil {
			return nil, err
		}
		node, err = luc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luc.check(); err != nil {
				return nil, err
			}
			luc.mutation = mutation
			if node, err = luc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(luc.hooks) - 1; i >= 0; i-- {
			if luc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = luc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (luc *LinUserCreate) SaveX(ctx context.Context) *LinUser {
	v, err := luc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (luc *LinUserCreate) Exec(ctx context.Context) error {
	_, err := luc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luc *LinUserCreate) ExecX(ctx context.Context) {
	if err := luc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luc *LinUserCreate) defaults() {
	if _, ok := luc.mutation.CreateTime(); !ok {
		v := linuser.DefaultCreateTime()
		luc.mutation.SetCreateTime(v)
	}
	if _, ok := luc.mutation.UpdateTime(); !ok {
		v := linuser.DefaultUpdateTime()
		luc.mutation.SetUpdateTime(v)
	}
	if _, ok := luc.mutation.Avatar(); !ok {
		v := linuser.DefaultAvatar
		luc.mutation.SetAvatar(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luc *LinUserCreate) check() error {
	if _, ok := luc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "create_time"`)}
	}
	if _, ok := luc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "update_time"`)}
	}
	if _, ok := luc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`model: missing required field "username"`)}
	}
	if _, ok := luc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`model: missing required field "nickname"`)}
	}
	if _, ok := luc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`model: missing required field "avatar"`)}
	}
	if _, ok := luc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`model: missing required field "email"`)}
	}
	return nil
}

func (luc *LinUserCreate) sqlSave(ctx context.Context) (*LinUser, error) {
	_node, _spec := luc.createSpec()
	if err := sqlgraph.CreateNode(ctx, luc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (luc *LinUserCreate) createSpec() (*LinUser, *sqlgraph.CreateSpec) {
	var (
		_node = &LinUser{config: luc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: linuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: linuser.FieldID,
			},
		}
	)
	if value, ok := luc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: linuser.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := luc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: linuser.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := luc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: linuser.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := luc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: linuser.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := luc.mutation.Nickname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: linuser.FieldNickname,
		})
		_node.Nickname = value
	}
	if value, ok := luc.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: linuser.FieldAvatar,
		})
		_node.Avatar = value
	}
	if value, ok := luc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: linuser.FieldEmail,
		})
		_node.Email = value
	}
	if nodes := luc.mutation.LinUserIdentiyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   linuser.LinUserIdentiyTable,
			Columns: []string{linuser.LinUserIdentiyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linuseridentiy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := luc.mutation.LinGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   linuser.LinGroupTable,
			Columns: linuser.LinGroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lingroup.FieldID,
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

// LinUserCreateBulk is the builder for creating many LinUser entities in bulk.
type LinUserCreateBulk struct {
	config
	builders []*LinUserCreate
}

// Save creates the LinUser entities in the database.
func (lucb *LinUserCreateBulk) Save(ctx context.Context) ([]*LinUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lucb.builders))
	nodes := make([]*LinUser, len(lucb.builders))
	mutators := make([]Mutator, len(lucb.builders))
	for i := range lucb.builders {
		func(i int, root context.Context) {
			builder := lucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinUserMutation)
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
					_, err = mutators[i+1].Mutate(root, lucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lucb *LinUserCreateBulk) SaveX(ctx context.Context) []*LinUser {
	v, err := lucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lucb *LinUserCreateBulk) Exec(ctx context.Context) error {
	_, err := lucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lucb *LinUserCreateBulk) ExecX(ctx context.Context) {
	if err := lucb.Exec(ctx); err != nil {
		panic(err)
	}
}
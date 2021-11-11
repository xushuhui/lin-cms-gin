// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"lin-cms-go/internal/data/model"
)

// The BookFunc type is an adapter to allow the use of ordinary
// function as Book mutator.
type BookFunc func(context.Context, *model.BookMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f BookFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.BookMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.BookMutation", m)
	}
	return f(ctx, mv)
}

// The LinFileFunc type is an adapter to allow the use of ordinary
// function as LinFile mutator.
type LinFileFunc func(context.Context, *model.LinFileMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f LinFileFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.LinFileMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.LinFileMutation", m)
	}
	return f(ctx, mv)
}

// The LinGroupFunc type is an adapter to allow the use of ordinary
// function as LinGroup mutator.
type LinGroupFunc func(context.Context, *model.LinGroupMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f LinGroupFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.LinGroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.LinGroupMutation", m)
	}
	return f(ctx, mv)
}

// The LinLogFunc type is an adapter to allow the use of ordinary
// function as LinLog mutator.
type LinLogFunc func(context.Context, *model.LinLogMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f LinLogFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.LinLogMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.LinLogMutation", m)
	}
	return f(ctx, mv)
}

// The LinPermissionFunc type is an adapter to allow the use of ordinary
// function as LinPermission mutator.
type LinPermissionFunc func(context.Context, *model.LinPermissionMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f LinPermissionFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.LinPermissionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.LinPermissionMutation", m)
	}
	return f(ctx, mv)
}

// The LinUserFunc type is an adapter to allow the use of ordinary
// function as LinUser mutator.
type LinUserFunc func(context.Context, *model.LinUserMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f LinUserFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.LinUserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.LinUserMutation", m)
	}
	return f(ctx, mv)
}

// The LinUserIdentiyFunc type is an adapter to allow the use of ordinary
// function as LinUserIdentiy mutator.
type LinUserIdentiyFunc func(context.Context, *model.LinUserIdentiyMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f LinUserIdentiyFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	mv, ok := m.(*model.LinUserIdentiyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *model.LinUserIdentiyMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, model.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op model.Op) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk model.Hook, cond Condition) model.Hook {
	return func(next model.Mutator) model.Mutator {
		return model.MutateFunc(func(ctx context.Context, m model.Mutation) (model.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, model.Delete|model.Create)
//
func On(hk model.Hook, op model.Op) model.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, model.Update|model.UpdateOne)
//
func Unless(hk model.Hook, op model.Op) model.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) model.Hook {
	return func(model.Mutator) model.Mutator {
		return model.MutateFunc(func(context.Context, model.Mutation) (model.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []model.Hook {
//		return []model.Hook{
//			Reject(model.Delete|model.Update),
//		}
//	}
//
func Reject(op model.Op) model.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []model.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...model.Hook) Chain {
	return Chain{append([]model.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() model.Hook {
	return func(mutator model.Mutator) model.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...model.Hook) Chain {
	newHooks := make([]model.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
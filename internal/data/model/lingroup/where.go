// Code generated by entc, DO NOT EDIT.

package lingroup

import (
	"lin-cms-go/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Info applies equality check predicate on the "info" field. It's identical to InfoEQ.
func Info(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfo), v))
	})
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int8) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.LinGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.LinGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// InfoEQ applies the EQ predicate on the "info" field.
func InfoEQ(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfo), v))
	})
}

// InfoNEQ applies the NEQ predicate on the "info" field.
func InfoNEQ(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInfo), v))
	})
}

// InfoIn applies the In predicate on the "info" field.
func InfoIn(vs ...string) predicate.LinGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInfo), v...))
	})
}

// InfoNotIn applies the NotIn predicate on the "info" field.
func InfoNotIn(vs ...string) predicate.LinGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInfo), v...))
	})
}

// InfoGT applies the GT predicate on the "info" field.
func InfoGT(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInfo), v))
	})
}

// InfoGTE applies the GTE predicate on the "info" field.
func InfoGTE(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInfo), v))
	})
}

// InfoLT applies the LT predicate on the "info" field.
func InfoLT(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInfo), v))
	})
}

// InfoLTE applies the LTE predicate on the "info" field.
func InfoLTE(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInfo), v))
	})
}

// InfoContains applies the Contains predicate on the "info" field.
func InfoContains(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInfo), v))
	})
}

// InfoHasPrefix applies the HasPrefix predicate on the "info" field.
func InfoHasPrefix(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInfo), v))
	})
}

// InfoHasSuffix applies the HasSuffix predicate on the "info" field.
func InfoHasSuffix(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInfo), v))
	})
}

// InfoEqualFold applies the EqualFold predicate on the "info" field.
func InfoEqualFold(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInfo), v))
	})
}

// InfoContainsFold applies the ContainsFold predicate on the "info" field.
func InfoContainsFold(v string) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInfo), v))
	})
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int8) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int8) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLevel), v))
	})
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int8) predicate.LinGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLevel), v...))
	})
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int8) predicate.LinGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLevel), v...))
	})
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int8) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLevel), v))
	})
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int8) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLevel), v))
	})
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int8) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLevel), v))
	})
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int8) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLevel), v))
	})
}

// HasLinUser applies the HasEdge predicate on the "lin_user" edge.
func HasLinUser() predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LinUserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, LinUserTable, LinUserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLinUserWith applies the HasEdge predicate on the "lin_user" edge with a given conditions (other predicates).
func HasLinUserWith(preds ...predicate.LinUser) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LinUserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, LinUserTable, LinUserPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLinPermission applies the HasEdge predicate on the "lin_permission" edge.
func HasLinPermission() predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LinPermissionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, LinPermissionTable, LinPermissionPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLinPermissionWith applies the HasEdge predicate on the "lin_permission" edge with a given conditions (other predicates).
func HasLinPermissionWith(preds ...predicate.LinPermission) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LinPermissionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, LinPermissionTable, LinPermissionPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LinGroup) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LinGroup) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LinGroup) predicate.LinGroup {
	return predicate.LinGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}

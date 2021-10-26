// Code generated by entc, DO NOT EDIT.

package lingrouppermission

import (
	"lin-cms-go/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
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
func IDGT(id int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// PermissionID applies equality check predicate on the "permission_id" field. It's identical to PermissionIDEQ.
func PermissionID(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermissionID), v))
	})
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroupID), v))
	})
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int) predicate.LinGroupPermission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGroupID), v...))
	})
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int) predicate.LinGroupPermission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGroupID), v...))
	})
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGroupID), v))
	})
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGroupID), v))
	})
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGroupID), v))
	})
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGroupID), v))
	})
}

// PermissionIDEQ applies the EQ predicate on the "permission_id" field.
func PermissionIDEQ(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermissionID), v))
	})
}

// PermissionIDNEQ applies the NEQ predicate on the "permission_id" field.
func PermissionIDNEQ(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPermissionID), v))
	})
}

// PermissionIDIn applies the In predicate on the "permission_id" field.
func PermissionIDIn(vs ...int) predicate.LinGroupPermission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPermissionID), v...))
	})
}

// PermissionIDNotIn applies the NotIn predicate on the "permission_id" field.
func PermissionIDNotIn(vs ...int) predicate.LinGroupPermission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPermissionID), v...))
	})
}

// PermissionIDGT applies the GT predicate on the "permission_id" field.
func PermissionIDGT(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPermissionID), v))
	})
}

// PermissionIDGTE applies the GTE predicate on the "permission_id" field.
func PermissionIDGTE(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPermissionID), v))
	})
}

// PermissionIDLT applies the LT predicate on the "permission_id" field.
func PermissionIDLT(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPermissionID), v))
	})
}

// PermissionIDLTE applies the LTE predicate on the "permission_id" field.
func PermissionIDLTE(v int) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPermissionID), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LinGroupPermission) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LinGroupPermission) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
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
func Not(p predicate.LinGroupPermission) predicate.LinGroupPermission {
	return predicate.LinGroupPermission(func(s *sql.Selector) {
		p(s.Not())
	})
}

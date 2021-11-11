// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"lin-cms-go/internal/data/model/linfile"
	"lin-cms-go/internal/data/model/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinFileQuery is the builder for querying LinFile entities.
type LinFileQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.LinFile
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LinFileQuery builder.
func (lfq *LinFileQuery) Where(ps ...predicate.LinFile) *LinFileQuery {
	lfq.predicates = append(lfq.predicates, ps...)
	return lfq
}

// Limit adds a limit step to the query.
func (lfq *LinFileQuery) Limit(limit int) *LinFileQuery {
	lfq.limit = &limit
	return lfq
}

// Offset adds an offset step to the query.
func (lfq *LinFileQuery) Offset(offset int) *LinFileQuery {
	lfq.offset = &offset
	return lfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lfq *LinFileQuery) Unique(unique bool) *LinFileQuery {
	lfq.unique = &unique
	return lfq
}

// Order adds an order step to the query.
func (lfq *LinFileQuery) Order(o ...OrderFunc) *LinFileQuery {
	lfq.order = append(lfq.order, o...)
	return lfq
}

// First returns the first LinFile entity from the query.
// Returns a *NotFoundError when no LinFile was found.
func (lfq *LinFileQuery) First(ctx context.Context) (*LinFile, error) {
	nodes, err := lfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{linfile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lfq *LinFileQuery) FirstX(ctx context.Context) *LinFile {
	node, err := lfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LinFile ID from the query.
// Returns a *NotFoundError when no LinFile ID was found.
func (lfq *LinFileQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{linfile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lfq *LinFileQuery) FirstIDX(ctx context.Context) int {
	id, err := lfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Last returns the last LinFile entity from the query.
// Returns a *NotFoundError when no LinFile was found.
func (lfq *LinFileQuery) Last(ctx context.Context) (*LinFile, error) {
	nodes, err := lfq.All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{linfile.Label}
	}
	return nodes[len(nodes)-1], nil
}

// LastX is like Last, but panics if an error occurs.
func (lfq *LinFileQuery) LastX(ctx context.Context) *LinFile {
	node, err := lfq.Last(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// LastID returns the last LinFile ID from the query.
// Returns a *NotFoundError when no LinFile ID was found.
func (lfq *LinFileQuery) LastID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lfq.IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{linfile.Label}
		return
	}
	return ids[len(ids)-1], nil
}

// LastIDX is like LastID, but panics if an error occurs.
func (lfq *LinFileQuery) LastIDX(ctx context.Context) int {
	id, err := lfq.LastID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LinFile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one LinFile entity is not found.
// Returns a *NotFoundError when no LinFile entities are found.
func (lfq *LinFileQuery) Only(ctx context.Context) (*LinFile, error) {
	nodes, err := lfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{linfile.Label}
	default:
		return nil, &NotSingularError{linfile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lfq *LinFileQuery) OnlyX(ctx context.Context) *LinFile {
	node, err := lfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LinFile ID in the query.
// Returns a *NotSingularError when exactly one LinFile ID is not found.
// Returns a *NotFoundError when no entities are found.
func (lfq *LinFileQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{linfile.Label}
	default:
		err = &NotSingularError{linfile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lfq *LinFileQuery) OnlyIDX(ctx context.Context) int {
	id, err := lfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LinFiles.
func (lfq *LinFileQuery) All(ctx context.Context) ([]*LinFile, error) {
	if err := lfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lfq *LinFileQuery) AllX(ctx context.Context) []*LinFile {
	nodes, err := lfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LinFile IDs.
func (lfq *LinFileQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := lfq.Select(linfile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lfq *LinFileQuery) IDsX(ctx context.Context) []int {
	ids, err := lfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lfq *LinFileQuery) Count(ctx context.Context) (int, error) {
	if err := lfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lfq *LinFileQuery) CountX(ctx context.Context) int {
	count, err := lfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lfq *LinFileQuery) Exist(ctx context.Context) (bool, error) {
	if err := lfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lfq *LinFileQuery) ExistX(ctx context.Context) bool {
	exist, err := lfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LinFileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lfq *LinFileQuery) Clone() *LinFileQuery {
	if lfq == nil {
		return nil
	}
	return &LinFileQuery{
		config:     lfq.config,
		limit:      lfq.limit,
		offset:     lfq.offset,
		order:      append([]OrderFunc{}, lfq.order...),
		predicates: append([]predicate.LinFile{}, lfq.predicates...),
		// clone intermediate query.
		sql:  lfq.sql.Clone(),
		path: lfq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Path string `json:"path,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LinFile.Query().
//		GroupBy(linfile.FieldPath).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (lfq *LinFileQuery) GroupBy(field string, fields ...string) *LinFileGroupBy {
	group := &LinFileGroupBy{config: lfq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lfq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Path string `json:"path,omitempty"`
//	}
//
//	client.LinFile.Query().
//		Select(linfile.FieldPath).
//		Scan(ctx, &v)
//
func (lfq *LinFileQuery) Select(fields ...string) *LinFileSelect {
	lfq.fields = append(lfq.fields, fields...)
	return &LinFileSelect{LinFileQuery: lfq}
}

func (lfq *LinFileQuery) prepareQuery(ctx context.Context) error {
	for _, f := range lfq.fields {
		if !linfile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if lfq.path != nil {
		prev, err := lfq.path(ctx)
		if err != nil {
			return err
		}
		lfq.sql = prev
	}
	return nil
}

func (lfq *LinFileQuery) sqlAll(ctx context.Context) ([]*LinFile, error) {
	var (
		nodes = []*LinFile{}
		_spec = lfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &LinFile{config: lfq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("model: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, lfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lfq *LinFileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lfq.querySpec()
	return sqlgraph.CountNodes(ctx, lfq.driver, _spec)
}

func (lfq *LinFileQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := lfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (lfq *LinFileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   linfile.Table,
			Columns: linfile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: linfile.FieldID,
			},
		},
		From:   lfq.sql,
		Unique: true,
	}
	if unique := lfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, linfile.FieldID)
		for i := range fields {
			if fields[i] != linfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lfq *LinFileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lfq.driver.Dialect())
	t1 := builder.Table(linfile.Table)
	columns := lfq.fields
	if len(columns) == 0 {
		columns = linfile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lfq.sql != nil {
		selector = lfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range lfq.predicates {
		p(selector)
	}
	for _, p := range lfq.order {
		p(selector)
	}
	if offset := lfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LinFileGroupBy is the group-by builder for LinFile entities.
type LinFileGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lfgb *LinFileGroupBy) Aggregate(fns ...AggregateFunc) *LinFileGroupBy {
	lfgb.fns = append(lfgb.fns, fns...)
	return lfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (lfgb *LinFileGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lfgb.path(ctx)
	if err != nil {
		return err
	}
	lfgb.sql = query
	return lfgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lfgb *LinFileGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := lfgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (lfgb *LinFileGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(lfgb.fields) > 1 {
		return nil, errors.New("model: LinFileGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := lfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lfgb *LinFileGroupBy) StringsX(ctx context.Context) []string {
	v, err := lfgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lfgb *LinFileGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lfgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linfile.Label}
	default:
		err = fmt.Errorf("model: LinFileGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lfgb *LinFileGroupBy) StringX(ctx context.Context) string {
	v, err := lfgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (lfgb *LinFileGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(lfgb.fields) > 1 {
		return nil, errors.New("model: LinFileGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := lfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lfgb *LinFileGroupBy) IntsX(ctx context.Context) []int {
	v, err := lfgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lfgb *LinFileGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lfgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linfile.Label}
	default:
		err = fmt.Errorf("model: LinFileGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lfgb *LinFileGroupBy) IntX(ctx context.Context) int {
	v, err := lfgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (lfgb *LinFileGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(lfgb.fields) > 1 {
		return nil, errors.New("model: LinFileGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := lfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lfgb *LinFileGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := lfgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lfgb *LinFileGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lfgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linfile.Label}
	default:
		err = fmt.Errorf("model: LinFileGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lfgb *LinFileGroupBy) Float64X(ctx context.Context) float64 {
	v, err := lfgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (lfgb *LinFileGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(lfgb.fields) > 1 {
		return nil, errors.New("model: LinFileGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := lfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lfgb *LinFileGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := lfgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lfgb *LinFileGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lfgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linfile.Label}
	default:
		err = fmt.Errorf("model: LinFileGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lfgb *LinFileGroupBy) BoolX(ctx context.Context) bool {
	v, err := lfgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lfgb *LinFileGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lfgb.fields {
		if !linfile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lfgb *LinFileGroupBy) sqlQuery() *sql.Selector {
	selector := lfgb.sql.Select()
	aggregation := make([]string, 0, len(lfgb.fns))
	for _, fn := range lfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lfgb.fields)+len(lfgb.fns))
		for _, f := range lfgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lfgb.fields...)...)
}

// LinFileSelect is the builder for selecting fields of LinFile entities.
type LinFileSelect struct {
	*LinFileQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (lfs *LinFileSelect) Scan(ctx context.Context, v interface{}) error {
	if err := lfs.prepareQuery(ctx); err != nil {
		return err
	}
	lfs.sql = lfs.LinFileQuery.sqlQuery(ctx)
	return lfs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lfs *LinFileSelect) ScanX(ctx context.Context, v interface{}) {
	if err := lfs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (lfs *LinFileSelect) Strings(ctx context.Context) ([]string, error) {
	if len(lfs.fields) > 1 {
		return nil, errors.New("model: LinFileSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := lfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lfs *LinFileSelect) StringsX(ctx context.Context) []string {
	v, err := lfs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (lfs *LinFileSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lfs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linfile.Label}
	default:
		err = fmt.Errorf("model: LinFileSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lfs *LinFileSelect) StringX(ctx context.Context) string {
	v, err := lfs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (lfs *LinFileSelect) Ints(ctx context.Context) ([]int, error) {
	if len(lfs.fields) > 1 {
		return nil, errors.New("model: LinFileSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := lfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lfs *LinFileSelect) IntsX(ctx context.Context) []int {
	v, err := lfs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (lfs *LinFileSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lfs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linfile.Label}
	default:
		err = fmt.Errorf("model: LinFileSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lfs *LinFileSelect) IntX(ctx context.Context) int {
	v, err := lfs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (lfs *LinFileSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(lfs.fields) > 1 {
		return nil, errors.New("model: LinFileSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := lfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lfs *LinFileSelect) Float64sX(ctx context.Context) []float64 {
	v, err := lfs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (lfs *LinFileSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lfs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linfile.Label}
	default:
		err = fmt.Errorf("model: LinFileSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lfs *LinFileSelect) Float64X(ctx context.Context) float64 {
	v, err := lfs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (lfs *LinFileSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(lfs.fields) > 1 {
		return nil, errors.New("model: LinFileSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := lfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lfs *LinFileSelect) BoolsX(ctx context.Context) []bool {
	v, err := lfs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (lfs *LinFileSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lfs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linfile.Label}
	default:
		err = fmt.Errorf("model: LinFileSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lfs *LinFileSelect) BoolX(ctx context.Context) bool {
	v, err := lfs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lfs *LinFileSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := lfs.sql.Query()
	if err := lfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
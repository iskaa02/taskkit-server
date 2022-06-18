// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/iskaa02/taskkit-server/ent/list"
	"github.com/iskaa02/taskkit-server/ent/predicate"
	"github.com/iskaa02/taskkit-server/ent/task"
	"github.com/iskaa02/taskkit-server/ent/theme"
)

// ListQuery is the builder for querying List entities.
type ListQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.List
	// eager-loading edges.
	withTheme *ThemeQuery
	withTasks *TaskQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ListQuery builder.
func (lq *ListQuery) Where(ps ...predicate.List) *ListQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit adds a limit step to the query.
func (lq *ListQuery) Limit(limit int) *ListQuery {
	lq.limit = &limit
	return lq
}

// Offset adds an offset step to the query.
func (lq *ListQuery) Offset(offset int) *ListQuery {
	lq.offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *ListQuery) Unique(unique bool) *ListQuery {
	lq.unique = &unique
	return lq
}

// Order adds an order step to the query.
func (lq *ListQuery) Order(o ...OrderFunc) *ListQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryTheme chains the current query on the "theme" edge.
func (lq *ListQuery) QueryTheme() *ThemeQuery {
	query := &ThemeQuery{config: lq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(list.Table, list.FieldID, selector),
			sqlgraph.To(theme.Table, theme.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, list.ThemeTable, list.ThemeColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTasks chains the current query on the "tasks" edge.
func (lq *ListQuery) QueryTasks() *TaskQuery {
	query := &TaskQuery{config: lq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(list.Table, list.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, list.TasksTable, list.TasksColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first List entity from the query.
// Returns a *NotFoundError when no List was found.
func (lq *ListQuery) First(ctx context.Context) (*List, error) {
	nodes, err := lq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{list.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *ListQuery) FirstX(ctx context.Context) *List {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first List ID from the query.
// Returns a *NotFoundError when no List ID was found.
func (lq *ListQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = lq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{list.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *ListQuery) FirstIDX(ctx context.Context) string {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single List entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one List entity is found.
// Returns a *NotFoundError when no List entities are found.
func (lq *ListQuery) Only(ctx context.Context) (*List, error) {
	nodes, err := lq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{list.Label}
	default:
		return nil, &NotSingularError{list.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *ListQuery) OnlyX(ctx context.Context) *List {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only List ID in the query.
// Returns a *NotSingularError when more than one List ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *ListQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = lq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{list.Label}
	default:
		err = &NotSingularError{list.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *ListQuery) OnlyIDX(ctx context.Context) string {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Lists.
func (lq *ListQuery) All(ctx context.Context) ([]*List, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lq *ListQuery) AllX(ctx context.Context) []*List {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of List IDs.
func (lq *ListQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := lq.Select(list.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *ListQuery) IDsX(ctx context.Context) []string {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *ListQuery) Count(ctx context.Context) (int, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lq *ListQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *ListQuery) Exist(ctx context.Context) (bool, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *ListQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ListQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *ListQuery) Clone() *ListQuery {
	if lq == nil {
		return nil
	}
	return &ListQuery{
		config:     lq.config,
		limit:      lq.limit,
		offset:     lq.offset,
		order:      append([]OrderFunc{}, lq.order...),
		predicates: append([]predicate.List{}, lq.predicates...),
		withTheme:  lq.withTheme.Clone(),
		withTasks:  lq.withTasks.Clone(),
		// clone intermediate query.
		sql:    lq.sql.Clone(),
		path:   lq.path,
		unique: lq.unique,
	}
}

// WithTheme tells the query-builder to eager-load the nodes that are connected to
// the "theme" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *ListQuery) WithTheme(opts ...func(*ThemeQuery)) *ListQuery {
	query := &ThemeQuery{config: lq.config}
	for _, opt := range opts {
		opt(query)
	}
	lq.withTheme = query
	return lq
}

// WithTasks tells the query-builder to eager-load the nodes that are connected to
// the "tasks" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *ListQuery) WithTasks(opts ...func(*TaskQuery)) *ListQuery {
	query := &TaskQuery{config: lq.config}
	for _, opt := range opts {
		opt(query)
	}
	lq.withTasks = query
	return lq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.List.Query().
//		GroupBy(list.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (lq *ListQuery) GroupBy(field string, fields ...string) *ListGroupBy {
	group := &ListGroupBy{config: lq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.List.Query().
//		Select(list.FieldName).
//		Scan(ctx, &v)
//
func (lq *ListQuery) Select(fields ...string) *ListSelect {
	lq.fields = append(lq.fields, fields...)
	return &ListSelect{ListQuery: lq}
}

func (lq *ListQuery) prepareQuery(ctx context.Context) error {
	for _, f := range lq.fields {
		if !list.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *ListQuery) sqlAll(ctx context.Context) ([]*List, error) {
	var (
		nodes       = []*List{}
		_spec       = lq.querySpec()
		loadedTypes = [2]bool{
			lq.withTheme != nil,
			lq.withTasks != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &List{config: lq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := lq.withTheme; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*List)
		for i := range nodes {
			fk := nodes[i].ThemeID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(theme.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "theme_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Theme = n
			}
		}
	}

	if query := lq.withTasks; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*List)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Tasks = []*Task{}
		}
		query.Where(predicate.Task(func(s *sql.Selector) {
			s.Where(sql.InValues(list.TasksColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ListID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "list_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.Tasks = append(node.Edges.Tasks, n)
		}
	}

	return nodes, nil
}

func (lq *ListQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.fields
	if len(lq.fields) > 0 {
		_spec.Unique = lq.unique != nil && *lq.unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *ListQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := lq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (lq *ListQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   list.Table,
			Columns: list.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: list.FieldID,
			},
		},
		From:   lq.sql,
		Unique: true,
	}
	if unique := lq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, list.FieldID)
		for i := range fields {
			if fields[i] != list.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *ListQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(list.Table)
	columns := lq.fields
	if len(columns) == 0 {
		columns = list.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.unique != nil && *lq.unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ListGroupBy is the group-by builder for List entities.
type ListGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *ListGroupBy) Aggregate(fns ...AggregateFunc) *ListGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the group-by query and scans the result into the given value.
func (lgb *ListGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lgb.path(ctx)
	if err != nil {
		return err
	}
	lgb.sql = query
	return lgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lgb *ListGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := lgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (lgb *ListGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(lgb.fields) > 1 {
		return nil, errors.New("ent: ListGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := lgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lgb *ListGroupBy) StringsX(ctx context.Context) []string {
	v, err := lgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lgb *ListGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{list.Label}
	default:
		err = fmt.Errorf("ent: ListGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lgb *ListGroupBy) StringX(ctx context.Context) string {
	v, err := lgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (lgb *ListGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(lgb.fields) > 1 {
		return nil, errors.New("ent: ListGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := lgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lgb *ListGroupBy) IntsX(ctx context.Context) []int {
	v, err := lgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lgb *ListGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{list.Label}
	default:
		err = fmt.Errorf("ent: ListGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lgb *ListGroupBy) IntX(ctx context.Context) int {
	v, err := lgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (lgb *ListGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(lgb.fields) > 1 {
		return nil, errors.New("ent: ListGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := lgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lgb *ListGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := lgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lgb *ListGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{list.Label}
	default:
		err = fmt.Errorf("ent: ListGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lgb *ListGroupBy) Float64X(ctx context.Context) float64 {
	v, err := lgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (lgb *ListGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(lgb.fields) > 1 {
		return nil, errors.New("ent: ListGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := lgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lgb *ListGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := lgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lgb *ListGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{list.Label}
	default:
		err = fmt.Errorf("ent: ListGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lgb *ListGroupBy) BoolX(ctx context.Context) bool {
	v, err := lgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lgb *ListGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lgb.fields {
		if !list.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lgb *ListGroupBy) sqlQuery() *sql.Selector {
	selector := lgb.sql.Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lgb.fields)+len(lgb.fns))
		for _, f := range lgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lgb.fields...)...)
}

// ListSelect is the builder for selecting fields of List entities.
type ListSelect struct {
	*ListQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ls *ListSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	ls.sql = ls.ListQuery.sqlQuery(ctx)
	return ls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ls *ListSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ls *ListSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ls.fields) > 1 {
		return nil, errors.New("ent: ListSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ls *ListSelect) StringsX(ctx context.Context) []string {
	v, err := ls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ls *ListSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{list.Label}
	default:
		err = fmt.Errorf("ent: ListSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ls *ListSelect) StringX(ctx context.Context) string {
	v, err := ls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ls *ListSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ls.fields) > 1 {
		return nil, errors.New("ent: ListSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ls *ListSelect) IntsX(ctx context.Context) []int {
	v, err := ls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ls *ListSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{list.Label}
	default:
		err = fmt.Errorf("ent: ListSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ls *ListSelect) IntX(ctx context.Context) int {
	v, err := ls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ls *ListSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ls.fields) > 1 {
		return nil, errors.New("ent: ListSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ls *ListSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ls *ListSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{list.Label}
	default:
		err = fmt.Errorf("ent: ListSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ls *ListSelect) Float64X(ctx context.Context) float64 {
	v, err := ls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ls *ListSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ls.fields) > 1 {
		return nil, errors.New("ent: ListSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ls *ListSelect) BoolsX(ctx context.Context) []bool {
	v, err := ls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ls *ListSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{list.Label}
	default:
		err = fmt.Errorf("ent: ListSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ls *ListSelect) BoolX(ctx context.Context) bool {
	v, err := ls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ls *ListSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ls.sql.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

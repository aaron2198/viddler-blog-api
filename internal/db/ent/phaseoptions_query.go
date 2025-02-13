// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/db/ent/phaseoptions"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/db/ent/predicate"
	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/db/ent/useroptions"
)

// PhaseOptionsQuery is the builder for querying PhaseOptions entities.
type PhaseOptionsQuery struct {
	config
	ctx             *QueryContext
	order           []phaseoptions.OrderOption
	inters          []Interceptor
	predicates      []predicate.PhaseOptions
	withUserOptions *UserOptionsQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PhaseOptionsQuery builder.
func (poq *PhaseOptionsQuery) Where(ps ...predicate.PhaseOptions) *PhaseOptionsQuery {
	poq.predicates = append(poq.predicates, ps...)
	return poq
}

// Limit the number of records to be returned by this query.
func (poq *PhaseOptionsQuery) Limit(limit int) *PhaseOptionsQuery {
	poq.ctx.Limit = &limit
	return poq
}

// Offset to start from.
func (poq *PhaseOptionsQuery) Offset(offset int) *PhaseOptionsQuery {
	poq.ctx.Offset = &offset
	return poq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (poq *PhaseOptionsQuery) Unique(unique bool) *PhaseOptionsQuery {
	poq.ctx.Unique = &unique
	return poq
}

// Order specifies how the records should be ordered.
func (poq *PhaseOptionsQuery) Order(o ...phaseoptions.OrderOption) *PhaseOptionsQuery {
	poq.order = append(poq.order, o...)
	return poq
}

// QueryUserOptions chains the current query on the "user_options" edge.
func (poq *PhaseOptionsQuery) QueryUserOptions() *UserOptionsQuery {
	query := (&UserOptionsClient{config: poq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := poq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := poq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(phaseoptions.Table, phaseoptions.FieldID, selector),
			sqlgraph.To(useroptions.Table, useroptions.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, phaseoptions.UserOptionsTable, phaseoptions.UserOptionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(poq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PhaseOptions entity from the query.
// Returns a *NotFoundError when no PhaseOptions was found.
func (poq *PhaseOptionsQuery) First(ctx context.Context) (*PhaseOptions, error) {
	nodes, err := poq.Limit(1).All(setContextOp(ctx, poq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{phaseoptions.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (poq *PhaseOptionsQuery) FirstX(ctx context.Context) *PhaseOptions {
	node, err := poq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PhaseOptions ID from the query.
// Returns a *NotFoundError when no PhaseOptions ID was found.
func (poq *PhaseOptionsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = poq.Limit(1).IDs(setContextOp(ctx, poq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{phaseoptions.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (poq *PhaseOptionsQuery) FirstIDX(ctx context.Context) int {
	id, err := poq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PhaseOptions entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PhaseOptions entity is found.
// Returns a *NotFoundError when no PhaseOptions entities are found.
func (poq *PhaseOptionsQuery) Only(ctx context.Context) (*PhaseOptions, error) {
	nodes, err := poq.Limit(2).All(setContextOp(ctx, poq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{phaseoptions.Label}
	default:
		return nil, &NotSingularError{phaseoptions.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (poq *PhaseOptionsQuery) OnlyX(ctx context.Context) *PhaseOptions {
	node, err := poq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PhaseOptions ID in the query.
// Returns a *NotSingularError when more than one PhaseOptions ID is found.
// Returns a *NotFoundError when no entities are found.
func (poq *PhaseOptionsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = poq.Limit(2).IDs(setContextOp(ctx, poq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{phaseoptions.Label}
	default:
		err = &NotSingularError{phaseoptions.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (poq *PhaseOptionsQuery) OnlyIDX(ctx context.Context) int {
	id, err := poq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PhaseOptionsSlice.
func (poq *PhaseOptionsQuery) All(ctx context.Context) ([]*PhaseOptions, error) {
	ctx = setContextOp(ctx, poq.ctx, ent.OpQueryAll)
	if err := poq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PhaseOptions, *PhaseOptionsQuery]()
	return withInterceptors[[]*PhaseOptions](ctx, poq, qr, poq.inters)
}

// AllX is like All, but panics if an error occurs.
func (poq *PhaseOptionsQuery) AllX(ctx context.Context) []*PhaseOptions {
	nodes, err := poq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PhaseOptions IDs.
func (poq *PhaseOptionsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if poq.ctx.Unique == nil && poq.path != nil {
		poq.Unique(true)
	}
	ctx = setContextOp(ctx, poq.ctx, ent.OpQueryIDs)
	if err = poq.Select(phaseoptions.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (poq *PhaseOptionsQuery) IDsX(ctx context.Context) []int {
	ids, err := poq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (poq *PhaseOptionsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, poq.ctx, ent.OpQueryCount)
	if err := poq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, poq, querierCount[*PhaseOptionsQuery](), poq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (poq *PhaseOptionsQuery) CountX(ctx context.Context) int {
	count, err := poq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (poq *PhaseOptionsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, poq.ctx, ent.OpQueryExist)
	switch _, err := poq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (poq *PhaseOptionsQuery) ExistX(ctx context.Context) bool {
	exist, err := poq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PhaseOptionsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (poq *PhaseOptionsQuery) Clone() *PhaseOptionsQuery {
	if poq == nil {
		return nil
	}
	return &PhaseOptionsQuery{
		config:          poq.config,
		ctx:             poq.ctx.Clone(),
		order:           append([]phaseoptions.OrderOption{}, poq.order...),
		inters:          append([]Interceptor{}, poq.inters...),
		predicates:      append([]predicate.PhaseOptions{}, poq.predicates...),
		withUserOptions: poq.withUserOptions.Clone(),
		// clone intermediate query.
		sql:  poq.sql.Clone(),
		path: poq.path,
	}
}

// WithUserOptions tells the query-builder to eager-load the nodes that are connected to
// the "user_options" edge. The optional arguments are used to configure the query builder of the edge.
func (poq *PhaseOptionsQuery) WithUserOptions(opts ...func(*UserOptionsQuery)) *PhaseOptionsQuery {
	query := (&UserOptionsClient{config: poq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	poq.withUserOptions = query
	return poq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PhaseName string `json:"phase_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PhaseOptions.Query().
//		GroupBy(phaseoptions.FieldPhaseName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (poq *PhaseOptionsQuery) GroupBy(field string, fields ...string) *PhaseOptionsGroupBy {
	poq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PhaseOptionsGroupBy{build: poq}
	grbuild.flds = &poq.ctx.Fields
	grbuild.label = phaseoptions.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PhaseName string `json:"phase_name,omitempty"`
//	}
//
//	client.PhaseOptions.Query().
//		Select(phaseoptions.FieldPhaseName).
//		Scan(ctx, &v)
func (poq *PhaseOptionsQuery) Select(fields ...string) *PhaseOptionsSelect {
	poq.ctx.Fields = append(poq.ctx.Fields, fields...)
	sbuild := &PhaseOptionsSelect{PhaseOptionsQuery: poq}
	sbuild.label = phaseoptions.Label
	sbuild.flds, sbuild.scan = &poq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PhaseOptionsSelect configured with the given aggregations.
func (poq *PhaseOptionsQuery) Aggregate(fns ...AggregateFunc) *PhaseOptionsSelect {
	return poq.Select().Aggregate(fns...)
}

func (poq *PhaseOptionsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range poq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, poq); err != nil {
				return err
			}
		}
	}
	for _, f := range poq.ctx.Fields {
		if !phaseoptions.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if poq.path != nil {
		prev, err := poq.path(ctx)
		if err != nil {
			return err
		}
		poq.sql = prev
	}
	return nil
}

func (poq *PhaseOptionsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PhaseOptions, error) {
	var (
		nodes       = []*PhaseOptions{}
		withFKs     = poq.withFKs
		_spec       = poq.querySpec()
		loadedTypes = [1]bool{
			poq.withUserOptions != nil,
		}
	)
	if poq.withUserOptions != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, phaseoptions.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PhaseOptions).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PhaseOptions{config: poq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, poq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := poq.withUserOptions; query != nil {
		if err := poq.loadUserOptions(ctx, query, nodes, nil,
			func(n *PhaseOptions, e *UserOptions) { n.Edges.UserOptions = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (poq *PhaseOptionsQuery) loadUserOptions(ctx context.Context, query *UserOptionsQuery, nodes []*PhaseOptions, init func(*PhaseOptions), assign func(*PhaseOptions, *UserOptions)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PhaseOptions)
	for i := range nodes {
		if nodes[i].user_options_phase_options == nil {
			continue
		}
		fk := *nodes[i].user_options_phase_options
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(useroptions.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_options_phase_options" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (poq *PhaseOptionsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := poq.querySpec()
	_spec.Node.Columns = poq.ctx.Fields
	if len(poq.ctx.Fields) > 0 {
		_spec.Unique = poq.ctx.Unique != nil && *poq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, poq.driver, _spec)
}

func (poq *PhaseOptionsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(phaseoptions.Table, phaseoptions.Columns, sqlgraph.NewFieldSpec(phaseoptions.FieldID, field.TypeInt))
	_spec.From = poq.sql
	if unique := poq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if poq.path != nil {
		_spec.Unique = true
	}
	if fields := poq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, phaseoptions.FieldID)
		for i := range fields {
			if fields[i] != phaseoptions.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := poq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := poq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := poq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := poq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (poq *PhaseOptionsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(poq.driver.Dialect())
	t1 := builder.Table(phaseoptions.Table)
	columns := poq.ctx.Fields
	if len(columns) == 0 {
		columns = phaseoptions.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if poq.sql != nil {
		selector = poq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if poq.ctx.Unique != nil && *poq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range poq.predicates {
		p(selector)
	}
	for _, p := range poq.order {
		p(selector)
	}
	if offset := poq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := poq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PhaseOptionsGroupBy is the group-by builder for PhaseOptions entities.
type PhaseOptionsGroupBy struct {
	selector
	build *PhaseOptionsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pogb *PhaseOptionsGroupBy) Aggregate(fns ...AggregateFunc) *PhaseOptionsGroupBy {
	pogb.fns = append(pogb.fns, fns...)
	return pogb
}

// Scan applies the selector query and scans the result into the given value.
func (pogb *PhaseOptionsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pogb.build.ctx, ent.OpQueryGroupBy)
	if err := pogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PhaseOptionsQuery, *PhaseOptionsGroupBy](ctx, pogb.build, pogb, pogb.build.inters, v)
}

func (pogb *PhaseOptionsGroupBy) sqlScan(ctx context.Context, root *PhaseOptionsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pogb.fns))
	for _, fn := range pogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pogb.flds)+len(pogb.fns))
		for _, f := range *pogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PhaseOptionsSelect is the builder for selecting fields of PhaseOptions entities.
type PhaseOptionsSelect struct {
	*PhaseOptionsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pos *PhaseOptionsSelect) Aggregate(fns ...AggregateFunc) *PhaseOptionsSelect {
	pos.fns = append(pos.fns, fns...)
	return pos
}

// Scan applies the selector query and scans the result into the given value.
func (pos *PhaseOptionsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pos.ctx, ent.OpQuerySelect)
	if err := pos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PhaseOptionsQuery, *PhaseOptionsSelect](ctx, pos.PhaseOptionsQuery, pos, pos.inters, v)
}

func (pos *PhaseOptionsSelect) sqlScan(ctx context.Context, root *PhaseOptionsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pos.fns))
	for _, fn := range pos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

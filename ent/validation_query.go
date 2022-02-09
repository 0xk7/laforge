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
	"github.com/gen0cide/laforge/ent/agenttask"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/script"
	"github.com/gen0cide/laforge/ent/validation"
	"github.com/google/uuid"
)

// ValidationQuery is the builder for querying Validation entities.
type ValidationQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Validation
	// eager-loading edges.
	withValidationToAgentTask *AgentTaskQuery
	withValidationToScript    *ScriptQuery
	withFKs                   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ValidationQuery builder.
func (vq *ValidationQuery) Where(ps ...predicate.Validation) *ValidationQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit adds a limit step to the query.
func (vq *ValidationQuery) Limit(limit int) *ValidationQuery {
	vq.limit = &limit
	return vq
}

// Offset adds an offset step to the query.
func (vq *ValidationQuery) Offset(offset int) *ValidationQuery {
	vq.offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *ValidationQuery) Unique(unique bool) *ValidationQuery {
	vq.unique = &unique
	return vq
}

// Order adds an order step to the query.
func (vq *ValidationQuery) Order(o ...OrderFunc) *ValidationQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryValidationToAgentTask chains the current query on the "ValidationToAgentTask" edge.
func (vq *ValidationQuery) QueryValidationToAgentTask() *AgentTaskQuery {
	query := &AgentTaskQuery{config: vq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(validation.Table, validation.FieldID, selector),
			sqlgraph.To(agenttask.Table, agenttask.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, validation.ValidationToAgentTaskTable, validation.ValidationToAgentTaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryValidationToScript chains the current query on the "ValidationToScript" edge.
func (vq *ValidationQuery) QueryValidationToScript() *ScriptQuery {
	query := &ScriptQuery{config: vq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(validation.Table, validation.FieldID, selector),
			sqlgraph.To(script.Table, script.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, validation.ValidationToScriptTable, validation.ValidationToScriptColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Validation entity from the query.
// Returns a *NotFoundError when no Validation was found.
func (vq *ValidationQuery) First(ctx context.Context) (*Validation, error) {
	nodes, err := vq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{validation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *ValidationQuery) FirstX(ctx context.Context) *Validation {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Validation ID from the query.
// Returns a *NotFoundError when no Validation ID was found.
func (vq *ValidationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{validation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *ValidationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Validation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Validation entity is not found.
// Returns a *NotFoundError when no Validation entities are found.
func (vq *ValidationQuery) Only(ctx context.Context) (*Validation, error) {
	nodes, err := vq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{validation.Label}
	default:
		return nil, &NotSingularError{validation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *ValidationQuery) OnlyX(ctx context.Context) *Validation {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Validation ID in the query.
// Returns a *NotSingularError when exactly one Validation ID is not found.
// Returns a *NotFoundError when no entities are found.
func (vq *ValidationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{validation.Label}
	default:
		err = &NotSingularError{validation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *ValidationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Validations.
func (vq *ValidationQuery) All(ctx context.Context) ([]*Validation, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vq *ValidationQuery) AllX(ctx context.Context) []*Validation {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Validation IDs.
func (vq *ValidationQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := vq.Select(validation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *ValidationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *ValidationQuery) Count(ctx context.Context) (int, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vq *ValidationQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *ValidationQuery) Exist(ctx context.Context) (bool, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *ValidationQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ValidationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *ValidationQuery) Clone() *ValidationQuery {
	if vq == nil {
		return nil
	}
	return &ValidationQuery{
		config:                    vq.config,
		limit:                     vq.limit,
		offset:                    vq.offset,
		order:                     append([]OrderFunc{}, vq.order...),
		predicates:                append([]predicate.Validation{}, vq.predicates...),
		withValidationToAgentTask: vq.withValidationToAgentTask.Clone(),
		withValidationToScript:    vq.withValidationToScript.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithValidationToAgentTask tells the query-builder to eager-load the nodes that are connected to
// the "ValidationToAgentTask" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *ValidationQuery) WithValidationToAgentTask(opts ...func(*AgentTaskQuery)) *ValidationQuery {
	query := &AgentTaskQuery{config: vq.config}
	for _, opt := range opts {
		opt(query)
	}
	vq.withValidationToAgentTask = query
	return vq
}

// WithValidationToScript tells the query-builder to eager-load the nodes that are connected to
// the "ValidationToScript" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *ValidationQuery) WithValidationToScript(opts ...func(*ScriptQuery)) *ValidationQuery {
	query := &ScriptQuery{config: vq.config}
	for _, opt := range opts {
		opt(query)
	}
	vq.withValidationToScript = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HclID string `json:"hcl_id,omitempty" hcl:"id,label"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Validation.Query().
//		GroupBy(validation.FieldHclID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vq *ValidationQuery) GroupBy(field string, fields ...string) *ValidationGroupBy {
	group := &ValidationGroupBy{config: vq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HclID string `json:"hcl_id,omitempty" hcl:"id,label"`
//	}
//
//	client.Validation.Query().
//		Select(validation.FieldHclID).
//		Scan(ctx, &v)
//
func (vq *ValidationQuery) Select(fields ...string) *ValidationSelect {
	vq.fields = append(vq.fields, fields...)
	return &ValidationSelect{ValidationQuery: vq}
}

func (vq *ValidationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vq.fields {
		if !validation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *ValidationQuery) sqlAll(ctx context.Context) ([]*Validation, error) {
	var (
		nodes       = []*Validation{}
		withFKs     = vq.withFKs
		_spec       = vq.querySpec()
		loadedTypes = [2]bool{
			vq.withValidationToAgentTask != nil,
			vq.withValidationToScript != nil,
		}
	)
	if vq.withValidationToAgentTask != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, validation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Validation{config: vq.config}
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
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := vq.withValidationToAgentTask; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Validation)
		for i := range nodes {
			if nodes[i].agent_task_agent_task_to_validation == nil {
				continue
			}
			fk := *nodes[i].agent_task_agent_task_to_validation
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(agenttask.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "agent_task_agent_task_to_validation" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ValidationToAgentTask = n
			}
		}
	}

	if query := vq.withValidationToScript; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Validation)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ValidationToScript = []*Script{}
		}
		query.withFKs = true
		query.Where(predicate.Script(func(s *sql.Selector) {
			s.Where(sql.InValues(validation.ValidationToScriptColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.script_script_to_validation
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "script_script_to_validation" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "script_script_to_validation" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.ValidationToScript = append(node.Edges.ValidationToScript, n)
		}
	}

	return nodes, nil
}

func (vq *ValidationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *ValidationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (vq *ValidationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   validation.Table,
			Columns: validation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: validation.FieldID,
			},
		},
		From:   vq.sql,
		Unique: true,
	}
	if unique := vq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, validation.FieldID)
		for i := range fields {
			if fields[i] != validation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *ValidationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(validation.Table)
	columns := vq.fields
	if len(columns) == 0 {
		columns = validation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ValidationGroupBy is the group-by builder for Validation entities.
type ValidationGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *ValidationGroupBy) Aggregate(fns ...AggregateFunc) *ValidationGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vgb *ValidationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vgb.path(ctx)
	if err != nil {
		return err
	}
	vgb.sql = query
	return vgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vgb *ValidationGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := vgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *ValidationGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: ValidationGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vgb *ValidationGroupBy) StringsX(ctx context.Context) []string {
	v, err := vgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *ValidationGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{validation.Label}
	default:
		err = fmt.Errorf("ent: ValidationGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vgb *ValidationGroupBy) StringX(ctx context.Context) string {
	v, err := vgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *ValidationGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: ValidationGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vgb *ValidationGroupBy) IntsX(ctx context.Context) []int {
	v, err := vgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *ValidationGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{validation.Label}
	default:
		err = fmt.Errorf("ent: ValidationGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vgb *ValidationGroupBy) IntX(ctx context.Context) int {
	v, err := vgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *ValidationGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: ValidationGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vgb *ValidationGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := vgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *ValidationGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{validation.Label}
	default:
		err = fmt.Errorf("ent: ValidationGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vgb *ValidationGroupBy) Float64X(ctx context.Context) float64 {
	v, err := vgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (vgb *ValidationGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(vgb.fields) > 1 {
		return nil, errors.New("ent: ValidationGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := vgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vgb *ValidationGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := vgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vgb *ValidationGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{validation.Label}
	default:
		err = fmt.Errorf("ent: ValidationGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vgb *ValidationGroupBy) BoolX(ctx context.Context) bool {
	v, err := vgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vgb *ValidationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vgb.fields {
		if !validation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vgb *ValidationGroupBy) sqlQuery() *sql.Selector {
	selector := vgb.sql.Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vgb.fields)+len(vgb.fns))
		for _, f := range vgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vgb.fields...)...)
}

// ValidationSelect is the builder for selecting fields of Validation entities.
type ValidationSelect struct {
	*ValidationQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vs *ValidationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	vs.sql = vs.ValidationQuery.sqlQuery(ctx)
	return vs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vs *ValidationSelect) ScanX(ctx context.Context, v interface{}) {
	if err := vs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (vs *ValidationSelect) Strings(ctx context.Context) ([]string, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: ValidationSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vs *ValidationSelect) StringsX(ctx context.Context) []string {
	v, err := vs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (vs *ValidationSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{validation.Label}
	default:
		err = fmt.Errorf("ent: ValidationSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vs *ValidationSelect) StringX(ctx context.Context) string {
	v, err := vs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (vs *ValidationSelect) Ints(ctx context.Context) ([]int, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: ValidationSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vs *ValidationSelect) IntsX(ctx context.Context) []int {
	v, err := vs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (vs *ValidationSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{validation.Label}
	default:
		err = fmt.Errorf("ent: ValidationSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vs *ValidationSelect) IntX(ctx context.Context) int {
	v, err := vs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (vs *ValidationSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: ValidationSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vs *ValidationSelect) Float64sX(ctx context.Context) []float64 {
	v, err := vs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (vs *ValidationSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{validation.Label}
	default:
		err = fmt.Errorf("ent: ValidationSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vs *ValidationSelect) Float64X(ctx context.Context) float64 {
	v, err := vs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (vs *ValidationSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(vs.fields) > 1 {
		return nil, errors.New("ent: ValidationSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := vs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vs *ValidationSelect) BoolsX(ctx context.Context) []bool {
	v, err := vs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (vs *ValidationSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{validation.Label}
	default:
		err = fmt.Errorf("ent: ValidationSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vs *ValidationSelect) BoolX(ctx context.Context) bool {
	v, err := vs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vs *ValidationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vs.sql.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

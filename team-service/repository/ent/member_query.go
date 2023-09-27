// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"team-service/repository/ent/member"
	"team-service/repository/ent/predicate"
	"team-service/repository/ent/team"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberQuery is the builder for querying Member entities.
type MemberQuery struct {
	config
	ctx          *QueryContext
	order        []member.OrderOption
	inters       []Interceptor
	predicates   []predicate.Member
	withTeam     *TeamQuery
	withTeams    *TeamQuery
	withAssigned *MemberQuery
	withAssign   *MemberQuery
	withApproved *MemberQuery
	withApprove  *MemberQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MemberQuery builder.
func (mq *MemberQuery) Where(ps ...predicate.Member) *MemberQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MemberQuery) Limit(limit int) *MemberQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MemberQuery) Offset(offset int) *MemberQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MemberQuery) Unique(unique bool) *MemberQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MemberQuery) Order(o ...member.OrderOption) *MemberQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryTeam chains the current query on the "team" edge.
func (mq *MemberQuery) QueryTeam() *TeamQuery {
	query := (&TeamClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, member.TeamTable, member.TeamColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeams chains the current query on the "teams" edge.
func (mq *MemberQuery) QueryTeams() *TeamQuery {
	query := (&TeamClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, member.TeamsTable, member.TeamsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssigned chains the current query on the "assigned" edge.
func (mq *MemberQuery) QueryAssigned() *MemberQuery {
	query := (&MemberClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, member.AssignedTable, member.AssignedColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssign chains the current query on the "member" edge.
func (mq *MemberQuery) QueryAssign() *MemberQuery {
	query := (&MemberClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.AssignTable, member.AssignColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryApproved chains the current query on the "approved" edge.
func (mq *MemberQuery) QueryApproved() *MemberQuery {
	query := (&MemberClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, member.ApprovedTable, member.ApprovedColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryApprove chains the current query on the "approve" edge.
func (mq *MemberQuery) QueryApprove() *MemberQuery {
	query := (&MemberClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.ApproveTable, member.ApproveColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Member entity from the query.
// Returns a *NotFoundError when no Member was found.
func (mq *MemberQuery) First(ctx context.Context) (*Member, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{member.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MemberQuery) FirstX(ctx context.Context) *Member {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Member ID from the query.
// Returns a *NotFoundError when no Member ID was found.
func (mq *MemberQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{member.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MemberQuery) FirstIDX(ctx context.Context) string {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Member entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Member entity is found.
// Returns a *NotFoundError when no Member entities are found.
func (mq *MemberQuery) Only(ctx context.Context) (*Member, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{member.Label}
	default:
		return nil, &NotSingularError{member.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MemberQuery) OnlyX(ctx context.Context) *Member {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Member ID in the query.
// Returns a *NotSingularError when more than one Member ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MemberQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = &NotSingularError{member.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MemberQuery) OnlyIDX(ctx context.Context) string {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Members.
func (mq *MemberQuery) All(ctx context.Context) ([]*Member, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Member, *MemberQuery]()
	return withInterceptors[[]*Member](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MemberQuery) AllX(ctx context.Context) []*Member {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Member IDs.
func (mq *MemberQuery) IDs(ctx context.Context) (ids []string, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(member.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MemberQuery) IDsX(ctx context.Context) []string {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MemberQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MemberQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MemberQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MemberQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MemberQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MemberQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MemberQuery) Clone() *MemberQuery {
	if mq == nil {
		return nil
	}
	return &MemberQuery{
		config:       mq.config,
		ctx:          mq.ctx.Clone(),
		order:        append([]member.OrderOption{}, mq.order...),
		inters:       append([]Interceptor{}, mq.inters...),
		predicates:   append([]predicate.Member{}, mq.predicates...),
		withTeam:     mq.withTeam.Clone(),
		withTeams:    mq.withTeams.Clone(),
		withAssigned: mq.withAssigned.Clone(),
		withAssign:   mq.withAssign.Clone(),
		withApproved: mq.withApproved.Clone(),
		withApprove:  mq.withApprove.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithTeam tells the query-builder to eager-load the nodes that are connected to
// the "team" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithTeam(opts ...func(*TeamQuery)) *MemberQuery {
	query := (&TeamClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withTeam = query
	return mq
}

// WithTeams tells the query-builder to eager-load the nodes that are connected to
// the "teams" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithTeams(opts ...func(*TeamQuery)) *MemberQuery {
	query := (&TeamClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withTeams = query
	return mq
}

// WithAssigned tells the query-builder to eager-load the nodes that are connected to
// the "assigned" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithAssigned(opts ...func(*MemberQuery)) *MemberQuery {
	query := (&MemberClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withAssigned = query
	return mq
}

// WithAssign tells the query-builder to eager-load the nodes that are connected to
// the "member" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithAssign(opts ...func(*MemberQuery)) *MemberQuery {
	query := (&MemberClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withAssign = query
	return mq
}

// WithApproved tells the query-builder to eager-load the nodes that are connected to
// the "approved" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithApproved(opts ...func(*MemberQuery)) *MemberQuery {
	query := (&MemberClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withApproved = query
	return mq
}

// WithApprove tells the query-builder to eager-load the nodes that are connected to
// the "approve" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithApprove(opts ...func(*MemberQuery)) *MemberQuery {
	query := (&MemberClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withApprove = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TeamID string `json:"teamId"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Member.Query().
//		GroupBy(member.FieldTeamID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MemberQuery) GroupBy(field string, fields ...string) *MemberGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MemberGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = member.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TeamID string `json:"teamId"`
//	}
//
//	client.Member.Query().
//		Select(member.FieldTeamID).
//		Scan(ctx, &v)
func (mq *MemberQuery) Select(fields ...string) *MemberSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MemberSelect{MemberQuery: mq}
	sbuild.label = member.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MemberSelect configured with the given aggregations.
func (mq *MemberQuery) Aggregate(fns ...AggregateFunc) *MemberSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MemberQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !member.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MemberQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Member, error) {
	var (
		nodes       = []*Member{}
		_spec       = mq.querySpec()
		loadedTypes = [6]bool{
			mq.withTeam != nil,
			mq.withTeams != nil,
			mq.withAssigned != nil,
			mq.withAssign != nil,
			mq.withApproved != nil,
			mq.withApprove != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Member).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Member{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withTeam; query != nil {
		if err := mq.loadTeam(ctx, query, nodes,
			func(n *Member) { n.Edges.Team = []*Team{} },
			func(n *Member, e *Team) { n.Edges.Team = append(n.Edges.Team, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withTeams; query != nil {
		if err := mq.loadTeams(ctx, query, nodes, nil,
			func(n *Member, e *Team) { n.Edges.Teams = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withAssigned; query != nil {
		if err := mq.loadAssigned(ctx, query, nodes, nil,
			func(n *Member, e *Member) { n.Edges.Assigned = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withAssign; query != nil {
		if err := mq.loadAssign(ctx, query, nodes,
			func(n *Member) { n.Edges.Assign = []*Member{} },
			func(n *Member, e *Member) { n.Edges.Assign = append(n.Edges.Assign, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withApproved; query != nil {
		if err := mq.loadApproved(ctx, query, nodes, nil,
			func(n *Member, e *Member) { n.Edges.Approved = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withApprove; query != nil {
		if err := mq.loadApprove(ctx, query, nodes,
			func(n *Member) { n.Edges.Approve = []*Member{} },
			func(n *Member, e *Member) { n.Edges.Approve = append(n.Edges.Approve, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MemberQuery) loadTeam(ctx context.Context, query *TeamQuery, nodes []*Member, init func(*Member), assign func(*Member, *Team)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(team.FieldCreatedBy)
	}
	query.Where(predicate.Team(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(member.TeamColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CreatedBy
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "created_by" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MemberQuery) loadTeams(ctx context.Context, query *TeamQuery, nodes []*Member, init func(*Member), assign func(*Member, *Team)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Member)
	for i := range nodes {
		if nodes[i].TeamID == nil {
			continue
		}
		fk := *nodes[i].TeamID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(team.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "team_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MemberQuery) loadAssigned(ctx context.Context, query *MemberQuery, nodes []*Member, init func(*Member), assign func(*Member, *Member)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Member)
	for i := range nodes {
		if nodes[i].AssignedBy == nil {
			continue
		}
		fk := *nodes[i].AssignedBy
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(member.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "assigned_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MemberQuery) loadAssign(ctx context.Context, query *MemberQuery, nodes []*Member, init func(*Member), assign func(*Member, *Member)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(member.FieldAssignedBy)
	}
	query.Where(predicate.Member(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(member.AssignColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AssignedBy
		if fk == nil {
			return fmt.Errorf(`foreign-key "assigned_by" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "assigned_by" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MemberQuery) loadApproved(ctx context.Context, query *MemberQuery, nodes []*Member, init func(*Member), assign func(*Member, *Member)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Member)
	for i := range nodes {
		if nodes[i].ApprovedBy == nil {
			continue
		}
		fk := *nodes[i].ApprovedBy
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(member.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "approved_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MemberQuery) loadApprove(ctx context.Context, query *MemberQuery, nodes []*Member, init func(*Member), assign func(*Member, *Member)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(member.FieldApprovedBy)
	}
	query.Where(predicate.Member(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(member.ApproveColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ApprovedBy
		if fk == nil {
			return fmt.Errorf(`foreign-key "approved_by" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "approved_by" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MemberQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MemberQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeString))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for i := range fields {
			if fields[i] != member.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mq.withTeams != nil {
			_spec.Node.AddColumnOnce(member.FieldTeamID)
		}
		if mq.withAssigned != nil {
			_spec.Node.AddColumnOnce(member.FieldAssignedBy)
		}
		if mq.withApproved != nil {
			_spec.Node.AddColumnOnce(member.FieldApprovedBy)
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MemberQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(member.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = member.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MemberGroupBy is the group-by builder for Member entities.
type MemberGroupBy struct {
	selector
	build *MemberQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MemberGroupBy) Aggregate(fns ...AggregateFunc) *MemberGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MemberGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberQuery, *MemberGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MemberGroupBy) sqlScan(ctx context.Context, root *MemberQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MemberSelect is the builder for selecting fields of Member entities.
type MemberSelect struct {
	*MemberQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MemberSelect) Aggregate(fns ...AggregateFunc) *MemberSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MemberSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberQuery, *MemberSelect](ctx, ms.MemberQuery, ms, ms.inters, v)
}

func (ms *MemberSelect) sqlScan(ctx context.Context, root *MemberQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"team-service/repository/ent/migrate"

	"team-service/repository/ent/member"
	"team-service/repository/ent/team"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Member is the client for interacting with the Member builders.
	Member *MemberClient
	// Team is the client for interacting with the Team builders.
	Team *TeamClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Member = NewMemberClient(c.config)
	c.Team = NewTeamClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Member: NewMemberClient(cfg),
		Team:   NewTeamClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Member: NewMemberClient(cfg),
		Team:   NewTeamClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Member.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In sorting to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Member.Use(hooks...)
	c.Team.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In sorting to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Member.Intercept(interceptors...)
	c.Team.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *MemberMutation:
		return c.Member.mutate(ctx, m)
	case *TeamMutation:
		return c.Team.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// MemberClient is a client for the Member schema.
type MemberClient struct {
	config
}

// NewMemberClient returns a client for the Member from the given config.
func NewMemberClient(c config) *MemberClient {
	return &MemberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `member.Hooks(f(g(h())))`.
func (c *MemberClient) Use(hooks ...Hook) {
	c.hooks.Member = append(c.hooks.Member, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `member.Intercept(f(g(h())))`.
func (c *MemberClient) Intercept(interceptors ...Interceptor) {
	c.inters.Member = append(c.inters.Member, interceptors...)
}

// Create returns a builder for creating a Member entity.
func (c *MemberClient) Create() *MemberCreate {
	mutation := newMemberMutation(c.config, OpCreate)
	return &MemberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Member entities.
func (c *MemberClient) CreateBulk(builders ...*MemberCreate) *MemberCreateBulk {
	return &MemberCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MemberClient) MapCreateBulk(slice any, setFunc func(*MemberCreate, int)) *MemberCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MemberCreateBulk{err: fmt.Errorf("calling to MemberClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MemberCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MemberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Member.
func (c *MemberClient) Update() *MemberUpdate {
	mutation := newMemberMutation(c.config, OpUpdate)
	return &MemberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MemberClient) UpdateOne(m *Member) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMember(m))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MemberClient) UpdateOneID(id string) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMemberID(id))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Member.
func (c *MemberClient) Delete() *MemberDelete {
	mutation := newMemberMutation(c.config, OpDelete)
	return &MemberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MemberClient) DeleteOne(m *Member) *MemberDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MemberClient) DeleteOneID(id string) *MemberDeleteOne {
	builder := c.Delete().Where(member.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MemberDeleteOne{builder}
}

// Query returns a query builder for Member.
func (c *MemberClient) Query() *MemberQuery {
	return &MemberQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMember},
		inters: c.Interceptors(),
	}
}

// Get returns a Member entity by its id.
func (c *MemberClient) Get(ctx context.Context, id string) (*Member, error) {
	return c.Query().Where(member.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MemberClient) GetX(ctx context.Context, id string) *Member {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTeam queries the team edge of a Member.
func (c *MemberClient) QueryTeam(m *Member) *TeamQuery {
	query := (&TeamClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, member.TeamTable, member.TeamColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTeams queries the teams edge of a Member.
func (c *MemberClient) QueryTeams(m *Member) *TeamQuery {
	query := (&TeamClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, member.TeamsTable, member.TeamsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAssigned queries the assigned edge of a Member.
func (c *MemberClient) QueryAssigned(m *Member) *MemberQuery {
	query := (&MemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, member.AssignedTable, member.AssignedColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAssign queries the member edge of a Member.
func (c *MemberClient) QueryAssign(m *Member) *MemberQuery {
	query := (&MemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.AssignTable, member.AssignColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryApproved queries the approved edge of a Member.
func (c *MemberClient) QueryApproved(m *Member) *MemberQuery {
	query := (&MemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, member.ApprovedTable, member.ApprovedColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryApprove queries the approve edge of a Member.
func (c *MemberClient) QueryApprove(m *Member) *MemberQuery {
	query := (&MemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.ApproveTable, member.ApproveColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MemberClient) Hooks() []Hook {
	return c.hooks.Member
}

// Interceptors returns the client interceptors.
func (c *MemberClient) Interceptors() []Interceptor {
	return c.inters.Member
}

func (c *MemberClient) mutate(ctx context.Context, m *MemberMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MemberCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MemberUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MemberDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Member mutation op: %q", m.Op())
	}
}

// TeamClient is a client for the Team schema.
type TeamClient struct {
	config
}

// NewTeamClient returns a client for the Team from the given config.
func NewTeamClient(c config) *TeamClient {
	return &TeamClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `team.Hooks(f(g(h())))`.
func (c *TeamClient) Use(hooks ...Hook) {
	c.hooks.Team = append(c.hooks.Team, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `team.Intercept(f(g(h())))`.
func (c *TeamClient) Intercept(interceptors ...Interceptor) {
	c.inters.Team = append(c.inters.Team, interceptors...)
}

// Create returns a builder for creating a Team entity.
func (c *TeamClient) Create() *TeamCreate {
	mutation := newTeamMutation(c.config, OpCreate)
	return &TeamCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Team entities.
func (c *TeamClient) CreateBulk(builders ...*TeamCreate) *TeamCreateBulk {
	return &TeamCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TeamClient) MapCreateBulk(slice any, setFunc func(*TeamCreate, int)) *TeamCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TeamCreateBulk{err: fmt.Errorf("calling to TeamClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TeamCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TeamCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Team.
func (c *TeamClient) Update() *TeamUpdate {
	mutation := newTeamMutation(c.config, OpUpdate)
	return &TeamUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TeamClient) UpdateOne(t *Team) *TeamUpdateOne {
	mutation := newTeamMutation(c.config, OpUpdateOne, withTeam(t))
	return &TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TeamClient) UpdateOneID(id string) *TeamUpdateOne {
	mutation := newTeamMutation(c.config, OpUpdateOne, withTeamID(id))
	return &TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Team.
func (c *TeamClient) Delete() *TeamDelete {
	mutation := newTeamMutation(c.config, OpDelete)
	return &TeamDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TeamClient) DeleteOne(t *Team) *TeamDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TeamClient) DeleteOneID(id string) *TeamDeleteOne {
	builder := c.Delete().Where(team.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TeamDeleteOne{builder}
}

// Query returns a query builder for Team.
func (c *TeamClient) Query() *TeamQuery {
	return &TeamQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTeam},
		inters: c.Interceptors(),
	}
}

// Get returns a Team entity by its id.
func (c *TeamClient) Get(ctx context.Context, id string) (*Team, error) {
	return c.Query().Where(team.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TeamClient) GetX(ctx context.Context, id string) *Team {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMembers queries the members edge of a Team.
func (c *TeamClient) QueryMembers(t *Team) *MemberQuery {
	query := (&MemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, team.MembersTable, team.MembersColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTeams queries the teams edge of a Team.
func (c *TeamClient) QueryTeams(t *Team) *MemberQuery {
	query := (&MemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, team.TeamsTable, team.TeamsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TeamClient) Hooks() []Hook {
	return c.hooks.Team
}

// Interceptors returns the client interceptors.
func (c *TeamClient) Interceptors() []Interceptor {
	return c.inters.Team
}

func (c *TeamClient) mutate(ctx context.Context, m *TeamMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TeamCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TeamUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TeamUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TeamDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Team mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Member, Team []ent.Hook
	}
	inters struct {
		Member, Team []ent.Interceptor
	}
)

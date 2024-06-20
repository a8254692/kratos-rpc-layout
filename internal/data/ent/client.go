// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent/account"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent/player"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Player is the client for interacting with the Player builders.
	Player *PlayerClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.Player = NewPlayerClient(c.config)
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

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

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
		ctx:     ctx,
		config:  cfg,
		Account: NewAccountClient(cfg),
		Player:  NewPlayerClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		Account: NewAccountClient(cfg),
		Player:  NewPlayerClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
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
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.Player.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Account.Intercept(interceptors...)
	c.Player.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AccountMutation:
		return c.Account.mutate(ctx, m)
	case *PlayerMutation:
		return c.Player.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `account.Intercept(f(g(h())))`.
func (c *AccountClient) Intercept(interceptors ...Interceptor) {
	c.inters.Account = append(c.inters.Account, interceptors...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AccountClient) MapCreateBulk(slice any, setFunc func(*AccountCreate, int)) *AccountCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AccountCreateBulk{err: fmt.Errorf("calling to AccountClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AccountCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id int64) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id int64) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAccount},
		inters: c.Interceptors(),
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id int64) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id int64) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// Interceptors returns the client interceptors.
func (c *AccountClient) Interceptors() []Interceptor {
	return c.inters.Account
}

func (c *AccountClient) mutate(ctx context.Context, m *AccountMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AccountCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AccountDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Account mutation op: %q", m.Op())
	}
}

// PlayerClient is a client for the Player schema.
type PlayerClient struct {
	config
}

// NewPlayerClient returns a client for the Player from the given config.
func NewPlayerClient(c config) *PlayerClient {
	return &PlayerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `player.Hooks(f(g(h())))`.
func (c *PlayerClient) Use(hooks ...Hook) {
	c.hooks.Player = append(c.hooks.Player, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `player.Intercept(f(g(h())))`.
func (c *PlayerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Player = append(c.inters.Player, interceptors...)
}

// Create returns a builder for creating a Player entity.
func (c *PlayerClient) Create() *PlayerCreate {
	mutation := newPlayerMutation(c.config, OpCreate)
	return &PlayerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Player entities.
func (c *PlayerClient) CreateBulk(builders ...*PlayerCreate) *PlayerCreateBulk {
	return &PlayerCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PlayerClient) MapCreateBulk(slice any, setFunc func(*PlayerCreate, int)) *PlayerCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PlayerCreateBulk{err: fmt.Errorf("calling to PlayerClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PlayerCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PlayerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Player.
func (c *PlayerClient) Update() *PlayerUpdate {
	mutation := newPlayerMutation(c.config, OpUpdate)
	return &PlayerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlayerClient) UpdateOne(pl *Player) *PlayerUpdateOne {
	mutation := newPlayerMutation(c.config, OpUpdateOne, withPlayer(pl))
	return &PlayerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlayerClient) UpdateOneID(id int64) *PlayerUpdateOne {
	mutation := newPlayerMutation(c.config, OpUpdateOne, withPlayerID(id))
	return &PlayerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Player.
func (c *PlayerClient) Delete() *PlayerDelete {
	mutation := newPlayerMutation(c.config, OpDelete)
	return &PlayerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PlayerClient) DeleteOne(pl *Player) *PlayerDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PlayerClient) DeleteOneID(id int64) *PlayerDeleteOne {
	builder := c.Delete().Where(player.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlayerDeleteOne{builder}
}

// Query returns a query builder for Player.
func (c *PlayerClient) Query() *PlayerQuery {
	return &PlayerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePlayer},
		inters: c.Interceptors(),
	}
}

// Get returns a Player entity by its id.
func (c *PlayerClient) Get(ctx context.Context, id int64) (*Player, error) {
	return c.Query().Where(player.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlayerClient) GetX(ctx context.Context, id int64) *Player {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PlayerClient) Hooks() []Hook {
	return c.hooks.Player
}

// Interceptors returns the client interceptors.
func (c *PlayerClient) Interceptors() []Interceptor {
	return c.inters.Player
}

func (c *PlayerClient) mutate(ctx context.Context, m *PlayerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PlayerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PlayerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PlayerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PlayerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Player mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Account, Player []ent.Hook
	}
	inters struct {
		Account, Player []ent.Interceptor
	}
)

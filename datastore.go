package polycode

import (
	"context"
	"time"
)

type WriteConfig struct {
	TTL    *int64
	ETag   *string
	Unsafe bool
	Upsert bool
}

type WriteOption func(*WriteConfig)

func WithExpireIn(expireIn time.Duration) WriteOption {
	expireInSeconds := int64(expireIn.Seconds())
	return func(cfg *WriteConfig) { cfg.TTL = &expireInSeconds }
}

func WithUnsafe() WriteOption {
	return func(cfg *WriteConfig) { cfg.Unsafe = true }
}

func WithUpsert() WriteOption {
	return func(cfg *WriteConfig) { cfg.Upsert = true }
}

type ReadOnlyDataStoreBuilder interface {
	WithTenantId(tenantId string) ReadOnlyDataStoreBuilder
	Get() ReadOnlyDataStore
}

type DataStoreBuilder interface {
	WithTenantId(tenantId string) DataStoreBuilder
	Get() DataStore
}

type ReadOnlyDataStore interface {
	Collection(name string) ReadOnlyCollection
}

type DataStore interface {
	Collection(name string) Collection
}

type ReadOnlyCollection interface {
	GetOne(id string) (ReadOnlyDoc, error)
	Query() ReadOnlyQuery

	Name() string
	Path() string
}

type Collection interface {
	GetOne(id string) (Doc, error)
	Query() Query
	InsertOne(id string, item interface{}, opts ...WriteOption) (Doc, error)

	Name() string
	Path() string
}

type ReadOnlyDoc interface {
	Unmarshal(item interface{}) error
	ChildCollection(name string) ReadOnlyCollection

	Parent() ReadOnlyDoc
	Collection() ReadOnlyCollection

	Id() string
	Path() string
}

type Doc interface {
	Unmarshal(item interface{}) error
	ExpireIn(expireIn time.Duration) error
	Update(item interface{}, opts ...WriteOption) error
	Delete(opts ...WriteOption) error
	ChildCollection(name string) Collection

	Parent() Doc
	Collection() Collection
	Id() string
	Path() string
}

type ReadOnlyDocList interface {
	Docs() []ReadOnlyDoc
}

type DocList interface {
	Docs() []Doc
}

type ReadOnlyQuery interface {
	Filter(expr string, args ...interface{}) ReadOnlyQuery
	Limit(limit int) ReadOnlyQuery
	GetOne(ctx context.Context) (ReadOnlyDoc, error)
	GetAll(ctx context.Context) (ReadOnlyDocList, error)

	// Optional future support
	// Iter() PagingIter
	// AllWithNextToken(ctx context.Context, ret interface{}) (PageToken, error)
	// Count(ctx context.Context) (int, error)
}

type Query interface {
	Filter(expr string, args ...interface{}) Query
	Limit(limit int) Query
	GetOne(ctx context.Context) (Doc, error)
	GetAll(ctx context.Context) (DocList, error)

	// Optional future support
	// Iter() PagingIter
	// AllWithNextToken(ctx context.Context, ret interface{}) (PageToken, error)
	// Count(ctx context.Context) (int, error)
}

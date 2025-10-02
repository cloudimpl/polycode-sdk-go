package polycode

import (
	"context"
	"time"
)

type WriteConfig struct {
	TTL    *time.Duration
	ETag   *string
	Unsafe bool
	Upsert bool
}

type WriteOption func(*WriteConfig)

func WithTTL(ttl time.Duration) WriteOption {
	return func(cfg *WriteConfig) { cfg.TTL = &ttl }
}

func WithUnsafe() WriteOption {
	return func(cfg *WriteConfig) { cfg.Unsafe = true }
}

func WithUpsert() WriteOption {
	return func(cfg *WriteConfig) { cfg.Upsert = true }
}

type ReadOnlyDoc interface {
	Unmarshal(item interface{}) error
	ExpireIn(expireIn time.Duration) error
	Update(item interface{}, opts ...WriteOption) error
	Delete(opts ...WriteOption) error
	Collection(name string) ReadOnlyCollection
	Parent() (ReadOnlyDoc, error)
	Path() (string, error)
}

type Doc interface {
	Unmarshal(item interface{}) error
	ExpireIn(expireIn time.Duration) error
	Update(item interface{}, opts ...WriteOption) error
	Delete(opts ...WriteOption) error
	Collection(name string) Collection
	Parent() (Doc, error)
	Path() (string, error)
}

type ReadOnlyCollection interface {
	GetOne(id string) (ReadOnlyDoc, error)
	Query() ReadOnlyQuery
}

type Collection interface {
	GetOne(id string) (Doc, error)
	Query() Query
	InsertOne(id string, item interface{}, opts ...WriteOption) (Doc, error)
}

type ReadOnlyDataStore interface {
	Collection(name string) ReadOnlyCollection
}

type DataStore interface {
	Collection(name string) Collection
}

type ReadOnlyDataStoreBuilder interface {
	WithTenantId(tenantId string) ReadOnlyDataStoreBuilder
	Get() ReadOnlyDataStore
}

type DataStoreBuilder interface {
	WithTenantId(tenantId string) DataStoreBuilder
	Get() DataStore
}

//type PageToken string

//type Iter interface {
//	Next(ctx context.Context, out interface{}) bool
//	Err() error
//}

//type PagingIter interface {
//	Iter
//	NextToken(ctx context.Context) (PageToken, error)
//}

type ReadOnlyQuery interface {
	Filter(expr string, args ...interface{}) ReadOnlyQuery
	Limit(limit int) ReadOnlyQuery
	GetOne(ctx context.Context) (ReadOnlyDoc, error)
	GetAll(ctx context.Context) ([]ReadOnlyDoc, error)

	// Optional future support
	// Iter() PagingIter
	// AllWithNextToken(ctx context.Context, ret interface{}) (PageToken, error)
	// Count(ctx context.Context) (int, error)
}

type Query interface {
	Filter(expr string, args ...interface{}) Query
	Limit(limit int) Query
	GetOne(ctx context.Context) (Doc, error)
	GetAll(ctx context.Context) ([]Doc, error)

	// Optional future support
	// Iter() PagingIter
	// AllWithNextToken(ctx context.Context, ret interface{}) (PageToken, error)
	// Count(ctx context.Context) (int, error)
}

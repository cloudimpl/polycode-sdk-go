package polycode

import (
	"context"
	"time"
)

type ReadOnlyDoc interface {
	Unmarshal(item interface{}) error
	ExpireIn(expireIn time.Duration) error
	Update(item interface{}) error
	Delete() error
	Collection(name string) ReadOnlyCollection
}

type Doc interface {
	Unmarshal(item interface{}) error
	ExpireIn(expireIn time.Duration) error
	Update(item interface{}) error
	Delete() error
	Collection(name string) Collection
}

type ReadOnlyCollection interface {
	GetOne(id string) (ReadOnlyDoc, bool, error)
	Query() ReadOnlyQuery
}

type Collection interface {
	GetOne(id string) (Doc, bool, error)
	Query() Query
	InsertOne(id string, item interface{}) (Doc, error)
	UpdateOne(id string, item interface{}) (Doc, error)
	UpsertOne(id string, item interface{}) (Doc, error)
	DeleteOne(id string) (Doc, error)
}

type ReadOnlyDataStore interface {
	Collection(name string) ReadOnlyCollection
	GlobalCollection(name string) ReadOnlyCollection
}

type DataStore interface {
	Collection(name string) Collection
	GlobalCollection(name string) Collection
}

type ReadOnlyDataStoreBuilder interface {
	WithTenantId(tenantId string) ReadOnlyDataStoreBuilder
	Get() ReadOnlyDataStore
}

type DataStoreBuilder interface {
	WithTenantId(tenantId string) DataStoreBuilder
	Get() DataStore
}

type PageToken string

type Iter interface {
	Next(ctx context.Context, out interface{}) bool
	Err() error
}

type PagingIter interface {
	Iter
	NextToken(ctx context.Context) (PageToken, error)
}

type ReadOnlyQuery interface {
	Filter(expr string, args ...interface{}) Query
	Limit(limit int) Query
	GetOne(ctx context.Context) (ReadOnlyDoc, bool, error)
	GetAll(ctx context.Context) ([]ReadOnlyDoc, error)

	// Optional future support
	// Iter() PagingIter
	// AllWithNextToken(ctx context.Context, ret interface{}) (PageToken, error)
	// Count(ctx context.Context) (int, error)
}

type Query interface {
	Filter(expr string, args ...interface{}) Query
	Limit(limit int) Query
	GetOne(ctx context.Context) (Doc, bool, error)
	GetAll(ctx context.Context) ([]Doc, error)

	// Optional future support
	// Iter() PagingIter
	// AllWithNextToken(ctx context.Context, ret interface{}) (PageToken, error)
	// Count(ctx context.Context) (int, error)
}

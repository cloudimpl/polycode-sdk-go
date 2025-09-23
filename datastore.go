package polycode

import "time"

type ReadOnlyCollection interface {
	GetOne(key string, ret interface{}) (bool, error)
	Query() Query
}

type Collection interface {
	ReadOnlyCollection
	InsertOne(item interface{}) error
	InsertOneWithTTL(item interface{}, expireIn time.Duration) error
	UpdateOne(item interface{}) error
	UpdateOneWithTTL(item interface{}, expireIn time.Duration) error
	UpsertOne(item interface{}) error
	UpsertOneWithTTL(item interface{}, expireIn time.Duration) error
	DeleteOne(key string) error
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
	WithPartitionKey(partitionKey string) ReadOnlyDataStoreBuilder
	Get() ReadOnlyDataStore
}

type DataStoreBuilder interface {
	WithTenantId(tenantId string) DataStoreBuilder
	WithPartitionKey(partitionKey string) DataStoreBuilder
	Get() DataStore
}

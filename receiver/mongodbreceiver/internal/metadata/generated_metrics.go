// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/model/pdata"
)

// Type is the component type name.
const Type config.Type = "mongodbreceiver"

// MetricIntf is an interface to generically interact with generated metric.
type MetricIntf interface {
	Name() string
	New() pdata.Metric
	Init(metric pdata.Metric)
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name     string
	initFunc func(pdata.Metric)
}

// Name returns the metric name.
func (m *metricImpl) Name() string {
	return m.name
}

// New creates a metric object preinitialized.
func (m *metricImpl) New() pdata.Metric {
	metric := pdata.NewMetric()
	m.Init(metric)
	return metric
}

// Init initializes the provided metric object.
func (m *metricImpl) Init(metric pdata.Metric) {
	m.initFunc(metric)
}

type metricStruct struct {
	MongodbCacheOperations MetricIntf
	MongodbCollectionCount MetricIntf
	MongodbConnectionCount MetricIntf
	MongodbDataSize        MetricIntf
	MongodbExtentCount     MetricIntf
	MongodbGlobalLockTime  MetricIntf
	MongodbIndexCount      MetricIntf
	MongodbIndexSize       MetricIntf
	MongodbMemoryUsage     MetricIntf
	MongodbObjectCount     MetricIntf
	MongodbOperationCount  MetricIntf
	MongodbStorageSize     MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"mongodb.cache.operations",
		"mongodb.collection.count",
		"mongodb.connection.count",
		"mongodb.data.size",
		"mongodb.extent.count",
		"mongodb.global_lock.time",
		"mongodb.index.count",
		"mongodb.index.size",
		"mongodb.memory.usage",
		"mongodb.object.count",
		"mongodb.operation.count",
		"mongodb.storage.size",
	}
}

var metricsByName = map[string]MetricIntf{
	"mongodb.cache.operations": Metrics.MongodbCacheOperations,
	"mongodb.collection.count": Metrics.MongodbCollectionCount,
	"mongodb.connection.count": Metrics.MongodbConnectionCount,
	"mongodb.data.size":        Metrics.MongodbDataSize,
	"mongodb.extent.count":     Metrics.MongodbExtentCount,
	"mongodb.global_lock.time": Metrics.MongodbGlobalLockTime,
	"mongodb.index.count":      Metrics.MongodbIndexCount,
	"mongodb.index.size":       Metrics.MongodbIndexSize,
	"mongodb.memory.usage":     Metrics.MongodbMemoryUsage,
	"mongodb.object.count":     Metrics.MongodbObjectCount,
	"mongodb.operation.count":  Metrics.MongodbOperationCount,
	"mongodb.storage.size":     Metrics.MongodbStorageSize,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"mongodb.cache.operations",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.cache.operations")
			metric.SetDescription("The number of cache operations of the instance.")
			metric.SetUnit("{operations}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodb.collection.count",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.collection.count")
			metric.SetDescription("The number of collections.")
			metric.SetUnit("{collections}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodb.connection.count",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.connection.count")
			metric.SetDescription("The number of connections.")
			metric.SetUnit("{connections}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodb.data.size",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.data.size")
			metric.SetDescription("The size of the collection. Data compression does not affect this value.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodb.extent.count",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.extent.count")
			metric.SetDescription("The number of extents.")
			metric.SetUnit("{extents}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodb.global_lock.time",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.global_lock.time")
			metric.SetDescription("The time the global lock has been held.")
			metric.SetUnit("ms")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodb.index.count",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.index.count")
			metric.SetDescription("The number of indexes.")
			metric.SetUnit("{indexes}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodb.index.size",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.index.size")
			metric.SetDescription("Sum of the space allocated to all indexes in the database, including free index space.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodb.memory.usage",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.memory.usage")
			metric.SetDescription("The amount of memory used.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodb.object.count",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.object.count")
			metric.SetDescription("The number of objects.")
			metric.SetUnit("{objects}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodb.operation.count",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.operation.count")
			metric.SetDescription("The number of operations executed.")
			metric.SetUnit("{operations}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"mongodb.storage.size",
		func(metric pdata.Metric) {
			metric.SetName("mongodb.storage.size")
			metric.SetDescription("The total amount of storage allocated to this collection.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// ConnectionType (The status of the connection.)
	ConnectionType string
	// Database (The name of a database.)
	Database string
	// MemoryType (The type of memory used.)
	MemoryType string
	// Operation (The MongoDB operation being counted.)
	Operation string
	// Type (The result of a cache request.)
	Type string
}{
	"type",
	"database",
	"type",
	"operation",
	"type",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeConnectionType are the possible values that the attribute "connection_type" can have.
var AttributeConnectionType = struct {
	Active    string
	Available string
	Current   string
}{
	"active",
	"available",
	"current",
}

// AttributeMemoryType are the possible values that the attribute "memory_type" can have.
var AttributeMemoryType = struct {
	Resident string
	Virtual  string
}{
	"resident",
	"virtual",
}

// AttributeOperation are the possible values that the attribute "operation" can have.
var AttributeOperation = struct {
	Insert  string
	Query   string
	Update  string
	Delete  string
	Getmore string
	Command string
}{
	"insert",
	"query",
	"update",
	"delete",
	"getmore",
	"command",
}

// AttributeType are the possible values that the attribute "type" can have.
var AttributeType = struct {
	Hit  string
	Miss string
}{
	"hit",
	"miss",
}

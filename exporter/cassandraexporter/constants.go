package cassandraexporter

const (
	// language=SQL
	createDatabaseSQL = `CREATE KEYSPACE %s WITH REPLICATION = { 'class' : '%s', 'replication_factor' : %d };`
	// language=SQL
	createEventTypeSql = `CREATE TYPE IF NOT EXISTS %s.Events (Timestamp Date, Name text, Attributes map<text, text>);`
	// language=SQL
	createLinksTypeSql = `CREATE TYPE IF NOT EXISTS %s.Links (TraceId text, SpanId text, TraceState text, Attributes map<text, text>);`
	// language=SQL
	createSpanTableSQL = `CREATE TABLE IF NOT EXISTS %s.%s (TimeStamp DATE, TraceId text, SpanId text, ParentSpanId text, TraceState text, SpanName text, SpanKind text, ServiceName text, ResourceAttributes map<text, text>, SpanAttributes map<text, text>, Duration int, StatusCode text, StatusMessage text, Events frozen<Events>, Links frozen<Links>, PRIMARY KEY (SpanId)) WITH COMPRESSION = {'class': '%s'}`
	// language=SQL
	insertSpanSQL = `INSERT INTO %s.%s (timestamp, traceid, spanid, parentspanid, tracestate, spanname, spankind, servicename, resourceattributes, spanattributes, duration, statuscode, statusmessage) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	//language=SQL
	createLogTableSQL = `CREATE TABLE IF NOT EXISTS %s.%s (TimeStamp DATE, TraceId text, SpanId text, TraceFlags int, SeverityText text, SeverityNumber int, ServiceName text, Body text, ResourceAttributes map<text, text>, LogAttributes map<text, text>, PRIMARY KEY (SpanId, SeverityNumber)) WITH COMPRESSION = {'class': '%s'}`
	//language=SQL
	insertLogTableSQL = `INSERT INTO %s.%s (timestamp, traceid, spanid, traceflags, severitytext, severitynumber, servicename, body, resourceattributes, logattributes) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
)

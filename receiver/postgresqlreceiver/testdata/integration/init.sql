CREATE USER otel WITH PASSWORD 'otel';
GRANT SELECT ON pg_stat_database TO otel;

CREATE TABLE table1 (
    id serial PRIMARY KEY
);
CREATE TABLE table2 (
    id serial PRIMARY KEY
);

CREATE DATABASE otel2;
\c otel2
CREATE TABLE test1 (
    id serial PRIMARY KEY
);
CREATE TABLE test2 (
    id serial PRIMARY KEY
);

CREATE INDEX otelindex ON test1(id);
CREATE INDEX otel2index ON test2(id);

-- Generating usage of index
SELECT * FROM test2; 

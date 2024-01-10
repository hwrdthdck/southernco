#!/bin/sh
# Copyright The OpenTelemetry Authors
# SPDX-License-Identifier: Apache-2.0


redis-server /etc/redis-cluster-6379.conf && \
    redis-server /etc/redis-cluster-6380.conf && \
    redis-server /etc/redis-cluster-6381.conf && \
    redis-server /etc/redis-cluster-6382.conf && \
    redis-server /etc/redis-cluster-6383.conf && \
    redis-server /etc/redis-cluster-6384.conf && \
    redis-server /etc/redis-cluster-6385.conf

# Simple health check: ping all nodes in cluster
redis-cli -p 6379 ping
redis-cli -p 6380 ping
redis-cli -p 6381 ping
redis-cli -p 6382 ping
redis-cli -p 6383 ping
redis-cli -p 6384 ping
redis-cli -p 6385 ping

# Now, create and adopt a cluster.  DO NOT ADD 6385 YET
redis-cli --cluster create localhost:6379 localhost:6380 localhost:6381 localhost:6382 localhost:6383 localhost:6384 --cluster-replicas 1 --cluster-yes
while true; do
    if redis-cli -p 6379 cluster info | grep -q "cluster_state:ok" ; then
        break
    fi
    echo "awaiting for cluster to be ready"
    sleep 2
done

redis-cli --cluster add-node 127.0.0.1:6385 127.0.0.1:6379 --cluster-slave
sleep 5

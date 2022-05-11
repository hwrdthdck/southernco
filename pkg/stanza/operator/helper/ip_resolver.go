// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helper

import (
	"net"
	"sync"
	"time"
)

// cacheEntry keeps information about host and expiration time
type cacheEntry struct {
	hostname   string
	expireTime time.Time
}

const (
	defaultInvalidationInterval time.Duration = 5 * time.Minute
)

type IPResolver struct {
	cache                map[string]cacheEntry
	mutex                sync.RWMutex
	done                 chan bool
	stopped              bool
	invalidationInterval time.Duration
}

// Create new resolver
func NewIpResolver() *IPResolver {
	r := &IPResolver{
		cache:                make(map[string]cacheEntry),
		stopped:              false,
		done:                 make(chan bool),
		invalidationInterval: defaultInvalidationInterval,
	}
	r.start()
	return r
}

// Stop cache invalidation
func (r *IPResolver) Stop() {
	r.mutex.Lock()
	if r.stopped {
		r.mutex.Unlock()
		return
	}

	r.stopped = true
	r.mutex.Unlock()
	r.done <- true
}

// start runs cache invalidation every 5 minutes
func (r *IPResolver) start() {
	ticker := time.NewTicker(r.invalidationInterval)
	go func() {
		for {
			select {
			case <-r.done:
				ticker.Stop()
				return
			case <-ticker.C:
				r.mutex.Lock()
				r.invalidateCache()
				r.mutex.Unlock()
			}
		}
	}()
}

// invalidateCache removes not longer valid entries from cache
func (r *IPResolver) invalidateCache() {
	now := time.Now()
	for key, entry := range r.cache {
		if entry.expireTime.Before(now) {
			delete(r.cache, key)
		}
	}
}

// GetHostFromIp returns hostname for given ip
// It is taken from cache if exists,
// otherwise lookup is performed and result is put into cache
func (r *IPResolver) GetHostFromIp(ip string) (host string) {
	r.mutex.RLock()
	entry, ok := r.cache[ip]
	if ok {
		host = entry.hostname
		defer r.mutex.RUnlock()
		return host
	}
	r.mutex.RUnlock()

	host = r.lookupIpAddr(ip)

	r.mutex.Lock()
	r.cache[ip] = cacheEntry{
		hostname:   host,
		expireTime: time.Now().Add(5 * time.Minute),
	}
	r.mutex.Unlock()

	return host
}

// lookupIpAddr resturns hostname based on ip address
func (r *IPResolver) lookupIpAddr(ip string) (host string) {
	res, err := net.LookupAddr(ip)
	if err != nil || len(res) == 0 {
		return ip
	}

	host = res[0]
	// Trim one trailing '.'.
	if last := len(host) - 1; last >= 0 && host[last] == '.' {
		host = host[:last]
	}
	return host
}

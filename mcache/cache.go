// Package mcache provides a global memory cache which can be
// used to store necessary data in local memory. It
// uses freecache as underlying engine.
// For more details about freecache, please refer to
// https://github.com/coocood/freecache
package mcache

import (
	"fmt"
	"runtime/debug"
	"strconv"

	"github.com/coocood/freecache"
)

var c *freecache.Cache

const (
	// SizeKB means 1 kilobyte.
	SizeKB = 1024
	// SizeMB means 1 megabyte.
	SizeMB = 1024 * SizeKB
	// SizeGB means 1 gigabyte.
	SizeGB = 1024 * SizeMB
)

// Init is the first function that needs to be invoked when
// this package is used. It will initialized freecache according
// to the given parameters.
// memSize: the initial size of the cache in Byte. The memory will
// be preallocated. Example: 100 * SizeMB
// setGCPercent: whether to set GC percentage.
// gcPercent: the GC percentage to set.
// Memory is preallocated.If you allocate large amount of memory,
// you may need to set debug.SetGCPercent() to a much lower
// percentage to get a normal GC frequency.
func Init(memSize int, setGCPercent bool, gcPercent int) {
	fmt.Printf("local cache initialized, size: %d, setGCPercent: %v, GCPercent: %d\n",
		memSize, setGCPercent, gcPercent)
	c = freecache.NewCache(memSize)
	if !setGCPercent {
		return
	}

	debug.SetGCPercent(gcPercent)
}

// IsUsable can be used to check if c has been initialized.
func IsUsable() bool {
	if c != nil {
		return true
	}

	return false
}

// Set sets a key, value and expiration for a cache entry and stores it in the cache.
// If the key is larger than 65535 or value is larger than 1/1024 of the cache size,
// the entry will not be written to the cache. expireSeconds <= 0 means no expire,
// but it can be evicted when cache is full.
func Set(key string, value []byte, expireSeconds int) (err error) {
	return c.Set([]byte(key), value, expireSeconds)
}

// Get returns the value or not found error.
func Get(key string) (value []byte, err error) {
	return c.Get([]byte(key))
}

// SetInt stores in integer value in the cahe.
func SetInt(key string, value int, expireSeconds int) error {
	return Set(key, []byte(fmt.Sprintf("%d", value)), expireSeconds)
}

// GetInt returns the value for an integer within the cache or a not found error.
func GetInt(key string) (int, error) {
	raw, err := Get(key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(raw))
}

// SetString stores in string value in the cache.
func SetString(key, value string, expireSeconds int) error {
	return Set(key, []byte(value), expireSeconds)
}

// GetString will get the cached value and try to convert it to string.
func GetString(key string) (string, error) {
	raw, err := Get(key)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// GetWithBuf copies the value to the buf or returns not found error.
// This method doesn't allocate memory when the capacity of buf is greater or equal to value.
func GetWithBuf(key string, buf []byte) (value []byte, err error) {
	return c.GetWithBuf([]byte(key), buf)
}

// GetWithExpiration returns the value with expiration or not found error.
func GetWithExpiration(key string) (value []byte, expireAt uint32, err error) {
	return c.GetWithExpiration([]byte(key))
}

// TTL returns the TTL time left for a given key or a not found error.
func TTL(key string) (timeLeft uint32, err error) {
	return c.TTL([]byte(key))
}

// Del deletes an item in the cache by key and returns true or false if a delete occurred.
func Del(key string) (affected bool) {
	return c.Del([]byte(key))
}

// GetIntWithExpiration returns the value and expiration or a not found error.
func GetIntWithExpiration(key string) (value int, expireAt uint32, err error) {
	var v []byte
	v, expireAt, err = GetWithExpiration(key)
	if err != nil {
		return
	}
	value, err = strconv.Atoi(string(v))
	return
}

// EvacuateCount is a metric indicating the number of times an eviction occurred.
func EvacuateCount() (count int64) {
	return c.EvacuateCount()
}

// ExpiredCount is a metric indicating the number of times an expire occurred.
func ExpiredCount() (count int64) {
	return c.ExpiredCount()
}

// EntryCount returns the number of items currently in the cache.
func EntryCount() (entryCount int64) {
	return c.EntryCount()
}

// AverageAccessTime returns the average unix timestamp when a entry being accessed.
// Entries have greater access time will be evacuated when it
// is about to be overwritten by new value.
func AverageAccessTime() int64 {
	return c.AverageAccessTime()
}

// HitCount is a metric that returns number of times a key was found in the cache.
func HitCount() (count int64) {
	return c.HitCount()
}

// MissCount is a metric that returns the number of times a miss occurred in the cache.
func MissCount() (count int64) {
	return c.MissCount()
}

// LookupCount is a metric that returns the number of times a lookup for a given key occurred.
func LookupCount() int64 {
	return c.LookupCount()
}

// HitRate is the ratio of hits over lookups.
func HitRate() float64 {
	return c.HitRate()
}

// OverwriteCount indicates the number of times entries have been overriden.
func OverwriteCount() (overwriteCount int64) {
	return c.OverwriteCount()
}

// Clear clears the cache.
func Clear() {
	c.Clear()
}

// ResetStatistics refreshes the current state of the statistics.
func ResetStatistics() {
	c.ResetStatistics()
}

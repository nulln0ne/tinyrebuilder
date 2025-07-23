package tinyrebuilder

import (
	"fmt"
	"regexp"

	lru "github.com/hashicorp/golang-lru"
)

var (
	defaultCacheSize = 128
	cache, _         = lru.New(defaultCacheSize)
)

// MustCompileWithCache is like MustCompile but uses a package-level LRU cache
// to store and retrieve compiled regular expressions. This is highly recommended
// for performance in high-load applications where the same regex patterns are
// built frequently.
func (r *RegexBuilder) MustCompileWithCache() *Regexp {
	pattern := r.Build()
	if val, ok := cache.Get(pattern); ok {
		if re, ok := val.(*Regexp); ok {
			return re
		}
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	wrappedRe := &Regexp{re: re}
	cache.Add(pattern, wrappedRe)

	return wrappedRe
}

// PurgeCache completely clears the regex cache.
func PurgeCache() {
	cache.Purge()
}

// SetCacheSize changes the size of the LRU cache. Note that this will purge
// the existing cache.
func SetCacheSize(size int) error {
	if size <= 0 {
		return fmt.Errorf("cache size must be positive")
	}
	var err error
	cache, err = lru.New(size)
	if err != nil {
		// Fallback to default if something goes wrong
		cache, _ = lru.New(defaultCacheSize)
		return err
	}
	return nil
}

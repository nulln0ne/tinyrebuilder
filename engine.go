package tinyrebuilder

import (
	"regexp"
	"strings"
	"sync"
)

var builderPool = sync.Pool{
	New: func() any {
		return &strings.Builder{}
	},
}

// RegexBuilder is a fluent interface for building regular expressions.
type RegexBuilder struct {
	builder *strings.Builder
}

// New creates a new RegexBuilder from the pool.
func New() *RegexBuilder {
	return &RegexBuilder{builder: builderPool.Get().(*strings.Builder)}
}

// NewWithCapacity creates a new RegexBuilder with a given initial capacity.
// Note: This does not use the pool as a specific capacity is requested.
func NewWithCapacity(capacity int) *RegexBuilder {
	b := &strings.Builder{}
	b.Grow(capacity)
	return &RegexBuilder{builder: b}
}

// Build returns the string representation of the regular expression.
func (r *RegexBuilder) Build() string {
	return r.builder.String()
}

// Compile compiles the regular expression and returns the builder to the pool.
func (r *RegexBuilder) Compile() (*Regexp, error) {
	s := r.builder.String()
	r.builder.Reset()
	builderPool.Put(r.builder)

	re, err := regexp.Compile(s)
	if err != nil {
		return nil, err
	}
	return &Regexp{re: re}, nil
}

// MustCompile compiles the regular expression, panicking if it fails,
// and returns the builder to the pool.
func (r *RegexBuilder) MustCompile() *Regexp {
	re, err := r.Compile()
	if err != nil {
		panic(err)
	}
	return re
}

# Tiny Rebuilder

[![Go Reference](https://pkg.go.dev/badge/github.com/nulln0ne/tinyrebuilder.svg)](https://pkg.go.dev/github.com/nulln0ne/tinyrebuilder)

`tinyrebuilder` is a fluent, high-performance, and cacheable regular expression builder for Go. It wraps the standard library's `regexp` package, offering a more expressive and readable way to construct complex regex patterns while adding performance optimizations like caching.

## Features

- **Fluent Interface:** Chain methods together to build complex regex patterns in a readable way.
- **High Performance:** Uses `strings.Builder` to efficiently build regex strings.
- **Cacheable:** Built-in LRU cache for compiled regex patterns to avoid redundant compilations in high-load applications.
- **Pre-defined Patterns:** A collection of common regex patterns is available in the `patterns` sub-package.
- **Extensible:** Easily create your own reusable patterns.

## Installation

```bash
go get github.com/nulln0ne/tinyrebuilder
```

## Usage

Here's a simple example of how to build a regex to validate a username:

```go
package main

import (
	"fmt"
	"github.com/nulln0ne/tinyrebuilder"
)

func main() {
	// Build a regex that matches a username with 3-16 alphanumeric characters or underscores.
	usernameValidator := tinyrebuilder.New().
		StartAnchor().
		Raw(`[a-zA-Z0-9_]{3,16}`).
		EndAnchor().
		MustCompile()

	fmt.Println(usernameValidator.IsMatch("valid_user")) // true
	fmt.Println(usernameValidator.IsMatch("invalid-user")) // false
}
```

### Using Pre-defined Patterns

The `patterns` package provides a set of common regex patterns.

```go
package main

import (
	"fmt"
	"github.com/nulln0ne/tinyrebuilder/patterns"
)

func main() {
	emailValidator := patterns.Email().MustCompile()

	fmt.Println(emailValidator.IsMatch("test@example.com")) // true
	fmt.Println(emailValidator.IsMatch("not-an-email"))      // false
}
```

### Caching

For performance-critical applications, you can use the cached compiler. This is useful when you are repeatedly building the same regex pattern in different parts of your application.

```go
package main

import (
	"fmt"
	"github.com/nulln0ne/tinyrebuilder"
)

func main() {
	// The first call compiles and caches the regex.
	re1 := tinyrebuilder.New().Literal("some_pattern").MustCompileWithCache()

	// The second call with the same pattern will retrieve the regex from the cache.
	re2 := tinyrebuilder.New().Literal("some_pattern").MustCompileWithCache()

	// re1 and re2 are the same underlying regexp.Regexp object.
	fmt.Println(re1 == re2) // true

	// You can also purge the cache if needed.
	tinyrebuilder.PurgeCache()

	// After purging, a new regex object is compiled and cached.
	re3 := tinyrebuilder.New().Literal("some_pattern").MustCompileWithCache()

	fmt.Println(re1 == re3) // false
}
```

## Performance

The library is designed to be efficient. Here are some benchmark results to give you an idea of the performance and allocation overhead.

| Benchmark                                 | Operations | ns/op (lower is better) | B/op (lower is better) | allocs/op (lower is better) |
| ----------------------------------------- | ---------- | ----------------------- | ---------------------- | --------------------------- |
| `BenchmarkSimpleRegexCompilation`         | 523200     | 2262                    | 4713                   | 66                          |
| `BenchmarkComplexRegexCompilation`        | 590811     | 1796                    | 4041                   | 47                          |
| `BenchmarkEmailPatternCompilation`        | 436510     | 2719                    | 5697                   | 69                          |
| `BenchmarkEmailPatternCompilationWithCache` | 11769547   | 101.6                   | 104                    | 4                           |
| `BenchmarkURLPatternCompilation`          | 287083     | 4140                    | 9715                   | 105                         |

As you can see, using the cache (`MustCompileWithCache`) dramatically reduces allocations and improves performance for repeated compilations of the same pattern.

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License.
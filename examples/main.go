package main

import (
	"fmt"

	"github.com/nulln0ne/tinyrebuilder"
	"github.com/nulln0ne/tinyrebuilder/patterns"
)

func main() {
	// Example 1: Basic username validation
	usernameValidator := tinyrebuilder.New().
		StartAnchor().
		Raw(`[a-zA-Z0-9_]{3,16}`).
		EndAnchor().
		MustCompile()

	fmt.Println("--- Username Validation ---")
	fmt.Printf("Is 'valid_user' a valid username? %v\n", usernameValidator.IsMatch("valid_user"))
	fmt.Printf("Is 'invalid-user' a valid username? %v\n", usernameValidator.IsMatch("invalid-user"))
	fmt.Println()

	// Example 2: Using a pre-defined pattern for email validation
	emailValidator := patterns.Email().MustCompile()

	fmt.Println("--- Email Validation ---")
	fmt.Printf("Is 'test@example.com' a valid email? %v\n", emailValidator.IsMatch("test@example.com"))
	fmt.Printf("Is 'not-an-email' a valid email? %v\n", emailValidator.IsMatch("not-an-email"))
	fmt.Println()

	// Example 3: Building a more complex regex for a URL
	urlValidator := tinyrebuilder.New().
		StartAnchor().
		Literal("http").
		Maybe().
		Literal("s").
		Maybe().
		Literal("://").
		NonCapturingGroup(tinyrebuilder.New().Literal("www.")).Maybe().
		Raw(`[a-zA-Z0-9.-]+`).
		Literal(".").
		Raw(`[a-zA-Z]{2,}`).
		EndAnchor().
		MustCompile()

	fmt.Println("--- URL Validation ---")
	fmt.Printf("Is 'https://example.com' a valid URL? %v\n", urlValidator.IsMatch("https://example.com"))
	fmt.Printf("Is 'http://www.google.com' a valid URL? %v\n", urlValidator.IsMatch("http://www.google.com"))
	fmt.Printf("Is 'ftp://invalid.com' a valid URL? %v\n", urlValidator.IsMatch("ftp://invalid.com"))
	fmt.Println()

	// Example 4: Using the cached compiler
	fmt.Println("--- Cache Demonstration ---")
	// The first call compiles and caches the regex.
	re1 := tinyrebuilder.New().Literal("some_pattern").MustCompileWithCache()
	// The second call retrieves the same regex object from the cache.
	re2 := tinyrebuilder.New().Literal("some_pattern").MustCompileWithCache()
	fmt.Printf("Are re1 and re2 the same object? %v\n", re1 == re2)

	tinyrebuilder.PurgeCache()
	// After purging, a new regex object is compiled and cached.
	re3 := tinyrebuilder.New().Literal("some_pattern").MustCompileWithCache()
	fmt.Printf("Are re1 and re3 the same object after purging cache? %v\n", re1 == re3)
	fmt.Println()
}

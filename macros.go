package tinyrebuilder

// Digit is a helper that returns a new builder matching any digit (`\d`).
func Digit() *RegexBuilder {
	return New().Raw(`\d`)
}

// WordChar is a helper that returns a new builder matching any word character (`\w`).
func WordChar() *RegexBuilder {
	return New().Raw(`\w`)
}

// Whitespace is a helper that returns a new builder matching any whitespace character (`\s`).
func Whitespace() *RegexBuilder {
	return New().Raw(`\s`)
}

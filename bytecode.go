package tinyrebuilder

import (
	"regexp"
)

// Regexp is a wrapper around the standard library's *regexp.Regexp.
// It provides all the methods of the original, allowing it to be used as a
// drop-in replacement where a *regexp.Regexp is expected.
type Regexp struct {
	re *regexp.Regexp
}

// IsMatch checks if the compiled regular expression matches the string.
func (r *Regexp) IsMatch(s string) bool {
	return r.re.MatchString(s)
}

// FindStringSubmatch returns a slice of strings holding the text of the
// leftmost match of the regular expression in s and the matches, if any, for
// its subexpressions.
func (r *Regexp) FindStringSubmatch(s string) []string {
	return r.re.FindStringSubmatch(s)
}

// FindAllString finds all successive non-overlapping matches of the Regexp in a string.
func (r *Regexp) FindAllString(s string, n int) []string {
	return r.re.FindAllString(s, n)
}

// FindAllStringIndex finds all successive non-overlapping matches of the Regexp in a string
// and returns a slice of pairs of indices.
func (r *Regexp) FindAllStringIndex(s string, n int) [][]int {
	return r.re.FindAllStringIndex(s, n)
}

// FindAllStringSubmatch finds all successive non-overlapping matches of the Regexp in a string
// and returns a slice of slices of strings.
func (r *Regexp) FindAllStringSubmatch(s string, n int) [][]string {
	return r.re.FindAllStringSubmatch(s, n)
}

// FindString finds the text of the leftmost match in a string.
func (r *Regexp) FindString(s string) string {
	return r.re.FindString(s)
}

// FindStringIndex returns a two-element slice of integers defining the location of
// the leftmost match in a string.
func (r *Regexp) FindStringIndex(s string) []int {
	return r.re.FindStringIndex(s)
}

// SubexpNames returns the names of the parenthesized subexpressions in this Regexp.
func (r *Regexp) SubexpNames() []string {
	return r.re.SubexpNames()
}

// Expand appends template to dst and returns the result.
func (r *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte {
	return r.re.Expand(dst, template, src, match)
}

// ExpandString is like Expand but the template and source are strings.
func (r *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte {
	return r.re.ExpandString(dst, template, src, match)
}

// LiteralPrefix returns a literal string that must begin any match of the Regexp.
func (r *Regexp) LiteralPrefix() (prefix string, complete bool) {
	return r.re.LiteralPrefix()
}

// MatchString reports whether the Regexp matches the string s.
func (r *Regexp) MatchString(s string) bool {
	return r.re.MatchString(s)
}

// NumSubexp returns the number of parenthesized subexpressions in this Regexp.
func (r *Regexp) NumSubexp() int {
	return r.re.NumSubexp()
}

// String returns the source text of the regular expression.
func (r *Regexp) String() string {
	return r.re.String()
}

// Unwrap returns the underlying *regexp.Regexp object.
func (r *Regexp) Unwrap() *regexp.Regexp {
	return r.re
}

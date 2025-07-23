package tinyrebuilder

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	// Common character classes
	charClassWhitespace      = `\s`
	charClassNotWhitespace   = `\S`
	charClassDigit           = `\d`
	charClassNotDigit        = `\D`
	charClassWord            = `\w`
	charClassNotWord         = `\W`
	charClassWordBoundary    = `\b`
	charClassNotWordBoundary = `\B`

	// Other special characters
	charTab             = `\t`
	charNewline         = `\n`
	charCarriageReturn  = `\r`
	anchorStartOfString = `\A`
	anchorEndOfString   = `\z`
)

// Raw adds a raw string to the regular expression.
func (r *RegexBuilder) Raw(s string) *RegexBuilder {
	r.builder.WriteString(s)
	return r
}

// Literal adds a literal string to the regular expression, escaping any special characters.
func (r *RegexBuilder) Literal(s string) *RegexBuilder {
	r.builder.WriteString(regexp.QuoteMeta(s))
	return r
}

// StartAnchor adds a start anchor (^) to the regular expression.
func (r *RegexBuilder) StartAnchor() *RegexBuilder {
	r.builder.WriteString("^")
	return r
}

// EndAnchor adds an end anchor ($) to the regular expression.
func (r *RegexBuilder) EndAnchor() *RegexBuilder {
	r.builder.WriteString("$")
	return r
}

// StartOfString adds a start of string anchor (\A) to the regular expression.
func (r *RegexBuilder) StartOfString() *RegexBuilder {
	r.builder.WriteString(anchorStartOfString)
	return r
}

// EndOfString adds an end of string anchor (\z) to the regular expression.
func (r *RegexBuilder) EndOfString() *RegexBuilder {
	r.builder.WriteString(anchorEndOfString)
	return r
}

// Group creates a capturing group from another RegexBuilder.
func (r *RegexBuilder) Group(group *RegexBuilder) *RegexBuilder {
	r.builder.WriteString("(")
	r.builder.WriteString(group.Build())
	r.builder.WriteString(")")
	return r
}

// NonCapturingGroup creates a non-capturing group from another RegexBuilder.
func (r *RegexBuilder) NonCapturingGroup(group *RegexBuilder) *RegexBuilder {
	r.builder.WriteString("(?:")
	r.builder.WriteString(group.Build())
	r.builder.WriteString(")")
	return r
}

// NamedGroup creates a named capturing group from another RegexBuilder.
func (r *RegexBuilder) NamedGroup(name string, group *RegexBuilder) *RegexBuilder {
	r.builder.WriteString("(?P<")
	r.builder.WriteString(name)
	r.builder.WriteString(">")
	r.builder.WriteString(group.Build())
	r.builder.WriteString(")")

	return r
}

// Or creates an OR condition with other RegexBuilders.
// It joins the builders with a "|" and wraps them in a non-capturing group.
func (r *RegexBuilder) Or(groups ...*RegexBuilder) *RegexBuilder {
	if len(groups) == 0 {
		return r
	}

	// To avoid extra allocations, we create a new builder, build the Or expression,
	// then reset the original builder and write the new expression to it.
	// This is more efficient than reading the original builder's string.
	newBuilder := builderPool.Get().(*strings.Builder)
	newBuilder.WriteString("(?:")
	newBuilder.WriteString(r.builder.String())

	for _, group := range groups {
		newBuilder.WriteString("|")
		newBuilder.WriteString(group.Build())
	}
	newBuilder.WriteString(")")

	r.builder.Reset()
	r.builder.WriteString(newBuilder.String())

	newBuilder.Reset()
	builderPool.Put(newBuilder)

	return r
}

// Exactly matches the previous element exactly n times.
func (r *RegexBuilder) Exactly(n int) *RegexBuilder {
	r.builder.WriteString("{")
	r.builder.WriteString(strconv.Itoa(n))
	r.builder.WriteString("}")
	return r
}

// Maybe makes the previous element optional (zero or one time).
func (r *RegexBuilder) Maybe() *RegexBuilder {
	r.builder.WriteString("?")
	return r
}

// OneOrMore matches the previous element one or more times.
func (r *RegexBuilder) OneOrMore() *RegexBuilder {
	r.builder.WriteString("+")
	return r
}

// ZeroOrMore matches the previous element zero or more times.
func (r *RegexBuilder) ZeroOrMore() *RegexBuilder {
	r.builder.WriteString("*")
	return r
}

// AtLeast matches the previous element at least n times.
func (r *RegexBuilder) AtLeast(n int) *RegexBuilder {
	r.builder.WriteString("{")
	r.builder.WriteString(strconv.Itoa(n))
	r.builder.WriteString(",}")
	return r
}

// Between matches the previous element between n and m times.
func (r *RegexBuilder) Between(n, m int) *RegexBuilder {
	r.builder.WriteString("{")
	r.builder.WriteString(strconv.Itoa(n))
	r.builder.WriteString(",")
	r.builder.WriteString(strconv.Itoa(m))
	r.builder.WriteString("}")
	return r
}

// NonGreedy makes the previous quantifier non-greedy.
func (r *RegexBuilder) NonGreedy() *RegexBuilder {
	r.builder.WriteString("?")
	return r
}

// Whitespace adds a whitespace character class (`\s`) to the expression.
func (r *RegexBuilder) Whitespace() *RegexBuilder {
	r.builder.WriteString(charClassWhitespace)
	return r
}

// NotWhitespace adds a non-whitespace character class (`\S`) to the expression.
func (r *RegexBuilder) NotWhitespace() *RegexBuilder {
	r.builder.WriteString(charClassNotWhitespace)
	return r
}

// Digit adds a digit character class (`\d`) to the expression.
func (r *RegexBuilder) Digit() *RegexBuilder {
	r.builder.WriteString(charClassDigit)
	return r
}

// NotDigit adds a non-digit character class (`\D`) to the expression.
func (r *RegexBuilder) NotDigit() *RegexBuilder {
	r.builder.WriteString(charClassNotDigit)
	return r
}

// WordChar adds a word character class (`\w`) to the expression.
func (r *RegexBuilder) WordChar() *RegexBuilder {
	r.builder.WriteString(charClassWord)
	return r
}

// NotWordChar adds a non-word character class (`\W`) to the expression.
func (r *RegexBuilder) NotWordChar() *RegexBuilder {
	r.builder.WriteString(charClassNotWord)
	return r
}

// WordBoundary adds a word boundary (`\b`) to the expression.
func (r *RegexBuilder) WordBoundary() *RegexBuilder {
	r.builder.WriteString(charClassWordBoundary)
	return r
}

// NotWordBoundary adds a non-word boundary (`\B`) to the expression.
func (r *RegexBuilder) NotWordBoundary() *RegexBuilder {
	r.builder.WriteString(charClassNotWordBoundary)
	return r
}

// Tab adds a tab character (`\t`) to the expression.
func (r *RegexBuilder) Tab() *RegexBuilder {
	r.builder.WriteString(charTab)
	return r
}

// Newline adds a newline character (`\n`) to the expression.
func (r *RegexBuilder) Newline() *RegexBuilder {
	r.builder.WriteString(charNewline)
	return r
}

// CarriageReturn adds a carriage return character (`\r`) to the expression.
func (r *RegexBuilder) CarriageReturn() *RegexBuilder {
	r.builder.WriteString(charCarriageReturn)
	return r
}

// Quote escapes all special characters in the given string.
func (r *RegexBuilder) Quote(s string) *RegexBuilder {
	r.builder.WriteString(regexp.QuoteMeta(s))
	return r
}

// AnyOf creates a character set that matches any of the characters in the string.
func (r *RegexBuilder) AnyOf(s string) *RegexBuilder {
	r.builder.WriteString("[")
	r.builder.WriteString(s)
	r.builder.WriteString("]")
	return r
}

// NotAnyOf creates a negated character set that matches any character not in the string.
func (r *RegexBuilder) NotAnyOf(s string) *RegexBuilder {
	r.builder.WriteString("[^")
	r.builder.WriteString(s)
	r.builder.WriteString("]")
	return r
}

// Range creates a character range.
func (r *RegexBuilder) Range(from, to rune) *RegexBuilder {
	r.builder.WriteString("[")
	r.builder.WriteRune(from)
	r.builder.WriteString("-")
	r.builder.WriteRune(to)
	r.builder.WriteString("]")
	return r
}

// WithFlags adds flags to the expression.
func (r *RegexBuilder) WithFlags(flags string) *RegexBuilder {
	r.builder.WriteString("(?")
	r.builder.WriteString(flags)
	r.builder.WriteString(")")
	return r
}

// GroupWithFlags creates a group with flags.
func (r *RegexBuilder) GroupWithFlags(flags string, group *RegexBuilder) *RegexBuilder {
	r.builder.WriteString("(?")
	r.builder.WriteString(flags)
	r.builder.WriteString(":")
	r.builder.WriteString(group.Build())
	r.builder.WriteString(")")
	return r
}

// POSIXClass adds a POSIX character class (e.g., "[:alnum:]").
func (r *RegexBuilder) POSIXClass(class string) *RegexBuilder {
	r.builder.WriteString("[[:")
	r.builder.WriteString(class)
	r.builder.WriteString(":]]")
	return r
}

// NotPOSIXClass adds a negated POSIX character class (e.g., "[^[:alnum:]]").
func (r *RegexBuilder) NotPOSIXClass(class string) *RegexBuilder {
	r.builder.WriteString("[^[:")
	r.builder.WriteString(class)
	r.builder.WriteString(":]]")
	return r
}

// UnicodeProperty adds a Unicode character property (e.g., `\p{Greek}`).
func (r *RegexBuilder) UnicodeProperty(property string) *RegexBuilder {
	r.builder.WriteString(`\p{`)
	r.builder.WriteString(property)
	r.builder.WriteString(`}`)
	return r
}

// NotUnicodeProperty adds a negated Unicode character property (e.g., `\P{Greek}`).
func (r *RegexBuilder) NotUnicodeProperty(property string) *RegexBuilder {
	r.builder.WriteString(`\P{`)
	r.builder.WriteString(property)
	r.builder.WriteString(`}`)
	return r
}

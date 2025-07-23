package tinyrebuilder_test

import (
	"testing"

	"github.com/nulln0ne/tinyrebuilder"
	"github.com/nulln0ne/tinyrebuilder/patterns"
)

func TestBasicLiterals(t *testing.T) {
	re := tinyrebuilder.New().
		Literal("hello").
		MustCompile()
	if !re.IsMatch("hello") {
		t.Error("Expected to match 'hello'")
	}
	if re.IsMatch("Hello") {
		t.Error("Expected NOT to match 'Hello'")
	}
}

func TestCompileError(t *testing.T) {
	_, err := tinyrebuilder.New().Raw("(?P<name>a").Compile()
	if err == nil {
		t.Error("Expected an error for invalid regex, got nil")
	}
}

func TestNewWithCapacity(t *testing.T) {
	b := tinyrebuilder.NewWithCapacity(100)
	b.Literal("test")
	// This test mainly ensures the constructor works.
	// Actual capacity check is hard without exposing internal state.
	if b.Build() != "test" {
		t.Errorf("Expected build to work with pre-allocated capacity")
	}
}

func TestAnyOf(t *testing.T) {
	re := tinyrebuilder.New().
		AnyOf("abc").
		OneOrMore().
		MustCompile()
	if !re.IsMatch("abc") {
		t.Error("Expected to match 'abc'")
	}
	if re.IsMatch("def") {
		t.Error("Expected NOT to match 'def'")
	}
}

func TestNotAnyOf(t *testing.T) {
	re := tinyrebuilder.New().
		NotAnyOf("abc").
		OneOrMore().
		MustCompile()
	if !re.IsMatch("def") {
		t.Error("Expected to match 'def'")
	}
	if re.IsMatch("abc") {
		t.Error("Expected NOT to match 'abc'")
	}
}

func TestRange(t *testing.T) {
	re := tinyrebuilder.New().
		Range('a', 'z').
		OneOrMore().
		MustCompile()
	if !re.IsMatch("hello") {
		t.Error("Expected to match 'hello'")
	}
	if re.IsMatch("123") {
		t.Error("Expected NOT to match '123'")
	}
}

func TestWithFlags(t *testing.T) {
	re := tinyrebuilder.New().
		WithFlags("i").
		Literal("hello").
		MustCompile()
	if !re.IsMatch("Hello") {
		t.Error("Expected to match 'Hello' with case-insensitivity")
	}
}

func TestGroupsAndQuantifiers(t *testing.T) {
	re := tinyrebuilder.New().
		Group(tinyrebuilder.New().Literal("go").Maybe()).
		OneOrMore().
		MustCompile()
	if !re.IsMatch("go") {
		t.Error("Expected to match 'go'")
	}
	if !re.IsMatch("gogo") {
		t.Error("Expected to match 'gogo'")
	}
}

func TestOrOperator(t *testing.T) {
	re := tinyrebuilder.New().
		Literal("cat").
		Or(
			tinyrebuilder.New().Literal("dog"),
			tinyrebuilder.New().Literal("bird"),
		).MustCompile()

	tests := map[string]bool{
		"cat":  true,
		"dog":  true,
		"bird": true,
		"fish": false,
	}

	for str, want := range tests {
		got := re.IsMatch(str)
		if got != want {
			t.Errorf("IsMatch(%q) = %v; want %v", str, got, want)
		}
	}
}

func TestAnchors(t *testing.T) {
	re := tinyrebuilder.New().
		StartAnchor().
		Literal("abc").
		EndAnchor().
		MustCompile()

	if !re.IsMatch("abc") {
		t.Error("Expected to match full string 'abc'")
	}
	if re.IsMatch(" abc ") {
		t.Error("Expected NOT to match ' abc '")
	}
}

func TestNamedGroup(t *testing.T) {
	re := tinyrebuilder.New().
		NamedGroup("word", tinyrebuilder.New().AnyOf("a-zA-Z").OneOrMore()).
		MustCompile()

	matches := re.FindStringSubmatch("hello")
	if len(matches) < 2 {
		t.Fatal("Expected to find match in 'hello'")
	}
	if matches[1] != "hello" {
		t.Errorf("Expected named group 'word' to capture 'hello', but got %s", matches[1])
	}
}

func TestCharacterClasses(t *testing.T) {
	testCases := []struct {
		name    string
		builder *tinyrebuilder.RegexBuilder
		good    string
		bad     string
	}{
		{"Digit", tinyrebuilder.Digit(), "1", "a"},
		{"NotDigit", tinyrebuilder.New().NotDigit(), "a", "1"},
		{"WordChar", tinyrebuilder.WordChar(), "a", "-"},
		{"NotWordChar", tinyrebuilder.New().NotWordChar(), "-", "a"},
		{"Whitespace", tinyrebuilder.Whitespace(), " ", "a"},
		{"NotWhitespace", tinyrebuilder.New().NotWhitespace(), "a", " "},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			re := tc.builder.MustCompile()
			if !re.IsMatch(tc.good) {
				t.Errorf("Expected to match %q", tc.good)
			}
			if re.IsMatch(tc.bad) {
				t.Errorf("Expected NOT to match %q", tc.bad)
			}
		})
	}
}

func TestPOSIXClasses(t *testing.T) {
	testCases := []struct {
		name    string
		builder *tinyrebuilder.RegexBuilder
		good    string
		bad     string
	}{
		{"alnum", tinyrebuilder.New().POSIXClass("alnum"), "a", "-"},
		{"alpha", tinyrebuilder.New().POSIXClass("alpha"), "Z", "5"},
		{"digit", tinyrebuilder.New().POSIXClass("digit"), "9", "a"},
		{"^alnum", tinyrebuilder.New().NotPOSIXClass("alnum"), "-", "a"},
		{"^alpha", tinyrebuilder.New().NotPOSIXClass("alpha"), "5", "Z"},
		{"^digit", tinyrebuilder.New().NotPOSIXClass("digit"), "a", "9"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			re := tc.builder.MustCompile()
			if !re.IsMatch(tc.good) {
				t.Errorf("Expected to match %q", tc.good)
			}
			if re.IsMatch(tc.bad) {
				t.Errorf("Expected NOT to match %q", tc.bad)
			}
		})
	}
}

func TestGroupWithFlags(t *testing.T) {
	re := tinyrebuilder.New().
		GroupWithFlags("i", tinyrebuilder.New().Literal("hello")).
		Literal(" world").
		MustCompile()
	if !re.IsMatch("HELLO world") {
		t.Error("Expected to match 'HELLO world'")
	}
	if re.IsMatch("HELLO WORLD") {
		t.Error("Expected NOT to match 'HELLO WORLD'")
	}
}

func TestEscapeSequences(t *testing.T) {
	t.Run("Tab", func(t *testing.T) {
		re := tinyrebuilder.New().Tab().MustCompile()
		if !re.IsMatch("\t") {
			t.Error("Expected to match tab character")
		}
	})
	t.Run("Newline", func(t *testing.T) {
		re := tinyrebuilder.New().Newline().MustCompile()
		if !re.IsMatch("\n") {
			t.Error("Expected to match newline character")
		}
	})
	t.Run("CarriageReturn", func(t *testing.T) {
		re := tinyrebuilder.New().CarriageReturn().MustCompile()
		if !re.IsMatch("\r") {
			t.Error("Expected to match carriage return character")
		}
	})
	t.Run("Quote", func(t *testing.T) {
		re := tinyrebuilder.New().Quote(".+*?^${}()|[]\\").MustCompile()
		if !re.IsMatch(".+*?^${}()|[]\\") {
			t.Error("Expected to match quoted special characters literally")
		}
	})
}

func TestExtraAnchors(t *testing.T) {
	t.Run("StartOfString", func(t *testing.T) {
		re := tinyrebuilder.New().StartOfString().Literal("a").MustCompile()
		if !re.IsMatch("a\nb") {
			t.Error("Expected to match 'a' at the start of the string")
		}
		if re.IsMatch("b\na") {
			t.Error("Expected NOT to match 'a' when not at the start")
		}
	})
	t.Run("EndOfString", func(t *testing.T) {
		re := tinyrebuilder.New().Literal("b").EndOfString().MustCompile()
		if !re.IsMatch("a\nb") {
			t.Error("Expected to match 'b' at the end of the string")
		}
		if re.IsMatch("b\na") {
			t.Error("Expected NOT to match 'b' when not at the end")
		}
	})
}

func TestUnicodeProperties(t *testing.T) {
	t.Run("Greek", func(t *testing.T) {
		re := tinyrebuilder.New().UnicodeProperty("Greek").OneOrMore().MustCompile()
		if !re.IsMatch("αβγ") {
			t.Error("Expected to match Greek letters")
		}
		if re.IsMatch("abc") {
			t.Error("Expected NOT to match Latin letters")
		}
	})
	t.Run("NotLatin", func(t *testing.T) {
		re := tinyrebuilder.New().NotUnicodeProperty("Latin").OneOrMore().MustCompile()
		if !re.IsMatch("αβγ") {
			t.Error("Expected to match non-Latin letters")
		}
		if re.IsMatch("abc") {
			t.Error("Expected NOT to match Latin letters")
		}
	})
}

func TestCache(t *testing.T) {
	tinyrebuilder.PurgeCache()
	defer tinyrebuilder.PurgeCache()

	// Compile for the first time
	re1 := tinyrebuilder.New().Literal("hello").MustCompileWithCache()

	// Should retrieve from cache this time
	re2 := tinyrebuilder.New().Literal("hello").MustCompileWithCache()

	// The pointers should be the same
	if re1 != re2 {
		t.Error("Expected the same regex object to be returned from cache")
	}

	// Compile a different regex
	re3 := tinyrebuilder.New().Literal("world").MustCompileWithCache()
	if re1 == re3 {
		t.Error("Expected different regex objects for different patterns")
	}

	// Purge the cache and compile again
	tinyrebuilder.PurgeCache()
	re4 := tinyrebuilder.New().Literal("hello").MustCompileWithCache()
	if re1 == re4 {
		t.Error("Expected a new regex object after purging the cache")
	}
}

func TestExample(t *testing.T) {
	// Build a regex to match "hello world" at the start and end of a string.
	re := tinyrebuilder.New().
		StartAnchor().
		Literal("hello").
		Whitespace().
		Literal("world").
		EndAnchor().
		MustCompile()

	if !re.IsMatch("hello world") {
		t.Error("Expected to match 'hello world'")
	}
	if re.IsMatch("prefix hello world suffix") {
		t.Error("Expected NOT to match 'prefix hello world suffix'")
	}
}

func TestExample_patterns(t *testing.T) {
	// Use a pre-defined pattern to validate an email address.
	emailValidator := patterns.Email().MustCompile()

	if !emailValidator.IsMatch("test@example.com") {
		t.Error("Expected to match 'test@example.com'")
	}
	if emailValidator.IsMatch("invalid-email") {
		t.Error("Expected NOT to match 'invalid-email'")
	}
}

func TestExampleRegexp_FindStringSubmatch(t *testing.T) {
	// Build a regex to capture the year, month, and day from a date string.
	datePattern := tinyrebuilder.New().
		NamedGroup("year", tinyrebuilder.Digit().Exactly(4)).
		Literal("-").
		NamedGroup("month", tinyrebuilder.Digit().Exactly(2)).
		Literal("-").
		NamedGroup("day", tinyrebuilder.Digit().Exactly(2))

	re := datePattern.MustCompile()
	matches := re.FindStringSubmatch("2025-07-21")

	if len(matches) < 4 {
		t.Fatal("Expected to find 3 submatches")
	}
	if matches[1] != "2025" {
		t.Errorf("Expected year to be '2025', got %s", matches[1])
	}
	if matches[2] != "07" {
		t.Errorf("Expected month to be '07', got %s", matches[2])
	}
	if matches[3] != "21" {
		t.Errorf("Expected day to be '21', got %s", matches[3])
	}
}

func TestExampleRegexBuilder_MustCompileWithCache(t *testing.T) {
	// Purge cache for predictable test output.
	tinyrebuilder.PurgeCache()
	defer tinyrebuilder.PurgeCache()

	// The first compilation will compile and add the regex to the cache.
	re1 := tinyrebuilder.New().Literal("cache-me").MustCompileWithCache()

	// The second compilation will retrieve the regex from the cache.
	re2 := tinyrebuilder.New().Literal("cache-me").MustCompileWithCache()

	// The pointers will be the same, indicating it's the same object.
	if re1 != re2 {
		t.Error("Expected re1 and re2 to be the same instance")
	}

	// After purging the cache, a new object will be created.
	tinyrebuilder.PurgeCache()
	re3 := tinyrebuilder.New().Literal("cache-me").MustCompileWithCache()
	if re1 == re3 {
		t.Error("Expected re1 and re3 to be different instances")
	}
}

func TestQuantifiers(t *testing.T) {
	testCases := []struct {
		name    string
		builder func() *tinyrebuilder.RegexBuilder
		good    []string
		bad     []string
	}{
		{
			name: "ZeroOrMore",
			builder: func() *tinyrebuilder.RegexBuilder {
				return tinyrebuilder.New().StartAnchor().Literal("a").ZeroOrMore().EndAnchor()
			},
			good: []string{"", "a", "aa", "aaaaa"},
			bad:  []string{"b", "aab"},
		},
		{
			name: "AtLeast",
			builder: func() *tinyrebuilder.RegexBuilder {
				return tinyrebuilder.New().StartAnchor().Literal("a").AtLeast(2).EndAnchor()
			},
			good: []string{"aa", "aaaaa"},
			bad:  []string{"", "a", "aab"},
		},
		{
			name: "Between",
			builder: func() *tinyrebuilder.RegexBuilder {
				return tinyrebuilder.New().StartAnchor().Literal("a").Between(2, 4).EndAnchor()
			},
			good: []string{"aa", "aaa", "aaaa"},
			bad:  []string{"", "a", "aaaaa"},
		},
		{
			name: "NonGreedy",
			builder: func() *tinyrebuilder.RegexBuilder {
				return tinyrebuilder.New().StartAnchor().Literal("a").OneOrMore().NonGreedy().EndAnchor()
			},
			good: []string{"a", "aa"},
			bad:  []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			re := tc.builder().MustCompile()
			for _, s := range tc.good {
				if !re.IsMatch(s) {
					t.Errorf("Expected to match %q", s)
				}
			}
			for _, s := range tc.bad {
				if re.IsMatch(s) {
					t.Errorf("Expected NOT to match %q", s)
				}
			}
		})
	}
}

func TestWordBoundaries(t *testing.T) {
	t.Run("WordBoundary", func(t *testing.T) {
		re := tinyrebuilder.New().WordBoundary().Literal("cat").WordBoundary().MustCompile()
		if !re.IsMatch("the cat is cute") {
			t.Error("Expected to match 'cat' as a whole word")
		}
		if re.IsMatch("thecatis a monster") {
			t.Error("Expected NOT to match 'cat' when not on a word boundary")
		}
	})

	t.Run("NotWordBoundary", func(t *testing.T) {
		re := tinyrebuilder.New().NotWordBoundary().Literal("cat").NotWordBoundary().MustCompile()
		if !re.IsMatch("thecatis a monster") {
			t.Error("Expected to match 'cat' when not on a word boundary")
		}
		if re.IsMatch("the cat is cute") {
			t.Error("Expected NOT to match 'cat' as a whole word")
		}
	})
}

func BenchmarkSimpleRegexCompilation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = tinyrebuilder.New().
			StartAnchor().
			Literal("hello").
			Whitespace().
			Literal("world").
			EndAnchor().
			MustCompile()
	}
}

func BenchmarkComplexRegexCompilation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = tinyrebuilder.New().
			StartAnchor().
			Group(tinyrebuilder.New().Literal("a").AtLeast(1)).
			Or(tinyrebuilder.New().Literal("b").Exactly(2)).
			EndAnchor().
			MustCompile()
	}
}

func BenchmarkEmailPatternCompilation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = patterns.Email().MustCompile()
	}
}

func BenchmarkEmailPatternCompilationWithCache(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = patterns.Email().MustCompileWithCache()
	}
}

func BenchmarkURLPatternCompilation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = tinyrebuilder.New().
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
	}
}

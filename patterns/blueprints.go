// Package patterns provides a collection of pre-defined, commonly used regular
// expression patterns built with tinyrebuilder. These patterns can be used directly
// or as building blocks for more complex expressions.
package patterns

import (
	"github.com/nulln0ne/tinyrebuilder"
)

func Email() *tinyrebuilder.RegexBuilder {
	return tinyrebuilder.New().
		StartAnchor().
		Raw(`[a-zA-Z0-9._%+\-]+`).
		Literal("@").
		Raw(`[a-zA-Z0-9.\-]+\.[a-zA-Z]{1,}`).
		EndAnchor()
}

func IPv4() *tinyrebuilder.RegexBuilder {
	octet := tinyrebuilder.New().Raw(`(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])`)
	return tinyrebuilder.New().
		StartAnchor().
		Group(octet).Literal(".").
		Group(octet).Literal(".").
		Group(octet).Literal(".").
		Group(octet).
		EndAnchor()
}

func UUID() *tinyrebuilder.RegexBuilder {
	return tinyrebuilder.New().
		StartAnchor().
		Group(tinyrebuilder.New().Raw(`[0-9a-fA-F]`).Exactly(8)).Literal("-").
		Group(tinyrebuilder.New().Raw(`[0-9a-fA-F]`).Exactly(4)).Literal("-").
		Group(tinyrebuilder.New().Raw(`[0-9a-fA-F]`).Exactly(4)).Literal("-").
		Group(tinyrebuilder.New().Raw(`[0-9a-fA-F]`).Exactly(4)).Literal("-").
		Group(tinyrebuilder.New().Raw(`[0-9a-fA-F]`).Exactly(12)).
		EndAnchor()
}

func HexColor() *tinyrebuilder.RegexBuilder {
	h3 := tinyrebuilder.New().Raw(`[0-9a-fA-F]`).Exactly(3)
	h6 := tinyrebuilder.New().Raw(`[0-9a-fA-F]`).Exactly(6)
	return tinyrebuilder.New().
		StartAnchor().
		Literal("#").
		NonCapturingGroup(
			h3.Or(h6),
		).
		EndAnchor()
}

func URL() *tinyrebuilder.RegexBuilder {
	protocol := tinyrebuilder.New().Literal("http").Maybe().Literal("s").Maybe().Literal("://")
	domain := tinyrebuilder.New().Raw(`[a-zA-Z0-9.-]+`)
	port := tinyrebuilder.New().NonCapturingGroup(tinyrebuilder.New().Literal(":").Raw(`[0-9]+`)).Maybe()
	path := tinyrebuilder.New().Raw(`(?:/[a-zA-Z0-9-._~:/?#\[\]@!$&'()*+,;=]*)?`)
	return tinyrebuilder.New().
		StartAnchor().
		Raw(protocol.Build()).
		Raw(domain.Build()).
		Raw(port.Build()).
		Raw(path.Build()).
		EndAnchor()
}

func Date_YYYYMMDD() *tinyrebuilder.RegexBuilder {
	year := tinyrebuilder.New().Raw(`\d`).Exactly(4)
	month := tinyrebuilder.New().Raw(`(0[1-9]|1[0-2])`)
	day := tinyrebuilder.New().Raw(`(0[1-9]|[12]\d|3[01])`)
	return tinyrebuilder.New().
		StartAnchor().
		Group(year).Literal("-").
		Group(month).Literal("-").
		Group(day).
		EndAnchor()
}

func Time_HHMMSS() *tinyrebuilder.RegexBuilder {
	hour := tinyrebuilder.New().Raw(`([01]\d|2[0-3])`)
	minute := tinyrebuilder.New().Raw(`[0-5]\d`)
	second := tinyrebuilder.New().Raw(`[0-5]\d`)
	return tinyrebuilder.New().
		StartAnchor().
		Group(hour).Literal(":").
		Group(minute).Literal(":").
		Group(second).
		EndAnchor()
}

func Username() *tinyrebuilder.RegexBuilder {
	return tinyrebuilder.New().
		StartAnchor().
		Raw(`[a-zA-Z0-9_]{3,16}`).
		EndAnchor()
}

func Slug() *tinyrebuilder.RegexBuilder {
	return tinyrebuilder.New().
		StartAnchor().
		Raw(`[a-z0-9]+(?:-[a-z0-9]+)*`).
		EndAnchor()
}

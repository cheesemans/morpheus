package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func revealablePlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: revealablePlaygroundDefaultHTML},
		{Label: "Text spoiler", HTML: revealableStickyHTML},
		{Label: "Image spoiler", HTML: revealableImageHTML, CSS: revealableImageCSS},
		{Label: "Blurred image spoiler", HTML: revealableBlurImageHTML, CSS: revealableBlurImageCSS},
		{Label: "Section spoiler", HTML: revealableSectionHTML, CSS: revealableSectionCSS},
		{Label: "Custom styling", HTML: revealableStyledHTML, CSS: revealableStyledCSS},
	}
}

//go:embed examples/revealable_default.html
var revealablePlaygroundDefaultHTML string

var revealableImageHTML = renderExampleHTML(examples.RevealableImage())

//go:embed examples/revealable_image.templ
var revealableImageTempl string

//go:embed examples/revealable_image.css
var revealableImageCSS string

var revealableBlurImageHTML = renderExampleHTML(examples.RevealableBlurImage())

//go:embed examples/revealable_blur_image.templ
var revealableBlurImageTempl string

//go:embed examples/revealable_blur_image.css
var revealableBlurImageCSS string

var revealableSectionHTML = renderExampleHTML(examples.RevealableSection())

//go:embed examples/revealable_section.templ
var revealableSectionTempl string

//go:embed examples/revealable_section.css
var revealableSectionCSS string

var revealableStyledHTML = renderExampleHTML(examples.RevealableStyled())

//go:embed examples/revealable_styled.templ
var revealableStyledTempl string

//go:embed examples/revealable_styled.css
var revealableStyledCSS string

var revealableStickyHTML = renderExampleHTML(examples.RevealableSticky())

//go:embed examples/revealable_sticky.templ
var revealableStickyTempl string

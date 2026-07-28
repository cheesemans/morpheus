package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func buttonPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: buttonPlaygroundDefaultHTML},
		{Label: "Variants & states", HTML: buttonVariantsHTML},
		{Label: "Composed content", HTML: buttonComposedHTML, CSS: buttonComposedCSS},
	}
}

//go:embed examples/button_default.html
var buttonPlaygroundDefaultHTML string

var buttonVariantsHTML = renderExampleHTML(examples.ButtonVariants())

//go:embed examples/button_variants.templ
var buttonVariantsTempl string

var buttonComposedHTML = renderExampleHTML(examples.ButtonComposed())

//go:embed examples/button_composed.templ
var buttonComposedTempl string

//go:embed examples/button_composed.css
var buttonComposedCSS string

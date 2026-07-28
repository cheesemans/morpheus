package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func breadcrumbPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: breadcrumbPlaygroundDefaultHTML},
		{Label: "Deep trail", HTML: breadcrumbDeepHTML},
		{Label: "Deep trail, compact labels", HTML: breadcrumbCompactHTML},
		{Label: "Slash separator", HTML: breadcrumbSlashHTML},
		{Label: "Dot separator", HTML: breadcrumbDotHTML},
		{Label: "Icon separator (sparkles)", HTML: breadcrumbSparklesHTML},
		{Label: "Chip wrappers", HTML: breadcrumbChipsHTML, CSS: breadcrumbChipsCSS},
	}
}

//go:embed examples/breadcrumb_default.html
var breadcrumbPlaygroundDefaultHTML string

var breadcrumbDeepHTML = renderExampleHTML(examples.BreadcrumbDeep())

//go:embed examples/breadcrumb_deep.templ
var breadcrumbDeepTempl string

var breadcrumbCompactHTML = renderExampleHTML(examples.BreadcrumbCompact())

//go:embed examples/breadcrumb_compact.templ
var breadcrumbCompactTempl string

var breadcrumbSlashHTML = renderExampleHTML(examples.BreadcrumbSlash())

//go:embed examples/breadcrumb_slash.templ
var breadcrumbSlashTempl string

var breadcrumbDotHTML = renderExampleHTML(examples.BreadcrumbDot())

//go:embed examples/breadcrumb_dot.templ
var breadcrumbDotTempl string

var breadcrumbSparklesHTML = renderExampleHTML(examples.BreadcrumbSparkles())

//go:embed examples/breadcrumb_sparkles.templ
var breadcrumbSparklesTempl string

var breadcrumbChipsHTML = renderExampleHTML(examples.BreadcrumbChips())

//go:embed examples/breadcrumb_chips.templ
var breadcrumbChipsTempl string

//go:embed examples/breadcrumb_chips.css
var breadcrumbChipsCSS string

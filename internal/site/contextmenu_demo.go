package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func contextMenuPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: contextMenuPlaygroundDefaultHTML,
			CSS: contextMenuPlaygroundDefaultCSS},
		{Label: "Simple", HTML: contextMenuSimpleHTML, CSS: contextMenuSimpleCSS},
		{Label: "Nested", HTML: contextMenuNestedHTML, CSS: contextMenuNestedCSS},
		{Label: "Custom content", HTML: contextMenuCustomHTML, CSS: contextMenuCustomCSS},
	}
}

//go:embed examples/context_menu_default.html
var contextMenuPlaygroundDefaultHTML string

//go:embed examples/context_menu_default.css
var contextMenuPlaygroundDefaultCSS string

var contextMenuSimpleHTML = renderExampleHTML(examples.ContextMenuSimple())

//go:embed examples/context_menu_simple.templ
var contextMenuSimpleTempl string

//go:embed examples/context_menu_simple.css
var contextMenuSimpleCSS string

var contextMenuNestedHTML = renderExampleHTML(examples.ContextMenuNested())

//go:embed examples/context_menu_nested.templ
var contextMenuNestedTempl string

//go:embed examples/context_menu_nested.css
var contextMenuNestedCSS string

var contextMenuCustomHTML = renderExampleHTML(examples.ContextMenuCustom())

//go:embed examples/context_menu_custom.templ
var contextMenuCustomTempl string

//go:embed examples/context_menu_custom.css
var contextMenuCustomCSS string

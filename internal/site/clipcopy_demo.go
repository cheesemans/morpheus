package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func clipcopyPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: clipcopyPlaygroundDefaultHTML},
		{Label: "Icon + target ref", HTML: clipcopyTargetHTML, CSS: clipcopyTargetCSS},
	}
}

//go:embed examples/clipcopy_default.html
var clipcopyPlaygroundDefaultHTML string

var clipcopyTargetHTML = renderExampleHTML(examples.ClipcopyTarget())

//go:embed examples/clipcopy_target.templ
var clipcopyTargetTempl string

//go:embed examples/clipcopy_target.css
var clipcopyTargetCSS string

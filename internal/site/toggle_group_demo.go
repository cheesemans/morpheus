package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func toggleGroupPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: toggleGroupPlaygroundDefaultHTML},
		{Label: "Formatting toolbar", HTML: toggleGroupFormattingHTML},
		{Label: "View options", HTML: toggleGroupAlignmentHTML},
		{Label: "Vertical", HTML: toggleGroupVerticalHTML},
	}
}

//go:embed examples/toggle_group_default.html
var toggleGroupPlaygroundDefaultHTML string

var toggleGroupFormattingHTML = renderExampleHTML(examples.ToggleGroupFormatting())

//go:embed examples/toggle_group_formatting.templ
var toggleGroupFormattingTempl string

var toggleGroupAlignmentHTML = renderExampleHTML(examples.ToggleGroupAlignment())

//go:embed examples/toggle_group_alignment.templ
var toggleGroupAlignmentTempl string

var toggleGroupVerticalHTML = renderExampleHTML(examples.ToggleGroupVertical())

//go:embed examples/toggle_group_vertical.templ
var toggleGroupVerticalTempl string

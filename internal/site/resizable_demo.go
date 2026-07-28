package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func resizablePlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: resizableStateDefaultHTML, CSS: resizableStateDefaultCSS},
		{Label: "Bottom-right corner", HTML: resizableCornerHTML, CSS: resizableCornerCSS},
		{Label: "Edges only", HTML: resizableEdgesHTML, CSS: resizableEdgesCSS},
		{Label: "All eight handles", HTML: resizableAllHTML, CSS: resizableAllCSS},
		{Label: "Vertical only", HTML: resizableVerticalHTML, CSS: resizableVerticalCSS},
		{Label: "Horizontal only", HTML: resizableHorizontalHTML, CSS: resizableHorizontalCSS},
		{Label: "Step grid", HTML: resizableStepHTML, CSS: resizableStepCSS},
		{Label: "Custom handle icon", HTML: resizableCustomIconHTML, CSS: resizableCustomIconCSS},
	}
}

//go:embed examples/resizable_state_default.html
var resizableStateDefaultHTML string

//go:embed examples/resizable_state_default.css
var resizableStateDefaultCSS string

var resizableCornerHTML = renderExampleHTML(examples.ResizableCorner())

//go:embed examples/resizable_corner.templ
var resizableCornerTempl string

//go:embed examples/resizable_corner.css
var resizableCornerCSS string

var resizableEdgesHTML = renderExampleHTML(examples.ResizableEdges())

//go:embed examples/resizable_edges.templ
var resizableEdgesTempl string

//go:embed examples/resizable_edges.css
var resizableEdgesCSS string

var resizableAllHTML = renderExampleHTML(examples.ResizableAll())

//go:embed examples/resizable_all.templ
var resizableAllTempl string

//go:embed examples/resizable_all.css
var resizableAllCSS string

var resizableVerticalHTML = renderExampleHTML(examples.ResizableVertical())

//go:embed examples/resizable_vertical.templ
var resizableVerticalTempl string

//go:embed examples/resizable_vertical.css
var resizableVerticalCSS string

var resizableHorizontalHTML = renderExampleHTML(examples.ResizableHorizontal())

//go:embed examples/resizable_horizontal.templ
var resizableHorizontalTempl string

//go:embed examples/resizable_horizontal.css
var resizableHorizontalCSS string

var resizableStepHTML = renderExampleHTML(examples.ResizableStep())

//go:embed examples/resizable_step.templ
var resizableStepTempl string

//go:embed examples/resizable_step.css
var resizableStepCSS string

var resizableCustomIconHTML = renderExampleHTML(examples.ResizableCustomIcon())

//go:embed examples/resizable_custom_icon.templ
var resizableCustomIconTempl string

//go:embed examples/resizable_custom_icon.css
var resizableCustomIconCSS string

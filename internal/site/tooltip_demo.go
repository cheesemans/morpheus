package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func tooltipPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: tooltipPlaygroundDefaultHTML},
		{Label: "Placements", HTML: tooltipPlacementsHTML, CSS: tooltipPlacementsCSS},
		{Label: "Shorthand body", HTML: tooltipShorthandHTML},
		{Label: "Rich content", HTML: tooltipRichHTML, CSS: tooltipRichCSS},
		{Label: "Delays", HTML: tooltipDelaysHTML},
	}
}

//go:embed examples/tooltip_default.html
var tooltipPlaygroundDefaultHTML string

var tooltipPlacementsHTML = renderExampleHTML(examples.TooltipPlacements())

//go:embed examples/tooltip_placements.templ
var tooltipPlacementsTempl string

//go:embed examples/tooltip_placements.css
var tooltipPlacementsCSS string

var tooltipShorthandHTML = renderExampleHTML(examples.TooltipShorthand())

//go:embed examples/tooltip_shorthand.templ
var tooltipShorthandTempl string

var tooltipRichHTML = renderExampleHTML(examples.TooltipRich())

//go:embed examples/tooltip_rich.templ
var tooltipRichTempl string

//go:embed examples/tooltip_rich.css
var tooltipRichCSS string

var tooltipDelaysHTML = renderExampleHTML(examples.TooltipDelays())

//go:embed examples/tooltip_delays.templ
var tooltipDelaysTempl string

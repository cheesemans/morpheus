package site

import (
	_ "embed"
	"github.com/romshark/morpheus/internal/site/examples"
)

//go:embed static/sim/popover/loadcontent.js
var popoverLoadContentScript string

//go:embed static/sim/popover/asyncload.js
var popoverAsyncLoadScript string

func popoverPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: popoverPlaygroundDefaultHTML},
		{Label: "Open", HTML: popoverPlaygroundOpenHTML},
		{Label: "Placements", HTML: popoverPlacementsHTML, CSS: popoverPlacementsCSS},
		{Label: "With actions", HTML: popoverActionsHTML, CSS: popoverActionsCSS},
		{Label: "Tall content (clamps + scrolls)", HTML: popoverTallHTML, CSS: popoverTallCSS},
		{Label: "Strict placement", HTML: popoverStrictHTML, CSS: popoverStrictCSS},
		{Label: "Wide content (horizontal clamp)", HTML: popoverWideHTML, CSS: popoverWideCSS},
		{Label: "Hover to open", HTML: popoverHoverHTML, CSS: popoverHoverCSS},
	}
}

//go:embed examples/popover_default.html
var popoverPlaygroundDefaultHTML string

//go:embed examples/popover_open.html
var popoverPlaygroundOpenHTML string

func popoverMorphStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Initial", HTML: popoverMorphInitialHTML, CSS: popoverMorphCSS},
		{Label: "Progress: 50%", HTML: popoverMorph50HTML, CSS: popoverMorphCSS},
		{Label: "Progress: 75%", HTML: popoverMorph75HTML, CSS: popoverMorphCSS},
		{Label: "Progress: 100%", HTML: popoverMorph100HTML, CSS: popoverMorphCSS},
	}
}

//go:embed examples/popover_morph_initial.html
var popoverMorphInitialHTML string

//go:embed examples/popover_morph_50.html
var popoverMorph50HTML string

//go:embed examples/popover_morph_75.html
var popoverMorph75HTML string

//go:embed examples/popover_morph_100.html
var popoverMorph100HTML string

//go:embed examples/popover_morph.css
var popoverMorphCSS string

var popoverTallHTML = renderExampleHTML(examples.PopoverTall())

//go:embed examples/popover_tall.templ
var popoverTallTempl string

//go:embed examples/popover_tall.css
var popoverTallCSS string

var popoverStrictHTML = renderExampleHTML(examples.PopoverStrict())

//go:embed examples/popover_strict.templ
var popoverStrictTempl string

//go:embed examples/popover_strict.css
var popoverStrictCSS string

var popoverWideHTML = renderExampleHTML(examples.PopoverWide())

//go:embed examples/popover_wide.templ
var popoverWideTempl string

//go:embed examples/popover_wide.css
var popoverWideCSS string

var popoverAutoFlipHTML = renderExampleHTML(examples.PopoverAutoFlip())

//go:embed examples/popover_auto_flip.templ
var popoverAutoFlipTempl string

//go:embed examples/popover_auto_flip.css
var popoverAutoFlipCSS string

var popoverPlacementsHTML = renderExampleHTML(examples.PopoverPlacements())

//go:embed examples/popover_placements.templ
var popoverPlacementsTempl string

//go:embed examples/popover_placements.css
var popoverPlacementsCSS string

var popoverActionsHTML = renderExampleHTML(examples.PopoverActions())

//go:embed examples/popover_actions.templ
var popoverActionsTempl string

//go:embed examples/popover_actions.css
var popoverActionsCSS string

var popoverCascadingHTML = renderExampleHTML(examples.PopoverCascading())

//go:embed examples/popover_cascading.templ
var popoverCascadingTempl string

//go:embed examples/popover_cascading.css
var popoverCascadingCSS string

var popoverHoverHTML = renderExampleHTML(examples.PopoverHover())

//go:embed examples/popover_hover.templ
var popoverHoverTempl string

//go:embed examples/popover_hover.css
var popoverHoverCSS string

var popoverLazyHTML = renderExampleHTML(examples.PopoverLazyDemo())

//go:embed examples/popover_lazy.templ
var popoverLazyTempl string

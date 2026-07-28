package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func tabsPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: tabsPlaygroundDefaultHTML},
		{Label: "Icons + labels", HTML: tabsIconsHTML},
		{Label: "No animation", HTML: tabsNoAnimationHTML},
		{Label: "Custom animation", HTML: tabsCustomAnimationHTML},
		{Label: "Custom styling", HTML: tabsUnderlineHTML, CSS: tabsUnderlineCSS},
		{Label: "Auto-activate", HTML: tabsAutoActivateHTML},
		{Label: "Vertical", HTML: tabsVerticalHTML},
	}
}

//go:embed examples/tabs_default.html
var tabsPlaygroundDefaultHTML string

var tabsIconsHTML = renderExampleHTML(examples.TabsIcons())

//go:embed examples/tabs_icons.templ
var tabsIconsTempl string

var tabsNoAnimationHTML = renderExampleHTML(examples.TabsNoAnimation())

//go:embed examples/tabs_no_animation.templ
var tabsNoAnimationTempl string

var tabsCustomAnimationHTML = renderExampleHTML(examples.TabsCustomAnimation())

//go:embed examples/tabs_custom_animation.templ
var tabsCustomAnimationTempl string

var tabsUnderlineHTML = renderExampleHTML(examples.TabsUnderline())

//go:embed examples/tabs_underline.templ
var tabsUnderlineTempl string

//go:embed examples/tabs_underline.css
var tabsUnderlineCSS string

var tabsAutoActivateHTML = renderExampleHTML(examples.TabsAutoActivate())

//go:embed examples/tabs_auto_activate.templ
var tabsAutoActivateTempl string

var tabsVerticalHTML = renderExampleHTML(examples.TabsVertical())

//go:embed examples/tabs_vertical.templ
var tabsVerticalTempl string

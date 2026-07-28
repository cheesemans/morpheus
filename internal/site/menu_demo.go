package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func menuPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: menuPlaygroundDefaultHTML},
		{Label: "Push mode", HTML: menuPushModeHTML},
		{Label: "Open above", HTML: menuOpenAboveHTML, CSS: menuOpenAboveCSS},
		{Label: "Disabled rows", HTML: menuDisabledRowsHTML},
	}
}

func menuMorphStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Initial", HTML: menuMorphInitialHTML},
		{Label: "Disabled options", HTML: menuMorphDisabledHTML},
		{Label: "More options", HTML: menuMorphMoreHTML},
	}
}

//go:embed examples/menu_morph_initial.html
var menuMorphInitialHTML string

//go:embed examples/menu_morph_disabled.html
var menuMorphDisabledHTML string

//go:embed examples/menu_morph_more.html
var menuMorphMoreHTML string

//go:embed examples/menu_default.html
var menuPlaygroundDefaultHTML string

var menuDisabledRowsHTML = renderExampleHTML(examples.MenuDisabledRows())

//go:embed examples/menu_disabled_rows.templ
var menuDisabledRowsTempl string

var menuPushModeHTML = renderExampleHTML(examples.MenuPushMode())

//go:embed examples/menu_push_mode.templ
var menuPushModeTempl string

var menuOpenAboveHTML = renderExampleHTML(examples.MenuOpenAbove())

//go:embed examples/menu_open_above.templ
var menuOpenAboveTempl string

//go:embed examples/menu_open_above.css
var menuOpenAboveCSS string

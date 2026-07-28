package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func drawerPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: drawerPlaygroundDefaultHTML},
		{Label: "Sides", HTML: drawerSidesHTML},
		{Label: "Frost glass", HTML: drawerFrostGlassHTML, CSS: drawerFrostGlassCSS},
		{Label: "Non-dismissible", HTML: drawerNonDismissibleHTML},
		{Label: "Scrolling body", HTML: drawerScrollingBodyPlaygroundHTML},
		{Label: "Touch-dismiss", HTML: drawerTouchDismissHTML, CSS: drawerTouchDismissCSS},
	}
}

//go:embed examples/drawer_default.html
var drawerPlaygroundDefaultHTML string

//go:embed examples/async_drawer_default.html
var asyncDrawerDefaultHTML string

var drawerSidesHTML = renderExampleHTML(examples.DrawerSides())

//go:embed examples/drawer_sides.templ
var drawerSidesTempl string

var drawerFrostGlassHTML = renderExampleHTML(examples.DrawerFrostGlass())

//go:embed examples/drawer_frost_glass.templ
var drawerFrostGlassTempl string

//go:embed examples/drawer_frost_glass.css
var drawerFrostGlassCSS string

var drawerNonDismissibleHTML = renderExampleHTML(examples.DrawerNonDismissible())

//go:embed examples/drawer_non_dismissible.templ
var drawerNonDismissibleTempl string

var drawerTouchDismissHTML = renderExampleHTML(examples.DrawerTouchDismiss())

//go:embed examples/drawer_touch_dismiss.templ
var drawerTouchDismissTempl string

//go:embed examples/drawer_touch_dismiss.css
var drawerTouchDismissCSS string

var drawerScrollingBodyHTML = renderExampleHTML(examples.DrawerScrollingBody())

//go:embed examples/drawer_scrolling_body.templ
var drawerScrollingBodyTempl string

var drawerFromDialogHTML = renderExampleHTML(examples.DrawerFromDialog())

//go:embed examples/drawer_from_dialog.templ
var drawerFromDialogTempl string

var drawerScrollingBodyPlaygroundHTML = renderExampleHTML(examples.DrawerScrollingBodyPlayground())

var drawerAsyncLoadingHTML = renderExampleHTML(examples.DrawerAsyncLoading())

//go:embed examples/drawer_async_loading.templ
var drawerAsyncLoadingTempl string

//go:embed examples/drawer_async_loading.css
var drawerAsyncLoadingCSS string

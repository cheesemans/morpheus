package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func sidebarPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: sidebarDefaultHTML, CSS: sidebarDefaultCSS},
		{Label: "Mixed breakpoints", HTML: sidebarMixedBreakpointsHTML,
			CSS: sidebarMixedBreakpointsCSS},
		{Label: "Resizable widths", HTML: sidebarResizableHTML,
			CSS: sidebarResizableCSS},
		{Label: "Touch-dismiss", HTML: sidebarTouchDismissHTML,
			CSS: sidebarTouchDismissCSS},
		{Label: "Minimized rail", HTML: sidebarMinimizedHTML, CSS: sidebarMinimizedCSS},
	}
}

//go:embed examples/sidebar_default.html
var sidebarDefaultHTML string

//go:embed examples/sidebar_default.css
var sidebarDefaultCSS string

var sidebarMixedBreakpointsHTML = renderExampleHTML(examples.SidebarMixedBreakpoints())

//go:embed examples/sidebar_mixed_breakpoints.templ
var sidebarMixedBreakpointsTempl string

//go:embed examples/sidebar_mixed_breakpoints.css
var sidebarMixedBreakpointsCSS string

var sidebarResizableHTML = renderExampleHTML(examples.SidebarResizable())

//go:embed examples/sidebar_resizable.templ
var sidebarResizableTempl string

//go:embed examples/sidebar_resizable.css
var sidebarResizableCSS string

var sidebarTouchDismissHTML = renderExampleHTML(examples.SidebarTouchDismiss())

//go:embed examples/sidebar_touch_dismiss.templ
var sidebarTouchDismissTempl string

//go:embed examples/sidebar_touch_dismiss.css
var sidebarTouchDismissCSS string

//go:embed examples/sidebar_minimized.html
var sidebarMinimizedHTML string

//go:embed examples/sidebar_minimized.css
var sidebarMinimizedCSS string

var sidebarAsyncLoadingHTML = renderExampleHTML(examples.SidebarAsyncLoading())

//go:embed examples/sidebar_async_loading.templ
var sidebarAsyncLoadingTempl string

//go:embed examples/sidebar_async_loading.css
var sidebarAsyncLoadingCSS string

//go:embed static/sim/sidebar/async.js
var sidebarAsyncScript string

//go:embed examples/sidebar_async_failure.templ
var sidebarAsyncFailureTempl string

//go:embed examples/sidebar_async_failure.css
var sidebarAsyncFailureCSS string

//go:embed static/sim/sidebar/asyncfail.js
var sidebarAsyncFailureScript string

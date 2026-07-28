package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

//go:embed static/sim/toaster/app-shell.js
var toasterAppShellScript string

//go:embed static/sim/toaster/action.js
var toasterActionScript string

//go:embed static/sim/toaster/patch-update.js
var toasterPatchUpdateScript string

//go:embed static/sim/toaster/patch-append.js
var toasterPatchAppendScript string

//go:embed static/sim/toaster/patch-replace.js
var toasterPatchReplaceScript string

func toasterPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: toasterPlaygroundDefaultHTML,
			CSS: toasterPlaygroundDefaultCSS},
		{Label: "Variants", HTML: toasterVariantsHTML},
		{Label: "Title + description", HTML: toasterTitleDescHTML},
		{Label: "Stack + dismiss-all", HTML: toasterStackDismissHTML},
	}
}

//go:embed examples/toaster_default.html
var toasterPlaygroundDefaultHTML string

//go:embed examples/toaster_default.css
var toasterPlaygroundDefaultCSS string

var toasterAppShellHTML = renderExampleHTML(examples.ToasterAppShell())

//go:embed examples/toaster_app_shell.templ
var toasterAppShellTempl string

var toasterVanillaHTML = renderExampleHTML(examples.ToasterVanilla())

//go:embed examples/toaster_vanilla.templ
var toasterVanillaTempl string

//go:embed examples/toaster_vanilla.css
var toasterVanillaCSS string

var toasterLoadingAsyncHTML = renderExampleHTML(examples.ToasterLoadingAsync())

//go:embed examples/toaster_loading_async.templ
var toasterLoadingAsyncTempl string

var toasterPatchAppendHTML = renderExampleHTML(examples.ToasterPatchAppend())

//go:embed examples/toaster_patch_append.templ
var toasterPatchAppendTempl string

var toasterPatchReplaceHTML = renderExampleHTML(examples.ToasterPatchReplace())

//go:embed examples/toaster_patch_replace.templ
var toasterPatchReplaceTempl string

var toasterPatchUpdateHTML = renderExampleHTML(examples.ToasterPatchUpdate())

//go:embed examples/toaster_patch_update.templ
var toasterPatchUpdateTempl string

var toasterActionButtonHTML = renderExampleHTML(examples.ToasterActionButton())

//go:embed examples/toaster_action_button.templ
var toasterActionButtonTempl string

var toasterVariantsHTML = renderExampleHTML(examples.ToasterVariants())

//go:embed examples/toaster_variants.templ
var toasterVariantsTempl string

var toasterTitleDescHTML = renderExampleHTML(examples.ToasterTitleDesc())

//go:embed examples/toaster_title_desc.templ
var toasterTitleDescTempl string

var toasterStackDismissHTML = renderExampleHTML(examples.ToasterStackDismiss())

//go:embed examples/toaster_stack_dismiss.templ
var toasterStackDismissTempl string

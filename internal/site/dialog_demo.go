package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func dialogPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: dialogPlaygroundDefaultHTML},
		{Label: "Non-dismissible", HTML: dialogNonDismissibleHTML},
		{Label: "Custom surface", HTML: dialogCustomSurfaceHTML,
			CSS: dialogCustomSurfaceCSS},
	}
}

func dialogMorphStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Initial", HTML: dialogMorphInitialHTML, CSS: dialogMorphCSS},
		{Label: "Different content", HTML: dialogMorphDifferentHTML,
			CSS: dialogMorphCSS},
	}
}

//go:embed examples/dialog_morph_initial.html
var dialogMorphInitialHTML string

//go:embed examples/dialog_morph_different.html
var dialogMorphDifferentHTML string

//go:embed examples/dialog_morph.css
var dialogMorphCSS string

//go:embed examples/dialog_default.html
var dialogPlaygroundDefaultHTML string

//go:embed examples/async_dialog_default.html
var asyncDialogDefaultHTML string

var dialogNonDismissibleHTML = renderExampleHTML(examples.DialogNonDismissible())

//go:embed examples/dialog_non_dismissible.templ
var dialogNonDismissibleTempl string

var dialogCustomSurfaceHTML = renderExampleHTML(examples.DialogCustomSurface())

//go:embed examples/dialog_custom_surface.templ
var dialogCustomSurfaceTempl string

//go:embed examples/dialog_custom_surface.css
var dialogCustomSurfaceCSS string

var dialogAsyncLoadingHTML = renderExampleHTML(examples.DialogAsyncLoading())

//go:embed examples/dialog_async_loading.templ
var dialogAsyncLoadingTempl string

//go:embed examples/dialog_async_loading.css
var dialogAsyncLoadingCSS string

//go:embed examples/dialog_async_failure.templ
var dialogAsyncFailureTempl string

//go:embed examples/dialog_async_failure.css
var dialogAsyncFailureCSS string

var dialogContinuousPatchingHTML = renderExampleHTML(examples.DialogContinuousPatching())

//go:embed examples/dialog_continuous_patching.templ
var dialogContinuousPatchingTempl string

//go:embed examples/dialog_continuous_patching.css
var dialogContinuousPatchingCSS string

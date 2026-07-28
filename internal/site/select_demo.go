package site

import (
	_ "embed"
	"github.com/romshark/morpheus/internal/site/examples"
)

func selectPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: selectPlaygroundDefaultHTML},
		{Label: "Disabled control", HTML: selectDisabledControlHTML},
		{Label: "Disabled options", HTML: selectDisabledHTML},
		{Label: "Grouped options (time zones)", HTML: selectGroupedHTML},
		{Label: "Clearable with empty slot", HTML: selectClearableHTML, CSS: selectClearableCSS},
		{Label: "Rich options", HTML: selectRichOptionsHTML, CSS: selectRichOptionsCSS},
		{Label: "Compact preview (label override)", HTML: selectCompactPreviewHTML, CSS: selectCompactPreviewCSS},
		{Label: "Custom trigger face", HTML: selectCustomTriggerFaceHTML, CSS: selectCustomTriggerFaceCSS},
	}
}

func selectMorphStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: selectMorphDefaultHTML},
		{Label: "Add more options", HTML: selectMorphMoreHTML},
		{Label: "Disable some options", HTML: selectMorphDisabledHTML},
	}
}

//go:embed examples/select_morph_default.html
var selectMorphDefaultHTML string

//go:embed examples/select_morph_more.html
var selectMorphMoreHTML string

//go:embed examples/select_morph_disabled.html
var selectMorphDisabledHTML string

//go:embed examples/select_default.html
var selectPlaygroundDefaultHTML string

//go:embed static/sim/select/lazy.js
var selectLazyScript string

//go:embed static/sim/select/async-spinner.js
var selectAsyncSpinnerScript string

var selectDisabledHTML = renderExampleHTML(examples.SelectDisabled())

//go:embed examples/select_disabled.templ
var selectDisabledTempl string

var selectDisabledControlHTML = renderExampleHTML(examples.SelectDisabledControl())

//go:embed examples/select_disabled_control.templ
var selectDisabledControlTempl string

var selectFormSubmissionHTML = renderExampleHTML(examples.SelectFormSubmission())

//go:embed examples/select_form.templ
var selectFormSubmissionTempl string

//go:embed examples/select_form.css
var selectFormSubmissionCSS string

var selectGroupedHTML = renderExampleHTML(examples.SelectGrouped())

//go:embed examples/select_grouped.templ
var selectGroupedTempl string

var selectClearableHTML = renderExampleHTML(examples.SelectClearable())

//go:embed examples/select_clearable.templ
var selectClearableTempl string

//go:embed examples/select_clearable.css
var selectClearableCSS string

var selectRichOptionsHTML = renderExampleHTML(examples.SelectRichOptions())

//go:embed examples/select_rich_options.templ
var selectRichOptionsTempl string

//go:embed examples/select_rich_options.css
var selectRichOptionsCSS string

var selectCompactPreviewHTML = renderExampleHTML(examples.SelectCompactPreview())

//go:embed examples/select_compact_preview.templ
var selectCompactPreviewTempl string

//go:embed examples/select_compact_preview.css
var selectCompactPreviewCSS string

var selectCustomTriggerFaceHTML = renderExampleHTML(examples.SelectCustomTriggerFace())

//go:embed examples/select_custom_trigger_face.templ
var selectCustomTriggerFaceTempl string

//go:embed examples/select_custom_trigger_face.css
var selectCustomTriggerFaceCSS string

var selectAsyncSpinnerHTML = renderExampleHTML(examples.SelectAsyncSpinner())

//go:embed examples/select_async_spinner.templ
var selectAsyncSpinnerTempl string

//go:embed examples/select_async_spinner.css
var selectAsyncSpinnerCSS string

var selectLazyHTML = renderExampleHTML(examples.SelectLazy())

//go:embed examples/select_lazy.templ
var selectLazyTempl string

//go:embed examples/select_lazy.css
var selectLazyCSS string

var selectLongListHTML = renderExampleHTML(examples.SelectLongList())

//go:embed examples/select_long_list.templ
var selectLongListTempl string

var selectAutoFlipHTML = renderExampleHTML(examples.SelectAutoFlip())

//go:embed examples/select_auto_flip.templ
var selectAutoFlipTempl string

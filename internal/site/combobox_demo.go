package site

import (
	_ "embed"
	"github.com/romshark/morpheus/internal/site/examples"
)

//go:embed static/sim/combobox/async-load.js
var comboboxAsyncLoadScript string

//go:embed static/sim/combobox/asyncload.js
var comboboxAsyncFailureScript string

//go:embed static/sim/combobox/search.js
var comboboxSearchScript string

//go:embed static/sim/combobox/lazy-once.js
var comboboxLazyOnceScript string

func comboboxPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: comboboxPlaygroundDefaultHTML},
		{Label: "Disabled control", HTML: comboboxDisabledControlHTML},
		{Label: "Grouped options", HTML: comboboxGroupedHTML, CSS: comboboxGroupedCSS},
		{Label: "Multiple selection", HTML: comboboxMultipleHTML},
		{Label: "Rich options", HTML: comboboxRichOptionsHTML, CSS: comboboxRichOptionsCSS},
		{Label: "Custom trigger face", HTML: comboboxCustomTriggerFaceHTML, CSS: comboboxCustomTriggerFaceCSS},
		{Label: "Multi-select chips", HTML: comboboxMultiFaceHTML, CSS: comboboxMultiFaceCSS},
		{Label: "Open above", HTML: comboboxOpenAboveHTML},
	}
}

func comboboxMorphStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: comboboxMorphDefaultHTML},
		{Label: "Add more options", HTML: comboboxMorphMoreHTML},
		{Label: "Disable some options", HTML: comboboxMorphDisabledHTML},
	}
}

//go:embed examples/combobox_morph_default.html
var comboboxMorphDefaultHTML string

//go:embed examples/combobox_morph_more.html
var comboboxMorphMoreHTML string

//go:embed examples/combobox_morph_disabled.html
var comboboxMorphDisabledHTML string

//go:embed examples/combobox_default.html
var comboboxPlaygroundDefaultHTML string

var comboboxDisabledControlHTML = renderExampleHTML(examples.ComboboxDisabledControl())

//go:embed examples/combobox_disabled_control.templ
var comboboxDisabledControlTempl string

var comboboxFormSubmissionHTML = renderExampleHTML(examples.ComboboxFormSubmission())

//go:embed examples/combobox_form.templ
var comboboxFormSubmissionTempl string

//go:embed examples/combobox_form.css
var comboboxFormSubmissionCSS string

var comboboxGroupedHTML = renderExampleHTML(examples.ComboboxGrouped())

//go:embed examples/combobox_grouped.templ
var comboboxGroupedTempl string

//go:embed examples/combobox_grouped.css
var comboboxGroupedCSS string

var comboboxMultipleHTML = renderExampleHTML(examples.ComboboxMultiple())

//go:embed examples/combobox_multiple.templ
var comboboxMultipleTempl string

var comboboxAsyncHTML = renderExampleHTML(examples.ComboboxAsync())

//go:embed examples/combobox_async.templ
var comboboxAsyncTempl string

//go:embed examples/combobox_async.css
var comboboxAsyncCSS string

var comboboxLazyHTML = renderExampleHTML(examples.ComboboxLazy())

//go:embed examples/combobox_lazy.templ
var comboboxLazyTempl string

//go:embed examples/combobox_lazy.css
var comboboxLazyCSS string

var comboboxLiveSearchHTML = renderExampleHTML(examples.ComboboxLiveSearch())

//go:embed examples/combobox_live_search.templ
var comboboxLiveSearchTempl string

//go:embed examples/combobox_live_search.css
var comboboxLiveSearchCSS string

var comboboxOpenAboveHTML = renderExampleHTML(examples.ComboboxOpenAbove())

//go:embed examples/combobox_open_above.templ
var comboboxOpenAboveTempl string

var comboboxRichOptionsHTML = renderExampleHTML(examples.ComboboxRichOptions())

//go:embed examples/combobox_rich_options.templ
var comboboxRichOptionsTempl string

//go:embed examples/combobox_rich_options.css
var comboboxRichOptionsCSS string

var comboboxCustomTriggerFaceHTML = renderExampleHTML(examples.ComboboxCustomTriggerFace())

//go:embed examples/combobox_custom_trigger_face.templ
var comboboxCustomTriggerFaceTempl string

//go:embed examples/combobox_custom_trigger_face.css
var comboboxCustomTriggerFaceCSS string

var comboboxMultiFaceHTML = renderExampleHTML(examples.ComboboxMultiFace())

//go:embed examples/combobox_multi_face.templ
var comboboxMultiFaceTempl string

//go:embed examples/combobox_multi_face.css
var comboboxMultiFaceCSS string

package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func inputGroupPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: inputGroupDefaultHTML, CSS: inputGroupDefaultCSS},
		{Label: "Currency prefix", HTML: inputGroupCurrencyHTML},
		{Label: "Search", HTML: inputGroupSearchIconHTML},
		{Label: "Search with shortcut", HTML: inputGroupSearchShortcutHTML},
		{Label: "Search with submit", HTML: inputGroupSearchSubmitHTML},
		{Label: "Toolbar", HTML: inputGroupToolbarHTML, CSS: inputGroupToolbarCSS},
		{Label: "Disabled", HTML: inputGroupDisabledHTML},
	}
}

//go:embed examples/input_group_default.html
var inputGroupDefaultHTML string

//go:embed examples/input_group_default.css
var inputGroupDefaultCSS string

var inputGroupCurrencyHTML = renderExampleHTML(examples.InputGroupCurrency())

//go:embed examples/input_group_currency.templ
var inputGroupCurrencyTempl string

var inputGroupDomainHTML = renderExampleHTML(examples.InputGroupDomain())

//go:embed examples/input_group_domain.templ
var inputGroupDomainTempl string

var inputGroupRangeHTML = renderExampleHTML(examples.InputGroupRange())

//go:embed examples/input_group_range.templ
var inputGroupRangeTempl string

var inputGroupComparisonHTML = renderExampleHTML(examples.InputGroupComparison())

//go:embed examples/input_group_comparison.templ
var inputGroupComparisonTempl string

var inputGroupSearchIconHTML = renderExampleHTML(examples.InputGroupSearchIcon())

//go:embed examples/input_group_search_icon.templ
var inputGroupSearchIconTempl string

var inputGroupSearchShortcutHTML = renderExampleHTML(examples.InputGroupSearchShortcut())

//go:embed examples/input_group_search_shortcut.templ
var inputGroupSearchShortcutTempl string

var inputGroupSearchSubmitHTML = renderExampleHTML(examples.InputGroupSearchSubmit())

//go:embed examples/input_group_search_submit.templ
var inputGroupSearchSubmitTempl string

var inputGroupDisabledHTML = renderExampleHTML(examples.InputGroupDisabled())

//go:embed examples/input_group_disabled.templ
var inputGroupDisabledTempl string

var inputGroupToolbarHTML = renderExampleHTML(examples.InputGroupToolbar())

//go:embed examples/input_group_toolbar.templ
var inputGroupToolbarTempl string

//go:embed examples/input_group_toolbar.css
var inputGroupToolbarCSS string

package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func kbdPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: kbdPlaygroundDefaultHTML},
		{Label: "Aliases", HTML: kbdAliasesHTML},
		{Label: "Variants and sizes", HTML: kbdVariantsHTML},
		{Label: "Combinations", HTML: kbdCombinationsHTML},
		{Label: "Inside buttons", HTML: kbdInsideButtonsHTML},
		{Label: "Inside an input group", HTML: kbdInputGroupHTML, CSS: kbdInputGroupCSS},
		{Label: "Shortcut reference", HTML: kbdShortcutTableHTML, CSS: kbdShortcutTableCSS},
	}
}

//go:embed examples/kbd_default.html
var kbdPlaygroundDefaultHTML string

var kbdPlatformKeysHTML = renderExampleHTML(examples.KbdPlatformKeys())

//go:embed examples/kbd_platform_keys.templ
var kbdPlatformKeysTempl string

var kbdAliasesHTML = renderExampleHTML(examples.KbdAliases())

//go:embed examples/kbd_aliases.templ
var kbdAliasesTempl string

var kbdVariantsHTML = renderExampleHTML(examples.KbdVariants())

//go:embed examples/kbd_variants.templ
var kbdVariantsTempl string

var kbdCombinationsHTML = renderExampleHTML(examples.KbdCombinations())

//go:embed examples/kbd_combinations.templ
var kbdCombinationsTempl string

var kbdInsideButtonsHTML = renderExampleHTML(examples.KbdInsideButtons())

//go:embed examples/kbd_inside_buttons.templ
var kbdInsideButtonsTempl string

var kbdInputGroupHTML = renderExampleHTML(examples.KbdInputGroup())

//go:embed examples/kbd_input_group.templ
var kbdInputGroupTempl string

//go:embed examples/kbd_input_group.css
var kbdInputGroupCSS string

var kbdShortcutTableHTML = renderExampleHTML(examples.KbdShortcutTable())

//go:embed examples/kbd_shortcut_table.templ
var kbdShortcutTableTempl string

//go:embed examples/kbd_shortcut_table.css
var kbdShortcutTableCSS string

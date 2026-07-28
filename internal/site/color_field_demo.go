package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func colorFieldPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: colorFieldPlaygroundDefaultHTML},
		{Label: "Derived hue", HTML: colorFieldDerivedHueHTML},
		{Label: "Explicit hue", HTML: colorFieldExplicitHueHTML},
		{Label: "Disabled", HTML: colorFieldDisabledHTML},
		{Label: "Custom size", HTML: colorFieldCustomSizeHTML,
			CSS: colorFieldCustomSizeCSS},
	}
}

//go:embed examples/color_field_default.html
var colorFieldPlaygroundDefaultHTML string

func colorFieldMorphStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Initial", HTML: colorFieldMorphInitialHTML, CSS: colorFieldMorphCSS},
		{Label: "Different color", HTML: colorFieldMorphColorHTML,
			CSS: colorFieldMorphCSS},
		{Label: "Disabled", HTML: colorFieldMorphDisabledHTML,
			CSS: colorFieldMorphCSS},
	}
}

//go:embed examples/color_field_morph_initial.html
var colorFieldMorphInitialHTML string

//go:embed examples/color_field_morph_color.html
var colorFieldMorphColorHTML string

//go:embed examples/color_field_morph_disabled.html
var colorFieldMorphDisabledHTML string

//go:embed examples/color_field_morph.css
var colorFieldMorphCSS string

var colorFieldDerivedHueHTML = renderExampleHTML(examples.ColorFieldDerivedHue())

//go:embed examples/color_field_derived_hue.templ
var colorFieldDerivedHueTempl string

var colorFieldExplicitHueHTML = renderExampleHTML(examples.ColorFieldExplicitHue())

//go:embed examples/color_field_explicit_hue.templ
var colorFieldExplicitHueTempl string

var colorFieldDisabledHTML = renderExampleHTML(examples.ColorFieldDisabled())

//go:embed examples/color_field_disabled.templ
var colorFieldDisabledTempl string

var colorFieldCustomSizeHTML = renderExampleHTML(examples.ColorFieldCustomSize())

//go:embed examples/color_field_custom_size.templ
var colorFieldCustomSizeTempl string

//go:embed examples/color_field_custom_size.css
var colorFieldCustomSizeCSS string

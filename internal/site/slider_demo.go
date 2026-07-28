package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func sliderPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: sliderPlaygroundDefaultHTML},
		{Label: "Marks with labels", HTML: sliderMarksHTML},
		{Label: "Dense mark labels", HTML: sliderDenseMarksHTML},
		{Label: "Clipping to range", HTML: sliderClipHTML},
		{Label: "Decimal step + dots only", HTML: sliderDecimalHTML},
		{Label: "Custom easing", HTML: sliderEasingHTML},
		{Label: "Custom anchor + thumb", HTML: sliderStarsHTML, CSS: sliderStarsCSS},
		{Label: "Value bar", HTML: sliderValueBarHTML, CSS: sliderValueBarCSS},
		{Label: "Static marks", HTML: sliderStaticMarksHTML},
		{Label: "Bare rail (no header)", HTML: sliderBareRailHTML},
		{Label: "Vertical", HTML: sliderVerticalHTML},
		{Label: "Negative range", HTML: sliderNegativeHTML},
		{Label: "Disabled", HTML: sliderDisabledHTML},
	}
}

//go:embed examples/slider_default.html
var sliderPlaygroundDefaultHTML string

func sliderMorphStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "0%", HTML: sliderMorph0HTML},
		{Label: "30%", HTML: sliderMorph30HTML},
		{Label: "100%", HTML: sliderMorph100HTML},
	}
}

//go:embed examples/slider_morph_0.html
var sliderMorph0HTML string

//go:embed examples/slider_morph_30.html
var sliderMorph30HTML string

//go:embed examples/slider_morph_100.html
var sliderMorph100HTML string

var sliderMarksHTML = renderExampleHTML(examples.SliderMarks())

//go:embed examples/slider_marks.templ
var sliderMarksTempl string

var sliderDenseMarksHTML = renderExampleHTML(examples.SliderDenseMarks())

//go:embed examples/slider_dense_marks.templ
var sliderDenseMarksTempl string

var sliderClipHTML = renderExampleHTML(examples.SliderClip())

//go:embed examples/slider_clip.templ
var sliderClipTempl string

var sliderDecimalHTML = renderExampleHTML(examples.SliderDecimal())

//go:embed examples/slider_decimal.templ
var sliderDecimalTempl string

var sliderEasingHTML = renderExampleHTML(examples.SliderEasing())

//go:embed examples/slider_easing.templ
var sliderEasingTempl string

var sliderStarsHTML = renderExampleHTML(examples.SliderStars())

//go:embed examples/slider_stars.templ
var sliderStarsTempl string

//go:embed examples/slider_stars.css
var sliderStarsCSS string

var sliderValueBarHTML = renderExampleHTML(examples.SliderValueBar())

//go:embed examples/slider_value_bar.templ
var sliderValueBarTempl string

//go:embed examples/slider_value_bar.css
var sliderValueBarCSS string

var sliderStaticMarksHTML = renderExampleHTML(examples.SliderStaticMarks())

//go:embed examples/slider_static_marks.templ
var sliderStaticMarksTempl string

var sliderBareRailHTML = renderExampleHTML(examples.SliderBareRail())

//go:embed examples/slider_bare_rail.templ
var sliderBareRailTempl string

var sliderVerticalHTML = renderExampleHTML(examples.SliderVertical())

//go:embed examples/slider_vertical.templ
var sliderVerticalTempl string

var sliderNegativeHTML = renderExampleHTML(examples.SliderNegative())

//go:embed examples/slider_negative.templ
var sliderNegativeTempl string

var sliderDisabledHTML = renderExampleHTML(examples.SliderDisabled())

//go:embed examples/slider_disabled.templ
var sliderDisabledTempl string

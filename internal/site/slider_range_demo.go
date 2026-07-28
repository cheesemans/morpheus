package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func sliderRangePlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: sliderRangePlaygroundDefaultHTML},
		{Label: "Decimal step", HTML: sliderRangeDecimalHTML},
		{Label: "Dense mark labels", HTML: sliderRangeDenseMarksHTML},
		{Label: "Negative range", HTML: sliderRangeNegativeHTML},
		{Label: "Custom easing", HTML: sliderRangeEasingHTML},
		{Label: "Custom anchor + thumb", HTML: sliderRangeStarsHTML, CSS: sliderRangeStarsCSS},
		{Label: "Vertical temperature", HTML: sliderRangeVerticalHTML},
		{Label: "Bare rail (no header)", HTML: sliderRangeBareHTML},
		{Label: "Disabled", HTML: sliderRangeDisabledHTML},
	}
}

//go:embed examples/slider_range_default.html
var sliderRangePlaygroundDefaultHTML string

func sliderRangeMorphStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Low", HTML: sliderRangeMorphLowHTML},
		{Label: "Medium", HTML: sliderRangeMorphMediumHTML},
		{Label: "High", HTML: sliderRangeMorphHighHTML},
	}
}

//go:embed examples/slider_range_morph_low.html
var sliderRangeMorphLowHTML string

//go:embed examples/slider_range_morph_medium.html
var sliderRangeMorphMediumHTML string

//go:embed examples/slider_range_morph_high.html
var sliderRangeMorphHighHTML string

var sliderRangeDecimalHTML = renderExampleHTML(examples.SliderRangeDecimal())

//go:embed examples/slider_range_decimal.templ
var sliderRangeDecimalTempl string

var sliderRangeDenseMarksHTML = renderExampleHTML(examples.SliderRangeDenseMarks())

//go:embed examples/slider_range_dense_marks.templ
var sliderRangeDenseMarksTempl string

var sliderRangeNegativeHTML = renderExampleHTML(examples.SliderRangeNegative())

//go:embed examples/slider_range_negative.templ
var sliderRangeNegativeTempl string

var sliderRangeEasingHTML = renderExampleHTML(examples.SliderRangeEasing())

//go:embed examples/slider_range_easing.templ
var sliderRangeEasingTempl string

var sliderRangeVerticalHTML = renderExampleHTML(examples.SliderRangeVertical())

//go:embed examples/slider_range_vertical.templ
var sliderRangeVerticalTempl string

var sliderRangeBareHTML = renderExampleHTML(examples.SliderRangeBare())

//go:embed examples/slider_range_bare.templ
var sliderRangeBareTempl string

var sliderRangeDisabledHTML = renderExampleHTML(examples.SliderRangeDisabled())

//go:embed examples/slider_range_disabled.templ
var sliderRangeDisabledTempl string

var sliderRangeStarsHTML = renderExampleHTML(examples.SliderRangeStars())

//go:embed examples/slider_range_stars.templ
var sliderRangeStarsTempl string

//go:embed examples/slider_range_stars.css
var sliderRangeStarsCSS string

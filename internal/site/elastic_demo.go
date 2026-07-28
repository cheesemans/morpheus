package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

//go:embed static/sim/elastic/list.js
var elasticGrowingListScript string

//go:embed static/sim/elastic/async.js
var elasticAsyncScript string

func elasticPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: elasticPlaygroundDefaultHTML},
		{Label: "Bigger content", HTML: elasticPlaygroundBiggerHTML},
	}
}

func elasticMorphStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Initial", HTML: elasticMorphInitialHTML, CSS: elasticMorphCSS},
		{Label: "Taller content", HTML: elasticMorphTallerHTML, CSS: elasticMorphCSS},
		{Label: "Shorter content", HTML: elasticMorphShorterHTML, CSS: elasticMorphCSS},
	}
}

//go:embed examples/elastic_morph_initial.html
var elasticMorphInitialHTML string

//go:embed examples/elastic_morph_taller.html
var elasticMorphTallerHTML string

//go:embed examples/elastic_morph_shorter.html
var elasticMorphShorterHTML string

//go:embed examples/elastic_morph.css
var elasticMorphCSS string

//go:embed examples/elastic_default.html
var elasticPlaygroundDefaultHTML string

//go:embed examples/elastic_bigger.html
var elasticPlaygroundBiggerHTML string

var elasticToggleRevealHTML = renderExampleHTML(examples.ElasticToggleReveal())

//go:embed examples/elastic_toggle_reveal.templ
var elasticToggleRevealTempl string

var elasticGrowingListHTML = renderExampleHTML(examples.ElasticGrowingList())

//go:embed examples/elastic_growing_list.templ
var elasticGrowingListTempl string

var elasticContentSwapHTML = renderExampleHTML(examples.ElasticContentSwap())

//go:embed examples/elastic_content_swap.templ
var elasticContentSwapTempl string

//go:embed examples/elastic_content_swap.css
var elasticContentSwapCSS string

var elasticAsyncPlaceholderHTML = renderExampleHTML(examples.ElasticAsyncPlaceholder())

//go:embed examples/elastic_async_placeholder.templ
var elasticAsyncPlaceholderTempl string

//go:embed examples/elastic_async_placeholder.css
var elasticAsyncPlaceholderCSS string

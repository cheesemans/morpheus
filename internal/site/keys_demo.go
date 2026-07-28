package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

var keysScopedHTML = renderExampleHTML(examples.KeysScoped())

//go:embed examples/keys_scoped.templ
var keysScopedTempl string

var keysSequenceHTML = renderExampleHTML(examples.KeysSequence())

//go:embed examples/keys_sequence.templ
var keysSequenceTempl string

var keysReactiveHTML = renderExampleHTML(examples.KeysReactive())

//go:embed examples/keys_reactive.templ
var keysReactiveTempl string

var keysGlobalHTML = renderExampleHTML(examples.KeysGlobal())

//go:embed examples/keys_global.templ
var keysGlobalTempl string

var keysForHTML = renderExampleHTML(examples.KeysFor())

//go:embed examples/keys_for.templ
var keysForTempl string

//go:embed examples/keys_for.css
var keysForCSS string

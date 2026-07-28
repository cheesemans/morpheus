package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

var persistScrollHTML = renderExampleHTML(examples.PersistScroll())

//go:embed examples/persist_scroll.templ
var persistScrollTempl string

var persistTextareaHTML = renderExampleHTML(examples.PersistTextarea())

//go:embed examples/persist_textarea.templ
var persistTextareaTempl string

var persistMultiHTML = renderExampleHTML(examples.PersistMulti())

//go:embed examples/persist_multi.templ
var persistMultiTempl string

var persistVideoHTML = renderExampleHTML(examples.PersistVideo())

//go:embed examples/persist_video.templ
var persistVideoTempl string

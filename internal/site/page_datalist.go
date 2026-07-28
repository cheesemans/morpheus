package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

var datalistSharedHTML = renderExampleHTML(examples.DatalistShared())

//go:embed examples/datalist_shared.templ
var datalistSharedTempl string

//go:embed examples/datalist_shared.css
var datalistSharedCSS string

package site

import _ "embed"

//go:embed examples/select_async_failure.templ
var selectAsyncFailureTempl string

//go:embed examples/select_async_failure.css
var selectAsyncFailureCSS string

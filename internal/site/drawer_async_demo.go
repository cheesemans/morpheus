package site

import _ "embed"

//go:embed examples/drawer_async_failure.templ
var drawerAsyncFailureTempl string

//go:embed examples/drawer_async_failure.css
var drawerAsyncFailureCSS string

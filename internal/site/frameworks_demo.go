package site

import (
	_ "embed"

	"github.com/a-h/templ"
)

func frameworksRaw(raw string) templ.Component { return templ.Raw(raw) }

//go:embed static/sim/frameworks/accent.js
var frameworksAccentScript string

//go:embed examples/datastar_local.html
var datastarLocalHTML string

//go:embed examples/datastar_morph.html
var datastarMorphHTML string

//go:embed examples/alpine_demo.html
var alpineDemoHTML string

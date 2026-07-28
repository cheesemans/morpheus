package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func avatarPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: avatarPlaygroundDefaultHTML},
		{Label: "Image", HTML: avatarImageHTML},
		{Label: "Fallback", HTML: avatarFallbackHTML},
	}
}

//go:embed examples/avatar_default.html
var avatarPlaygroundDefaultHTML string

var avatarImageHTML = renderExampleHTML(examples.AvatarImage())

//go:embed examples/avatar_image.templ
var avatarImageTempl string

var avatarFallbackHTML = renderExampleHTML(examples.AvatarFallback())

//go:embed examples/avatar_fallback.templ
var avatarFallbackTempl string

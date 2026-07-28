package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func skeletonMorphStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Loading", HTML: skeletonMorphLoadingHTML, CSS: skeletonMorphCSS},
		{Label: "Content loaded", HTML: skeletonMorphLoadedHTML,
			CSS: skeletonMorphCSS},
	}
}

//go:embed examples/skeleton_morph_loading.html
var skeletonMorphLoadingHTML string

//go:embed examples/skeleton_morph_loaded.html
var skeletonMorphLoadedHTML string

//go:embed examples/skeleton_morph.css
var skeletonMorphCSS string

var skeletonVariantsHTML = renderExampleHTML(examples.SkeletonVariants())

//go:embed examples/skeleton_variants.templ
var skeletonVariantsTempl string

//go:embed examples/skeleton_variants.css
var skeletonVariantsCSS string

var skeletonCardHTML = renderExampleHTML(examples.SkeletonCard())

//go:embed examples/skeleton_card.templ
var skeletonCardTempl string

//go:embed examples/skeleton_card.css
var skeletonCardCSS string

var skeletonInlineHTML = renderExampleHTML(examples.SkeletonInline())

//go:embed examples/skeleton_inline.templ
var skeletonInlineTempl string

//go:embed examples/skeleton_inline.css
var skeletonInlineCSS string

var skeletonAvatarRowHTML = renderExampleHTML(examples.SkeletonAvatarRow())

//go:embed examples/skeleton_avatar_row.templ
var skeletonAvatarRowTempl string

//go:embed examples/skeleton_avatar_row.css
var skeletonAvatarRowCSS string

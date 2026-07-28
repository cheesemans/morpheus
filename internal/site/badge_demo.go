package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func badgePlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: badgePlaygroundDefaultHTML},
		{Label: "Variants", HTML: badgeVariantsHTML},
		{Label: "With icon", HTML: badgeWithIconHTML},
		{Label: "With status dot", HTML: badgeStatusDotHTML, CSS: badgeStatusDotCSS},
		{Label: "As link", HTML: badgeAsLinkHTML, CSS: badgeAsLinkCSS},
		{Label: "Counts", HTML: badgeCountsHTML, CSS: badgeCountsCSS},
	}
}

//go:embed examples/badge_default.html
var badgePlaygroundDefaultHTML string

var badgeVariantsHTML = renderExampleHTML(examples.BadgeVariants())

//go:embed examples/badge_variants.templ
var badgeVariantsTempl string

var badgeWithIconHTML = renderExampleHTML(examples.BadgeWithIcon())

//go:embed examples/badge_with_icon.templ
var badgeWithIconTempl string

var badgeStatusDotHTML = renderExampleHTML(examples.BadgeStatusDot())

//go:embed examples/badge_status_dot.templ
var badgeStatusDotTempl string

//go:embed examples/badge_status_dot.css
var badgeStatusDotCSS string

var badgeAsLinkHTML = renderExampleHTML(examples.BadgeAsLink())

//go:embed examples/badge_as_link.templ
var badgeAsLinkTempl string

//go:embed examples/badge_as_link.css
var badgeAsLinkCSS string

var badgeCountsHTML = renderExampleHTML(examples.BadgeCounts())

//go:embed examples/badge_counts.templ
var badgeCountsTempl string

//go:embed examples/badge_counts.css
var badgeCountsCSS string

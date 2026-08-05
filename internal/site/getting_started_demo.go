package site

import (
	_ "embed"
	"strings"
)

//go:embed examples/getting_started_install.html
var gettingStartedInstallSrc string

const gettingStartedInstallGo = `go get github.com/romshark/morpheus`

// gettingStartedInstall resolves the {VERSION} placeholders in the install snippet,
// pinning its CDN URLs to the release the site was built from.
func gettingStartedInstall(version string) string {
	return strings.ReplaceAll(gettingStartedInstallSrc, "{VERSION}", version)
}

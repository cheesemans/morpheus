package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

func progressPlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: progressPlaygroundDefaultHTML},
		{Label: "Marks with labels", HTML: progressMarksHTML},
		{Label: "Dense mark labels", HTML: progressDenseMarksHTML},
		{Label: "Bare rail", HTML: progressBareRailHTML},
		{Label: "Indeterminate", HTML: progressIndeterminateHTML},
		{Label: "Custom indeterminate animation", HTML: progressPingPongHTML, CSS: progressPingPongCSS},
		{Label: "Vertical", HTML: progressVerticalHTML},
	}
}

//go:embed examples/progress_default.html
var progressPlaygroundDefaultHTML string

func progressMorphStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "0%", HTML: progressMorph0HTML},
		{Label: "30%", HTML: progressMorph30HTML},
		{Label: "40%", HTML: progressMorph40HTML},
	}
}

//go:embed examples/progress_morph_0.html
var progressMorph0HTML string

//go:embed examples/progress_morph_30.html
var progressMorph30HTML string

//go:embed examples/progress_morph_40.html
var progressMorph40HTML string

var progressMarksHTML = renderExampleHTML(examples.ProgressMarks())

//go:embed examples/progress_marks.templ
var progressMarksTempl string

var progressDenseMarksHTML = renderExampleHTML(examples.ProgressDenseMarks())

//go:embed examples/progress_dense_marks.templ
var progressDenseMarksTempl string

var progressBareRailHTML = renderExampleHTML(examples.ProgressBareRail())

//go:embed examples/progress_bare_rail.templ
var progressBareRailTempl string

var progressIndeterminateHTML = renderExampleHTML(examples.ProgressIndeterminate())

//go:embed examples/progress_indeterminate.templ
var progressIndeterminateTempl string

var progressPingPongHTML = renderExampleHTML(examples.ProgressPingPong())

//go:embed examples/progress_ping_pong.templ
var progressPingPongTempl string

//go:embed examples/progress_ping_pong.css
var progressPingPongCSS string

var progressVerticalHTML = renderExampleHTML(examples.ProgressVertical())

//go:embed examples/progress_vertical.templ
var progressVerticalTempl string

var progressDownloadHTML = renderExampleHTML(examples.ProgressDownload())

//go:embed examples/progress_download.templ
var progressDownloadTempl string

const progressDownloadScript = `import sim from "/static/datasim.js";

sim.post("/progress/download", async (_ctx, sse) => {
  sse.patchSignals({ dl_value: 0, dl_label: "Uploading…" });

  for (let pct = 25; pct <= 100; pct += 25) {
    await sse.delay(1000);
    sse.patchSignals({ dl_value: pct });
  }

  sse.patchSignals({ dl_label: "Done ✓", dl_running: false });
});`

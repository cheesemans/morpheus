package site

import (
	_ "embed"

	"github.com/romshark/morpheus/internal/site/examples"
)

//go:embed static/sim/tree/loadnode.js
var treeLoadNodeScript string

type LazyNode struct {
	Path     string     `json:"path"`
	Label    string     `json:"label"`
	Children []LazyNode `json:"children,omitempty"`
}

var LazyTreeData = []LazyNode{
	{Path: "src", Label: "src", Children: []LazyNode{
		{Path: "src/components", Label: "components", Children: []LazyNode{
			{Path: "src/components/Button.tsx", Label: "Button.tsx"},
			{Path: "src/components/Card.tsx", Label: "Card.tsx"},
			{Path: "src/components/Modal.tsx", Label: "Modal.tsx"},
		}},
		{Path: "src/hooks", Label: "hooks", Children: []LazyNode{
			{Path: "src/hooks/useDebounce.ts", Label: "useDebounce.ts"},
			{Path: "src/hooks/useMediaQuery.ts", Label: "useMediaQuery.ts"},
		}},
		{Path: "src/App.tsx", Label: "App.tsx"},
	}},
	{Path: "docs", Label: "docs", Children: []LazyNode{
		{Path: "docs/intro.md", Label: "intro.md"},
		{Path: "docs/guide", Label: "guide", Children: []LazyNode{
			{Path: "docs/guide/setup.md", Label: "setup.md"},
			{Path: "docs/guide/usage.md", Label: "usage.md"},
		}},
		{Path: "docs/api.md", Label: "api.md"},
	}},
	{Path: "tests", Label: "tests", Children: []LazyNode{
		{Path: "tests/unit", Label: "unit", Children: []LazyNode{
			{Path: "tests/unit/parser.test.ts", Label: "parser.test.ts"},
			{Path: "tests/unit/router.test.ts", Label: "router.test.ts"},
		}},
		{Path: "tests/integration.test.ts", Label: "integration.test.ts"},
	}},
}

var treeFileSystemHTML = renderExampleHTML(examples.TreeFileSystem())

//go:embed examples/tree_file_system.templ
var treeFileSystemTempl string

var treeCategoriesHTML = renderExampleHTML(examples.TreeCategories())

//go:embed examples/tree_categories.templ
var treeCategoriesTempl string

var treeSingleRootHTML = renderExampleHTML(examples.TreeSingleRoot())

//go:embed examples/tree_single_root.templ
var treeSingleRootTempl string

var treeExplorerHTML = renderExampleHTML(examples.TreeExplorer())

//go:embed examples/tree_explorer.templ
var treeExplorerTempl string

var treeAsyncHTML = renderExampleHTML(examples.TreeAsync())

//go:embed examples/tree_async.templ
var treeAsyncTempl string

//go:embed examples/tree_async.css
var treeAsyncCSS string

var treeBookmarksHTML = renderExampleHTML(examples.TreeBookmarks())

//go:embed examples/tree_bookmarks.templ
var treeBookmarksTempl string

func treePlaygroundStates() []PlaygroundState {
	return []PlaygroundState{
		{Label: "Default", HTML: treePlaygroundDefaultHTML},
		{Label: "File system", HTML: treeFileSystemHTML},
		{Label: "Single root, all collapsed", HTML: treeSingleRootHTML},
	}
}

//go:embed examples/tree_default.html
var treePlaygroundDefaultHTML string

---
name: templ
description: Use when creating or editing any Templ (`.templ`) file in this repo.
paths:
  - "**/*.templ"
  - "*.templ"
---

Two parts: parser gotchas that apply to every `.templ` file, and the naming rule for wrapper `Opts` fields.

Component doc pages (`internal/site/page_*.templ`), their examples, and their playgrounds have their own conventions: see the `doc-page` skill.

## Parser gotchas

### No Go keywords at line start

Inside `.templ` element content, the words `for`, `if`, `switch`, `case`, `default`, `else`, `range` are parsed as control-flow keywords when one is the first token on a line (ignoring leading whitespace): `<p>X to do.\n\tfor a Y</p>`. Mid-line is fine, even right after an inline tag close (`<p>add <code>flag</code> for the Y</p>` parses).

Symptom: `expected nodes, but none were found: line N, col M`, often pointing at EOF.

Make sure Go keywords aren't the first token in a template line.

### Literal `=` / `"` inside `<code>`

`<code>foo="bar"</code>` is parsed as a `<code>` tag with attribute `foo="bar"` and breaks the same way. Wrap with backticks: `<code>{ `foo="bar"` }</code>`.

## Opts field naming

A wrapper `Opts` field has the same name as the custom-element attribute it sets, exported (PascalCase, since it's Go) and typed `Attr[T]`. Use the most specific type, not `Attr[string]`: a named enum for a fixed set of values, `CSSUnit` for a CSS length, `int` for a whole number (count, index, ms delay), `float64` for a value that can be fractional (a slider `value`/`min`/`max`/`step`).

- attribute `touch` becomes `Touch Attr[bool]`
- attribute `align` becomes `Align Attr[CarouselAlign]`
- attribute `spacing` becomes `Spacing Attr[CSSUnit]`
- attribute `page` becomes `Page Attr[int]`

For a layout-axis field, type it `Attr[Axis]` when the component supports only horizontal and vertical, or `Attr[Orientation]` when it also supports `grid` (e.g. navgroup, sortable). Using `Axis` keeps an unsupported `grid` unrepresentable.

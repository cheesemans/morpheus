// Runtime-derived state a component exposes to CSS lives in the element's
// CustomStateSet (`:state(name)`), never in an attribute. A morph reconciles
// the light DOM against the server template and strips every attribute the
// template doesn't carry, while keeping the same element instance: a marker
// the component derived at runtime (overlay mode, drag in flight, animation
// direction) disappears with no callback to restore it, and every rule keyed
// on it stops matching. Custom states are not part of the light DOM, so a
// morph cannot touch them, and nothing outside the component can set them.
//
// Attributes remain the right home for the authored surface: configuration the
// server sets, and interactive state (`open`, `checked`, `value`) the template
// is expected to render.
export function setState(internals: ElementInternals, name: string, on: boolean): void {
	if (on) internals.states.add(name);
	else internals.states.delete(name);
}

"use strict";

(function exposeRuntimeLabUtils(root) {
	function escapeHTML(value) {
		return String(value)
			.replaceAll("&", "&amp;")
			.replaceAll("<", "&lt;")
			.replaceAll(">", "&gt;")
			.replaceAll('"', "&quot;")
			.replaceAll("'", "&#039;");
	}

	function renderFieldRow({
		offset,
		name,
		width,
		value,
		active = false,
		label = "",
	}) {
		return `
			<div
				class="field-row${active ? " is-active" : ""}"
				data-field="${escapeHTML(name)}"
				data-offset="${offset}"
				data-width="${width}"
			>
				<span class="field-offset">+${offset}</span>
				<strong class="field-name">${escapeHTML(name)}</strong>
				<code class="field-value">${escapeHTML(value)}</code>
				<small class="field-width">${width} B${label ? ` · ${escapeHTML(label)}` : ""}</small>
				${active ? '<span class="field-state">当前字段</span>' : ""}
			</div>
		`;
	}

	const api = { escapeHTML, renderFieldRow };
	root.RuntimeLabUtils = api;
	if (typeof module !== "undefined" && module.exports) {
		module.exports = api;
	}
})(typeof globalThis !== "undefined" ? globalThis : window);

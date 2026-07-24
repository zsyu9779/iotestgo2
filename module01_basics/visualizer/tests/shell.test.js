"use strict";

const test = require("node:test");
const assert = require("node:assert/strict");
const fs = require("node:fs");
const path = require("node:path");

const {
	createInitialState,
	reduceState,
	renderAppMarkup,
	renderLabFragments,
} = require("../app.js");
const { escapeHTML } = require("../render-utils.js");

const visualizerRoot = path.resolve(__dirname, "..");

test("escapes text used in runtime structure markup", () => {
	assert.equal(
		escapeHTML('<img src=x onerror="x">'),
		"&lt;img src=x onerror=&quot;x&quot;&gt;",
	);
});

test("offline shell loads only local classic runtime-lab scripts in order", () => {
	const html = fs.readFileSync(
		path.join(visualizerRoot, "index.html"),
		"utf8",
	);
	const files = [
		"render-utils.js",
		"scenes.js",
		"slice-lab.js",
		"defer-lab.js",
		"app.js",
	];
	let previousIndex = -1;
	for (const file of files) {
		const index = html.indexOf(`src="./${file}"`);
		assert.ok(index > previousIndex, `${file} must load in dependency order`);
		previousIndex = index;
	}
	assert.doesNotMatch(html, /type="module"|https?:\/\//);
});

test("menu exposes only the two runtime laboratories", () => {
	assert.equal(typeof renderAppMarkup, "function");
	const markup = renderAppMarkup(createInitialState());
	assert.equal((markup.match(/data-lab-id=/g) || []).length, 2);
	assert.match(markup, /Slice Descriptor 与扩容/);
	assert.match(markup, /经典 _defer 链表与返回/);
	assert.doesNotMatch(markup, /先预测|揭晓答案/);
});

test("lesson shell is organized around runtime structures", () => {
	const state = reduceState(createInitialState(), {
		type: "SELECT_LAB",
		labId: "slice-runtime",
	});
	const markup = renderAppMarkup(state);
	for (const id of [
		"lab-nav",
		"structure-inspector",
		"runtime-stage",
		"operation-panel",
		"previous-button",
		"replay-button",
		"next-button",
	]) {
		assert.match(markup, new RegExp(`id="${id}"`));
	}
	assert.match(markup, /第 1 \/ 8 步/);
	assert.doesNotMatch(
		markup,
		/code-panel|prediction-choice|data-choice-id|揭晓答案/,
	);
});

test("lab renderer routes three fragments to the three workbench regions", () => {
	for (const labId of ["slice-runtime", "defer-runtime"]) {
		const state = reduceState(createInitialState(), {
			type: "SELECT_LAB",
			labId,
		});
		const fragments = renderLabFragments(state);
		assert.deepEqual(Object.keys(fragments), [
			"inspector",
			"stage",
			"operation",
		]);
		const markup = renderAppMarkup(state);
		assert.match(
			markup,
			new RegExp(
				`id="structure-inspector"[\\s\\S]*${labId === "slice-runtime" ? "runtime\\.slice" : "runtime\\._defer"}`,
			),
		);
		assert.match(markup, /id="runtime-stage"/);
		assert.match(markup, /id="operation-panel"/);
	}
});

test("first step disables previous and final step returns to directory", () => {
	let state = reduceState(createInitialState(), {
		type: "SELECT_LAB",
		labId: "defer-runtime",
	});
	let markup = renderAppMarkup(state);
	assert.match(markup, /id="previous-button"[^>]*disabled/);

	state = { ...state, stepIndex: 7 };
	markup = renderAppMarkup(state);
	assert.match(markup, /id="next-button"[\s\S]*返回实验室目录/);
});

test("stylesheet defines three-column and narrow-screen runtime layouts", () => {
	const css = fs.readFileSync(
		path.join(visualizerRoot, "styles.css"),
		"utf8",
	);
	assert.match(css, /\.runtime-workbench/);
	assert.match(
		css,
		/grid-template-columns:[^;]*minmax\(16rem,[^;]*minmax\(30rem,[^;]*minmax\(16rem/s,
	);
	for (const selector of [
		".field-row",
		".memory-cell",
		".descriptor-bounds",
		".defer-node",
		".nil-node",
		".runtime-pipeline",
	]) {
		assert.match(css, new RegExp(selector.replace(".", "\\.")));
	}
	assert.match(css, /@media\s*\(max-width:\s*900px\)/);
	assert.match(css, /prefers-reduced-motion:\s*reduce/);
	assert.match(css, /:focus-visible/);
});

test("animation styles run only during controller-owned transitions", () => {
	const css = fs.readFileSync(
		path.join(visualizerRoot, "styles.css"),
		"utf8",
	);
	assert.match(css, /#app\[data-animating="true"\]/);
});

test("module README documents the runtime structures and implementation boundary", () => {
	const readme = fs.readFileSync(
		path.resolve(visualizerRoot, "..", "README.md"),
		"utf8",
	);
	for (const term of [
		"Slice Descriptor",
		"runtime.growslice",
		"runtime._defer",
		"open-coded defer",
		"示意地址",
	]) {
		assert.match(readme, new RegExp(term.replace(".", "\\.")));
	}
	assert.doesNotMatch(readme, /四个逐步场景|先预测|揭晓答案/);
});

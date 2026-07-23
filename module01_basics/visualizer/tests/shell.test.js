"use strict";

const test = require("node:test");
const assert = require("node:assert/strict");
const fs = require("node:fs");
const path = require("node:path");

const {
	createInitialState,
	reduceState,
	renderAppMarkup,
	escapeHTML,
} = require("../app.js");

const visualizerRoot = path.resolve(__dirname, "..");

test("escapes text before inserting it into markup", () => {
	assert.equal(
		escapeHTML('<img src=x onerror="x">'),
		"&lt;img src=x onerror=&quot;x&quot;&gt;",
	);
});

test("offline HTML shell uses local classic scripts", () => {
	const html = fs.readFileSync(path.join(visualizerRoot, "index.html"), "utf8");
	assert.match(html, /lang="zh-CN"/);
	assert.match(html, /id="app"/);
	assert.match(html, /src="\.\/scenes\.js"/);
	assert.match(html, /src="\.\/app\.js"/);
	assert.doesNotMatch(html, /type="module"/);
	assert.doesNotMatch(html, /https?:\/\//);
});

test("menu markup exposes four native scene buttons", () => {
	const markup = renderAppMarkup({
		sceneId: null,
		stepIndex: 0,
		predictionChoice: null,
		predictionRevealed: false,
		isAnimating: false,
		replayNonce: 0,
		error: null,
	});
	assert.equal((markup.match(/data-scene-id=/g) || []).length, 4);
	assert.equal((markup.match(/<button/g) || []).length, 4);
	assert.match(markup, /Slice/);
	assert.match(markup, /Defer/);
});

test("lesson shell contains code, stage, note, progress, and controls", () => {
	const state = reduceState(createInitialState(), {
		type: "SELECT_SCENE",
		sceneId: "slice-shared",
	});
	const markup = renderAppMarkup(state);

	for (const id of [
		"scene-nav",
		"code-panel",
		"stage",
		"teaching-note",
		"previous-button",
		"replay-button",
		"next-button",
	]) {
		assert.match(markup, new RegExp(`id="${id}"`));
	}
	assert.match(markup, /aria-current="step"/);
	assert.match(markup, /1 \/ 6/);
});

test("prediction markup separates choice, reveal, and next actions", () => {
	let state = reduceState(createInitialState(), {
		type: "SELECT_SCENE",
		sceneId: "slice-shared",
	});
	state = { ...state, stepIndex: 2 };

	let markup = renderAppMarkup(state);
	assert.equal((markup.match(/data-choice-id=/g) || []).length, 2);
	assert.match(markup, /data-action="reveal"[^>]*disabled/);
	assert.match(markup, /id="next-button"[^>]*disabled/);

	state = reduceState(state, {
		type: "SELECT_PREDICTION",
		choiceId: "changes",
	});
	markup = renderAppMarkup(state);
	assert.doesNotMatch(markup, /data-action="reveal"[^>]*disabled/);
	assert.match(markup, /id="next-button"[^>]*disabled/);

	state = reduceState(state, { type: "REVEAL" });
	markup = renderAppMarkup(state);
	assert.match(markup, /回答正确/);
	assert.match(markup, /同一个底层数组单元/);
	assert.doesNotMatch(markup, /id="next-button"[^>]*disabled/);
});

test("animation lock is controller-owned so unlocking does not rerender the stage", () => {
	let state = reduceState(createInitialState(), {
		type: "SELECT_SCENE",
		sceneId: "slice-shared",
	});
	state = { ...state, stepIndex: 1 };
	state = reduceState(state, { type: "ANIMATION_START" });
	const markup = renderAppMarkup(state);

	assert.doesNotMatch(markup, /id="previous-button"[^>]*disabled/);
	assert.doesNotMatch(markup, /id="replay-button"[^>]*disabled/);
	assert.doesNotMatch(markup, /id="next-button"[^>]*disabled/);
});

test("stylesheet defines projection and narrow-screen layouts", () => {
	const css = fs.readFileSync(path.join(visualizerRoot, "styles.css"), "utf8");
	assert.match(css, /\.lesson-grid/);
	assert.match(css, /grid-template-columns:\s*minmax\(18rem,\s*34%\)/);
	assert.match(
		css,
		/\.code-line\s*\{[^}]*column-gap:\s*0\.75rem/s,
	);
	assert.match(css, /@media\s*\(max-width:\s*880px\)/);
	assert.match(css, /prefers-reduced-motion:\s*reduce/);
	assert.match(css, /:focus-visible/);
});

test("module README documents the offline classroom entry point", () => {
	const readme = fs.readFileSync(
		path.resolve(visualizerRoot, "..", "README.md"),
		"utf8",
	);
	assert.match(readme, /## Slice 与 Defer 原理动画/);
	assert.match(readme, /visualizer\/index\.html/);
	assert.match(readme, /无需后端或网络/);
	assert.match(readme, /→.*下一步/);
	assert.match(readme, /←.*上一步/);
	assert.match(readme, /R.*重播/);
	assert.match(readme, /先预测/);
});

"use strict";

const test = require("node:test");
const assert = require("node:assert/strict");

const scenes = require("../scenes.js");
const { renderStageMarkup } = require("../app.js");

function scene(id) {
	const result = scenes.find((item) => item.id === id);
	assert.ok(result, `scene ${id} must exist`);
	return result;
}

function stepAtPhase(target, phase) {
	const result = target.steps.find((item) => item.stage.phase === phase);
	assert.ok(result, `phase ${phase} must exist`);
	return result;
}

test("shared Slice reveal renders one mutated backing cell and two headers", () => {
	const target = scene("slice-shared");
	const markup = renderStageMarkup(target, stepAtPhase(target, "mutated"));

	assert.match(markup, /role="img"/);
	assert.match(markup, /data-slice-id="base"/);
	assert.match(markup, /data-slice-id="view"/);
	assert.match(
		markup,
		/data-index="1"\s+data-state="mutated"/,
	);
	assert.match(markup, />9</);
	assert.match(markup, /base\[1\]/);
	assert.match(markup, /view\[0\]/);
});

test("shared Slice view exposes independent len and cap values", () => {
	const target = scene("slice-shared");
	const markup = renderStageMarkup(target, stepAtPhase(target, "view"));

	assert.match(markup, /base/);
	assert.match(markup, /len\s*=\s*4/);
	assert.match(markup, /cap\s*=\s*4/);
	assert.match(markup, /view/);
	assert.match(markup, /len\s*=\s*2/);
	assert.match(markup, /cap\s*=\s*3/);
});

test("append reallocation keeps old and new arrays distinct", () => {
	const target = scene("slice-append");
	const markup = renderStageMarkup(
		target,
		stepAtPhase(target, "reallocated"),
	);

	assert.match(markup, /data-array-id="full-array"/);
	assert.match(markup, /data-array-id="grown-array"/);
	assert.match(markup, /容量 ≥ 3/);
	assert.doesNotMatch(markup, /容量 = 4/);
	assert.match(markup, /复制到新存储/);
});

test("append verification shows independent old and new values", () => {
	const target = scene("slice-append");
	const markup = renderStageMarkup(target, stepAtPhase(target, "verified"));

	assert.match(markup, /data-array-id="full-array"/);
	assert.match(markup, />7</);
	assert.match(markup, /data-array-id="grown-array"/);
	assert.match(markup, />9</);
	assert.match(markup, /不再共享/);
});

test("every Slice step renders a bounded accessible stage", () => {
	for (const id of ["slice-shared", "slice-append"]) {
		const target = scene(id);
		for (const step of target.steps) {
			const markup = renderStageMarkup(target, step);
			assert.match(markup, /class="stage-visual/);
			assert.match(markup, /aria-label="/);
			assert.doesNotMatch(markup, /undefined|null/);
		}
	}
});

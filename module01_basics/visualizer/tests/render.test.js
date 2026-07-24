"use strict";

const test = require("node:test");
const assert = require("node:assert/strict");

const labs = require("../scenes.js");

let renderSliceLab;
try {
	({ renderSliceLab } = require("../slice-lab.js"));
} catch {
	renderSliceLab = undefined;
}

function lab(id) {
	const result = labs.find((item) => item.id === id);
	assert.ok(result, `lab ${id} must exist`);
	return result;
}

function step(labId, stepId) {
	const result = lab(labId).steps.find((item) => item.id === stepId);
	assert.ok(result, `step ${labId}/${stepId} must exist`);
	return result;
}

function markup(rendered) {
	return `${rendered.inspector}${rendered.stage}${rendered.operation}`;
}

test("exports the Slice runtime laboratory renderer", () => {
	assert.equal(typeof renderSliceLab, "function");
});

test("descriptor step renders the real arm64 three-word layout", () => {
	const html = markup(renderSliceLab(step("slice-runtime", "descriptor")));
	assert.match(html, /data-structure="runtime\.slice"/);
	assert.match(html, /data-byte-size="24"/);
	assert.match(
		html,
		/data-field="array"[^>]*data-offset="0"[^>]*data-width="8"/,
	);
	assert.match(
		html,
		/data-field="len"[^>]*data-offset="8"[^>]*data-width="8"/,
	);
	assert.match(
		html,
		/data-field="cap"[^>]*data-offset="16"[^>]*data-width="8"/,
	);
	assert.match(html, /Go 1\.25\.11 · darwin\/arm64/);
	assert.match(html, /示意地址/);
});

test("memory cells expose address, offset, value, and concrete targets", () => {
	const html = markup(renderSliceLab(step("slice-runtime", "address")));
	assert.match(html, /data-memory-target="base-array-2"/);
	assert.match(html, /data-address="0x1010"/);
	assert.match(html, /\+16 B/);
	assert.match(html, />3</);
	assert.match(html, /0x1000 \+ 2 × 8 B = 0x1010/);
});

test("reslice step exposes the pointer and bound formulas", () => {
	const html = markup(renderSliceLab(step("slice-runtime", "reslice")));
	assert.match(html, /new\.array = old\.array \+ low × elementSize/);
	assert.match(html, /new\.len = high - low/);
	assert.match(html, /new\.cap = oldCap - low/);
	assert.match(html, /data-descriptor-id="old"/);
	assert.match(html, /data-descriptor-id="view"/);
	assert.match(html, /0x1008/);
	assert.match(html, /data-memory-target="base-array-1"/);
});

test("append fast path explicitly bypasses growslice", () => {
	const html = markup(
		renderSliceLab(step("slice-runtime", "append-in-place")),
	);
	assert.match(html, /data-branch="reuse"/);
	assert.match(html, /newLen &lt;= cap/);
	assert.match(html, /不调用 runtime\.growslice/);
	assert.match(html, /data-state="written"/);
	assert.match(html, /data-memory-target="spare-array-2"/);
});

test("growth path names the real runtime operations", () => {
	for (const [id, token] of [
		["growslice", "runtime.growslice"],
		["allocate", "mallocgc"],
		["copy", "memmove"],
	]) {
		assert.match(
			markup(renderSliceLab(step("slice-runtime", id))),
			new RegExp(token),
		);
	}
});

test("pipeline identifies active, completed, and pending runtime calls", () => {
	const html = markup(renderSliceLab(step("slice-runtime", "allocate")));
	assert.match(
		html,
		/data-call="runtime\.growslice"[^>]*data-status="done"/,
	);
	assert.match(
		html,
		/data-call="mallocgc"[^>]*data-status="active"/,
	);
	assert.match(
		html,
		/data-call="memmove"[^>]*data-status="pending"/,
	);
	assert.match(html, /不是语言保证/);
});

test("final growth step keeps old and new allocations distinct", () => {
	const html = markup(
		renderSliceLab(step("slice-runtime", "return-descriptor")),
	);
	assert.match(html, /data-allocation="old"/);
	assert.match(html, /data-allocation="new"/);
	assert.match(html, /data-descriptor-id="old"/);
	assert.match(html, /data-descriptor-id="grown"/);
	assert.match(html, /旧内存仅在不可达后才有资格被 GC 回收/);
});

test("every Slice step returns all three named fragments", () => {
	for (const current of lab("slice-runtime").steps) {
		const rendered = renderSliceLab(current);
		assert.deepEqual(Object.keys(rendered), [
			"inspector",
			"stage",
			"operation",
		]);
		for (const fragment of Object.values(rendered)) {
			assert.equal(typeof fragment, "string");
			assert.doesNotMatch(fragment, /undefined|null/);
		}
	}
});

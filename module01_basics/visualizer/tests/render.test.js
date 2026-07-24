"use strict";

const test = require("node:test");
const assert = require("node:assert/strict");

const labs = require("../scenes.js");

let renderSliceLab;
let renderDeferLab;
try {
	({ renderSliceLab } = require("../slice-lab.js"));
} catch {
	renderSliceLab = undefined;
}
try {
	({ renderDeferLab } = require("../defer-lab.js"));
} catch {
	renderDeferLab = undefined;
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
	assert.match(html, /mallocgc\(32 B, nil, false\)/);
});

test("growslice copies old values before the append caller writes the new value", () => {
	const copy = markup(renderSliceLab(step("slice-runtime", "copy")));
	assert.doesNotMatch(copy, /class="cell-value">30</);
	assert.match(copy, /复制源/);
	assert.match(copy, /已复制/);

	const returned = markup(
		renderSliceLab(step("slice-runtime", "return-descriptor")),
	);
	assert.match(returned, /caller store/);
	assert.match(returned, /s\[2\] = 30/);
	assert.match(returned, /class="memory-cell is-written/);
});

test("Slice bounds use half-open intervals", () => {
	const html = markup(renderSliceLab(step("slice-runtime", "descriptor")));
	assert.match(html, /len 范围：\[0, 4\)/);
	assert.match(html, /cap 范围：\[0, 4\)/);
	assert.doesNotMatch(html, /0…4/);
});

test("declared active targets receive visible non-color labels", () => {
	const address = markup(renderSliceLab(step("slice-runtime", "address")));
	assert.match(
		address,
		/data-memory-target="base-array-2"[\s\S]*正在读取/,
	);

	const allocate = markup(renderSliceLab(step("slice-runtime", "allocate")));
	assert.match(
		allocate,
		/data-array-id="new-array"[^>]*data-active-target="true"/,
	);
	assert.match(allocate, /新分配/);
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

test("exports the classic Defer linked-list renderer", () => {
	assert.equal(typeof renderDeferLab, "function");
});

test("classic defer nodes expose the selected runtime fields", () => {
	const html = markup(
		renderDeferLab(step("defer-runtime", "register-d1")),
	);
	assert.match(html, /data-structure="runtime\._defer"/);
	for (const field of ["heap", "sp", "pc", "fn", "link"]) {
		assert.match(html, new RegExp(`data-field="${field}"`));
	}
	assert.match(html, /rangefunc/);
	assert.match(html, /head/);
	assert.match(html, /本场景不展开/);
	assert.match(html, /data-field="heap"[\s\S]*true/);
});

test("empty defer frame renders a single terminal nil node", () => {
	const html = renderDeferLab(step("defer-runtime", "frame")).stage;
	assert.equal((html.match(/class="nil-node"/g) || []).length, 1);
});

test("second registration forms a real head-linked list", () => {
	const html = markup(
		renderDeferLab(step("defer-runtime", "register-d2")),
	);
	assert.match(html, /data-g-defer-head="0x3100"/);
	assert.match(
		html,
		/data-node-id="D2"[^>]*data-address="0x3100"/,
	);
	assert.match(html, /data-link-target="0x3000"/);
	assert.match(html, /data-chain="D2-&gt;D1-&gt;nil"/);
	assert.match(html, /D2[\s\S]*D1[\s\S]*nil/);
});

test("stack frame binds defer records by SP", () => {
	const html = markup(
		renderDeferLab(step("defer-runtime", "deferreturn")),
	);
	assert.match(html, /data-frame="work"/);
	assert.match(html, /data-register="sp"[\s\S]*0x8f00/);
	assert.match(html, /data-register="return-pc"[\s\S]*0x401240/);
	assert.match(html, /match head\.sp with frame\.SP/);
});

test("deferreturn removes the head before invoking fn", () => {
	const d2 = markup(
		renderDeferLab(step("defer-runtime", "execute-d2")),
	);
	assert.match(d2, /data-g-defer-head="0x3000"/);
	assert.match(
		d2,
		/data-node-id="D2"[^>]*data-status="executing"/,
	);
	assert.match(d2, /head = D2\.link/);
	assert.match(d2, /call D2\.fn/);
	assert.match(d2, /历史快照：_defer 已由 popDefer 清理/);

	const d1 = markup(
		renderDeferLab(step("defer-runtime", "execute-d1")),
	);
	assert.match(d1, /data-g-defer-head="nil"/);
	assert.match(
		d1,
		/data-node-id="D2"[^>]*data-status="detached"/,
	);
	assert.match(
		d1,
		/data-node-id="D1"[^>]*data-status="executing"/,
	);
});

test("completed return explains the structural source of LIFO", () => {
	const html = markup(renderDeferLab(step("defer-runtime", "return")));
	assert.match(html, /data-g-defer-head="nil"/);
	assert.match(html, /链表已清空/);
	assert.match(html, /头插、头取/);
});

test("panic step is visibly bounded without expanding recover", () => {
	const html = markup(
		renderDeferLab(step("defer-runtime", "panic-entry")),
	);
	assert.match(html, /data-panic-entry="true"/);
	assert.match(html, /_panic\.nextDefer\(\)/);
	assert.doesNotMatch(html, /_panic\.next\(\)/);
	assert.match(html, /不展开 _panic 与 recover/);
});

test("every Defer step carries the classic implementation boundary", () => {
	for (const current of lab("defer-runtime").steps) {
		const rendered = renderDeferLab(current);
		const html = markup(rendered);
		assert.deepEqual(Object.keys(rendered), [
			"inspector",
			"stage",
			"operation",
		]);
		assert.match(html, /经典 _defer 链表实现模型/);
		assert.match(html, /open-coded defer/);
		assert.doesNotMatch(html, /undefined/);
	}
});

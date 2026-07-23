# Slice 与 Defer 运行时结构动画 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 将现有代码示例辅助动画替换为两个可离线、由讲师逐步控制的运行时结构实验室，直接展示 Slice Descriptor、扩容调用链和经典 `_defer` 链表。

**Architecture:** 保留无构建依赖的经典 JavaScript 页面，但把当前超大的渲染器拆为纯状态控制器、Slice 渲染器和 Defer 渲染器。`scenes.js` 保存两个实验室的完整声明式状态；每一步都从完整状态重新生成结构舞台，避免依赖前一步 DOM。

**Tech Stack:** HTML5、CSS3、原生经典 JavaScript、SVG、Node.js 内置 `node:test`、Go 1.25.11 runtime 源码作为结构核对基准、浏览器人工验收

## Global Constraints

- Slice 结构以 Go 1.25.11、darwin/arm64 为基准：`array unsafe.Pointer`、`len int`、`cap int`，偏移为 0、8、16，总大小 24 字节。
- 页面地址必须标注为示意地址；字段布局和指针计算必须与上述基准一致。
- 扩容展示 `runtime.growslice`、容量计算、`mallocgc`、`memmove` 和新 Descriptor，不把具体扩容倍率描述为语言保证。
- Defer 使用用户选定的经典 `_defer` 链表模型，并持续提示现代 Go 可能使用 open-coded defer。
- `_defer` 主体字段为 `heap / sp / pc / fn / link`；`rangefunc / head` 标注为本场景不展开。
- 页面只保留 Slice 和 Defer 两个实验室，不保留代码面板、代码高亮、预测题或输出结果流程。
- 页面无需后端、网络、包安装或构建命令即可打开。
- 单步动画不自动连播；支持上一步、下一步、重播、返回目录和 `prefers-reduced-motion`。
- 1280×720 与 1440×900 使用三栏；390×844 纵向排列且无页面横向溢出。
- 颜色不能成为唯一信息来源；地址、字段、指针和状态必须同时有文字或形状标签。

---

## File Map

- Modify `module01_basics/visualizer/scenes.js`: 用两个运行时实验室及其完整步骤替换四个代码场景。
- Modify `module01_basics/visualizer/app.js`: 简化为实验室状态机、通用壳渲染和浏览器控制器。
- Create `module01_basics/visualizer/render-utils.js`: 浏览器与 Node 共用的安全转义和结构字段渲染工具。
- Create `module01_basics/visualizer/slice-lab.js`: Slice Descriptor、连续内存和扩容调用链渲染器。
- Create `module01_basics/visualizer/defer-lab.js`: goroutine、栈帧和 `_defer` 链表渲染器。
- Modify `module01_basics/visualizer/index.html`: 按依赖顺序加载五个本地经典脚本。
- Replace `module01_basics/visualizer/styles.css`: 三栏运行时实验室、内存网格、链表和响应式样式。
- Modify `module01_basics/visualizer/tests/state.test.js`: 两实验室契约和无预测状态流。
- Replace `module01_basics/visualizer/tests/render.test.js`: Slice 与 Defer 真实结构标记测试。
- Modify `module01_basics/visualizer/tests/shell.test.js`: 离线脚本、页面壳和旧 UI 移除测试。
- Modify `module01_basics/README.md`: 更新入口说明、结构真实性边界和课堂使用方法。

---

### Task 1: 替换场景契约与状态机

**Files:**
- Modify: `module01_basics/visualizer/scenes.js`
- Modify: `module01_basics/visualizer/app.js`
- Modify: `module01_basics/visualizer/tests/state.test.js`

**Interfaces:**
- Produces: `RUNTIME_LABS: RuntimeLab[]`
- Produces: `createInitialState(): { labId, stepIndex, isAnimating, replayNonce, error }`
- Produces: `reduceState(state, action): AppState`
- Produces: `getCurrentLab(state): RuntimeLab | null`
- Produces: `getCurrentStep(state): RuntimeStep | null`
- `Action.type` is one of `SELECT_LAB`, `NEXT`, `PREVIOUS`, `REPLAY`, `RESET_TO_MENU`, `ANIMATION_START`, `ANIMATION_END`, `SET_ERROR`.

- [ ] **Step 1: 用两实验室契约测试替换四场景测试**

在 `tests/state.test.js` 中建立以下断言：

```js
test("defines only the two runtime laboratories", () => {
	assert.deepEqual(
		labs.map((lab) => lab.id),
		["slice-runtime", "defer-runtime"],
	);
	assert.deepEqual(
		labs.map((lab) => lab.kind),
		["slice", "defer"],
	);
	for (const lab of labs) {
		assert.equal(lab.steps.length, 8);
		assert.equal(new Set(lab.steps.map((step) => step.id)).size, 8);
		for (const step of lab.steps) {
			assert.equal(typeof step.title, "string");
			assert.equal(typeof step.operation.name, "string");
			assert.equal(typeof step.operation.summary, "string");
			assert.equal(typeof step.state.phase, "string");
		}
	}
});

test("runtime labs contain no code-example or prediction contract", () => {
	const serialized = JSON.stringify(labs);
	assert.doesNotMatch(serialized, /prediction|choices|correctChoiceId/);
	for (const lab of labs) assert.equal("code" in lab, false);
});
```

把 reducer 测试改为：

```js
test("teacher controls move deterministically through a lab", () => {
	let state = reduceState(createInitialState(), {
		type: "SELECT_LAB",
		labId: "slice-runtime",
	});
	assert.equal(state.stepIndex, 0);
	state = reduceState(state, { type: "NEXT" });
	assert.equal(state.stepIndex, 1);
	state = reduceState(state, { type: "PREVIOUS" });
	assert.equal(state.stepIndex, 0);
	state = {
		...state,
		stepIndex: getCurrentLab(state).steps.length - 1,
	};
	assert.equal(reduceState(state, { type: "NEXT" }).labId, null);
});
```

- [ ] **Step 2: 运行状态测试并确认 RED**

Run:

```bash
/Users/zhangshiyu/.cache/codex-runtimes/codex-primary-runtime/dependencies/node/bin/node \
  --test module01_basics/visualizer/tests/state.test.js
```

Expected: FAIL，旧数据仍包含四个场景和预测字段，且没有 `SELECT_LAB`。

- [ ] **Step 3: 写入两个实验室的步骤骨架**

`scenes.js` 使用现有浏览器/Node 暴露模式，Slice 步骤 ID 必须是：

```js
[
	"descriptor",
	"address",
	"reslice",
	"append-in-place",
	"growslice",
	"allocate",
	"copy",
	"return-descriptor",
]
```

Defer 步骤 ID 必须是：

```js
[
	"frame",
	"register-d1",
	"register-d2",
	"deferreturn",
	"execute-d2",
	"execute-d1",
	"return",
	"panic-entry",
]
```

每一步使用完整状态。例如 Slice 重新切片步骤：

```js
{
	id: "reslice",
	title: "生成新的 Slice Descriptor",
	operation: {
		name: "reslice",
		call: "new.array = old.array + low × elementSize",
		summary: "Descriptor 被复制并重新计算；底层连续内存没有复制。",
	},
	state: {
		phase: "reslice",
		elementSize: 8,
		descriptors: [
			{ id: "old", address: "0x7000", array: "0x1000", length: 4, capacity: 4 },
			{ id: "view", address: "0x7020", array: "0x1008", length: 2, capacity: 3 },
		],
		arrays: [
			{ id: "old-array", baseAddress: "0x1000", values: [1, 2, 3, 4], capacity: 4 },
		],
		activeTargets: ["view.array", "old-array[1]"],
	},
}
```

Defer 第二次注册步骤的完整链表状态：

```js
{
	id: "register-d2",
	title: "D2 插入链表头部",
	operation: {
		name: "deferproc",
		call: "D2.link = g._defer; g._defer = D2",
		summary: "头插法让后注册的记录成为第一个待处理节点。",
	},
	state: {
		phase: "register-d2",
		goroutine: { id: "g17", deferHead: "0x3100" },
		frame: { id: "work", sp: "0x8f00", returnPC: "0x401240" },
		nodes: [
			{ id: "D2", address: "0x3100", heap: false, sp: "0x8f00", pc: "0x401180", fn: "deferredFnB", link: "0x3000", status: "head" },
			{ id: "D1", address: "0x3000", heap: false, sp: "0x8f00", pc: "0x401140", fn: "deferredFnA", link: null, status: "linked" },
		],
	},
}
```

- [ ] **Step 4: 简化 reducer**

`app.js` 中删除预测状态和动作。状态必须为：

```js
function createInitialState() {
	return {
		labId: null,
		stepIndex: 0,
		isAnimating: false,
		replayNonce: 0,
		error: null,
	};
}
```

`SELECT_LAB` 校验实验室 ID，`NEXT` 在末步返回目录，动画锁只允许
`ANIMATION_END` 和 `SELECT_LAB`：

```js
if (
	state.isAnimating &&
	!["ANIMATION_END", "SELECT_LAB"].includes(action.type)
) {
	return state;
}
```

- [ ] **Step 5: 运行状态测试并确认 GREEN**

Run:

```bash
/Users/zhangshiyu/.cache/codex-runtimes/codex-primary-runtime/dependencies/node/bin/node \
  --test module01_basics/visualizer/tests/state.test.js
```

Expected: 全部通过。

- [ ] **Step 6: 提交数据模型**

```bash
git add module01_basics/visualizer/scenes.js \
  module01_basics/visualizer/app.js \
  module01_basics/visualizer/tests/state.test.js
git commit -m "refactor: model runtime internals labs"
```

---

### Task 2: 实现 Slice 真实结构渲染器

**Files:**
- Create: `module01_basics/visualizer/render-utils.js`
- Create: `module01_basics/visualizer/slice-lab.js`
- Replace: `module01_basics/visualizer/tests/render.test.js`

**Interfaces:**
- Produces: `escapeHTML(value): string`
- Produces: `renderFieldRow({ offset, name, width, value, active }): string`
- Produces: `renderSliceLab(step): { inspector: string, stage: string, operation: string }`
- Consumes: Task 1 中 `step.state.descriptors`、`step.state.arrays` 和 `step.operation`。

- [ ] **Step 1: 写 Slice 结构失败测试**

`tests/render.test.js` 引入 `renderSliceLab`，并添加：

```js
function markup(rendered) {
	return `${rendered.inspector}${rendered.stage}${rendered.operation}`;
}

test("descriptor step renders the real arm64 three-word layout", () => {
	const html = markup(renderSliceLab(step("slice-runtime", "descriptor")));
	assert.match(html, /data-structure="runtime\\.slice"/);
	assert.match(html, /data-byte-size="24"/);
	assert.match(html, /data-field="array"[^>]*data-offset="0"[^>]*data-width="8"/);
	assert.match(html, /data-field="len"[^>]*data-offset="8"[^>]*data-width="8"/);
	assert.match(html, /data-field="cap"[^>]*data-offset="16"[^>]*data-width="8"/);
	assert.match(html, /示意地址/);
});

test("reslice step exposes the pointer and bound formulas", () => {
	const html = markup(renderSliceLab(step("slice-runtime", "reslice")));
	assert.match(html, /new\\.array = old\\.array \\+ low × elementSize/);
	assert.match(html, /new\\.len = high - low/);
	assert.match(html, /new\\.cap = oldCap - low/);
	assert.match(html, /0x1008/);
	assert.match(html, /data-memory-target="old-array-1"/);
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
	const finalMarkup = markup(
		renderSliceLab(step("slice-runtime", "return-descriptor")),
	);
	assert.match(finalMarkup, /data-allocation="old"/);
	assert.match(finalMarkup, /data-allocation="new"/);
	assert.match(finalMarkup, /不可达后.*GC/);
});
```

- [ ] **Step 2: 运行渲染测试并确认 RED**

Run:

```bash
/Users/zhangshiyu/.cache/codex-runtimes/codex-primary-runtime/dependencies/node/bin/node \
  --test module01_basics/visualizer/tests/render.test.js
```

Expected: FAIL，因为 `slice-lab.js` 和新导出不存在。

- [ ] **Step 3: 实现公共渲染工具**

`render-utils.js` 使用 UMD 风格暴露：

```js
(function exposeRuntimeLabUtils(root) {
	function escapeHTML(value) {
		return String(value)
			.replaceAll("&", "&amp;")
			.replaceAll("<", "&lt;")
			.replaceAll(">", "&gt;")
			.replaceAll('"', "&quot;")
			.replaceAll("'", "&#039;");
	}

	function renderFieldRow({ offset, name, width, value, active = false }) {
		return `<div class="field-row${active ? " is-active" : ""}"
			data-field="${escapeHTML(name)}"
			data-offset="${offset}"
			data-width="${width}">
			<span>+${offset}</span><strong>${escapeHTML(name)}</strong>
			<code>${escapeHTML(value)}</code><small>${width} B</small>
		</div>`;
	}

	const api = { escapeHTML, renderFieldRow };
	root.RuntimeLabUtils = api;
	if (typeof module !== "undefined" && module.exports) module.exports = api;
})(typeof globalThis !== "undefined" ? globalThis : window);
```

- [ ] **Step 4: 实现 Descriptor 和连续内存**

`slice-lab.js` 提供独立 UMD 导出。每个 Descriptor 必须生成：

```html
<section
	class="descriptor"
	data-structure="runtime.slice"
	data-byte-size="24"
	data-descriptor-id="old"
>
	<!-- array / len / cap field rows -->
</section>
```

每个内存槽必须带稳定 ID 和地址：

```html
<div
	id="old-array-1"
	class="memory-cell"
	data-memory-target="old-array-1"
	data-address="0x1008"
>
	<span>+8 B</span><strong>2</strong><code>0x1008</code>
</div>
```

使用内联 SVG `viewBox="0 0 1000 420"` 绘制 Descriptor 到具体内存槽的连接线，
同时在文字标签中重复目标地址，保证窄屏或无动画时仍能理解。

- [ ] **Step 5: 实现 append 两条路径**

容量足够时渲染：

```html
<div class="runtime-decision" data-branch="reuse">
	<code>newLen &lt;= cap</code>
	<strong>不调用 runtime.growslice</strong>
</div>
```

容量不足步骤依次渲染：

```html
<ol class="runtime-pipeline">
	<li data-call="runtime.growslice">runtime.growslice</li>
	<li data-call="nextslicecap">计算 newCap ≥ newLen</li>
	<li data-call="mallocgc">mallocgc</li>
	<li data-call="memmove">memmove</li>
	<li data-call="return">返回新 Descriptor</li>
</ol>
```

当前步骤只激活一个 `data-status="active"`，已完成步骤为
`data-status="done"`，后续步骤为 `data-status="pending"`。

- [ ] **Step 6: 运行渲染测试并确认 GREEN**

Run:

```bash
/Users/zhangshiyu/.cache/codex-runtimes/codex-primary-runtime/dependencies/node/bin/node \
  --test module01_basics/visualizer/tests/render.test.js
```

Expected: Slice 测试全部通过。

- [ ] **Step 7: 提交 Slice 渲染器**

```bash
git add module01_basics/visualizer/render-utils.js \
  module01_basics/visualizer/slice-lab.js \
  module01_basics/visualizer/tests/render.test.js
git commit -m "feat: visualize slice runtime internals"
```

---

### Task 3: 实现经典 `_defer` 链表渲染器

**Files:**
- Create: `module01_basics/visualizer/defer-lab.js`
- Modify: `module01_basics/visualizer/tests/render.test.js`

**Interfaces:**
- Produces: `renderDeferLab(step): { inspector: string, stage: string, operation: string }`
- Consumes: Task 1 中 `step.state.goroutine`、`step.state.frame`、`step.state.nodes` 和 `step.operation`。

- [ ] **Step 1: 写 `_defer` 结构失败测试**

追加：

```js
test("classic defer nodes expose the selected runtime fields", () => {
	const html = markup(renderDeferLab(step("defer-runtime", "register-d1")));
	assert.match(html, /data-structure="runtime\\._defer"/);
	for (const field of ["heap", "sp", "pc", "fn", "link"]) {
		assert.match(html, new RegExp(`data-field="${field}"`));
	}
	assert.match(html, /rangefunc/);
	assert.match(html, /head/);
	assert.match(html, /本场景不展开/);
});

test("second registration forms a real head-linked list", () => {
	const html = markup(renderDeferLab(step("defer-runtime", "register-d2")));
	assert.match(html, /data-g-defer-head="0x3100"/);
	assert.match(html, /data-node-id="D2"[^>]*data-address="0x3100"/);
	assert.match(html, /data-link-target="0x3000"/);
	assert.match(html, /D2[\s\S]*D1[\s\S]*nil/);
});

test("deferreturn removes the head before invoking fn", () => {
	const d2 = markup(renderDeferLab(step("defer-runtime", "execute-d2")));
	assert.match(d2, /data-g-defer-head="0x3000"/);
	assert.match(d2, /data-node-id="D2"[^>]*data-status="executing"/);
	assert.match(d2, /head = D2\\.link/);
	assert.match(d2, /call D2\\.fn/);

	const done = markup(renderDeferLab(step("defer-runtime", "return")));
	assert.match(done, /data-g-defer-head="nil"/);
	assert.match(done, /链表已清空/);
});

test("every defer step carries the classic implementation boundary", () => {
	for (const current of lab("defer-runtime").steps) {
		const html = markup(renderDeferLab(current));
		assert.match(html, /经典 _defer 链表实现模型/);
		assert.match(html, /open-coded defer/);
	}
});
```

- [ ] **Step 2: 运行 Defer 测试并确认 RED**

Run:

```bash
/Users/zhangshiyu/.cache/codex-runtimes/codex-primary-runtime/dependencies/node/bin/node \
  --test module01_basics/visualizer/tests/render.test.js
```

Expected: FAIL，因为 `defer-lab.js` 不存在。

- [ ] **Step 3: 渲染 goroutine 和函数栈帧**

`defer-lab.js` 使用 UMD 风格导出。goroutine 区必须包含：

```html
<section
	class="goroutine-panel"
	data-goroutine="g17"
	data-g-defer-head="0x3100"
>
	<strong>g17</strong>
	<div><code>g._defer = 0x3100</code></div>
</section>
```

栈帧必须同时显示符号名和地址：

```html
<section class="stack-frame" data-frame="work">
	<div data-register="sp"><span>SP</span><code>0x8f00</code></div>
	<div data-register="return-pc">
		<span>return PC</span><code>0x401240</code>
	</div>
</section>
```

- [ ] **Step 4: 渲染 `_defer` 节点与链线**

每个节点使用：

```html
<article
	class="defer-node"
	data-structure="runtime._defer"
	data-node-id="D2"
	data-address="0x3100"
	data-status="head"
>
	<!-- heap / sp / pc / fn / link -->
</article>
```

`link` 字段同时生成 `data-link-target` 和指向目标节点的 SVG 箭头。空链使用
独立的 `nil` 终点节点，而不是空白区域。

- [ ] **Step 5: 渲染注册与返回操作**

注册步骤的右栏操作必须展示：

```text
D2.link = g._defer
g._defer = D2
```

执行节点时必须按以下顺序显示三个原子操作：

```text
fn = D2.fn
g._defer = D2.link
call fn
```

当前节点使用 `data-status="executing"`；已经脱链的节点保留灰色轮廓并使用
`data-status="detached"`，让学生看到它不再属于链表。

- [ ] **Step 6: 运行全部渲染测试并确认 GREEN**

Run:

```bash
/Users/zhangshiyu/.cache/codex-runtimes/codex-primary-runtime/dependencies/node/bin/node \
  --test module01_basics/visualizer/tests/render.test.js
```

Expected: Slice 与 Defer 测试全部通过。

- [ ] **Step 7: 提交 Defer 渲染器**

```bash
git add module01_basics/visualizer/defer-lab.js \
  module01_basics/visualizer/tests/render.test.js
git commit -m "feat: visualize classic defer linked list"
```

---

### Task 4: 重建页面壳、三栏布局和讲师控制

**Files:**
- Modify: `module01_basics/visualizer/index.html`
- Modify: `module01_basics/visualizer/app.js`
- Replace: `module01_basics/visualizer/styles.css`
- Modify: `module01_basics/visualizer/tests/shell.test.js`

**Interfaces:**
- Consumes: `RUNTIME_LABS`、`renderSliceLab(step)`、`renderDeferLab(step)` 的三个命名片段。
- Produces: `renderAppMarkup(state): string`
- Produces DOM IDs: `lab-nav`、`structure-inspector`、`runtime-stage`、`operation-panel`、`previous-button`、`replay-button`、`next-button`。

- [ ] **Step 1: 写离线壳和旧 UI 移除失败测试**

在 `shell.test.js` 中断言脚本顺序和新布局：

```js
test("offline shell loads only local classic runtime-lab scripts", () => {
	const html = fs.readFileSync(path.join(root, "index.html"), "utf8");
	for (const file of [
		"render-utils.js",
		"scenes.js",
		"slice-lab.js",
		"defer-lab.js",
		"app.js",
	]) {
		assert.match(html, new RegExp(`src="\\\\./${file}"`));
	}
	assert.doesNotMatch(html, /type="module"|https?:\\/\\//);
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
	assert.doesNotMatch(markup, /code-panel|prediction-choice|揭晓答案/);
});
```

- [ ] **Step 2: 运行 shell 测试并确认 RED**

Run:

```bash
/Users/zhangshiyu/.cache/codex-runtimes/codex-primary-runtime/dependencies/node/bin/node \
  --test module01_basics/visualizer/tests/shell.test.js
```

Expected: FAIL，入口仍只加载旧的三个脚本，页面仍包含代码区和预测 UI。

- [ ] **Step 3: 更新 HTML 脚本顺序**

`index.html` 底部必须是：

```html
<script src="./render-utils.js"></script>
<script src="./scenes.js"></script>
<script src="./slice-lab.js"></script>
<script src="./defer-lab.js"></script>
<script src="./app.js"></script>
```

- [ ] **Step 4: 实现目录与三栏实验室壳**

目录页只生成两个 `data-lab-id` 原生按钮。实验室页结构固定为：

```html
<section class="runtime-workbench">
	<aside id="structure-inspector" class="structure-inspector"></aside>
	<section id="runtime-stage" class="runtime-stage"></section>
	<aside id="operation-panel" class="operation-panel"></aside>
</section>
```

渲染器统一返回 `{ inspector, stage, operation }`。`app.js` 必须把三个字符串
分别插入上述三个区域，不要把整个实验室再次嵌套到单一右栏中：

```js
const fragments =
	lab.kind === "slice" ? renderSliceLab(step) : renderDeferLab(step);
return `
	<section class="runtime-workbench">
		<aside id="structure-inspector">${fragments.inspector}</aside>
		<section id="runtime-stage">${fragments.stage}</section>
		<aside id="operation-panel">${fragments.operation}</aside>
	</section>
`;
```

- [ ] **Step 5: 保留确定性课堂控制**

键盘保持：

```js
function actionFromKey(key) {
	if (key === "ArrowRight" || key === " ") return { type: "NEXT" };
	if (key === "ArrowLeft") return { type: "PREVIOUS" };
	if (key === "r" || key === "R") return { type: "REPLAY" };
	return null;
}
```

动画锁时长为 900 ms；减弱动画为 0。选择实验室、前进、后退和重播才设置
`data-animating="true"`，普通焦点变化不能重播舞台动画。

- [ ] **Step 6: 重写样式**

桌面布局：

```css
.runtime-workbench {
	display: grid;
	grid-template-columns:
		minmax(16rem, 0.82fr)
		minmax(30rem, 1.75fr)
		minmax(16rem, 0.9fr);
	gap: 1rem;
}
```

样式必须包含：

- `.field-row` 的固定字节网格。
- `.memory-cell` 的地址、偏移和值三层标签。
- `.descriptor-bounds` 的 len/cap 括号线。
- `.defer-node`、`.nil-node` 和 `.link-arrow`。
- `.runtime-pipeline [data-status]`。
- `.is-reading`、`.is-writing`、`.is-allocated`、`.is-detached` 的文字徽标。
- `@media (max-width: 900px)` 切换为单栏。
- `@media (prefers-reduced-motion: reduce)` 清除动画和过渡。

- [ ] **Step 7: 运行所有 Node 测试并确认 GREEN**

Run:

```bash
/Users/zhangshiyu/.cache/codex-runtimes/codex-primary-runtime/dependencies/node/bin/node \
  --test module01_basics/visualizer/tests/*.test.js
```

Expected: 全部通过。

- [ ] **Step 8: 提交新页面壳**

```bash
git add module01_basics/visualizer/index.html \
  module01_basics/visualizer/app.js \
  module01_basics/visualizer/styles.css \
  module01_basics/visualizer/tests/shell.test.js
git commit -m "feat: rebuild runtime laboratory interface"
```

---

### Task 5: 更新课堂文档并完成浏览器验收

**Files:**
- Modify: `module01_basics/README.md`
- Modify: `module01_basics/visualizer/tests/shell.test.js`

**Interfaces:**
- Documents: 离线入口、两个实验室、经典 Defer 边界、键盘控制。
- Verifies: 1280×720、1440×900、390×844 三档视口和所有 16 个步骤。

- [ ] **Step 1: 写 README 失败测试**

```js
test("README describes the runtime internals laboratories", () => {
	const readme = fs.readFileSync(
		path.resolve(root, "..", "README.md"),
		"utf8",
	);
	assert.match(readme, /Slice Descriptor/);
	assert.match(readme, /runtime\\.growslice/);
	assert.match(readme, /_defer/);
	assert.match(readme, /经典.*链表/);
	assert.match(readme, /open-coded defer/);
	assert.doesNotMatch(readme, /先预测|揭晓答案/);
});
```

- [ ] **Step 2: 运行测试并确认 RED**

Run:

```bash
/Users/zhangshiyu/.cache/codex-runtimes/codex-primary-runtime/dependencies/node/bin/node \
  --test module01_basics/visualizer/tests/shell.test.js
```

Expected: FAIL，README 仍描述四个代码语义场景和预测流程。

- [ ] **Step 3: 更新 README**

说明必须包含：

- 直接打开 `visualizer/index.html`，无需后端或网络。
- Slice 实验室展示 24 字节 Descriptor、指针计算和
  `runtime.growslice → mallocgc → memmove`。
- Defer 实验室展示经典 `_defer` 链表；现代 Go 可能使用
  open-coded defer。
- `→/空格`、`←` 和 `R` 的课堂控制。
- 十六进制地址是示意地址。

- [ ] **Step 4: 运行自动化回归**

Run:

```bash
/Users/zhangshiyu/.cache/codex-runtimes/codex-primary-runtime/dependencies/node/bin/node \
  --test module01_basics/visualizer/tests/*.test.js
go test ./...
make module01-audit
git diff --check
```

Expected: 四条命令均退出 0，Node 输出 0 failures，Module01 输出
`module01 demo contracts: PASS` 和 `module01 teaching failures: PASS`。

- [ ] **Step 5: 浏览器逐步验收**

使用本地静态服务器打开页面并完成：

1. 1280×720：
   - 进入 Slice，逐步播放 8 步。
   - Descriptor 字段、内存地址和 runtime pipeline 无遮挡。
   - `return-descriptor` 同时显示新旧内存。
   - 进入 Defer，逐步播放 8 步。
   - `D2 → D1 → nil` 和两次头指针移动可见。
2. 1440×900：
   - 三栏充分利用宽度，无超出或内部不必要滚动。
3. 390×844：
   - 三个区域纵向排列。
   - `document.documentElement.scrollWidth === clientWidth`。
   - Descriptor、内存槽和节点字段可以换行，不被裁切。
4. 控制：
   - 快速双击下一步只前进一次。
   - 左箭头、右箭头、空格和 `R` 有效。
   - 减弱动画时控制器锁延迟为 0。
   - 场景切换清除当前步骤和动画状态。
5. 浏览器控制台没有应用错误。

- [ ] **Step 6: 提交文档与验收修复**

```bash
git add module01_basics/README.md \
  module01_basics/visualizer
git commit -m "docs: document runtime internals labs"
```

---

## Completion Checklist

- [ ] Slice 数据结构、指针计算、重新切片和 append 两条路径均有自动化结构断言。
- [ ] Defer 经典链表注册、头插、头取和版本边界均有自动化断言。
- [ ] 旧四场景、代码面板和预测状态已删除，不只是隐藏。
- [ ] 全部 16 个步骤都能前进、后退和重播。
- [ ] 三档视口浏览器验收通过。
- [ ] `go test ./...`、`make module01-audit`、Node 测试和 `git diff --check` 均通过。
- [ ] 用户原有未提交文件和 `.superpowers/` 没有被纳入提交。

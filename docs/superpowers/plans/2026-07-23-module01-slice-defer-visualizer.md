# Module01 Slice 与 Defer 原理动画页 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 在 `module01_basics` 中增加一个可离线打开、由讲师逐步控制的 Slice 与 Defer 底层原理动画页。

**Architecture:** 使用无构建依赖的 HTML、CSS 和经典 JavaScript。`scenes.js` 保存四个场景的声明式步骤；`app.js` 提供可在 Node 中测试的纯状态机，并在浏览器中将当前完整状态渲染为 DOM/SVG，避免依赖上一步残留 DOM。

**Tech Stack:** HTML5、CSS3、原生 JavaScript、SVG、Node.js 内置 `node:test`、Playwright（最终浏览器验证）

## Global Constraints

- 页面必须无需后端、网络、包安装或构建命令即可打开。
- 第一版只包含 Slice 共享、Slice append、Defer LIFO、Defer 求值时机四个场景。
- 每个场景包含一次预测，预测选择与揭晓是两个独立动作。
- 不执行用户输入的 Go 代码，不使用 `eval`，不加载 CDN，不持久化浏览器数据。
- 新数组容量只表达为“不小于新长度”，不展示固定扩容倍率或虚构地址。
- Defer 卡片区明确为语义模型，不承诺特定编译器或运行时实现。
- 主视口为 1280×720，同时支持 1440×900 和 390×844。
- 单步动画不自动连播，并支持 `prefers-reduced-motion: reduce`。
- 颜色不能成为唯一信息来源；状态必须同时使用标签、文字、形状或连线表达。

---

## File Map

- Create `module01_basics/visualizer/index.html`: 静态入口、语义化页面骨架、四场景导航和控制区。
- Create `module01_basics/visualizer/styles.css`: 投屏布局、深色主题、响应式规则、状态样式和动画。
- Create `module01_basics/visualizer/scenes.js`: 四个场景的代码、步骤、预测题、讲师提示和舞台状态。
- Create `module01_basics/visualizer/app.js`: 纯状态机、事件控制、HTML/SVG 舞台渲染、动画锁和错误恢复。
- Create `module01_basics/visualizer/tests/state.test.js`: 场景契约、状态流、预测门控和键盘动作的 Node 自动化测试。
- Create `module01_basics/visualizer/tests/render.test.js`: Slice/Defer 关键舞台标记和安全转义测试。
- Modify `module01_basics/README.md`: 增加动画页入口、使用方法和键盘说明。

---

### Task 1: 场景数据契约与纯状态机

**Files:**
- Create: `module01_basics/visualizer/scenes.js`
- Create: `module01_basics/visualizer/app.js`
- Create: `module01_basics/visualizer/tests/state.test.js`

**Interfaces:**
- Produces: `VISUALIZER_SCENES: Scene[]`
- Produces: `createInitialState(): AppState`
- Produces: `reduceState(state: AppState, action: Action): AppState`
- Produces: `getCurrentScene(state): Scene | null`
- Produces: `getCurrentStep(state): Step | null`
- Produces: `isPredictionReady(state): boolean`
- `Action.type` is one of `SELECT_SCENE`, `NEXT`, `PREVIOUS`, `SELECT_PREDICTION`, `REVEAL`, `RESET_TO_MENU`.

- [ ] **Step 1: Write the failing scene and state tests**

Create `module01_basics/visualizer/tests/state.test.js` with Node's built-in test runner:

```js
const test = require("node:test");
const assert = require("node:assert/strict");

const scenes = require("../scenes.js");
const {
  createInitialState,
  reduceState,
  getCurrentScene,
  getCurrentStep,
  isPredictionReady,
} = require("../app.js");

test("defines the four approved classroom scenes", () => {
  assert.deepEqual(
    scenes.map((scene) => scene.id),
    ["slice-shared", "slice-append", "defer-lifo", "defer-evaluation"],
  );
  for (const scene of scenes) {
    assert.ok(scene.code.length >= 4);
    assert.ok(scene.steps.length >= 6);
    assert.equal(scene.steps.filter((step) => step.kind === "prediction").length, 1);
    assert.equal(scene.steps.at(-1).kind, "conclusion");
  }
});

test("prediction cannot advance until it is selected and revealed", () => {
  let state = reduceState(createInitialState(), {
    type: "SELECT_SCENE",
    sceneId: "slice-shared",
  });
  const scene = getCurrentScene(state);
  const predictionIndex = scene.steps.findIndex((step) => step.kind === "prediction");
  state = { ...state, stepIndex: predictionIndex };

  assert.equal(isPredictionReady(state), false);
  assert.deepEqual(reduceState(state, { type: "NEXT" }), state);

  state = reduceState(state, { type: "SELECT_PREDICTION", choiceId: "changes" });
  assert.equal(isPredictionReady(state), true);
  assert.equal(reduceState(state, { type: "NEXT" }).predictionRevealed, false);

  state = reduceState(state, { type: "REVEAL" });
  assert.equal(state.predictionRevealed, true);
  assert.equal(reduceState(state, { type: "NEXT" }).stepIndex, predictionIndex + 1);
});

test("previous and scene switching reset transient prediction state", () => {
  let state = reduceState(createInitialState(), {
    type: "SELECT_SCENE",
    sceneId: "defer-lifo",
  });
  state = { ...state, stepIndex: 3, predictionChoice: "b-first", predictionRevealed: true };
  state = reduceState(state, { type: "PREVIOUS" });
  assert.equal(state.stepIndex, 2);
  assert.equal(state.predictionChoice, null);
  assert.equal(state.predictionRevealed, false);

  state = reduceState(state, {
    type: "SELECT_SCENE",
    sceneId: "slice-append",
  });
  assert.equal(state.stepIndex, 0);
  assert.equal(state.predictionChoice, null);
});
```

- [ ] **Step 2: Run the tests and verify RED**

Run:

```bash
node --test module01_basics/visualizer/tests/state.test.js
```

Expected: FAIL because `scenes.js` and `app.js` do not exist.

- [ ] **Step 3: Implement the scene contract and reducer**

Create four scene objects in `scenes.js`. Use a browser/Node wrapper:

```js
(function exposeScenes(root) {
  const scenes = [
    {
      id: "slice-shared",
      chapter: "Slice",
      title: "两个 Slice，共享一个数组",
      code: [
        "base := []int{1, 2, 3, 4}",
        "view := base[1:3]",
        "view[0] = 9",
        "fmt.Println(base)",
      ],
      steps: [
        { kind: "observe", line: 0, title: "创建 base", narration: "一个 Slice Header 指向底层数组。", stage: { phase: "base" } },
        { kind: "observe", line: 1, title: "创建 view", narration: "新的 Header 从数组下标 1 开始观察。", stage: { phase: "view" } },
        {
          kind: "prediction",
          line: 2,
          title: "先预测",
          narration: "修改 view[0] 后，base 会变化吗？",
          prediction: {
            choices: [
              { id: "changes", label: "base 会变化" },
              { id: "unchanged", label: "base 不会变化" },
            ],
            correctChoiceId: "changes",
            explanation: "base[1] 与 view[0] 对应同一个底层数组单元。",
          },
          stage: { phase: "view" },
        },
        { kind: "reveal", line: 2, title: "修改共享单元", narration: "值 2 变为 9，两条指针保持不动。", stage: { phase: "mutated" } },
        { kind: "observe", line: 3, title: "对应关系", narration: "base[1] 与 view[0] 是同一个位置。", stage: { phase: "aliases" } },
        { kind: "conclusion", line: 3, title: "结论", narration: "Header 是值，底层数组元素仍可共享。", stage: { phase: "aliases" } },
      ],
    },
    {
      id: "slice-append",
      chapter: "Slice",
      title: "append：复用还是搬家",
      code: [
        "spare := make([]int, 2, 4)",
        "reused := append(spare, 3)",
        "full := make([]int, 2, 2)",
        "grown := append(full, 3)",
        "full[0], grown[0] = 7, 9",
      ],
      steps: [
        { kind: "observe", line: 0, title: "并排观察", narration: "左侧容量有余，右侧容量已满。", stage: { phase: "initial" } },
        { kind: "observe", line: 1, title: "复用原数组", narration: "新元素落入原数组空槽，返回 Slice 的 len 增加。", stage: { phase: "reused" } },
        { kind: "prediction", line: 3, title: "先预测", narration: "容量已满时还能写入原数组吗？", prediction: { choices: [{ id: "move", label: "需要新的数组" }, { id: "reuse", label: "继续复用原数组" }], correctChoiceId: "move", explanation: "新长度超过容量，append 必须返回指向新存储的 Slice。" }, stage: { phase: "before-growth" } },
        { kind: "reveal", line: 3, title: "分配并复制", narration: "元素被复制到容量不小于 3 的新数组。", stage: { phase: "reallocated" } },
        { kind: "observe", line: 3, title: "指向不同数组", narration: "full 留在旧数组，grown 指向新数组。", stage: { phase: "separated" } },
        { kind: "observe", line: 4, title: "验证分离", narration: "分别修改后，旧数组是 7，新数组是 9。", stage: { phase: "verified" } },
        { kind: "conclusion", line: 4, title: "结论", narration: "是否继续共享，取决于 append 时容量是否足够。", stage: { phase: "verified" } },
      ],
    },
    {
      id: "defer-lifo",
      chapter: "Defer",
      title: "后注册，先执行",
      code: [
        "func work() {",
        "    defer print(\"A\")",
        "    defer print(\"B\")",
        "    print(\"work\")",
        "}",
      ],
      steps: [
        { kind: "observe", line: 0, title: "进入函数", narration: "待执行调用区现在为空。", stage: { phase: "empty" } },
        { kind: "observe", line: 1, title: "注册 A", narration: "A 进入待执行调用区。", stage: { phase: "register-a" } },
        { kind: "observe", line: 2, title: "注册 B", narration: "B 后注册，位于 A 上方。", stage: { phase: "register-b" } },
        { kind: "prediction", line: 4, title: "先预测", narration: "函数返回前会先执行 A 还是 B？", prediction: { choices: [{ id: "b-first", label: "先执行 B" }, { id: "a-first", label: "先执行 A" }], correctChoiceId: "b-first", explanation: "已注册的 Defer 调用按逆序执行。" }, stage: { phase: "register-b" } },
        { kind: "reveal", line: 3, title: "执行函数体", narration: "先输出 work，然后到达准备返回关口。", stage: { phase: "work" } },
        { kind: "observe", line: 4, title: "B 先执行", narration: "B 从待执行区弹出。", stage: { phase: "execute-b" } },
        { kind: "observe", line: 4, title: "A 再执行", narration: "A 随后弹出，函数才能返回。", stage: { phase: "execute-a" } },
        { kind: "conclusion", line: 4, title: "结论", narration: "多个 Defer 调用后注册的先执行。", stage: { phase: "complete" } },
      ],
    },
    {
      id: "defer-evaluation",
      chapter: "Defer",
      title: "参数快照与闭包读取",
      code: [
        "value := 1",
        "defer print(value)",
        "defer func() { print(value) }()",
        "value = 2",
        "// return",
      ],
      steps: [
        { kind: "observe", line: 0, title: "创建变量", narration: "变量格中的 value 是 1。", stage: { phase: "value-one" } },
        { kind: "observe", line: 1, title: "保存普通参数", narration: "调用参数现在求值，卡片保存 1。", stage: { phase: "save-argument" } },
        { kind: "observe", line: 2, title: "注册闭包", narration: "闭包体会在执行时读取 value。", stage: { phase: "register-closure" } },
        { kind: "observe", line: 3, title: "修改变量", narration: "变量格中的 value 变为 2。", stage: { phase: "value-two" } },
        { kind: "prediction", line: 4, title: "先预测", narration: "两个调用分别会打印什么？", prediction: { choices: [{ id: "closure-2-arg-1", label: "闭包 2，普通参数 1" }, { id: "both-2", label: "两者都是 2" }], correctChoiceId: "closure-2-arg-1", explanation: "普通参数已保存 1；闭包执行时读取到 2。" }, stage: { phase: "value-two" } },
        { kind: "reveal", line: 4, title: "闭包先执行", narration: "后注册的闭包先读取 value，输出 2。", stage: { phase: "execute-closure" } },
        { kind: "observe", line: 4, title: "普通调用再执行", narration: "普通调用使用保存的参数，输出 1。", stage: { phase: "complete" } },
        { kind: "conclusion", line: 4, title: "结论", narration: "注册时求参数，执行时运行闭包体。", stage: { phase: "complete" } },
      ],
    },
  ];

  root.VISUALIZER_SCENES = scenes;
  if (typeof module !== "undefined" && module.exports) module.exports = scenes;
})(typeof globalThis !== "undefined" ? globalThis : window);
```

Implement pure state helpers in `app.js`; prediction `NEXT` is a no-op before reveal, and terminal `NEXT` returns to menu:

```js
function createInitialState() {
  return {
    sceneId: null,
    stepIndex: 0,
    predictionChoice: null,
    predictionRevealed: false,
    isAnimating: false,
    replayNonce: 0,
    error: null,
  };
}

function reduceState(state, action) {
  if (state.isAnimating && !["ANIMATION_END", "SELECT_SCENE"].includes(action.type)) {
    return state;
  }
  if (action.type === "SELECT_SCENE") {
    return { ...createInitialState(), sceneId: action.sceneId };
  }
  if (action.type === "SELECT_PREDICTION") {
    return { ...state, predictionChoice: action.choiceId, predictionRevealed: false };
  }
  if (action.type === "REVEAL") {
    return state.predictionChoice
      ? { ...state, predictionRevealed: true }
      : state;
  }
  if (action.type === "PREVIOUS") {
    if (!state.sceneId || state.stepIndex === 0) return state;
    return {
      ...state,
      stepIndex: state.stepIndex - 1,
      predictionChoice: null,
      predictionRevealed: false,
    };
  }
  if (action.type === "NEXT") {
    const scene = getCurrentScene(state);
    const step = getCurrentStep(state);
    if (!scene || !step) return createInitialState();
    if (step.kind === "prediction" && !state.predictionRevealed) return state;
    if (state.stepIndex === scene.steps.length - 1) return createInitialState();
    return {
      ...state,
      stepIndex: state.stepIndex + 1,
      predictionChoice: null,
      predictionRevealed: false,
    };
  }
  if (action.type === "REPLAY") {
    return { ...state, replayNonce: state.replayNonce + 1 };
  }
  if (action.type === "RESET_TO_MENU") return createInitialState();
  if (action.type === "ANIMATION_START") return { ...state, isAnimating: true };
  if (action.type === "ANIMATION_END") return { ...state, isAnimating: false };
  if (action.type === "SET_ERROR") {
    return { ...state, isAnimating: false, error: String(action.error) };
  }
  return state;
}
```

- [ ] **Step 4: Run the tests and verify GREEN**

Run:

```bash
node --test module01_basics/visualizer/tests/state.test.js
```

Expected: 3 tests pass.

- [ ] **Step 5: Commit the model**

```bash
git add module01_basics/visualizer/scenes.js \
  module01_basics/visualizer/app.js \
  module01_basics/visualizer/tests/state.test.js
git commit -m "feat: add visualizer scenes and state model"
```

---

### Task 2: 页面骨架、导航和通用控制

**Files:**
- Create: `module01_basics/visualizer/index.html`
- Create: `module01_basics/visualizer/styles.css`
- Modify: `module01_basics/visualizer/app.js`
- Modify: `module01_basics/visualizer/tests/state.test.js`

**Interfaces:**
- Consumes: `VISUALIZER_SCENES`, `createInitialState`, `reduceState`.
- Produces: `renderApp(root: HTMLElement, state: AppState): void`
- Produces: `actionFromKeyboard(event): Action | null`
- Produces DOM IDs: `app`, `scene-nav`, `code-panel`, `stage`, `teaching-note`, `previous-button`, `replay-button`, `next-button`.

- [ ] **Step 1: Add failing tests for keyboard mapping and terminal flow**

Append:

```js
const { actionFromKey } = require("../app.js");

test("maps classroom keyboard controls", () => {
  assert.deepEqual(actionFromKey("ArrowRight"), { type: "NEXT" });
  assert.deepEqual(actionFromKey(" "), { type: "NEXT" });
  assert.deepEqual(actionFromKey("ArrowLeft"), { type: "PREVIOUS" });
  assert.deepEqual(actionFromKey("r"), { type: "REPLAY" });
  assert.equal(actionFromKey("Enter"), null);
});

test("next from a conclusion returns to the scene menu", () => {
  let state = reduceState(createInitialState(), {
    type: "SELECT_SCENE",
    sceneId: "defer-lifo",
  });
  state = { ...state, stepIndex: getCurrentScene(state).steps.length - 1 };
  state = reduceState(state, { type: "NEXT" });
  assert.equal(state.sceneId, null);
});
```

- [ ] **Step 2: Run the focused tests and verify RED**

Run:

```bash
node --test module01_basics/visualizer/tests/state.test.js
```

Expected: FAIL because `actionFromKey` and terminal flow are incomplete.

- [ ] **Step 3: Implement the semantic shell and controller**

Create `index.html` with no module scripts so it works over `file://`:

```html
<!doctype html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="color-scheme" content="dark">
  <title>Go Slice 与 Defer 原理演示</title>
  <link rel="stylesheet" href="./styles.css">
</head>
<body>
  <main id="app" class="app-shell" aria-live="polite"></main>
  <noscript>此演示需要启用 JavaScript。</noscript>
  <script src="./scenes.js"></script>
  <script src="./app.js"></script>
</body>
</html>
```

In `app.js`, boot only when `document` exists. Render a four-card menu when
`state.sceneId === null`; otherwise render chapter navigation, highlighted code,
stage, teaching note, progress and controls. Attach one delegated click listener
to `#app` and one keyboard listener to `document`.

`actionFromKey` must map only `ArrowRight`, space, `ArrowLeft`, `r`, and `R`.
Ignore keyboard shortcuts when the event target is an interactive control so a
focused button cannot trigger twice.

- [ ] **Step 4: Implement the base projection layout**

Create `styles.css` using CSS custom properties and these layout invariants:

```css
:root {
  color-scheme: dark;
  --background: #08111f;
  --surface: #101d2f;
  --surface-raised: #17263a;
  --text: #f4f7fb;
  --muted: #aab7c8;
  --border: #33445c;
  --active: #ffd166;
  --shared: #54d2a0;
  --allocated: #66a3ff;
  --deferred: #c792ea;
  --danger: #ff7b72;
}

.lesson-grid {
  display: grid;
  grid-template-columns: minmax(18rem, 34%) minmax(0, 66%);
  gap: clamp(1rem, 2vw, 2rem);
}

@media (max-width: 880px) {
  .lesson-grid { grid-template-columns: 1fr; }
}
```

Use a minimum 16px base font, visible `:focus-visible` outlines, wrapping
controls, no fixed viewport height, and no horizontal page overflow.

- [ ] **Step 5: Run tests and open the shell**

Run:

```bash
node --test module01_basics/visualizer/tests/state.test.js
python3 -m http.server 4173 --directory module01_basics/visualizer
```

Expected: all state tests pass; `http://localhost:4173` shows four scene cards
and opens a lesson without console errors.

- [ ] **Step 6: Commit the shell**

```bash
git add module01_basics/visualizer/index.html \
  module01_basics/visualizer/styles.css \
  module01_basics/visualizer/app.js \
  module01_basics/visualizer/tests/state.test.js
git commit -m "feat: add classroom visualizer shell"
```

---

### Task 3: Slice 舞台渲染

**Files:**
- Modify: `module01_basics/visualizer/app.js`
- Modify: `module01_basics/visualizer/styles.css`
- Create: `module01_basics/visualizer/tests/render.test.js`

**Interfaces:**
- Produces: `renderStageMarkup(scene: Scene, step: Step): string`
- Produces: `escapeHTML(value: unknown): string`
- Slice stage markers: `data-array-id`, `data-index`, `data-slice-id`, `data-pointer-target`, `data-state`.

- [ ] **Step 1: Write failing Slice render tests**

Create:

```js
const test = require("node:test");
const assert = require("node:assert/strict");
const scenes = require("../scenes.js");
const { renderStageMarkup, escapeHTML } = require("../app.js");

function scene(id) {
  return scenes.find((item) => item.id === id);
}

test("shared Slice reveal renders one mutated backing cell and two headers", () => {
  const target = scene("slice-shared");
  const step = target.steps.find((item) => item.stage.phase === "mutated");
  const markup = renderStageMarkup(target, step);
  assert.match(markup, /data-slice-id="base"/);
  assert.match(markup, /data-slice-id="view"/);
  assert.match(markup, /data-index="1"[^>]*data-state="mutated"/);
  assert.match(markup, />9</);
});

test("append reallocation keeps old and new arrays distinct", () => {
  const target = scene("slice-append");
  const step = target.steps.find((item) => item.stage.phase === "reallocated");
  const markup = renderStageMarkup(target, step);
  assert.match(markup, /data-array-id="full-array"/);
  assert.match(markup, /data-array-id="grown-array"/);
  assert.match(markup, /容量 ≥ 3/);
  assert.doesNotMatch(markup, /容量 = 4/);
});

test("escapes dynamic teaching copy", () => {
  assert.equal(escapeHTML('<img src=x onerror="x">'), "&lt;img src=x onerror=&quot;x&quot;&gt;");
});
```

- [ ] **Step 2: Run render tests and verify RED**

Run:

```bash
node --test module01_basics/visualizer/tests/render.test.js
```

Expected: FAIL because `renderStageMarkup` and `escapeHTML` are not exported.

- [ ] **Step 3: Implement Slice markup renderers**

Add `renderSliceShared(stage)` and `renderSliceAppend(stage)` and dispatch from
`renderStageMarkup`. Render:

- A labeled Header box per Slice with pointer label, len and cap.
- One cell per visible backing-array slot, including capacity-only slots.
- SVG or CSS connector lines with textual target labels.
- `data-state="mutated"` on the changed shared cell.
- Separate old/new backing arrays after reallocation.
- A visible `容量 ≥ 3` label for the grown array.
- A concise `role="img"` and `aria-label` summary on each stage.

All scene-derived text must pass through `escapeHTML`.

- [ ] **Step 4: Add movement and reduced-motion CSS**

Use state classes for:

- `.array-cell.is-mutated`
- `.array-block.is-new-allocation`
- `.slice-pointer`
- `.copy-token`

Animate transforms and background changes for 400–700ms. Add:

```css
@media (prefers-reduced-motion: reduce) {
  *, *::before, *::after {
    scroll-behavior: auto !important;
    transition-duration: 0.01ms !important;
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
  }
}
```

- [ ] **Step 5: Run all tests**

Run:

```bash
node --test module01_basics/visualizer/tests/*.test.js
```

Expected: all tests pass.

- [ ] **Step 6: Commit Slice visualization**

```bash
git add module01_basics/visualizer/app.js \
  module01_basics/visualizer/styles.css \
  module01_basics/visualizer/tests/render.test.js
git commit -m "feat: visualize slice sharing and growth"
```

---

### Task 4: Defer 舞台渲染与预测反馈

**Files:**
- Modify: `module01_basics/visualizer/app.js`
- Modify: `module01_basics/visualizer/styles.css`
- Modify: `module01_basics/visualizer/tests/render.test.js`

**Interfaces:**
- Consumes: `renderStageMarkup`, `escapeHTML`.
- Defer stage markers: `data-defer-id`, `data-status`, `data-value-cell`, `data-output`.

- [ ] **Step 1: Write failing Defer render tests**

Append:

```js
test("LIFO execution marks B before A", () => {
  const target = scene("defer-lifo");
  const step = target.steps.find((item) => item.stage.phase === "execute-b");
  const markup = renderStageMarkup(target, step);
  assert.match(markup, /data-defer-id="B"[^>]*data-status="executing"/);
  assert.match(markup, /data-defer-id="A"[^>]*data-status="pending"/);
  assert.match(markup, /data-output="work,B"/);
});

test("evaluation result distinguishes closure read from saved argument", () => {
  const target = scene("defer-evaluation");
  const step = target.steps.find((item) => item.stage.phase === "complete");
  const markup = renderStageMarkup(target, step);
  assert.match(markup, /普通参数[^<]*1/);
  assert.match(markup, /闭包[^<]*2/);
  assert.match(markup, /语义模型/);
});
```

- [ ] **Step 2: Run render tests and verify RED**

Run:

```bash
node --test module01_basics/visualizer/tests/render.test.js
```

Expected: the two new tests fail because Defer stage renderers are absent.

- [ ] **Step 3: Implement Defer renderers**

Add `renderDeferLIFO(stage)` and `renderDeferEvaluation(stage)`.

The LIFO stage includes:

- A labeled function frame.
- A “待执行调用（语义模型）” column.
- A and B cards with `pending`, `executing`, or `done` status.
- Output timeline `work → B → A` revealed incrementally.

The evaluation stage includes:

- A visible `value` variable cell changing from 1 to 2.
- A normal-call card labeled `已保存参数：1`.
- A closure card labeled `执行时读取 value`.
- Final outputs `闭包：2` and `普通参数：1`.
- The note “卡片栈是语义模型；编译器优化不能改变结果”.

Prediction feedback must display the selected answer, correct/incorrect label,
and explanation after `predictionRevealed === true`. Correctness must never be
expressed with color alone.

- [ ] **Step 4: Run all JavaScript tests**

Run:

```bash
node --test module01_basics/visualizer/tests/*.test.js
```

Expected: all tests pass.

- [ ] **Step 5: Commit Defer visualization**

```bash
git add module01_basics/visualizer/app.js \
  module01_basics/visualizer/styles.css \
  module01_basics/visualizer/tests/render.test.js
git commit -m "feat: visualize defer registration and evaluation"
```

---

### Task 5: Documentation, browser verification and visual QA

**Files:**
- Modify: `module01_basics/README.md`
- Modify: `module01_basics/visualizer/index.html`
- Modify: `module01_basics/visualizer/styles.css`
- Modify: `module01_basics/visualizer/app.js`
- Modify: `module01_basics/visualizer/scenes.js`
- Modify: `module01_basics/visualizer/tests/state.test.js`
- Modify: `module01_basics/visualizer/tests/render.test.js`

**Interfaces:**
- Consumes all finished visualizer files.
- Produces a documented classroom entry point and verified behavior at all target viewports.

- [ ] **Step 1: Document offline and classroom usage**

Add a “Slice 与 Defer 原理动画” section to `module01_basics/README.md` with:

```markdown
## Slice 与 Defer 原理动画

直接在浏览器中打开 `module01_basics/visualizer/index.html`。页面无需后端或网络，
包含 Slice 共享、append 扩容、Defer LIFO 和 Defer 求值时机四个逐步场景。

- `→` 或空格：下一步
- `←`：上一步
- `R`：重播当前步

课堂建议先停在每个场景的“先预测”步骤，收集学生判断后再点击“揭晓”。
```

- [ ] **Step 2: Run content correctness commands**

Run:

```bash
go run ./module01_basics/blocks/02_collections/demo/04_arrays_slices
go run ./module01_basics/blocks/02_collections/demo/06_slice_map_edges
go test ./module01_basics/blocks/04_functions_testing/demo/11_defer_edges
node --test module01_basics/visualizer/tests/*.test.js
```

Expected:

- Shared Slice output contains `[1 999 3 4 5]`.
- Capacity-full append output contains `original=[1 2] grown=[9 2 3]`.
- Defer test passes with arguments `[1 0]` and closures `[1 1]`.
- All visualizer Node tests pass.

- [ ] **Step 3: Start a local static server for browser QA**

Run:

```bash
python3 -m http.server 4173 --directory module01_basics/visualizer
```

Use Playwright to open `http://127.0.0.1:4173/`. Assert:

- Four `[data-scene-id]` buttons exist.
- Every scene reaches its conclusion using the Next/Reveal flow.
- Prediction cannot reveal before a choice.
- `ArrowRight`, space, `ArrowLeft`, and `R` produce the documented actions.
- Rapid input while an animation is locked does not skip steps.
- Switching scenes clears the old prediction result.
- Browser console contains no errors.

- [ ] **Step 4: Render and inspect target viewports**

Capture full-page screenshots at:

- 1280×720.
- 1440×900.
- 390×844.

For each viewport inspect the menu, one Slice scene, one Defer scene, a prediction
screen, and a conclusion screen. Fix any overlapping code, pointer lines, cards,
controls, clipped text or horizontal page overflow. Repeat screenshots after each
fix until all five screen types are clean.

- [ ] **Step 5: Verify local-file mode and reduced motion**

Open the absolute `file:///.../module01_basics/visualizer/index.html` URL and
complete one Slice and one Defer scene. Emulate `prefers-reduced-motion: reduce`
and verify the same final state is reached without long transitions.

- [ ] **Step 6: Run final regression suite**

Run:

```bash
go test ./module01_basics/...
node --test module01_basics/visualizer/tests/*.test.js
git diff --check
```

Expected: all Go and JavaScript tests pass and `git diff --check` produces no output.

- [ ] **Step 7: Commit documentation and QA fixes**

```bash
git add module01_basics/README.md module01_basics/visualizer
git commit -m "docs: add classroom visualizer instructions"
```

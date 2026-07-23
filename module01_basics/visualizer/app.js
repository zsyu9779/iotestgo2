"use strict";

(function exposeVisualizer(root) {
	const scenes =
		root.VISUALIZER_SCENES ||
		(typeof require === "function" ? require("./scenes.js") : []);

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

	function getCurrentScene(state) {
		if (!state || !state.sceneId) return null;
		return scenes.find((scene) => scene.id === state.sceneId) || null;
	}

	function getCurrentStep(state) {
		const scene = getCurrentScene(state);
		if (!scene) return null;
		return scene.steps[state.stepIndex] || null;
	}

	function isPredictionReady(state) {
		const step = getCurrentStep(state);
		return Boolean(
			step &&
				step.kind === "prediction" &&
				state.predictionChoice &&
				!state.predictionRevealed,
		);
	}

	function actionFromKey(key) {
		switch (key) {
			case "ArrowRight":
			case " ":
				return { type: "NEXT" };
			case "ArrowLeft":
				return { type: "PREVIOUS" };
			case "r":
			case "R":
				return { type: "REPLAY" };
			default:
				return null;
		}
	}

	function escapeHTML(value) {
		return String(value)
			.replaceAll("&", "&amp;")
			.replaceAll("<", "&lt;")
			.replaceAll(">", "&gt;")
			.replaceAll('"', "&quot;")
			.replaceAll("'", "&#039;");
	}

	function renderTopbar() {
		return `
			<header class="topbar">
				<p class="brand"><span class="brand-mark">Go</span> 底层原理演示</p>
				<p class="keyboard-hint">← 上一步　→ / 空格 下一步　R 重播</p>
			</header>
		`;
	}

	function renderMenuMarkup(state) {
		const cards = scenes
			.map(
				(scene) => `
					<button class="scene-card" type="button" data-scene-id="${escapeHTML(scene.id)}">
						<span>
							<span class="eyebrow">${escapeHTML(scene.chapter)} · ${escapeHTML(scene.eyebrow)}</span>
							<h2>${escapeHTML(scene.title)}</h2>
							<p>${escapeHTML(scene.description)}</p>
						</span>
						<span class="scene-meta">${scene.steps.length} 步 · 含一次课堂预测 →</span>
					</button>
				`,
			)
			.join("");

		return `
			${renderTopbar()}
			<section class="menu-hero">
				<p class="eyebrow">Module 01 · 讲师逐步播放</p>
				<h1>让运行时语义真正动起来</h1>
				<p class="menu-intro">选择一个场景。代码只负责定位，动画舞台负责解释共享、扩容、注册与求值时机。</p>
			</section>
			${state.error ? `<p class="error-banner" role="alert">${escapeHTML(state.error)}</p>` : ""}
			<section class="scene-grid" aria-label="演示场景">${cards}</section>
		`;
	}

	function renderChapterNavigation(activeScene) {
		return scenes
			.map(
				(scene) => `
					<button
						type="button"
						class="chapter-button"
						data-scene-id="${escapeHTML(scene.id)}"
						${scene.id === activeScene.id ? 'aria-current="page"' : ""}
					>${escapeHTML(scene.chapter)} · ${escapeHTML(scene.eyebrow)}</button>
				`,
			)
			.join("");
	}

	function renderCode(scene, step) {
		return scene.code
			.map(
				(line, index) => `
					<li
						class="code-line${index === step.line ? " is-active" : ""}"
						data-line-number="${index + 1}"
						${index === step.line ? 'aria-current="step"' : ""}
					><span>${escapeHTML(line)}</span></li>
				`,
			)
			.join("");
	}

	function renderArrayCells({
		id,
		values,
		capacity,
		mutatedIndex = -1,
		aliases = {},
		newAllocation = false,
	}) {
		const cells = Array.from({ length: capacity }, (_, index) => {
			const hasValue = index < values.length;
			const value = hasValue ? values[index] : "";
			const state =
				index === mutatedIndex
					? "mutated"
					: hasValue
						? "visible"
						: "capacity";
			const alias = aliases[index]
				? `<span class="cell-alias">${escapeHTML(aliases[index])}</span>`
				: "";
			return `
				<div class="array-slot">
					<div
						class="array-cell${state === "capacity" ? " is-capacity" : ""}${state === "mutated" ? " is-mutated" : ""}"
						data-index="${index}"
						data-state="${state}"
					>
						<span class="array-index">${index}</span>
						<span class="array-value">${escapeHTML(value)}</span>
					</div>
					${alias}
				</div>
			`;
		}).join("");

		return `
			<div
				class="array-block${newAllocation ? " is-new-allocation" : ""}"
				data-array-id="${escapeHTML(id)}"
			>
				<div class="array-cells">${cells}</div>
			</div>
		`;
	}

	function renderSliceHeader({ id, label, pointer, length, capacity, tone }) {
		return `
			<div class="slice-header tone-${escapeHTML(tone)}" data-slice-id="${escapeHTML(id)}">
				<div class="header-name">${escapeHTML(label)}</div>
				<div class="header-field" aria-label="ptr → index ${pointer}"><span>ptr</span><strong>→ index ${pointer}</strong></div>
				<div class="header-field" aria-label="len = ${length}"><span>len</span><strong>= ${length}</strong></div>
				<div class="header-field" aria-label="cap = ${capacity}"><span>cap</span><strong>= ${capacity}</strong></div>
				<svg
					class="pointer-arrow"
					data-pointer-target="${pointer}"
					viewBox="0 0 120 18"
					aria-hidden="true"
				>
					<path d="M2 9 H106"></path>
					<path d="M106 3 L118 9 L106 15 Z"></path>
				</svg>
			</div>
		`;
	}

	function renderSliceShared(stage) {
		const hasView = stage.phase !== "base";
		const isMutated = ["mutated", "aliases"].includes(stage.phase);
		const values = isMutated ? [1, 9, 3, 4] : [1, 2, 3, 4];
		const aliases = isMutated
			? {
					1: "base[1] = view[0]",
				}
			: {};
		const summary = isMutated
			? "base 和 view 仍指向同一底层数组；共享单元的值已经从 2 变为 9。"
			: hasView
				? "base 与 view 拥有独立的 len 和 cap，但它们指向同一个底层数组。"
				: "base 的 Slice Header 指向包含四个元素的底层数组。";

		return `
			<div class="stage-visual slice-shared-stage" role="img" aria-label="${escapeHTML(summary)}">
				<div class="stage-caption">
					<span class="stage-kicker">Slice Header</span>
					<span class="stage-status">${hasView ? "两个 Header · 一个数组" : "一个 Header · 一个数组"}</span>
				</div>
				<div class="shared-layout">
					<div class="header-stack">
						${renderSliceHeader({
							id: "base",
							label: "base",
							pointer: 0,
							length: 4,
							capacity: 4,
							tone: "shared",
						})}
						${
							hasView
								? renderSliceHeader({
										id: "view",
										label: "view",
										pointer: 1,
										length: 2,
										capacity: 3,
										tone: "active",
									})
								: '<div class="header-ghost">下一步将创建 view</div>'
						}
					</div>
					<div class="memory-zone">
						<p class="memory-label">底层数组 · 同一存储区域</p>
						${renderArrayCells({
							id: "shared-array",
							values,
							capacity: 4,
							mutatedIndex: isMutated ? 1 : -1,
							aliases,
						})}
						<p class="memory-note">${escapeHTML(
							isMutated
								? "修改发生在数组元素上；base[1] 与 view[0] 同时观察到 9。"
								: "Header 的 ptr 决定从数组哪里开始观察。",
						)}</p>
					</div>
				</div>
			</div>
		`;
	}

	function renderAppendLane({
		title,
		tag,
		headers,
		arrays,
		note,
		tone,
	}) {
		return `
			<section class="append-lane tone-${escapeHTML(tone)}">
				<header class="lane-heading">
					<span>${escapeHTML(tag)}</span>
					<h3>${escapeHTML(title)}</h3>
				</header>
				<div class="lane-headers">
					${headers
						.map((header) => renderSliceHeader(header))
						.join("")}
				</div>
				<div class="lane-arrays">${arrays.join("")}</div>
				<p class="lane-note">${escapeHTML(note)}</p>
			</section>
		`;
	}

	function renderSliceAppend(stage) {
		const phase = stage.phase;
		const hasReused = phase !== "initial";
		const hasGrown = [
			"reallocated",
			"separated",
			"verified",
		].includes(phase);
		const isVerified = phase === "verified";

		const spareHeaders = [
			{
				id: "spare",
				label: "spare",
				pointer: 0,
				length: 2,
				capacity: 4,
				tone: "shared",
			},
		];
		if (hasReused) {
			spareHeaders.push({
				id: "reused",
				label: "reused",
				pointer: 0,
				length: 3,
				capacity: 4,
				tone: "active",
			});
		}

		const fullHeaders = [
			{
				id: "full",
				label: "full",
				pointer: 0,
				length: 2,
				capacity: 2,
				tone: "shared",
			},
		];
		if (hasGrown) {
			fullHeaders.push({
				id: "grown",
				label: "grown",
				pointer: 0,
				length: 3,
				capacity: "≥ 3",
				tone: "allocated",
			});
		}

		const spareArray = renderArrayCells({
			id: "spare-array",
			values: hasReused ? [0, 0, 3] : [0, 0],
			capacity: 4,
			mutatedIndex: hasReused ? 2 : -1,
		});
		const fullArray = renderArrayCells({
			id: "full-array",
			values: isVerified ? [7, 0] : [0, 0],
			capacity: 2,
			mutatedIndex: isVerified ? 0 : -1,
		});
		const grownArray = hasGrown
			? `
				<div class="copy-bridge" aria-hidden="true">
					<span class="copy-token">0</span>
					<span class="copy-token">0</span>
					<strong>复制到新存储 →</strong>
				</div>
				<p class="capacity-contract">新数组：容量 ≥ 3（具体倍率不是语言契约）</p>
				${renderArrayCells({
					id: "grown-array",
					values: isVerified ? [9, 0, 3] : [0, 0, 3],
					capacity: 3,
					mutatedIndex: isVerified ? 0 : 2,
					newAllocation: true,
				})}
			`
			: "";

		const summary = hasGrown
			? "容量不足时，full 保留旧数组，grown 指向复制元素后的新数组。"
			: "容量有余时 append 可以复用原数组；容量已满的 full 正在等待预测。";

		return `
			<div class="stage-visual slice-append-stage" role="img" aria-label="${escapeHTML(summary)}">
				<div class="stage-caption">
					<span class="stage-kicker">两条 append 轨道</span>
					<span class="stage-status">${hasGrown ? "一条复用 · 一条搬家" : "先观察容量"}</span>
				</div>
				<div class="append-grid">
					${renderAppendLane({
						title: "容量有余",
						tag: "len 2 · cap 4",
						headers: spareHeaders,
						arrays: [spareArray],
						note: hasReused
							? "reused 与 spare 指向同一数组；新元素进入空槽。"
							: "数组中还有两个容量槽位。",
						tone: "shared",
					})}
					${renderAppendLane({
						title: "容量已满",
						tag: "len 2 · cap 2",
						headers: fullHeaders,
						arrays: [fullArray, grownArray],
						note: isVerified
							? "旧数组是 7，新数组是 9：二者不再共享。"
							: hasGrown
								? "full 留在旧数组，grown 指向新数组。"
								: "没有空槽；下一次 append 必须返回新的 Slice。",
						tone: "allocated",
					})}
				</div>
				${isVerified ? '<p class="separation-banner">验证完成：full 与 grown 不再共享底层数组</p>' : ""}
			</div>
		`;
	}

	function renderStageMarkup(scene, step) {
		if (!scene || !step || !step.stage) {
			throw new Error("缺少舞台步骤数据");
		}
		switch (step.stage.type) {
			case "slice-shared":
				return renderSliceShared(step.stage);
			case "slice-append":
				return renderSliceAppend(step.stage);
			default:
				return `
					<div class="stage-visual stage-placeholder" role="img" aria-label="${escapeHTML(step.title)}">
						舞台状态：${escapeHTML(step.stage.phase)}
					</div>
				`;
		}
	}

	function renderPrediction(step, state) {
		if (step.kind !== "prediction" || !step.prediction) return "";

		const choices = step.prediction.choices
			.map(
				(choice) => `
					<button
						type="button"
						class="prediction-choice"
						data-choice-id="${escapeHTML(choice.id)}"
						aria-pressed="${choice.id === state.predictionChoice}"
						${state.predictionRevealed ? "disabled" : ""}
					>${escapeHTML(choice.label)}</button>
				`,
			)
			.join("");

		let feedback = "";
		if (state.predictionRevealed) {
			const isCorrect =
				state.predictionChoice === step.prediction.correctChoiceId;
			feedback = `
				<p class="prediction-feedback" role="status">
					<strong>${isCorrect ? "回答正确" : "再想一步"}</strong>：
					${escapeHTML(step.prediction.explanation)}
				</p>
			`;
		}

		return `
			<section class="prediction-panel" aria-labelledby="prediction-title">
				<h3 id="prediction-title">先让学生预测</h3>
				<div class="prediction-options">${choices}</div>
				<button
					type="button"
					class="reveal-button"
					data-action="reveal"
					${state.predictionChoice && !state.predictionRevealed ? "" : "disabled"}
				>揭晓答案</button>
				${feedback}
			</section>
		`;
	}

	function renderLessonMarkup(state) {
		const scene = getCurrentScene(state);
		const step = getCurrentStep(state);
		if (!scene || !step) {
			return renderMenuMarkup({
				...createInitialState(),
				error: "找不到当前步骤，已返回场景目录。",
			});
		}

		const isPredictionBlocked =
			step.kind === "prediction" && !state.predictionRevealed;
		const isConclusion = step.kind === "conclusion";

		return `
			<section class="lesson-page">
				${renderTopbar()}
				<nav class="chapter-row" id="scene-nav" aria-label="场景导航">
					<div class="chapter-tabs">${renderChapterNavigation(scene)}</div>
					<button class="menu-button" type="button" data-action="menu">返回目录</button>
				</nav>
				<header class="lesson-heading">
					<div>
						<p class="eyebrow">${escapeHTML(scene.chapter)} · ${escapeHTML(scene.eyebrow)}</p>
						<h1>${escapeHTML(scene.title)}</h1>
					</div>
					<p class="progress" aria-label="当前步骤">第 ${state.stepIndex + 1} / ${scene.steps.length} 步</p>
				</header>
				${state.error ? `<p class="error-banner" role="alert">${escapeHTML(state.error)}</p>` : ""}
				<div class="lesson-grid">
					<section class="code-panel" id="code-panel" aria-label="精简 Go 代码">
						<p class="panel-label">精简代码 · 当前行高亮</p>
						<ol class="code-list">${renderCode(scene, step)}</ol>
					</section>
					<section class="stage" id="stage" aria-label="${escapeHTML(step.title)}">
						<p class="panel-label">动画舞台</p>
						${renderStageMarkup(scene, step)}
						${renderPrediction(step, state)}
					</section>
				</div>
				<section class="teaching-strip" id="teaching-note">
					<div>
						<p class="note-label">讲师提示</p>
						<h2>${escapeHTML(step.title)}</h2>
						<p>${escapeHTML(step.narration)}</p>
					</div>
					<span class="progress">${state.stepIndex + 1} / ${scene.steps.length}</span>
				</section>
				<footer class="lesson-controls" aria-label="播放控制">
					<button class="menu-button" type="button" data-action="menu">退出本场景</button>
					<div class="control-group">
						<button
							class="control-button"
							id="previous-button"
							type="button"
							data-action="previous"
							${state.stepIndex === 0 || state.isAnimating ? "disabled" : ""}
						>← 上一步</button>
						<button
							class="control-button"
							id="replay-button"
							type="button"
							data-action="replay"
							${state.isAnimating ? "disabled" : ""}
						>重播本步</button>
						<button
							class="control-button is-primary"
							id="next-button"
							type="button"
							data-action="next"
							${isPredictionBlocked || state.isAnimating ? "disabled" : ""}
						>${isConclusion ? "返回场景目录" : "下一步 →"}</button>
					</div>
				</footer>
			</section>
		`;
	}

	function renderAppMarkup(state) {
		return state.sceneId ? renderLessonMarkup(state) : renderMenuMarkup(state);
	}

	function reduceState(state, action) {
		if (!state || !action || typeof action.type !== "string") return state;

		if (
			state.isAnimating &&
			!["ANIMATION_END", "SELECT_SCENE", "SET_ERROR"].includes(action.type)
		) {
			return state;
		}

		switch (action.type) {
			case "SELECT_SCENE": {
				if (!scenes.some((scene) => scene.id === action.sceneId)) {
					return {
						...createInitialState(),
						error: "找不到这个演示场景，请返回目录重试。",
					};
				}
				return { ...createInitialState(), sceneId: action.sceneId };
			}
			case "SELECT_PREDICTION": {
				const step = getCurrentStep(state);
				if (
					!step ||
					step.kind !== "prediction" ||
					!step.prediction.choices.some(
						(choice) => choice.id === action.choiceId,
					)
				) {
					return state;
				}
				return {
					...state,
					predictionChoice: action.choiceId,
					predictionRevealed: false,
				};
			}
			case "REVEAL":
				return isPredictionReady(state)
					? { ...state, predictionRevealed: true }
					: state;
			case "PREVIOUS":
				if (!state.sceneId || state.stepIndex === 0) return state;
				return {
					...state,
					stepIndex: state.stepIndex - 1,
					predictionChoice: null,
					predictionRevealed: false,
					error: null,
				};
			case "NEXT": {
				const scene = getCurrentScene(state);
				const step = getCurrentStep(state);
				if (!scene || !step) return createInitialState();
				if (step.kind === "prediction" && !state.predictionRevealed) {
					return state;
				}
				if (state.stepIndex >= scene.steps.length - 1) {
					return createInitialState();
				}
				return {
					...state,
					stepIndex: state.stepIndex + 1,
					predictionChoice: null,
					predictionRevealed: false,
					error: null,
				};
			}
			case "REPLAY":
				return state.sceneId
					? { ...state, replayNonce: state.replayNonce + 1 }
					: state;
			case "RESET_TO_MENU":
				return createInitialState();
			case "ANIMATION_START":
				return { ...state, isAnimating: true };
			case "ANIMATION_END":
				return { ...state, isAnimating: false };
			case "SET_ERROR":
				return {
					...state,
					isAnimating: false,
					error: String(action.error || "演示渲染失败，请返回目录重试。"),
				};
			default:
				return state;
		}
	}

	const api = {
		createInitialState,
		reduceState,
		getCurrentScene,
		getCurrentStep,
		isPredictionReady,
		actionFromKey,
		escapeHTML,
		renderAppMarkup,
		renderStageMarkup,
	};

	root.VisualizerApp = api;
	if (typeof module !== "undefined" && module.exports) {
		module.exports = api;
	}

	function bootBrowser() {
		if (typeof document === "undefined") return;
		const appRoot = document.getElementById("app");
		if (!appRoot) return;

		let state = createInitialState();

		function render() {
			try {
				appRoot.innerHTML = renderAppMarkup(state);
			} catch (error) {
				state = reduceState(state, { type: "SET_ERROR", error });
				appRoot.innerHTML = renderMenuMarkup(state);
			}
		}

		function dispatch(action) {
			if (!action) return;
			state = reduceState(state, action);
			render();
		}

		appRoot.addEventListener("click", (event) => {
			const sceneButton = event.target.closest("[data-scene-id]");
			if (sceneButton) {
				dispatch({
					type: "SELECT_SCENE",
					sceneId: sceneButton.dataset.sceneId,
				});
				return;
			}

			const choiceButton = event.target.closest("[data-choice-id]");
			if (choiceButton) {
				dispatch({
					type: "SELECT_PREDICTION",
					choiceId: choiceButton.dataset.choiceId,
				});
				return;
			}

			const actionButton = event.target.closest("[data-action]");
			if (!actionButton || actionButton.disabled) return;
			const actions = {
				menu: { type: "RESET_TO_MENU" },
				previous: { type: "PREVIOUS" },
				replay: { type: "REPLAY" },
				next: { type: "NEXT" },
				reveal: { type: "REVEAL" },
			};
			dispatch(actions[actionButton.dataset.action] || null);
		});

		document.addEventListener("keydown", (event) => {
			if (
				event.target instanceof HTMLElement &&
				event.target.closest("button, input, select, textarea, a")
			) {
				return;
			}
			const action = actionFromKey(event.key);
			if (!action) return;
			event.preventDefault();
			dispatch(action);
		});

		render();
	}

	bootBrowser();
})(typeof globalThis !== "undefined" ? globalThis : window);

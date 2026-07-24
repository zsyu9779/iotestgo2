"use strict";

(function exposeRuntimeLabApp(root) {
	const labs =
		root.RUNTIME_LABS ||
		(typeof require === "function" ? require("./scenes.js") : []);
	const utils =
		root.RuntimeLabUtils ||
		(typeof require === "function" ? require("./render-utils.js") : null);
	const browserRenderers = root.RuntimeLabRenderers || {};
	const nodeRenderers =
		typeof require === "function"
			? {
					...require("./slice-lab.js"),
					...require("./defer-lab.js"),
				}
			: {};
	const renderers = { ...browserRenderers, ...nodeRenderers };
	const { escapeHTML } = utils;

	function createInitialState() {
		return {
			labId: null,
			stepIndex: 0,
			isAnimating: false,
			replayNonce: 0,
			error: null,
		};
	}

	function getCurrentLab(state) {
		if (!state || !state.labId) return null;
		return labs.find((lab) => lab.id === state.labId) || null;
	}

	function getCurrentStep(state) {
		const lab = getCurrentLab(state);
		if (!lab) return null;
		return lab.steps[state.stepIndex] || null;
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

	function shouldAnimateTransition(previousState, nextState, action) {
		if (!previousState || !nextState || previousState === nextState || !action) {
			return false;
		}
		return ["NEXT", "PREVIOUS", "REPLAY", "SELECT_LAB"].includes(
			action.type,
		);
	}

	function animationDelay(prefersReducedMotion) {
		return prefersReducedMotion ? 0 : 900;
	}

	function reduceState(state, action) {
		if (!state || !action || !action.type) return state;
		if (
			state.isAnimating &&
			!["ANIMATION_END", "SELECT_LAB"].includes(action.type)
		) {
			return state;
		}

		switch (action.type) {
			case "SELECT_LAB": {
				const exists = labs.some((lab) => lab.id === action.labId);
				return exists
					? { ...createInitialState(), labId: action.labId }
					: state;
			}
			case "PREVIOUS":
				if (!getCurrentLab(state) || state.stepIndex === 0) return state;
				return {
					...state,
					stepIndex: state.stepIndex - 1,
					error: null,
				};
			case "NEXT": {
				const lab = getCurrentLab(state);
				const step = getCurrentStep(state);
				if (!lab || !step) return state;
				if (state.stepIndex >= lab.steps.length - 1) {
					return createInitialState();
				}
				return {
					...state,
					stepIndex: state.stepIndex + 1,
					error: null,
				};
			}
			case "REPLAY":
				return state.labId
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
					error: String(
						action.error || "运行时结构渲染失败，请返回目录重试。",
					),
				};
			default:
				return state;
		}
	}

	function renderTopbar() {
		return `
			<header class="topbar">
				<p class="brand">
					<span class="brand-mark">Go</span>
					<span>Runtime 解剖实验室</span>
				</p>
				<p class="keyboard-hint">← 上一步　→ / 空格 下一步　R 重播</p>
			</header>
		`;
	}

	function renderMenuMarkup(state) {
		const cards = labs
			.map(
				(lab, index) => `
					<button
						class="lab-card"
						type="button"
						data-lab-id="${escapeHTML(lab.id)}"
					>
						<span class="lab-index">0${index + 1}</span>
						<span class="lab-card-copy">
							<span class="eyebrow">${escapeHTML(lab.chapter)} · ${escapeHTML(lab.eyebrow)}</span>
							<strong>${escapeHTML(lab.title)}</strong>
							<small>${escapeHTML(lab.description)}</small>
						</span>
						<span class="lab-enter">进入 8 步解剖 →</span>
					</button>
				`,
			)
			.join("");
		return `
			${renderTopbar()}
			<section class="menu-hero">
				<p class="eyebrow">MODULE 01 · 代码背后的真实结构</p>
				<h1>不要看输出，直接拆运行时</h1>
				<p>
					观察机器字、连续内存、指针运算、runtime 调用链和经典链表节点。
					地址是示意值，结构与运算规则按 Go 1.25.11 核对。
				</p>
			</section>
			${state.error ? `<p class="error-banner" role="alert">${escapeHTML(state.error)}</p>` : ""}
			<section class="lab-grid" aria-label="运行时实验室">${cards}</section>
		`;
	}

	function renderLabNavigation(activeLab) {
		return labs
			.map(
				(lab) => `
					<button
						type="button"
						class="lab-tab"
						data-lab-id="${escapeHTML(lab.id)}"
						${lab.id === activeLab.id ? 'aria-current="page"' : ""}
					>
						${escapeHTML(lab.chapter)} · ${escapeHTML(lab.eyebrow)}
					</button>
				`,
			)
			.join("");
	}

	function renderLabFragments(state) {
		const lab = getCurrentLab(state);
		const step = getCurrentStep(state);
		if (!lab || !step) {
			throw new RangeError("No active runtime laboratory step");
		}
		if (lab.kind === "slice") return renderers.renderSliceLab(step);
		if (lab.kind === "defer") return renderers.renderDeferLab(step);
		throw new TypeError(`Unsupported runtime laboratory kind: ${lab.kind}`);
	}

	function renderControls(state, lab) {
		const isLast = state.stepIndex === lab.steps.length - 1;
		return `
			<nav class="lesson-controls" aria-label="讲师播放控制">
				<button
					id="previous-button"
					class="control-button"
					type="button"
					data-action="previous"
					${state.stepIndex === 0 ? "disabled" : ""}
				>← 上一步</button>
				<button
					id="replay-button"
					class="control-button"
					type="button"
					data-action="replay"
				>重播本步</button>
				<button
					id="next-button"
					class="control-button is-primary"
					type="button"
					data-action="next"
				>${isLast ? "返回实验室目录" : "下一步 →"}</button>
			</nav>
		`;
	}

	function renderLessonMarkup(state) {
		const lab = getCurrentLab(state);
		const step = getCurrentStep(state);
		const fragments = renderLabFragments(state);
		return `
			${renderTopbar()}
			<section class="lesson-page" data-lab-kind="${escapeHTML(lab.kind)}">
				<nav id="lab-nav" class="lab-nav" aria-label="实验室导航">
					<div class="lab-tabs">${renderLabNavigation(lab)}</div>
					<button class="menu-button" type="button" data-action="menu">返回目录</button>
				</nav>
				<header class="lesson-heading">
					<div>
						<p class="eyebrow">${escapeHTML(lab.eyebrow)}</p>
						<h1>${escapeHTML(step.title)}</h1>
					</div>
					<p class="progress">第 ${state.stepIndex + 1} / ${lab.steps.length} 步</p>
				</header>
				<section class="runtime-workbench">
					<aside id="structure-inspector" class="structure-inspector">
						${fragments.inspector}
					</aside>
					<section id="runtime-stage" class="runtime-stage">
						<div class="region-label">运行时舞台</div>
						${fragments.stage}
					</section>
					<aside id="operation-panel" class="operation-panel">
						${fragments.operation}
					</aside>
				</section>
				<footer class="lesson-footer">
					<p class="step-summary">${escapeHTML(step.operation.summary)}</p>
					${renderControls(state, lab)}
				</footer>
			</section>
		`;
	}

	function renderAppMarkup(state) {
		return state.labId ? renderLessonMarkup(state) : renderMenuMarkup(state);
	}

	const api = {
		createInitialState,
		reduceState,
		getCurrentLab,
		getCurrentStep,
		actionFromKey,
		shouldAnimateTransition,
		animationDelay,
		renderLabFragments,
		renderAppMarkup,
	};

	root.RuntimeLabApp = api;
	root.VisualizerApp = api;
	if (typeof module !== "undefined" && module.exports) {
		module.exports = api;
	}

	function bootBrowser() {
		if (typeof document === "undefined") return;
		const appRoot = document.getElementById("app");
		if (!appRoot) return;

		let state = createInitialState();
		let animationTimer = null;

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
			const previousState = state;
			const nextState = reduceState(previousState, action);
			if (nextState === previousState) return;

			if (animationTimer !== null) {
				window.clearTimeout(animationTimer);
				animationTimer = null;
			}

			if (shouldAnimateTransition(previousState, nextState, action)) {
				state = reduceState(nextState, { type: "ANIMATION_START" });
				appRoot.dataset.animating = "true";
				appRoot.setAttribute("aria-busy", "true");
				render();
				const reducedMotion = window.matchMedia(
					"(prefers-reduced-motion: reduce)",
				).matches;
				animationTimer = window.setTimeout(() => {
					state = reduceState(state, { type: "ANIMATION_END" });
					animationTimer = null;
					delete appRoot.dataset.animating;
					appRoot.removeAttribute("aria-busy");
				}, animationDelay(reducedMotion));
				return;
			}

			state = nextState;
			render();
		}

		appRoot.addEventListener("click", (event) => {
			const labButton = event.target.closest("[data-lab-id]");
			if (labButton) {
				dispatch({
					type: "SELECT_LAB",
					labId: labButton.dataset.labId,
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

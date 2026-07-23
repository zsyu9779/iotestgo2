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
						<div class="stage-placeholder">舞台状态：${escapeHTML(step.stage.phase)}</div>
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

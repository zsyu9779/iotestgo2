"use strict";

(function exposeRuntimeLabApp(root) {
	const labs =
		root.RUNTIME_LABS ||
		(typeof require === "function" ? require("./scenes.js") : []);

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

	const api = {
		createInitialState,
		reduceState,
		getCurrentLab,
		getCurrentStep,
		actionFromKey,
		shouldAnimateTransition,
		animationDelay,
	};

	root.RuntimeLabApp = api;
	root.VisualizerApp = api;
	if (typeof module !== "undefined" && module.exports) {
		module.exports = api;
	}
})(typeof globalThis !== "undefined" ? globalThis : window);

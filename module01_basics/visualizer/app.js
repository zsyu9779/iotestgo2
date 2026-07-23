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
	};

	root.VisualizerApp = api;
	if (typeof module !== "undefined" && module.exports) {
		module.exports = api;
	}
})(typeof globalThis !== "undefined" ? globalThis : window);

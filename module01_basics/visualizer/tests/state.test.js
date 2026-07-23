"use strict";

const test = require("node:test");
const assert = require("node:assert/strict");

const scenes = require("../scenes.js");
const {
	createInitialState,
	reduceState,
	getCurrentScene,
	getCurrentStep,
	isPredictionReady,
	actionFromKey,
} = require("../app.js");

test("defines the four approved classroom scenes", () => {
	assert.deepEqual(
		scenes.map((scene) => scene.id),
		["slice-shared", "slice-append", "defer-lifo", "defer-evaluation"],
	);

	for (const scene of scenes) {
		assert.ok(scene.code.length >= 4);
		assert.ok(scene.steps.length >= 6);
		assert.equal(
			scene.steps.filter((step) => step.kind === "prediction").length,
			1,
		);
		assert.equal(scene.steps.at(-1).kind, "conclusion");
		for (const step of scene.steps) {
			assert.equal(typeof step.title, "string");
			assert.equal(typeof step.narration, "string");
			assert.ok(step.line >= 0 && step.line < scene.code.length);
		}
	}
});

test("prediction cannot advance until it is selected and revealed", () => {
	let state = reduceState(createInitialState(), {
		type: "SELECT_SCENE",
		sceneId: "slice-shared",
	});
	const scene = getCurrentScene(state);
	const predictionIndex = scene.steps.findIndex(
		(step) => step.kind === "prediction",
	);
	state = { ...state, stepIndex: predictionIndex };

	assert.equal(isPredictionReady(state), false);
	assert.deepEqual(reduceState(state, { type: "NEXT" }), state);

	state = reduceState(state, {
		type: "SELECT_PREDICTION",
		choiceId: "changes",
	});
	assert.equal(isPredictionReady(state), true);
	assert.equal(reduceState(state, { type: "NEXT" }).stepIndex, predictionIndex);

	state = reduceState(state, { type: "REVEAL" });
	assert.equal(state.predictionRevealed, true);
	assert.equal(
		reduceState(state, { type: "NEXT" }).stepIndex,
		predictionIndex + 1,
	);
});

test("previous and scene switching reset transient prediction state", () => {
	let state = reduceState(createInitialState(), {
		type: "SELECT_SCENE",
		sceneId: "defer-lifo",
	});
	state = {
		...state,
		stepIndex: 3,
		predictionChoice: "b-first",
		predictionRevealed: true,
	};

	state = reduceState(state, { type: "PREVIOUS" });
	assert.equal(state.stepIndex, 2);
	assert.equal(state.predictionChoice, null);
	assert.equal(state.predictionRevealed, false);

	state = reduceState(state, {
		type: "SELECT_SCENE",
		sceneId: "slice-append",
	});
	assert.equal(state.sceneId, "slice-append");
	assert.equal(state.stepIndex, 0);
	assert.equal(state.predictionChoice, null);
	assert.equal(state.predictionRevealed, false);
});

test("selectors return null while the scene menu is active", () => {
	const state = createInitialState();
	assert.equal(getCurrentScene(state), null);
	assert.equal(getCurrentStep(state), null);
});

test("maps classroom keyboard controls", () => {
	assert.deepEqual(actionFromKey("ArrowRight"), { type: "NEXT" });
	assert.deepEqual(actionFromKey(" "), { type: "NEXT" });
	assert.deepEqual(actionFromKey("ArrowLeft"), { type: "PREVIOUS" });
	assert.deepEqual(actionFromKey("r"), { type: "REPLAY" });
	assert.deepEqual(actionFromKey("R"), { type: "REPLAY" });
	assert.equal(actionFromKey("Enter"), null);
});

test("next from a conclusion returns to the scene menu", () => {
	let state = reduceState(createInitialState(), {
		type: "SELECT_SCENE",
		sceneId: "defer-lifo",
	});
	state = {
		...state,
		stepIndex: getCurrentScene(state).steps.length - 1,
	};

	state = reduceState(state, { type: "NEXT" });
	assert.equal(state.sceneId, null);
	assert.equal(state.stepIndex, 0);
});

test("animation lock ignores navigation but still permits scene switching", () => {
	let state = reduceState(createInitialState(), {
		type: "SELECT_SCENE",
		sceneId: "slice-shared",
	});
	state = reduceState(state, { type: "ANIMATION_START" });
	assert.equal(state.isAnimating, true);
	assert.deepEqual(reduceState(state, { type: "NEXT" }), state);

	state = reduceState(state, {
		type: "SELECT_SCENE",
		sceneId: "defer-lifo",
	});
	assert.equal(state.sceneId, "defer-lifo");
	assert.equal(state.isAnimating, false);
});

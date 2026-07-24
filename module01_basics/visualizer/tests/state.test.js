"use strict";

const test = require("node:test");
const assert = require("node:assert/strict");

const labs = require("../scenes.js");
const {
	createInitialState,
	reduceState,
	getCurrentLab,
	getCurrentStep,
	actionFromKey,
	shouldAnimateTransition,
	animationDelay,
} = require("../app.js");

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
	assert.doesNotMatch(
		serialized,
		/prediction|choices|correctChoiceId|predictionChoice/,
	);
	for (const lab of labs) assert.equal("code" in lab, false);
});

test("uses the approved eight-step Slice runtime sequence", () => {
	assert.deepEqual(
		labs[0].steps.map((step) => step.id),
		[
			"descriptor",
			"address",
			"reslice",
			"append-in-place",
			"growslice",
			"allocate",
			"copy",
			"return-descriptor",
		],
	);
});

test("uses the approved eight-step classic Defer sequence", () => {
	assert.deepEqual(
		labs[1].steps.map((step) => step.id),
		[
			"frame",
			"register-d1",
			"register-d2",
			"deferreturn",
			"execute-d2",
			"execute-d1",
			"return",
			"panic-entry",
		],
	);
});

test("teacher controls move deterministically through a lab", () => {
	let state = reduceState(createInitialState(), {
		type: "SELECT_LAB",
		labId: "slice-runtime",
	});
	assert.equal(state.stepIndex, 0);
	assert.equal(getCurrentLab(state).id, "slice-runtime");
	assert.equal(getCurrentStep(state).id, "descriptor");

	state = reduceState(state, { type: "NEXT" });
	assert.equal(state.stepIndex, 1);
	state = reduceState(state, { type: "PREVIOUS" });
	assert.equal(state.stepIndex, 0);

	state = {
		...state,
		stepIndex: getCurrentLab(state).steps.length - 1,
	};
	state = reduceState(state, { type: "NEXT" });
	assert.equal(state.labId, null);
	assert.equal(state.stepIndex, 0);
});

test("rejects unknown labs and resets transient state when switching", () => {
	const initial = createInitialState();
	assert.deepEqual(
		reduceState(initial, { type: "SELECT_LAB", labId: "missing" }),
		initial,
	);

	let state = reduceState(initial, {
		type: "SELECT_LAB",
		labId: "slice-runtime",
	});
	state = {
		...state,
		stepIndex: 5,
		replayNonce: 3,
		error: "old",
	};
	state = reduceState(state, {
		type: "SELECT_LAB",
		labId: "defer-runtime",
	});
	assert.deepEqual(state, {
		labId: "defer-runtime",
		stepIndex: 0,
		isAnimating: false,
		replayNonce: 0,
		error: null,
	});
});

test("animation lock ignores navigation but permits lab switching", () => {
	let state = reduceState(createInitialState(), {
		type: "SELECT_LAB",
		labId: "slice-runtime",
	});
	state = reduceState(state, { type: "ANIMATION_START" });
	assert.equal(state.isAnimating, true);
	assert.deepEqual(reduceState(state, { type: "NEXT" }), state);
	assert.deepEqual(reduceState(state, { type: "REPLAY" }), state);

	state = reduceState(state, {
		type: "SELECT_LAB",
		labId: "defer-runtime",
	});
	assert.equal(state.labId, "defer-runtime");
	assert.equal(state.isAnimating, false);
});

test("maps classroom keyboard controls", () => {
	assert.deepEqual(actionFromKey("ArrowRight"), { type: "NEXT" });
	assert.deepEqual(actionFromKey(" "), { type: "NEXT" });
	assert.deepEqual(actionFromKey("ArrowLeft"), { type: "PREVIOUS" });
	assert.deepEqual(actionFromKey("r"), { type: "REPLAY" });
	assert.deepEqual(actionFromKey("R"), { type: "REPLAY" });
	assert.equal(actionFromKey("Enter"), null);
});

test("only visible step transitions request animation", () => {
	let state = reduceState(createInitialState(), {
		type: "SELECT_LAB",
		labId: "slice-runtime",
	});
	const next = reduceState(state, { type: "NEXT" });
	assert.equal(
		shouldAnimateTransition(state, next, { type: "NEXT" }),
		true,
	);
	assert.equal(
		shouldAnimateTransition(state, state, { type: "NEXT" }),
		false,
	);
});

test("reduced motion removes the controller delay", () => {
	assert.equal(animationDelay(true), 0);
	assert.equal(animationDelay(false), 900);
});

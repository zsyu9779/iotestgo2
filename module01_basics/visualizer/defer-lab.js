"use strict";

(function exposeDeferLab(root) {
	const utils =
		root.RuntimeLabUtils ||
		(typeof require === "function" ? require("./render-utils.js") : null);
	const { escapeHTML } = utils;

	const statusLabels = {
		head: "链表头",
		linked: "已链接",
		selected: "当前选中",
		executing: "已脱链 · 正在执行",
		detached: "已脱链",
	};

	function renderNodeFields(node) {
		const fields = [
			["heap", String(node.heap), "记录是否位于堆上"],
			["sp", node.sp, "注册它的函数栈帧"],
			["pc", node.pc, "注册时程序位置"],
			["fn", node.fn, "等待执行的函数值"],
			["link", node.link || "nil", "更早注册的 _defer"],
		];
		return fields
			.map(
				([name, value, description]) => `
					<div class="defer-field" data-field="${escapeHTML(name)}">
						<strong>${escapeHTML(name)}</strong>
						<code>${escapeHTML(value)}</code>
						<small>${escapeHTML(description)}</small>
					</div>
				`,
			)
			.join("");
	}

	function renderNodeInspector(node) {
		return `
			<article
				class="defer-structure"
				data-structure="runtime._defer"
				data-inspector-node="${escapeHTML(node.id)}"
			>
				<header class="structure-heading">
					<span>runtime._defer · ${escapeHTML(node.id)}</span>
					<code>${escapeHTML(node.address)}</code>
				</header>
				<div class="defer-field-table">${renderNodeFields(node)}</div>
			</article>
		`;
	}

	function renderInspector(state) {
		const structures =
			state.nodes.length > 0
				? state.nodes.map(renderNodeInspector).join("")
				: `
					<div class="empty-structure">
						<strong>当前没有 _defer 节点</strong>
						<span>执行 deferproc 后，这里会出现真实字段。</span>
					</div>
				`;
		return `
			<div class="panel-heading">
				<span>结构检查器</span>
				<strong>runtime._defer</strong>
			</div>
			<p class="implementation-boundary">
				<strong>经典 _defer 链表实现模型。</strong>
				现代 Go 编译器可能使用 <code>open-coded defer</code>；本实验室专门展开链表数据结构。
			</p>
			<div class="defer-structure-list">${structures}</div>
			<div class="unexpanded-fields">
				<strong>完整结构中的其他字段</strong>
				<span><code>rangefunc</code>、<code>head</code>：本场景不展开</span>
			</div>
		`;
	}

	function renderGoroutine(goroutine) {
		const head = goroutine.deferHead || "nil";
		return `
			<section
				class="goroutine-panel"
				data-goroutine="${escapeHTML(goroutine.id)}"
				data-g-defer-head="${escapeHTML(head)}"
			>
				<header>
					<strong>${escapeHTML(goroutine.id)}</strong>
					<span>goroutine runtime state</span>
				</header>
				<div class="g-head-pointer">
					<code>g._defer = ${escapeHTML(head)}</code>
					<span>${head === "nil" ? "空链表" : "指向当前链表头"}</span>
				</div>
			</section>
		`;
	}

	function renderFrame(frame) {
		return `
			<section class="stack-frame" data-frame="${escapeHTML(frame.id)}">
				<header>
					<strong>stack frame · ${escapeHTML(frame.id)}</strong>
					<span>${escapeHTML(frame.status || "active")}</span>
				</header>
				<div class="register-value" data-register="sp">
					<span>SP</span><code>${escapeHTML(frame.sp)}</code>
				</div>
				<div class="register-value" data-register="return-pc">
					<span>return PC</span><code>${escapeHTML(frame.returnPC)}</code>
				</div>
			</section>
		`;
	}

	function renderChainNode(node) {
		const target = node.link || "nil";
		return `
			<article
				class="defer-node is-${escapeHTML(node.status)}"
				data-node-id="${escapeHTML(node.id)}"
				data-address="${escapeHTML(node.address)}"
				data-status="${escapeHTML(node.status)}"
				data-structure="runtime._defer"
				data-link-target="${escapeHTML(target)}"
			>
				<header>
					<strong>${escapeHTML(node.id)}</strong>
					<code>${escapeHTML(node.address)}</code>
				</header>
				<div><span>sp</span><code>${escapeHTML(node.sp)}</code></div>
				<div><span>pc</span><code>${escapeHTML(node.pc)}</code></div>
				<div><span>fn</span><code>${escapeHTML(node.fn)}</code></div>
				<div class="node-link">
					<span>link</span><code>${escapeHTML(target)}</code>
				</div>
				<small>${escapeHTML(statusLabels[node.status] || node.status)}</small>
			</article>
		`;
	}

	function chainLabel(nodes, headAddress) {
		if (!headAddress) return "nil";
		const byAddress = new Map(nodes.map((node) => [node.address, node]));
		const labels = [];
		let address = headAddress;
		const visited = new Set();
		while (address && !visited.has(address)) {
			visited.add(address);
			const node = byAddress.get(address);
			if (!node) break;
			labels.push(node.id);
			address = node.link;
		}
		labels.push("nil");
		return labels.join("->");
	}

	function renderLinkArrows(nodes) {
		const linked = nodes.filter(
			(node) => node.status !== "detached" && node.status !== "executing",
		);
		if (linked.length === 0) return "";
		const lines = linked
			.map((node, index) => {
				const y = 30 + index * 38;
				return `
					<g class="link-arrow" data-link-from="${escapeHTML(node.address)}" data-link-to="${escapeHTML(node.link || "nil")}">
						<text x="8" y="${y - 6}">${escapeHTML(node.id)}.link</text>
						<path d="M 8 ${y} H 520"></path>
						<path d="M 510 ${y - 6} L 520 ${y} L 510 ${y + 6}"></path>
						<text x="535" y="${y + 5}">${escapeHTML(node.link || "nil")}</text>
					</g>
				`;
			})
			.join("");
		return `
			<svg class="defer-link-map" viewBox="0 0 760 ${Math.max(72, linked.length * 44)}" role="img" aria-label="_defer link 指针关系">
				<title>_defer 链接关系</title>
				${lines}
			</svg>
		`;
	}

	function renderStage(step) {
		const { state } = step;
		const head = state.goroutine.deferHead;
		const chain = chainLabel(state.nodes, head);
		const nodes =
			state.nodes.length > 0
				? state.nodes.map(renderChainNode).join("")
				: '<div class="nil-node" data-node-id="nil">nil</div>';
		const completion = state.completion
			? `<p class="completion-note">${escapeHTML(state.completion)}</p>`
			: "";
		return `
			<div
				class="defer-runtime-stage"
				data-phase="${escapeHTML(state.phase)}"
				data-chain="${escapeHTML(chain)}"
				${state.panicBoundary ? 'data-panic-entry="true"' : ""}
				role="img"
				aria-label="${escapeHTML(step.title)}"
			>
				<div class="runtime-context">
					${renderGoroutine(state.goroutine)}
					${renderFrame(state.frame)}
				</div>
				<div class="linked-list-lane" aria-label="${escapeHTML(chain.replaceAll("->", " → "))}">
					${nodes}
					<div class="nil-node" data-node-id="nil">nil</div>
				</div>
				${renderLinkArrows(state.nodes)}
				${completion}
			</div>
		`;
	}

	function renderOperation(step) {
		const { operation } = step;
		const atomic = (operation.atomic || [])
			.map(
				(item, index) => `
					<li data-operation-index="${index}">
						<span>${index + 1}</span><code>${escapeHTML(item)}</code>
					</li>
				`,
			)
			.join("");
		return `
			<div class="panel-heading">
				<span>当前 runtime 操作</span>
				<strong>${escapeHTML(operation.name)}</strong>
			</div>
			<code class="operation-call">${escapeHTML(operation.call)}</code>
			<p class="operation-summary">${escapeHTML(operation.summary)}</p>
			${atomic ? `<ol class="atomic-operations">${atomic}</ol>` : ""}
			<p class="implementation-boundary compact">
				经典 _defer 链表实现模型 · 现代编译器可能使用 <code>open-coded defer</code>
			</p>
		`;
	}

	function renderDeferLab(step) {
		if (!step || !step.state || !step.operation) {
			throw new TypeError("Defer runtime step is incomplete");
		}
		return {
			inspector: renderInspector(step.state),
			stage: renderStage(step),
			operation: renderOperation(step),
		};
	}

	const api = { renderDeferLab, chainLabel };
	root.RuntimeLabRenderers = {
		...(root.RuntimeLabRenderers || {}),
		...api,
	};
	if (typeof module !== "undefined" && module.exports) {
		module.exports = api;
	}
})(typeof globalThis !== "undefined" ? globalThis : window);

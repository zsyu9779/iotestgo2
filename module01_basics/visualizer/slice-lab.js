"use strict";

(function exposeSliceLab(root) {
	const utils =
		root.RuntimeLabUtils ||
		(typeof require === "function" ? require("./render-utils.js") : null);
	const { escapeHTML, renderFieldRow } = utils;

	function addHexAddress(baseAddress, byteOffset) {
		const width = String(baseAddress).replace(/^0x/i, "").length;
		const value = Number.parseInt(baseAddress, 16) + byteOffset;
		return `0x${value.toString(16).padStart(width, "0")}`;
	}

	function renderDescriptor(descriptor) {
		const fields = [
			{
				offset: 0,
				name: "array",
				width: 8,
				value: descriptor.array,
				label: "unsafe.Pointer",
			},
			{
				offset: 8,
				name: "len",
				width: 8,
				value: descriptor.length,
				label: "int",
			},
			{
				offset: 16,
				name: "cap",
				width: 8,
				value: descriptor.capacity,
				label: "int",
			},
		];
		const rows = fields
			.map((field) =>
				renderFieldRow({
					...field,
					active: descriptor.activeField === field.name,
				}),
			)
			.join("");

		return `
			<section
				class="descriptor"
				data-structure="runtime.slice"
				data-byte-size="24"
				data-descriptor-id="${escapeHTML(descriptor.id)}"
			>
				<header class="structure-heading">
					<span>runtime.slice · ${escapeHTML(descriptor.id)}</span>
					<code>${escapeHTML(descriptor.address)} · 24 B</code>
				</header>
				<div class="field-table">${rows}</div>
			</section>
		`;
	}

	function renderInspector(state) {
		return `
			<div class="panel-heading">
				<span>结构检查器</span>
				<strong>Go 1.25.11 · darwin/arm64</strong>
			</div>
			<p class="boundary-note">
				字段布局真实；十六进制数值是<strong>示意地址</strong>，不是浏览器读取的 Go 进程地址。
			</p>
			<div class="descriptor-list">
				${state.descriptors.map(renderDescriptor).join("")}
			</div>
		`;
	}

	function renderMemoryArray(array, state) {
		const activeTargets = new Set(state.activeTargets || []);
		const isActiveArray = activeTargets.has(array.id);
		const cells = Array.from({ length: array.capacity }, (_, index) => {
			const hasValue = index < array.values.length;
			const value = hasValue ? array.values[index] : "空槽";
			const byteOffset = index * state.elementSize;
			const address = addHexAddress(array.baseAddress, byteOffset);
			const targetKey = `${array.id}[${index}]`;
			const isTargeted = activeTargets.has(targetKey);
			const isCopyCell =
				state.phase === "copy" &&
				index < (state.copyCount || 0) &&
				["old-array", "new-array"].includes(array.id);
			let stateName = hasValue ? "stored" : "capacity";
			let stateLabel = "";
			if (isTargeted) {
				stateName = "targeted";
				stateLabel = "当前目标";
			}
			if (isCopyCell) {
				stateName =
					array.id === "old-array" ? "copy-source" : "copied";
				stateLabel = array.id === "old-array" ? "复制源" : "已复制";
			}
			if (state.selectedIndex === index) {
				stateName = "reading";
				stateLabel = "正在读取";
			}
			if (array.writtenIndex === index) {
				stateName = "written";
				stateLabel = "调用方写入";
			}
			return `
				<div
					id="${escapeHTML(array.id)}-${index}"
					class="memory-cell is-${stateName}"
					data-memory-target="${escapeHTML(array.id)}-${index}"
					data-address="${escapeHTML(address)}"
					data-state="${escapeHTML(stateName)}"
					${isTargeted ? 'data-active-target="true"' : ""}
				>
					<span class="cell-offset">+${byteOffset} B</span>
					<strong class="cell-value">${escapeHTML(value)}</strong>
					<code class="cell-address">${escapeHTML(address)}</code>
					<small>index ${index}</small>
					${stateLabel ? `<span class="cell-state-label">${escapeHTML(stateLabel)}</span>` : ""}
				</div>
			`;
		}).join("");

		const capacityNote = array.capacityNote
			? `<p class="allocation-note">${escapeHTML(array.capacityNote)}</p>`
			: "";
		const arrayStateLabel =
			isActiveArray && array.allocation === "new"
				? '<span class="allocation-state">新分配</span>'
				: isActiveArray
					? '<span class="allocation-state">当前对象</span>'
					: "";
		return `
			<section
				class="memory-allocation"
				data-array-id="${escapeHTML(array.id)}"
				data-allocation="${escapeHTML(array.allocation)}"
				${isActiveArray ? 'data-active-target="true"' : ""}
			>
				<header>
					<strong>${escapeHTML(array.id)}</strong>
					<span class="allocation-meta">
						${arrayStateLabel}
						<code>base ${escapeHTML(array.baseAddress)}</code>
					</span>
				</header>
				<div class="memory-cells">${cells}</div>
				${capacityNote}
			</section>
		`;
	}

	function renderPointerMap(state) {
		const links = state.descriptors
			.map((descriptor, index) => {
				const targetArray = state.arrays.find((array) => {
					const start = Number.parseInt(array.baseAddress, 16);
					const end = start + array.capacity * state.elementSize;
					const pointer = Number.parseInt(descriptor.array, 16);
					return pointer >= start && pointer < end;
				});
				if (!targetArray) return "";
				const offset =
					Number.parseInt(descriptor.array, 16) -
					Number.parseInt(targetArray.baseAddress, 16);
				const targetIndex = offset / state.elementSize;
				const y = 38 + index * 44;
				return `
					<g
						class="pointer-link"
						data-pointer-from="${escapeHTML(descriptor.id)}.array"
						data-pointer-to="${escapeHTML(targetArray.id)}-${targetIndex}"
					>
						<text x="8" y="${y - 7}">${escapeHTML(descriptor.id)}.array</text>
						<path d="M 8 ${y} H 660"></path>
						<path d="M 650 ${y - 7} L 660 ${y} L 650 ${y + 7}"></path>
						<text x="675" y="${y + 5}">${escapeHTML(targetArray.id)}[${targetIndex}] · ${escapeHTML(descriptor.array)}</text>
					</g>
				`;
			})
			.join("");

		return `
			<svg
				class="pointer-map"
				viewBox="0 0 1000 ${Math.max(90, state.descriptors.length * 52)}"
				role="img"
				aria-label="Slice array 字段到具体连续内存单元的指向关系"
			>
				<title>Descriptor 指针连接</title>
				${links}
			</svg>
		`;
	}

	function renderBounds(state) {
		return state.descriptors
			.map(
				(descriptor) => `
					<div class="descriptor-bounds" data-bounds-for="${escapeHTML(descriptor.id)}">
						<strong>${escapeHTML(descriptor.id)}</strong>
						<span class="len-bound">len 范围：[0, ${descriptor.length})</span>
						<span class="cap-bound">cap 范围：[0, ${descriptor.capacity})</span>
					</div>
				`,
			)
			.join("");
	}

	function renderStage(step) {
		const { state } = step;
		const arrays = state.arrays
			.map((array) => renderMemoryArray(array, state))
			.join("");
		const gcNote = state.gcNote
			? `<p class="gc-note">${escapeHTML(state.gcNote)}</p>`
			: "";
		return `
			<div
				class="slice-runtime-stage"
				data-phase="${escapeHTML(state.phase)}"
				role="img"
				aria-label="${escapeHTML(step.title)}"
			>
				${renderPointerMap(state)}
				<div class="allocation-list">${arrays}</div>
				<div class="bounds-list">${renderBounds(state)}</div>
				${gcNote}
			</div>
		`;
	}

	function renderPipeline(pipeline = []) {
		if (pipeline.length === 0) return "";
		return `
			<ol class="runtime-pipeline" aria-label="runtime 调用链">
				${pipeline
					.map(
						([name, status]) => `
							<li
								data-call="${escapeHTML(name)}"
								data-status="${escapeHTML(status)}"
							>
								<strong>${escapeHTML(name)}</strong>
								<span>${status === "done" ? "已完成" : status === "active" ? "当前操作" : "等待"}</span>
							</li>
						`,
					)
					.join("")}
			</ol>
		`;
	}

	function renderOperation(step) {
		const { operation } = step;
		const formulas = (operation.formulas || [])
			.map((formula) => `<code class="formula">${escapeHTML(formula)}</code>`)
			.join("");
		const branch = operation.branch
			? `
				<div class="runtime-decision" data-branch="${escapeHTML(operation.branch)}">
					<code>newLen &lt;= cap</code>
					<strong>不调用 runtime.growslice</strong>
				</div>
			`
			: "";
		return `
			<div class="panel-heading">
				<span>当前 runtime 操作</span>
				<strong>${escapeHTML(operation.name)}</strong>
			</div>
			<code class="operation-call">${escapeHTML(operation.call)}</code>
			<p class="operation-summary">${escapeHTML(operation.summary)}</p>
			${branch}
			<div class="formula-list">${formulas}</div>
			${renderPipeline(operation.pipeline)}
		`;
	}

	function renderSliceLab(step) {
		if (!step || !step.state || !step.operation) {
			throw new TypeError("Slice runtime step is incomplete");
		}
		return {
			inspector: renderInspector(step.state),
			stage: renderStage(step),
			operation: renderOperation(step),
		};
	}

	const api = { renderSliceLab, addHexAddress };
	root.RuntimeLabRenderers = {
		...(root.RuntimeLabRenderers || {}),
		...api,
	};
	if (typeof module !== "undefined" && module.exports) {
		module.exports = api;
	}
})(typeof globalThis !== "undefined" ? globalThis : window);

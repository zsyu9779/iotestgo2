"use strict";

(function exposeRuntimeLabs(root) {
	const sliceBaseArray = {
		id: "base-array",
		baseAddress: "0x1000",
		values: [1, 2, 3, 4],
		capacity: 4,
		allocation: "shared",
	};

	const fullOldArray = {
		id: "old-array",
		baseAddress: "0x2000",
		values: [10, 20],
		capacity: 2,
		allocation: "old",
	};

	const grownArrayEmpty = {
		id: "new-array",
		baseAddress: "0x4000",
		values: [],
		capacity: 4,
		capacityNote: "本次 Go 1.25.11 路径得到 4；增长算法不是语言保证",
		allocation: "new",
	};

	const grownArrayCopied = {
		...grownArrayEmpty,
		values: [10, 20, 30],
	};

	const sliceSteps = [
		{
			id: "descriptor",
			title: "Slice 是一个 24 字节 Descriptor",
			operation: {
				name: "runtime.slice",
				call: "array | len | cap",
				summary:
					"arm64 下三个字段各占 8 字节。Descriptor 本身不包含数组元素。",
			},
			state: {
				phase: "descriptor",
				elementSize: 8,
				descriptors: [
					{
						id: "base",
						address: "0x7000",
						array: "0x1000",
						length: 4,
						capacity: 4,
						activeField: "array",
					},
				],
				arrays: [sliceBaseArray],
				activeTargets: ["base.array", "base-array[0]"],
			},
		},
		{
			id: "address",
			title: "array 指针参与元素地址计算",
			operation: {
				name: "element address",
				call: "0x1000 + 2 × 8 B = 0x1010",
				summary:
					"读取下标 2 时，运行时按元素宽度从 array 指针计算目标地址。",
			},
			state: {
				phase: "address",
				elementSize: 8,
				selectedIndex: 2,
				descriptors: [
					{
						id: "base",
						address: "0x7000",
						array: "0x1000",
						length: 4,
						capacity: 4,
						activeField: "array",
					},
				],
				arrays: [sliceBaseArray],
				activeTargets: ["base.array", "base-array[2]"],
			},
		},
		{
			id: "reslice",
			title: "重新切片只生成新的 Descriptor",
			operation: {
				name: "reslice",
				call: "new.array = old.array + low × elementSize",
				summary:
					"low=1、high=3；底层连续内存没有复制，只有三个字段被重新计算。",
				formulas: [
					"new.array = old.array + low × elementSize",
					"new.len = high - low",
					"new.cap = oldCap - low",
				],
			},
			state: {
				phase: "reslice",
				elementSize: 8,
				low: 1,
				high: 3,
				descriptors: [
					{
						id: "old",
						address: "0x7000",
						array: "0x1000",
						length: 4,
						capacity: 4,
					},
					{
						id: "view",
						address: "0x7020",
						array: "0x1008",
						length: 2,
						capacity: 3,
						activeField: "array",
					},
				],
				arrays: [sliceBaseArray],
				activeTargets: ["view.array", "base-array[1]"],
			},
		},
		{
			id: "append-in-place",
			title: "容量足够：直接写入容量槽",
			operation: {
				name: "append fast path",
				call: "newLen = 3; newLen <= cap",
				summary:
					"不调用 runtime.growslice；写入原连续内存，并返回 len 更新后的 Descriptor。",
				branch: "reuse",
			},
			state: {
				phase: "append-in-place",
				elementSize: 8,
				descriptors: [
					{
						id: "before",
						address: "0x7100",
						array: "0x1800",
						length: 2,
						capacity: 4,
					},
					{
						id: "returned",
						address: "0x7120",
						array: "0x1800",
						length: 3,
						capacity: 4,
						activeField: "len",
					},
				],
				arrays: [
					{
						id: "spare-array",
						baseAddress: "0x1800",
						values: [10, 20, 30],
						capacity: 4,
						allocation: "reused",
						writtenIndex: 2,
					},
				],
				activeTargets: ["returned.len", "spare-array[2]"],
			},
		},
		{
			id: "growslice",
			title: "容量不足：进入 runtime.growslice",
			operation: {
				name: "runtime.growslice",
				call:
					"growslice(oldPtr=0x2000, newLen=3, oldCap=2, num=1, et=int)",
				summary:
					"newLen 已超过 oldCap，编译器生成的慢路径必须请求新的存储。",
				pipeline: [
					["runtime.growslice", "active"],
					["nextslicecap", "pending"],
					["mallocgc", "pending"],
					["memmove", "pending"],
					["return", "pending"],
				],
			},
			state: {
				phase: "growslice",
				elementSize: 8,
				newLength: 3,
				descriptors: [
					{
						id: "old",
						address: "0x7200",
						array: "0x2000",
						length: 2,
						capacity: 2,
						activeField: "cap",
					},
				],
				arrays: [fullOldArray],
				activeTargets: ["old.cap"],
			},
		},
		{
			id: "allocate",
			title: "计算 newCap，并由 mallocgc 分配",
			operation: {
				name: "nextslicecap → mallocgc",
				call: "newCap = 4; mallocgc(4 × 8 B, intType, ...)",
				summary:
					"本次 Go 1.25.11 路径得到 newCap=4；不能把它推广成固定翻倍规则。",
				pipeline: [
					["runtime.growslice", "done"],
					["nextslicecap", "done"],
					["mallocgc", "active"],
					["memmove", "pending"],
					["return", "pending"],
				],
			},
			state: {
				phase: "allocate",
				elementSize: 8,
				newLength: 3,
				newCapacity: 4,
				descriptors: [
					{
						id: "old",
						address: "0x7200",
						array: "0x2000",
						length: 2,
						capacity: 2,
					},
				],
				arrays: [fullOldArray, grownArrayEmpty],
				activeTargets: ["new-array"],
			},
		},
		{
			id: "copy",
			title: "memmove 复制旧元素，再写入新元素",
			operation: {
				name: "memmove",
				call: "memmove(0x4000, 0x2000, 2 × 8 B)",
				summary:
					"旧数组的有效元素按字节复制到新分配；追加值写到下标 2。",
				pipeline: [
					["runtime.growslice", "done"],
					["nextslicecap", "done"],
					["mallocgc", "done"],
					["memmove", "active"],
					["return", "pending"],
				],
			},
			state: {
				phase: "copy",
				elementSize: 8,
				copyCount: 2,
				descriptors: [
					{
						id: "old",
						address: "0x7200",
						array: "0x2000",
						length: 2,
						capacity: 2,
					},
				],
				arrays: [fullOldArray, grownArrayCopied],
				activeTargets: ["old-array", "new-array"],
			},
		},
		{
			id: "return-descriptor",
			title: "growslice 返回新的 Descriptor",
			operation: {
				name: "return from runtime.growslice",
				call: "slice{array: 0x4000, len: 3, cap: 4}",
				summary:
					"旧 Descriptor 仍指向旧内存；旧内存只有在不可达后才有资格被 GC 回收。",
				pipeline: [
					["runtime.growslice", "done"],
					["nextslicecap", "done"],
					["mallocgc", "done"],
					["memmove", "done"],
					["return", "active"],
				],
			},
			state: {
				phase: "return-descriptor",
				elementSize: 8,
				descriptors: [
					{
						id: "old",
						address: "0x7200",
						array: "0x2000",
						length: 2,
						capacity: 2,
					},
					{
						id: "grown",
						address: "0x7240",
						array: "0x4000",
						length: 3,
						capacity: 4,
						activeField: "array",
					},
				],
				arrays: [fullOldArray, grownArrayCopied],
				activeTargets: ["grown.array", "new-array[0]"],
				gcNote: "旧内存仅在不可达后才有资格被 GC 回收。",
			},
		},
	];

	function deferNode({
		id,
		address,
		pc,
		fn,
		link,
		status,
	}) {
		return {
			id,
			address,
			heap: false,
			sp: "0x8f00",
			pc,
			fn,
			link,
			status,
		};
	}

	const d1Linked = deferNode({
		id: "D1",
		address: "0x3000",
		pc: "0x401140",
		fn: "deferredFnA",
		link: null,
		status: "linked",
	});

	const d2Head = deferNode({
		id: "D2",
		address: "0x3100",
		pc: "0x401180",
		fn: "deferredFnB",
		link: "0x3000",
		status: "head",
	});

	const deferFrame = {
		id: "work",
		sp: "0x8f00",
		returnPC: "0x401240",
	};

	const deferSteps = [
		{
			id: "frame",
			title: "函数栈帧与 goroutine 已建立",
			operation: {
				name: "enter frame",
				call: "g._defer = nil",
				summary:
					"当前函数拥有自己的 SP 和返回位置；goroutine 的 defer 链表头为空。",
			},
			state: {
				phase: "frame",
				goroutine: { id: "g17", deferHead: null },
				frame: deferFrame,
				nodes: [],
			},
		},
		{
			id: "register-d1",
			title: "创建 D1，并写入链表头",
			operation: {
				name: "deferproc",
				call: "D1.link = g._defer; g._defer = D1",
				summary:
					"节点保存注册栈帧、程序位置和函数值；D1.link 为 nil。",
				atomic: ["D1.link = nil", "g._defer = D1"],
			},
			state: {
				phase: "register-d1",
				goroutine: { id: "g17", deferHead: "0x3000" },
				frame: deferFrame,
				nodes: [{ ...d1Linked, status: "head" }],
			},
		},
		{
			id: "register-d2",
			title: "D2 以头插法链接到 D1",
			operation: {
				name: "deferproc",
				call: "D2.link = g._defer; g._defer = D2",
				summary:
					"后注册的 D2 成为链表头，D2.link 保存此前的头指针 0x3000。",
				atomic: ["D2.link = g._defer", "g._defer = D2"],
			},
			state: {
				phase: "register-d2",
				goroutine: { id: "g17", deferHead: "0x3100" },
				frame: deferFrame,
				nodes: [d2Head, d1Linked],
			},
		},
		{
			id: "deferreturn",
			title: "正常返回进入 deferreturn",
			operation: {
				name: "runtime.deferreturn",
				call: "find head node with sp == frame.SP",
				summary:
					"返回路径从 g._defer 头部检查属于当前栈帧的记录。",
				atomic: ["head = g._defer", "match head.sp with frame.SP"],
			},
			state: {
				phase: "deferreturn",
				goroutine: { id: "g17", deferHead: "0x3100" },
				frame: deferFrame,
				nodes: [
					{ ...d2Head, status: "selected" },
					d1Linked,
				],
			},
		},
		{
			id: "execute-d2",
			title: "先脱链 D2，再调用 D2.fn",
			operation: {
				name: "pop head D2",
				call: "head = D2.link; call D2.fn",
				summary:
					"头指针先移动到 D1，D2 从链表脱离，然后执行保存的函数值。",
				atomic: [
					"fn = D2.fn",
					"head = D2.link",
					"g._defer = head",
					"call D2.fn",
				],
			},
			state: {
				phase: "execute-d2",
				goroutine: { id: "g17", deferHead: "0x3000" },
				frame: deferFrame,
				nodes: [
					{ ...d2Head, status: "executing" },
					{ ...d1Linked, status: "head" },
				],
			},
		},
		{
			id: "execute-d1",
			title: "再脱链 D1，头指针变为 nil",
			operation: {
				name: "pop head D1",
				call: "head = D1.link; call D1.fn",
				summary:
					"D1.link 为 nil，因此移除 D1 后，goroutine 的 defer 链表清空。",
				atomic: [
					"fn = D1.fn",
					"head = D1.link",
					"g._defer = nil",
					"call D1.fn",
				],
			},
			state: {
				phase: "execute-d1",
				goroutine: { id: "g17", deferHead: null },
				frame: deferFrame,
				nodes: [
					{ ...d2Head, status: "detached" },
					{ ...d1Linked, status: "executing" },
				],
			},
		},
		{
			id: "return",
			title: "链表清空，函数完成返回",
			operation: {
				name: "return",
				call: "g._defer == nil",
				summary:
					"LIFO 不是抽象卡片规则，而是头插、头取数据结构产生的顺序。",
			},
			state: {
				phase: "return",
				goroutine: { id: "g17", deferHead: null },
				frame: { ...deferFrame, status: "returning" },
				nodes: [
					{ ...d2Head, status: "detached" },
					{ ...d1Linked, status: "detached" },
				],
				completion: "链表已清空",
			},
		},
		{
			id: "panic-entry",
			title: "panic 是另一条遍历入口",
			operation: {
				name: "panic unwind",
				call: "_panic.next() walks pending defers",
				summary:
					"异常返回也会处理待执行记录；本实验室不展开 _panic 与 recover 状态机。",
			},
			state: {
				phase: "panic-entry",
				goroutine: { id: "g17", deferHead: "0x3100" },
				frame: { ...deferFrame, status: "panicking" },
				nodes: [d2Head, d1Linked],
				panicBoundary: true,
			},
		},
	];

	const labs = [
		{
			id: "slice-runtime",
			kind: "slice",
			chapter: "Slice",
			eyebrow: "Descriptor · growslice",
			title: "Slice Descriptor 与扩容",
			description:
				"拆开三个机器字、连续内存、指针运算和 runtime.growslice 调用链。",
			steps: sliceSteps,
		},
		{
			id: "defer-runtime",
			kind: "defer",
			chapter: "Defer",
			eyebrow: "_defer · linked list",
			title: "经典 _defer 链表与返回",
			description:
				"展开 goroutine 头指针、真实节点字段、头插注册和头取执行。",
			steps: deferSteps,
		},
	];

	root.RUNTIME_LABS = labs;
	if (typeof module !== "undefined" && module.exports) {
		module.exports = labs;
	}
})(typeof globalThis !== "undefined" ? globalThis : window);

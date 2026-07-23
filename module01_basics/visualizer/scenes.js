"use strict";

(function exposeScenes(root) {
	const scenes = [
		{
			id: "slice-shared",
			chapter: "Slice",
			eyebrow: "视图与共享",
			title: "两个 Slice，共享一个数组",
			description: "观察两个独立的 Slice Header 如何指向同一段底层数组。",
			code: [
				"base := []int{1, 2, 3, 4}",
				"view := base[1:3]",
				"view[0] = 9",
				"fmt.Println(base)",
			],
			steps: [
				{
					kind: "observe",
					line: 0,
					title: "创建 base",
					narration: "一个 Slice Header 指向底层数组的第一个元素。",
					stage: { type: "slice-shared", phase: "base" },
				},
				{
					kind: "observe",
					line: 1,
					title: "创建 view",
					narration: "新的 Header 从数组下标 1 开始观察，Header 独立，数组仍然共享。",
					stage: { type: "slice-shared", phase: "view" },
				},
				{
					kind: "prediction",
					line: 2,
					title: "先预测",
					narration: "修改 view[0] 后，base 会变化吗？",
					prediction: {
						choices: [
							{ id: "changes", label: "base 会变化" },
							{ id: "unchanged", label: "base 不会变化" },
						],
						correctChoiceId: "changes",
						explanation:
							"base[1] 与 view[0] 对应同一个底层数组单元，所以修改会同时被两个视图观察到。",
					},
					stage: { type: "slice-shared", phase: "view" },
				},
				{
					kind: "reveal",
					line: 2,
					title: "修改共享单元",
					narration: "值 2 变为 9，两条指针保持不动。",
					stage: { type: "slice-shared", phase: "mutated" },
				},
				{
					kind: "observe",
					line: 3,
					title: "看见同一个位置",
					narration: "base[1] 与 view[0] 是两个索引表达式，却落在同一个数组单元。",
					stage: { type: "slice-shared", phase: "aliases" },
				},
				{
					kind: "conclusion",
					line: 3,
					title: "Header 是值，数组可共享",
					narration: "复制或分片会得到新的 Slice Header，不会自动复制底层数组元素。",
					stage: { type: "slice-shared", phase: "aliases" },
				},
			],
		},
		{
			id: "slice-append",
			chapter: "Slice",
			eyebrow: "容量与 append",
			title: "append：复用还是搬家",
			description: "把容量有余与容量已满放在同一舞台，比较共享关系的变化。",
			code: [
				"spare := make([]int, 2, 4)",
				"reused := append(spare, 3)",
				"full := make([]int, 2, 2)",
				"grown := append(full, 3)",
				"full[0], grown[0] = 7, 9",
			],
			steps: [
				{
					kind: "observe",
					line: 0,
					title: "并排观察两种容量",
					narration: "spare 还有两个空槽；full 的 len 已经等于 cap。",
					stage: { type: "slice-append", phase: "initial" },
				},
				{
					kind: "observe",
					line: 1,
					title: "容量有余：复用",
					narration: "新元素落入原数组空槽，reused 继续指向原数组。",
					stage: { type: "slice-append", phase: "reused" },
				},
				{
					kind: "prediction",
					line: 3,
					title: "先预测",
					narration: "full 已经没有空槽，append 还能继续写入原数组吗？",
					prediction: {
						choices: [
							{ id: "move", label: "需要新的数组" },
							{ id: "reuse", label: "继续复用原数组" },
						],
						correctChoiceId: "move",
						explanation:
							"新长度超过原容量，append 必须返回一个指向新存储区域的 Slice。",
					},
					stage: { type: "slice-append", phase: "before-growth" },
				},
				{
					kind: "reveal",
					line: 3,
					title: "容量不足：分配并复制",
					narration: "出现容量不小于新长度的数组，已有元素被复制过去，新元素随后落位。",
					stage: { type: "slice-append", phase: "reallocated" },
				},
				{
					kind: "observe",
					line: 3,
					title: "两个 Header 分开指向",
					narration: "full 留在旧数组，grown 指向新数组。",
					stage: { type: "slice-append", phase: "separated" },
				},
				{
					kind: "observe",
					line: 4,
					title: "修改验证分离",
					narration: "旧数组第一个元素变为 7，新数组第一个元素变为 9，彼此不再影响。",
					stage: { type: "slice-append", phase: "verified" },
				},
				{
					kind: "conclusion",
					line: 4,
					title: "必须接收 append 的返回值",
					narration: "返回的新 Slice 是否继续共享原数组，取决于 append 时容量是否足够。",
					stage: { type: "slice-append", phase: "verified" },
				},
			],
		},
		{
			id: "defer-lifo",
			chapter: "Defer",
			eyebrow: "注册与返回",
			title: "后注册，先执行",
			description: "跟随代码执行线，观察待执行调用如何在返回前逆序取出。",
			code: [
				"func work() {",
				"    defer print(\"A\")",
				"    defer print(\"B\")",
				"    print(\"work\")",
				"}",
			],
			steps: [
				{
					kind: "observe",
					line: 0,
					title: "进入函数",
					narration: "函数帧已建立，待执行调用区还是空的。",
					stage: { type: "defer-lifo", phase: "empty" },
				},
				{
					kind: "observe",
					line: 1,
					title: "注册 A",
					narration: "A 现在不会打印，它先进入待执行调用区。",
					stage: { type: "defer-lifo", phase: "register-a" },
				},
				{
					kind: "observe",
					line: 2,
					title: "注册 B",
					narration: "B 后注册，位于 A 上方。",
					stage: { type: "defer-lifo", phase: "register-b" },
				},
				{
					kind: "prediction",
					line: 4,
					title: "先预测",
					narration: "函数返回前会先执行 A 还是 B？",
					prediction: {
						choices: [
							{ id: "b-first", label: "先执行 B" },
							{ id: "a-first", label: "先执行 A" },
						],
						correctChoiceId: "b-first",
						explanation: "已注册的 Defer 调用按逆序执行，所以 B 先于 A。",
					},
					stage: { type: "defer-lifo", phase: "register-b" },
				},
				{
					kind: "reveal",
					line: 3,
					title: "先完成函数体",
					narration: "普通语句先输出 work，然后到达准备返回关口。",
					stage: { type: "defer-lifo", phase: "work" },
				},
				{
					kind: "observe",
					line: 4,
					title: "B 先执行",
					narration: "最后注册的 B 先从待执行区取出。",
					stage: { type: "defer-lifo", phase: "execute-b" },
				},
				{
					kind: "observe",
					line: 4,
					title: "A 再执行",
					narration: "A 随后执行，所有待执行调用完成后函数才真正返回。",
					stage: { type: "defer-lifo", phase: "execute-a" },
				},
				{
					kind: "conclusion",
					line: 4,
					title: "后注册的先执行",
					narration: "多个 Defer 调用的可观察执行顺序是 LIFO。",
					stage: { type: "defer-lifo", phase: "complete" },
				},
			],
		},
		{
			id: "defer-evaluation",
			chapter: "Defer",
			eyebrow: "求值时机",
			title: "参数快照与闭包读取",
			description: "区分 defer 语句执行时的参数求值与闭包体执行时的变量读取。",
			code: [
				"value := 1",
				"defer print(value)",
				"defer func() { print(value) }()",
				"value = 2",
				"// return",
			],
			steps: [
				{
					kind: "observe",
					line: 0,
					title: "创建变量",
					narration: "函数帧中的变量格 value 当前保存 1。",
					stage: { type: "defer-evaluation", phase: "value-one" },
				},
				{
					kind: "observe",
					line: 1,
					title: "普通参数立即求值",
					narration: "注册普通调用时，参数已经求值，卡片保存 arg = 1。",
					stage: { type: "defer-evaluation", phase: "save-argument" },
				},
				{
					kind: "observe",
					line: 2,
					title: "注册闭包",
					narration: "闭包卡片保留访问 value 的能力，函数体尚未执行。",
					stage: { type: "defer-evaluation", phase: "register-closure" },
				},
				{
					kind: "observe",
					line: 3,
					title: "变量变为 2",
					narration: "value 的变量格更新为 2；普通调用已经保存的参数不变。",
					stage: { type: "defer-evaluation", phase: "value-two" },
				},
				{
					kind: "prediction",
					line: 4,
					title: "先预测",
					narration: "两个调用分别会打印什么？",
					prediction: {
						choices: [
							{
								id: "closure-2-arg-1",
								label: "闭包 2，普通参数 1",
							},
							{ id: "both-2", label: "两者都是 2" },
						],
						correctChoiceId: "closure-2-arg-1",
						explanation:
							"普通参数已在注册时保存 1；闭包执行时读取变量格，得到 2。",
					},
					stage: { type: "defer-evaluation", phase: "value-two" },
				},
				{
					kind: "reveal",
					line: 4,
					title: "闭包先读取 2",
					narration: "后注册的闭包先执行，此刻读取 value，输出 2。",
					stage: { type: "defer-evaluation", phase: "execute-closure" },
				},
				{
					kind: "observe",
					line: 4,
					title: "普通调用输出 1",
					narration: "普通调用使用注册时保存的参数 1。",
					stage: { type: "defer-evaluation", phase: "complete" },
				},
				{
					kind: "conclusion",
					line: 4,
					title: "两个时间点",
					narration: "defer 语句执行时求参数；真正执行时才运行闭包体。",
					stage: { type: "defer-evaluation", phase: "complete" },
				},
			],
		},
	];

	root.VISUALIZER_SCENES = scenes;
	if (typeof module !== "undefined" && module.exports) {
		module.exports = scenes;
	}
})(typeof globalThis !== "undefined" ? globalThis : window);

# 六周课程 Module 标准

本标准用于六周 Go 后端课程的 Module 01–06。目标学员是有 Java 经验、无需 Go 前置知识的大学生。每个 Module 必须形成从讲师 Demo、学员实操到可重复验收的完整闭环；Module 01–02 已按本标准迁移，Module 03–06 后续迁移时继续遵守。

## 标准目录

每个 Module 只提供两个权威入口：Module 根目录的 `README.md` 是唯一学员入口，`instructor/RUNBOOK.md` 是唯一讲师入口。兼容性教案或小抄只能链接到这两个入口，不得维护第二套日程或命令。

```text
moduleXX_topic/
├── README.md
├── instructor/
│   ├── RUNBOOK.md
│   ├── DEMO_NOTES.md
│   └── RUBRIC.md
├── blocks/
│   └── NN_topic/{README.md,demo/,lab/}
├── integrated_lab/
├── homework/
├── assessments/
└── bonus/
```

核心内容按学习旅程组织为连续 Block，而不是按知识点保留重复入口。旧内容迁移时使用 `git mv` 保留历史；超出核心时间盒的内容进入 `bonus/`。

## 时间盒与讲练比例

每周授课一天：09:30–12:00、13:00–16:00；12:00–13:00 午休不计入课堂，上午和下午各有一次 10 分钟短休息，净教学与实操时间为 310 分钟。学员亲手编写、运行、测试和复盘的时间不得少于 155 分钟，即净课堂时间的 50%。

每 30–45 分钟应产生一个可运行、可验收的结果；较长 Block 必须设置中间检查点。讲师按“学习目标 → Java 对比 → Demo → Lab → 自动验收 → 错误复盘 → 迁移问题”推进。Bonus 不计入核心 310 分钟，不能挤压 Starter 实操、综合 Lab 或作业启动。

## Block README 模板

每个 Block 的 `README.md` 必须按以下顺序包含标题：

1. `学习结果`
2. `时间盒`
3. `前置知识`
4. `Java 对比`
5. `讲师 Demo`
6. `学员任务`
7. `验收命令`
8. `常见错误`
9. `三级提示`
10. `复盘问题`
11. `Bonus`

学习结果写成可观察行为；验收命令必须能从仓库约定位置直接复制执行。三级提示应逐步缩小问题，但不得粘贴 Solution。

## Demo、Lab、Homework 职责

- Demo 只展示本 Block 的最小核心概念，保持可运行；复杂模式、深层陷阱和后续周内容进入 Bonus。
- Lab 在课堂内完成，使用与 Demo 不同但连续的领域问题；每个阶段都有行为检查点，课堂 Solution 仅供教师验收。
- Integrated Lab 组合当天能力，但不提供 Homework 的直接答案。
- Homework 要求学员跨领域独立迁移，默认按一周里程碑完成；课堂最后只演示目标行为、Starter RED、验收和提交流程。

Week 1 Homework 只能使用 Go 标准库，并兼容 Gitee 官方 `build@golang` 的 Go 1.16 环境，不得要求 Generics。JSON 持久化、第三方 CLI/测试/评分依赖、复杂接口注入、隐藏测试、反作弊或相似度检测均不属于 Week 1 必做范围。

## Starter、公开测试与 Solution 分离

课堂 Starter 必须可编译，公开练习测试应在未实现时因明确的行为断言失败，并在学员完成后变绿；不能以编译失败、缺依赖或权限错误充当 RED。根 Module 的默认测试不运行故意失败的 Starter，Starter 测试通过显式 build tag 或独立 Module opt in。

教师仓库可以保存 Demo、Starter、公开测试、Solution、测验答案和讲师材料。Solution、答案与教师评分资料不得进入学员仓库；学员说明也不得泄露实现。每个行为修改遵循 RED–GREEN–REFACTOR，并保留可重复运行的验收证据。

## 本地验收与 Gitee 手动 CI

每份独立作业只有一个权威评分脚本：`scripts/grade.sh`。学员本地 `make grade` 和 Gitee `GradePipeline` 必须调用同一脚本，禁止在 Makefile 或 YAML 中复制另一套评分规则。

脚本依次检查 `gofmt`、`go vet ./...`、`go test ./...` 和 `go build ./...`，任一步失败即返回非零退出码，并给出当前阶段提示而不暴露答案。Gitee 流水线不配置 push 或 Pull Request 自动触发；教师在 Gitee UI 中选择学生报告的开发分支、核对最新提交 SHA 后手动运行，并保存运行结果。

## 教师仓库到学员仓库的披露流程

1. GitHub 教师仓库保存完整课程，是备课与答案来源，不直接作为无答案学员仓库公开。
2. 教师按周确认披露清单，只发布当前允许的 Demo 和作业 `student_pack/` 内容。
3. 复制 `student_pack/` 的内容到学员 Gitee 仓库根目录，包含 `.workflow/`，但不复制外层目录、`teacher/`、Solution、答案或发布检查表。
4. 教师在干净临时目录检查无答案泄漏，运行 Starter 验证预期行为 RED，再验证教师 Solution 为 GREEN。
5. 学员在各自开发分支提交；教师选择该分支手动运行与本地相同的评分脚本。

本标准不授权创建或管理真实 Gitee 学员仓库，也不包含自动批量导出、GitHub/Gitee 同步或平台侧触发器配置。

## Module 完成定义

一个 Module 只有同时满足以下条件才算完成：

- 唯一学员入口和唯一讲师入口能导航全部公开材料，且 Markdown 本地链接有效。
- 目录符合标准，旧入口已迁移且无重复；核心 Demo、Lab Solution、综合 Lab 和教师 Homework Solution 均可运行、测试、Vet 和构建。
- 每个 Block 的 README 结构完整，日程覆盖全部时间段，学员 hands-on 不少于净课堂时间的 50%。
- Starter 和学员作业只在显式调用时行为失败；默认仓库验收保持绿色。
- Homework 可独立复制，Week 1 仅使用标准库；本地与 Gitee 共享唯一 `scripts/grade.sh`，Gitee 由教师选分支后手动运行。
- Entry/Exit Quiz、评分语言、发布检查表和故障排查资料齐全；教师到学员的披露边界清楚。
- 核心时间盒之外的主题明确列为 Bonus 或 out of scope，不为“讲完内容”而挤占实操。

统一 out of scope 包括：借迁移某一 Module 改写其他 Module 源码、创建真实平台仓库、自动跨平台同步、引入第三方评分依赖，以及把文件持久化、隐藏测试或反作弊设为 Week 1 必做项。

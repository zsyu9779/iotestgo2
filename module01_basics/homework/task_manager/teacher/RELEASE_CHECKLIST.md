# Task Manager 发布检查清单

## 一、准备学生包

- [ ] 只把 `student_pack/` 的**内容**复制到学生 Gitee 仓库根目录，包括隐藏目录 `.workflow/`。
- [ ] 发布物中不存在 `teacher/`、答案实现、教师补充测试或评分记录。
- [ ] 根目录包含 `README.md`、`go.mod`、`Makefile`、`cmd/`、`taskmanager/`、`scripts/` 和 `.workflow/`。
- [ ] `go.mod` 为 `module taskmanager` 和 `go 1.16`，无第三方依赖、无 `go.sum`。
- [ ] 公开测试和规定的接口签名未被改动。

可在临时副本根目录检查教师文件未混入：

```sh
test ! -e teacher
find . -type f -print | sort
```

## 二、验证学生起始包

- [ ] `scripts/grade.sh` 保持可执行；Git 文件模式为 `100755`。
- [ ] `go test ./...` 能完成编译，并因第一次 `Add` 返回 `ErrNotImplemented` 而失败。
- [ ] `make grade` 依次通过 gofmt 和 Vet，然后在测试步骤行为性失败，而不是因依赖、权限或编译环境失败。
- [ ] `go build ./...` 可单独构建起始 CLI。
- [ ] 在一个不依赖课程仓库其他文件的临时目录中重复上述检查，证明学生包可独立复制。

## 三、验证教师答案

- [ ] 在 `teacher/solution` 目录运行 `../../student_pack/scripts/grade.sh`。
- [ ] gofmt、Vet、测试、CLI 构建四步全部通过，末行是 `Task Manager 作业验收通过`。
- [ ] 教师答案的公开接口、公共测试和命令语法与学生包一致。
- [ ] `List` 及其他公开方法均返回副本，不暴露内部 `*Task`。

## 四、检查 Gitee 配置

- [ ] `.workflow/GradePipeline.yml` 使用 `build@golang` 和 `golangVersion: '1.16'`。
- [ ] 流水线只调用 `./scripts/grade.sh`，没有复制另一套验收命令。
- [ ] stage 的触发方式为 `manual`，文件没有 push、PR 或 schedule 的 `triggers` 配置。
- [ ] 用测试分支人工运行一次：先选择分支，再点击“运行”，并保存运行链接。

## 五、发布说明

- [ ] 明确告知学生这是一周课后作业，而非课堂综合项目。
- [ ] 明确要求学生在自己的开发分支提交并推送。
- [ ] 明确告知教师必须在 Gitee UI 选择学生分支后手动运行流水线。

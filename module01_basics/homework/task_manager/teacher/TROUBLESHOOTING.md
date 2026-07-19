# Task Manager 验收排障

## `Permission denied: ./scripts/grade.sh`

先检查文件模式：

```sh
ls -l scripts/grade.sh
chmod +x scripts/grade.sh
git update-index --chmod=+x scripts/grade.sh
git commit -m "fix: restore grader permission"
```

流水线虽然会再次执行 `chmod +x`，仓库仍应提交可执行位，保证本地 `make grade` 与 CI 行为一致。

## 流水线运行了错误分支

本流水线不会在 push 或合并请求时自动运行。教师须进入 Gitee 流水线页面，在运行界面先选择学生报告的开发分支，核对最新提交 SHA，再点击“运行”。如果已在错误分支运行，切换到正确分支后重新手动运行，并在评分记录中保存新的运行链接。

## Go 版本过旧或不一致

本作业的最低和目标版本均为 Go 1.16。先运行：

```sh
go version
go env GOMOD
```

确认当前目录使用本作业的 `go.mod`。本地 Go 低于 1.16 时先升级；高版本本地环境也不得引入泛型或更新 `go` 指令。Gitee 配置必须保留 `golangVersion: '1.16'`，代码只能依赖标准库和 Go 1.16 可用能力。

## 本地通过但 CI 失败

先查看日志停在哪个编号步骤，再从学生仓库根目录运行与 CI 完全相同的入口：

```sh
./scripts/grade.sh
```

- `[1/4]` 失败：运行 `gofmt -w ./cmd ./taskmanager` 后重新检查改动。
- `[2/4]` 失败：运行 `go vet ./...`，修复报告的问题，不要跳过 Vet。
- `[3/4]` 失败：运行 `go test ./... -v` 定位行为差异。未完成的起始包在这里返回 `ErrNotImplemented` 属于预期；学生提交则必须通过。
- `[4/4]` 失败：运行 `go build ./...`，检查 CLI 导入路径和命令解析代码。

同时核对流水线选择的分支和提交 SHA、脚本可执行权限、仓库根目录是否就是 `go.mod` 所在目录，以及 `.workflow/GradePipeline.yml` 是否来自当前分支。不要在 YAML 中另写一套检查来掩盖本地失败。

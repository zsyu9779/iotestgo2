# 文件扫描器发布检查清单

1. 只复制 `student_pack/` 内的内容到学员仓库根目录。
2. 确认没有复制 `teacher/`、Solution 或评分说明。
3. 在干净目录运行 `make grade`，确认 Starter 在测试阶段因 `ErrNotImplemented` 失败。
4. 在教师 Solution 目录调用同一 `student_pack/scripts/grade.sh`，确认全部通过。
5. 检查公开测试没有泄露实现，只描述行为契约。

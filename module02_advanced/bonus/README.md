# Module 02 Bonus

Bonus 不计入核心 310 分钟，不得挤占 Starter、综合 Lab 和作业启动时间。

## Embed

运行：

```bash
go run ./module02_advanced/13_embed
```

主题包括静态资源嵌入、`embed.FS`、模板和单二进制部署。

## Generate

运行前需要额外安装 `stringer`、`mockgen` 等工具。默认仓库验收不会执行这些命令：

```bash
go generate -n ./module02_advanced/14_generate
```

`-n` 只预览命令，不生成文件。真实执行前应确认工具版本和生成结果已纳入版本控制。

## 其他扩展

- `08_os_interaction`：OS 命令、信号和平台差异；
- `09_file_io`：文件读写、缓冲和 Seek；
- `11_runtime_control`：调度和 runtime；
- `12_stdlib_utils`：正则、JSON、编码、哈希和时间；
- `07_testing`：pprof、DeepEqual 和性能暗角。

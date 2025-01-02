# WRollup


## ✨项目简介

WRollup 是一个用于管理 Elasticsearch Rollup 作业的命令行工具。与标准的 Rollup 工具相比，WRollup 提供了监视模式和一些插件功能，旨在帮助用户高效地处理和分析大规模数据。

## 🏗️ 功能

- **创建 Rollup 作业**：支持从原始索引创建 Rollup 作业，并配置聚合和分组。
- **删除 Rollup 作业**：可以删除指定的 Rollup 作业。
- **查询 Rollup 作业**：获取所有 Rollup 作业的信息或特定作业的详细信息。
- **清理旧数据**：删除指定索引中超过指定时间的数据，支持灵活的时间范围配置。
- **监视模式**：实时监控 Rollup 作业的执行状态和性能指标。

## 📦 安装

1. 确保你已经安装了 Go 语言环境（版本 1.16 及以上）。
2. 克隆项目到本地：

   ```bash
   git clone https://github.com/yourusername/wrollup.git
   cd wrollup
   ```

3. 安装依赖：

   ```bash
   go mod tidy
   ```

4. 编译项目：

   ```bash
   go build -o wrollup
   ```

## 使用说明

### 🚀启动工具

运行以下命令启动工具：

```bash
./wrollup
```

### 创建 Rollup 作业

使用以下命令创建 Rollup 作业：

```bash
./wrollup create --job <job_name> --indice <indice_name> --config <config_file>
```

### 删除 Rollup 作业

使用以下命令删除 Rollup 作业：

```bash
./wrollup delete --job <job_name>
```

### 查询 Rollup 作业

使用以下命令查询 Rollup 作业信息：

```bash
./wrollup get --job <job_name>  # 查询特定作业
./wrollup get                   # 查询所有作业
```

### 清理旧数据

使用以下命令清理旧数据：

```bash
./wrollup clean --indice <indice_name> --duration <duration>
```

- `--duration` 参数支持格式如 `1h`, `1d`, `3M`, `1y` 等。


## ⚙️配置

在项目根目录下创建配置文件 `vda.conf`，内容示例：

```plaintext
elasticsearchServiceUrl="http://localhost:9200"
```

## 🗺️开发计划
- [ ] 自适应各种 indice, 自动生成 tmp 文件
- [ ] 增加性能监控能力

##  🤝贡献

欢迎任何形式的贡献！请提交问题、建议或拉取请求。我们鼓励社区参与，以帮助改进这个项目。

## 📄许可证

该项目采用 MIT 许可证，详细信息请查看 [LICENSE](LICENSE) 文件。
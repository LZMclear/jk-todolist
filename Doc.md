# jk-todolist 项目说明文档

下面是针对本仓库（jk-todolist）的完整说明文件，包含技术选型、目录结构、运行与测试说明、接口文档、前端实现细节、已知问题与改进建议等。

## 1. 项目概述
这是一个用 Go 语言实现的简单 TODO 列表 Web 应用：后端使用 Gin 提供 HTTP API，持久化使用 MySQL（通过 database/sql 驱动），前端使用一个轻量静态页面（在 `web/index.html`，基于 Bootstrap）与后端 API 交互。

目标是提供一个干净、可扩展的示例工程：支持创建/查询/更新/删除任务（CRUD）、任务截止时间高亮、按关键字/分类过滤、简单排序等功能。

## 2. 技术选型
- 编程语言：Go（语法简单，高效并发，内置垃圾回收）。
- Web 框架：Gin（轻量且性能好，易上手）。
- 数据库：MySQL（通过 `github.com/go-sql-driver/mysql` 驱动）。
- 环境配置：`github.com/joho/godotenv`（用于加载 `.env` 文件的可选配置）。
- 前端：静态 HTML + 原生 JS + Bootstrap（方便查看、无需复杂构建流程）。

选用理由：项目以后端为主，Go + Gin 能快速搭建 REST API，使用 MySQL 保持一般生产环境的兼容性；前端保持极简，便于学习与演示。

## 3. 项目结构
- `cmd/todo/main.go`：程序入口，加载配置，初始化 DB，启动 HTTP 服务。
- `internal/server/router.go`：Gin 路由注册，静态文件和 API 路由。
- `internal/handler/task.go`：HTTP handler（请求校验、时间解析、调用 store 层）。
- `internal/store/sql.go`：与数据库交互（建表、CRUD 操作）。
- `internal/model/task.go`：Task 数据结构定义（用于 JSON 与 DB 映射）。
- `web/index.html`：前端静态页面（使用 Fetch API 与后端交互）。
- 其它：`go.mod`/`go.sum` 等 Go 模块文件。

## 4. 需求细节与决策
- 任务标题（title）为必填字段，不能为空。 描述（description）和分类（category）可选。
- 任务完成状态（completed）为布尔值，前端通过复选框控制。
- 截止时间（due_date）可选，前端使用 `datetime-local` 输入，后端以 UTC 存储。
- 任务列表默认按创建时间排序，前端可按截止时间排序。
- 任务过滤支持按分类和关键字（标题/描述）搜索。
- 已完成任务在前端显示为灰色，并可通过复选框切换状态。
- 截止时间高亮：已过期显示红色，1 小时内到期显示黄色。

## 5. AI使用说明
- 使用工具：Copilot（GPT5 mini）
- 使用环节：
  - 代码片段生成：如 CRUD 操作的 SQL 语句、Gin 路由注册等。
  - 文档初稿编写：生成本说明文档的结构和部分内容。
- 修改说明：AI 提供的代码片段经过手动调整以符合项目需求和最佳实践，例如调整错误处理逻辑、调整第三方库版本等。

## 6. 运行与测试方式
### 先决条件
- Go 1.20++
- MySQL 数据库
### 本地运行步骤
1. 克隆仓库并进入目录：
   ```bash
    git clone git@github.com:LZMclear/jk-todolist.git
    cd jk-todolist
    ```
2. 创建 `.env` 文件，内容示例：
   ```env
    MYSQL_DSN="user:password@tcp(127.0.0.1:3306)/dbname?parseTime=true"
    PORT=8080
   ```
3. 运行服务：
   ```bash
   go run ./cmd/todo
   ```
4. 访问：http://localhost:8080/

### 使用docker构建容器部署服务
1. 构建镜像
    ```bash
    docker build -t jk-todolist:latest .
    ```
2. 通过compose启动
    ```bash
    docker compose up -d
    ```

### 已测试环境
- Go 1.23.0，MySQL 8.0，Windows 10

## 7. 总结与反思
- 改进建议：
  - 增加用户认证功能，支持多用户任务管理。
  - 增加任务优先级字段，支持更复杂的排序。
  - 使用前端框架（如 React/Vue）重构前端页面，提升用户体验。
- 亮点：项目结构清晰，易于扩展和维护，适合作为学习 Go Web 开发的入门示例。


## 8. 数据模型
Task（tasks 表）字段：
- id (BIGINT, 自增)
- title (VARCHAR(255), 必填)
- description (TEXT)
- category (VARCHAR(100))
- completed (TINYINT(1))
- due_date (DATETIME, 可空)
- created_at (DATETIME)
- updated_at (DATETIME)

注意：仓库内的 `store.InitDB` 会在启动时自动创建 `tasks` 表（如果不存在）。

## 9. API 文档（HTTP）
基地址：/api/tasks/

1. 列表
- 方法：GET
- 路径：/api/tasks/
- 返回：JSON 数组，元素为 Task 对象。

2. 创建任务
- 方法：POST
- 路径：/api/tasks/
- 请求 JSON：
  - title (string, required)
  - description (string, optional)
  - category (string, optional)
  - due_date (string, optional) — 前端使用 `datetime-local` 值，如 `2025-11-25T14:30`（后端以本地时区解析为 UTC 存储）
- 返回：201 Created + 创建的 Task JSON

3. 获取单个任务
- 方法：GET
- 路径：/api/tasks/:id
- 返回：200 + Task JSON 或 404

4. 更新任务
- 方法：PUT
- 路径：/api/tasks/:id
- 请求 JSON：同创建，但可包含 `completed` (bool)
- 返回：200 + 更新后的 Task JSON

5. 删除任务
- 方法：DELETE
- 路径：/api/tasks/:id
- 返回：204 No Content

注意：以上接口行为由 `internal/handler/task.go` 实现，时间字段在后端统一以 UTC 存储，前端在展示时会按本地时间格式化。

## 10. 前端说明（`web/index.html`）
- 单文件静态页面，使用原生 JS + Fetch 调用后端 API。
- 提供新增、编辑、删除、标记完成、按分类/关键字搜索与排序功能。
- 截止时间高亮：
  - overdue（已过期）显示红色背景
  - soon（1 小时内到期）显示黄色背景
  - normal（其他）无特殊背景
- 时间格式化：前端显示使用 `YYYY-MM-DD HH:MM`，并与 `datetime-local` 输入互转（后端期望 `YYYY-MM-DDTHH:MM` 格式，无时区信息）。





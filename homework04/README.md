# Homework04 博客后端作业

本项目按 `homework04.md` 的 6 个题目完成实现，使用 Gin + GORM + SQLite 提供博客后端能力：
- 用户注册与登录（JWT 认证）
- 文章 CRUD（仅作者可更新/删除）
- 评论创建与查询
- 统一错误响应与请求日志

## 目录说明
- `01-project-init` ~ `06-error-logging`：分题实现与测试
- `internal/blog`：核心服务实现
- `internal/testkit`：接口测试辅助工具
- `cmd/blogserver`：服务启动入口
- `TESTING.md`：接口测试用例与测试结果汇总
- `test-results/go-test.txt`：完整测试执行结果

## 运行环境
- Go `1.21+`（与 `go.mod` 一致）
- 支持 Windows / Linux / macOS

## 依赖安装
```bash
go mod tidy
```

## 启动方式
默认会在当前目录创建并使用 `blog.db`（SQLite）。

PowerShell:
```powershell
$env:BLOG_DSN="blog.db"
$env:BLOG_JWT_SECRET="homework04-dev-secret"
$env:BLOG_ADDR=":8080"
go run ./cmd/blogserver
```

Bash:
```bash
export BLOG_DSN="blog.db"
export BLOG_JWT_SECRET="homework04-dev-secret"
export BLOG_ADDR=":8080"
go run ./cmd/blogserver
```

启动后可访问：
- `POST /register`
- `POST /login`
- `GET /posts`
- `GET /posts/:id`
- `GET /posts/:id/comments`
- `POST /posts`（需要 Bearer Token）
- `PUT /posts/:id`（需要 Bearer Token）
- `DELETE /posts/:id`（需要 Bearer Token）
- `POST /posts/:id/comments`（需要 Bearer Token）

## 测试
一次性运行全部测试：
```bash
go test ./... -v
```

分题测试命令：
```bash
go test ./01-project-init -v
go test ./02-models -v
go test ./03-auth -v
go test ./04-posts -v
go test ./05-comments -v
go test ./06-error-logging -v
```

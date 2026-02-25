# Homework04 实现说明

已按 `homework04.md` 的 6 个题目完成目录化实现：

- `01-project-init`
- `02-models`
- `03-auth`
- `04-posts`
- `05-comments`
- `06-error-logging`

核心复用代码位于：
- `internal/blog`：Gin + GORM 博客服务（用户认证、文章 CRUD、评论、统一错误和日志）
- `internal/testkit`：测试辅助工具

## 一次性运行全部测试
```bash
cd /mnt/hgfs/GL/lesson-04/homework04
go test ./... -v
```

## 分题测试命令
```bash
go test ./01-project-init -v
go test ./02-models -v
go test ./03-auth -v
go test ./04-posts -v
go test ./05-comments -v
go test ./06-error-logging -v
```

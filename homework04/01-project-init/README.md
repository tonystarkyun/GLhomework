# 题目1：项目初始化

## 编码答案
- `solution.go`：实现 `BuildApp(dsn)`，完成 Gin + GORM + SQLite 启动和自动迁移。

## 测试用例
1. 应用初始化成功并完成 `users/posts/comments` 三张表迁移。
2. 基础路由已挂载（`/register` 存在并能返回参数校验错误，而不是 404）。

## 测试命令
```bash
go test ./01-project-init -v
```

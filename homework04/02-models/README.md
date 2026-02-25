# 题目2：数据库设计与模型定义

## 编码答案
- `solution.go`：定义模型落库流程 `SeedCoreModels`，验证 `users -> posts -> comments` 的外键关联。
- 真实模型定义在 `internal/blog/models.go`。

## 测试用例
1. 创建用户、文章、评论并校验关联字段正确。
2. 使用 `Preload` 查询评论并携带关联用户和文章。
3. 校验 `username/email` 唯一约束生效。

## 测试命令
```bash
go test ./02-models -v
```

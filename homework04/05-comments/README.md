# 题目5：评论功能

## 编码答案
- `solution.go`：实现评论创建与查询接口调用封装。
- 服务端实现位于 `internal/blog/handlers.go`：
  - `POST /posts/:id/comments`（需要 JWT）
  - `GET /posts/:id/comments`

## 测试用例
1. 已认证用户可以对文章发表评论。
2. 可以查询某篇文章的评论列表。
3. 未认证发表评论返回 401。
4. 对不存在文章评论返回 404。

## 测试命令
```bash
go test ./05-comments -v
```

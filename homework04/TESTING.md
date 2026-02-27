# 接口测试用例与结果

本作业使用 `go test` + `net/http/httptest` 对 HTTP 接口进行自动化测试（符合“Postman 或其他工具”中的“其他工具”）。

## 测试用例
- 题目1（项目初始化）
  - 应用初始化成功并完成 `users/posts/comments` 表迁移
  - 基础路由存在（`/register` 非 404）
- 题目2（模型）
  - 用户、文章、评论创建成功
  - `Preload` 关联查询成功
  - `username/email` 唯一约束生效
- 题目3（认证）
  - 注册密码为 bcrypt 哈希
  - 登录返回 JWT
  - 无 token 访问受保护接口返回 401
  - 有 token 可创建文章
- 题目4（文章）
  - 创建/列表/详情/更新/删除链路可用
  - 非作者更新/删除返回 403
  - 删除后再查询返回 404
- 题目5（评论）
  - 已认证用户可发表评论
  - 查询文章评论列表成功
  - 未认证发表评论返回 401
  - 对不存在文章评论返回 404
- 题目6（错误与日志）
  - 错误返回统一 JSON 结构
  - 非法 token 返回 401 并记录错误日志
  - 资源不存在返回 404 并记录错误日志
  - 正常请求有访问日志

## 测试命令
```bash
go test ./... -v
```

## 测试结果
- 执行时间：2026-02-27
- 结果：全部通过
- 完整日志：`test-results/go-test.txt`

摘要：
```text
ok  	homework04/01-project-init
ok  	homework04/02-models
ok  	homework04/03-auth
ok  	homework04/04-posts
ok  	homework04/05-comments
ok  	homework04/06-error-logging
```

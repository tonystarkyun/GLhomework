# 题目6：错误处理与日志记录

## 编码答案
- `solution.go`：定义统一错误响应结构 `{"error": "..."}` 的解析。
- 服务端统一错误与日志在 `internal/blog/app.go`：
  - `abortError` 统一 HTTP 错误码和 JSON 错误体
  - 请求日志中间件记录 method/path/status/latency

## 测试用例
1. 参数错误返回统一 JSON 结构。
2. 非法 token 返回 401，并写入错误日志。
3. 资源不存在返回 404，并写入错误日志。
4. 普通请求日志会被记录。

## 测试命令
```bash
go test ./06-error-logging -v
```

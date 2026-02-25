# 题目3：用户认证与授权

## 编码答案
- `solution.go`：提供注册、登录、携带 JWT 调用受保护接口的请求封装。
- 认证核心实现位于 `internal/blog`：
  - 密码 bcrypt 加密
  - JWT 生成与校验
  - `Authorization: Bearer <token>` 中间件

## 测试用例
1. 注册成功后数据库保存的是 bcrypt 哈希密码。
2. 登录成功返回 JWT。
3. 无 token 访问受保护接口返回 401；携带 token 可成功创建文章。
4. 密码错误登录返回 401。

## 测试命令
```bash
go test ./03-auth -v
```

# homework02 题目答案

每个题目单独放在一个目录中，目录里包含：
- `solution.go`：题目实现
- `solution_test.go`：测试用例

## 目录说明

- `pointer_q1`：指针题 1
- `pointer_q2`：指针题 2
- `goroutine_q1`：Goroutine 题 1
- `goroutine_q2`：Goroutine 题 2
- `oop_q1`：面向对象题 1
- `oop_q2`：面向对象题 2
- `channel_q1`：Channel 题 1
- `channel_q2`：Channel 题 2
- `lock_q1`：锁机制题 1（Mutex）
- `lock_q2`：锁机制题 2（Atomic）

## 测试命令

在 `homework02/` 目录下执行：

```bash
go test ./pointer_q1 -v
go test ./pointer_q2 -v
go test ./goroutine_q1 -v
go test ./goroutine_q2 -v
go test ./oop_q1 -v
go test ./oop_q2 -v
go test ./channel_q1 -v
go test ./channel_q2 -v
go test ./lock_q1 -v
go test ./lock_q2 -v
```

全量测试：

```bash
go test ./... -v
```

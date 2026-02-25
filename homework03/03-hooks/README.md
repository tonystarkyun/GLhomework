# 题目3：钩子函数

## 实现内容
- 为 `Post` 模型实现 `AfterCreate`：文章创建成功后，用户 `post_count + 1`。
- 为 `Comment` 模型实现 `AfterDelete`：删除评论后若该文章评论数为 0，更新 `comment_status = "无评论"`。
- 补充 `Comment.AfterCreate`，使文章在有评论时状态为 `"有评论"`，便于状态流转测试。

## 目录文件
- `models.go`: 模型定义与钩子实现。
- `main.go`: 可运行演示。
- `hooks_test.go`: 钩子行为测试。

## 测试用例
1. `TestPostAfterCreateUpdatesUserPostCount`
- 连续创建两篇文章，验证用户 `PostCount` 自动变为 2。

2. `TestCommentAfterDeleteUpdatesCommentStatus`
- 创建两条评论后状态应为 `有评论`。
- 删除一条评论后仍为 `有评论`。
- 删除最后一条评论后变为 `无评论`。

## 测试命令
```bash
cd /mnt/hgfs/GL/lesson-03/homework03/03-hooks
GOTOOLCHAIN=go1.24.11 go test -v ./...
```

## 运行命令
```bash
cd /mnt/hgfs/GL/lesson-03/homework03/03-hooks
GOTOOLCHAIN=go1.24.11 go run .
```

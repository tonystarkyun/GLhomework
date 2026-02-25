# 题目2：关联查询

## 实现内容
- 基于博客模型实现两个查询函数：
  - `GetUserPostsWithComments`: 查询指定用户的所有文章及评论。
  - `GetMostCommentedPost`: 查询评论数最多的文章。

## 目录文件
- `models.go`: 模型定义和迁移。
- `query.go`: 题目要求的查询实现。
- `main.go`: 可运行示例（自动插入演示数据并查询）。
- `query_test.go`: 测试用例。

## 测试用例
1. `TestGetUserPostsWithComments`
- 验证仅返回目标用户的文章，并正确预加载评论。

2. `TestGetMostCommentedPost`
- 验证返回评论数量最多的文章及评论数。

3. `TestGetMostCommentedPostWhenNoComments`
- 验证在没有任何评论时返回 `nil`。

## 测试命令
```bash
cd /mnt/hgfs/GL/lesson-03/homework03/02-association-query
GOTOOLCHAIN=go1.24.11 go test -v ./...
```

## 运行命令
```bash
cd /mnt/hgfs/GL/lesson-03/homework03/02-association-query
GOTOOLCHAIN=go1.24.11 go run .
```

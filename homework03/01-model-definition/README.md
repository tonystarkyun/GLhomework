# 题目1：模型定义

## 实现内容
- 使用 GORM 定义 `User`、`Post`、`Comment` 三个模型。
- 建立一对多关系：`User -> Posts`、`Post -> Comments`。
- 通过 `AutoMigrate` 自动创建数据表。

## 目录文件
- `models.go`: 模型定义与 `AutoMigrateBlog`。
- `main.go`: 建表入口程序。
- `models_test.go`: 自动化测试。

## 测试用例
1. `TestAutoMigrateCreatesTables`
- 验证 `users`、`posts`、`comments` 三张表被成功创建。

2. `TestRelationsWorkWithPreload`
- 插入用户、文章、评论数据。
- 使用 `Preload("Posts.Comments")` 验证关系查询正确。

## 测试命令
```bash
cd /mnt/hgfs/GL/lesson-03/homework03/01-model-definition
GOTOOLCHAIN=go1.24.11 go test -v ./...
```

## 运行命令
```bash
cd /mnt/hgfs/GL/lesson-03/homework03/01-model-definition
GOTOOLCHAIN=go1.24.11 go run .
```

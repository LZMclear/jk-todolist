# jk-todolist
小宇宙校招笔试 —— 基于 Go + Gin + database/sql 的 TODO List 示例项目

# 提交步骤
## feat：项目初始化
- 初始化Go mod
- 创建基础项目结构（handler，model，server，store）
- 添加Gin框架依赖
- 使用ping测试Gin是否成功运行

## feat：定义数据类型
- 初始化数据库表结构
- 定义Task项的数据结构

## feat：定义API接口
- 定义增删改查Task的接口
- 添加静态文件路由

## feat：完善后端API核心功能
- 实现获取Task列表接口
- 实现创建Task接口
- 实现更新Task接口
- 实现删除Task接口
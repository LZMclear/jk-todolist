# jk-todolist
小宇宙校招笔试 —— 基于 Go + Gin + database/sql 的 TODO List 示例项目

## 本地运行（开发）

### 先决条件：Go 1.20+、MySQL。

1. 克隆仓库并进入目录
2. 在仓库根目录创建 `.env` 文件，格式：
```env
MYSQL_DSN="user:password@tcp(127.0.0.1:3306)/dbname?parseTime=true"
PORT=8080
```
### 运行服务

#### 开发模式（不编译二进制）
go run ./cmd/todo

#### 编译并运行
go build -o bin/todo ./cmd/todo ; .\bin\todo

### 使用docker构建容器部署服务

1. 构建镜像
    ```bash
    docker build -t jk-todolist:latest .
    ```
2. 通过compose启动
    ```bash
    docker compose up -d
    ```
   
启动后访问：http://localhost:8080/

注意：`store.InitDB` 会在启动时自动创建 `tasks` 表（如果数据库用户有足够权限，并且有 todolist 数据库）。

## 本地运行界面
![运行界面](https://blog-1316762285.cos.ap-beijing.myqcloud.com//todolist1.png)

## 服务器部署运行界面
可以通过访问 http://43.138.33.205:8083/ 查看运行效果
![服务器部署运行界面](https://blog-1316762285.cos.ap-beijing.myqcloud.com//20251126194510.png)
# fx-gin

一个基于 Go 语言的 Web 应用框架，结合了 Gin 和 Uber FX 依赖注入框架，提供了清晰的项目结构和强大的功能。

## 项目描述

fx-gin 是一个用 Go 语言编写的 Web 应用框架，它:

- 使用 Gin 作为 HTTP 路由引擎，提供高性能的 HTTP 请求处理
- 集成 Uber FX 作为依赖注入框架，实现松耦合的组件管理
- 提供结构化的日志记录（基于 zap）
- 支持环境变量配置（使用 godotenv 和 env）
- 包含 Swagger 文档支持
- 使用 SQLite 作为数据存储

## 项目结构

```
├── cmd/                    # 命令行入口点
│   ├── server/             # HTTP 服务器
│   └── swagger/            # Swagger 文档生成
├── internal/               # 内部应用代码
│   ├── config/             # 应用配置
│   ├── domain/             # 领域模型和业务规则
│   ├── infra/              # 基础设施代码
│   │   └── db/             # 数据库连接和管理
│   ├── repo/               # 数据访问层
│   ├── service/            # 业务逻辑服务
│   └── transport/          # 传输层
│       └── http/           # HTTP 相关代码
│           ├── handler/    # HTTP 处理器
│           └── middleware/ # HTTP 中间件
├── pkg/                    # 公共包
│   └── logger/             # 日志包
├── scripts/                # 脚本文件
├── test/                   # 测试相关文件
├── docs/                   # 文档
├── go.mod                  # Go 模块文件
├── go.sum                  # Go 依赖校验文件
├── .env                    # 环境变量文件
└── Makefile                # 构建脚本
```

## 核心功能

- **依赖注入**: 使用 Uber FX 管理应用组件间的依赖关系
- **HTTP 路由**: 基于 Gin 框架提供高性能的 HTTP 路由
- **配置管理**: 支持通过环境变量配置应用
- **结构化日志**: 使用 zap 提供高性能的结构化日志
- **API 文档**: 集成 Swagger 提供 API 文档
- **数据持久化**: 支持 SQLite 数据库存储

## 编码规范

本项目遵循以下编码规范，以确保代码的一致性和可读性：

### 接口和实现

1. **接口命名**: 
   - 接口命名采用 Go 的标准命名约定，不使用 "I" 前缀
   - 例如：使用 `UserRepo` 而不是 `IUserRepo`

2. **实现可见性**:
   - 接口的具体实现应为非导出类型（小写开头）
   - 例如：`userService`, `userRepo`
   - 对于需要在包外部访问的处理器类型，保持类型名称为导出类型（如 `UserHandler`）

3. **依赖注入**:
   - 使用 `fx.In` 包装构造函数参数
   - 参数命名简洁，推荐使用 `p` 作为参数名
   - 例如：`func NewUserRepo(p UserRepoParams) domain.UserRepo`

### 数据传输和模型

1. **请求与响应模型**:
   - 所有请求和响应模型都应在 `domain` 包中定义
   - 请求模型使用 `XxxRequest` 命名格式
   - 响应模型使用 `XxxResponse` 命名格式

2. **数据流转规范**:
   - Handler 层：接收 HTTP 请求并绑定到 `XxxRequest` 模型，返回 `XxxResponse` 模型
   - Service 层：接收 `XxxRequest` 模型作为参数，返回 `XxxResponse` 模型
   - Repository 层：接收和返回领域实体模型（如 `User`, `Profile` 等）

3. **模型转换**:
   - 在合适的层级进行模型转换（如 DTO 到实体，实体到 DTO）
   - 避免在多个层级重复相同的转换逻辑

### 代码组织

1. **分层结构**:
   - `domain`: 定义领域模型、接口、请求和响应模型
   - `repo`: 实现数据访问层
   - `service`: 实现业务逻辑
   - `transport/http/handler`: 实现HTTP处理逻辑

2. **函数命名**:
   - 导出的构造函数使用 `New` 前缀
   - 例如：`NewUserService`, `NewUserRepo`

3. **错误处理**:
   - 使用描述性错误消息
   - 错误应该被记录并正确传播

4. **注释**:
   - 所有导出的函数、类型和方法都应有注释
   - 复杂的非导出函数也应有注释说明

### 扩展性

1. **数据库扩展**:
   - 项目支持多种数据库后端，可通过实现相应的接口进行扩展
   - 默认使用 SQLite，但可以轻松替换为 MySQL、PostgreSQL 等
   - 数据访问层与业务逻辑解耦，便于切换数据存储方式

2. **缓存扩展**:
   - 项目设计支持添加 Redis 等缓存服务
   - 通过实现缓存接口可以无缝集成到现有系统中

3. **微服务扩展**:
   - 项目结构支持从单体应用扩展到微服务架构
   - 领域模型和业务逻辑可以方便地迁移到独立服务中

### 其他规范

1. **依赖管理**:
   - 在模块初始化函数 `init()` 中注册组件
   - 明确声明和管理依赖关系

2. **数据库操作**:
   - 使用参数化查询防止SQL注入
   - 确保正确关闭数据库资源

3. **日志记录**:
   - 使用结构化日志
   - 记录合适的日志级别（错误、警告、信息等）

## API文档与测试

系统提供了完整的API文档和测试指南：

1. **Swagger API文档**:
   - 访问 http://localhost:38080/swagger/index.html 查看交互式API文档

2. **API测试指南**:
   - 查看 [docs/api.md](docs/api.md) 获取详细的API测试命令和使用示例

## 快速开始

### 前置条件

- Go 1.24+
- SQLite3 (可选)

### 环境配置

1. 复制环境变量示例文件:
   ```bash
   cp .env.example .env
   ```

2. 根据需要修改 `.env` 文件中的配置:
   ```
   APP_HOST=localhost
   APP_PORT=38080
   APP_NAME=your-app-name
   APP_ENV=dev
   DATABASE_DATABASE=your-db-file.db
   ```

### 运行应用

```bash
# 使用 Makefile 构建并运行
make run

# 或直接使用 Go 命令
go run cmd/server/main.go
```

服务默认在 http://localhost:38080 启动

### API 文档

启动服务后可以通过浏览器访问 Swagger 文档:
http://localhost:38080/swagger/index.html

## 项目扩展

### 添加新的 API 端点

1. 在 `internal/domain` 中定义领域模型
2. 在 `internal/repo` 中实现数据访问层
3. 在 `internal/service` 中实现业务逻辑
4. 在 `internal/transport/http/handler` 中创建处理器
5. 在 `internal/transport/http/router.go` 中注册路由

### 添加新的配置项

在 `internal/config/config.go` 文件中修改配置结构体，并确保在 `.env` 文件中提供相应的环境变量。

## 数据结构

项目代码采用简洁清晰的分层架构:

- **领域层** (`domain`): 定义核心业务实体和规则
- **仓储层** (`repo`): 处理数据持久化逻辑
- **服务层** (`service`): 实现业务逻辑和用例
- **传输层** (`transport`): 处理 HTTP 请求和响应

## 数据库

本项目使用SQLite作为默认数据库，所有表结构在应用启动时自动创建。

### 数据库迁移

数据库迁移在应用启动时自动执行，定义在`internal/infra/db/migration.go`文件中。主要完成以下工作：

1. **表创建**：自动创建所有必要的数据库表
   - `users`: 用户信息表
   - `profiles`: 用户资料表
   - `roles`: 角色表
   - `user_roles`: 用户与角色的关联表
   - `configs`: 配置表

2. **初始数据**：默认插入基础数据
   - 创建默认角色：`admin`和`user`

### 数据库模型

项目的主要数据库模型包括：

1. **用户模型**：存储用户基本信息
2. **用户资料**：存储用户详细资料
3. **角色模型**：定义系统角色
4. **配置模型**：系统配置信息

### 手动执行迁移

系统会在启动时自动执行迁移，但如果需要手动触发迁移（例如在开发环境中），可以直接启动应用：

```bash
make run
```

或者：

```bash
go run cmd/server/main.go
```

## 贡献

欢迎提交 Issue 和 Pull Request 来改进项目。

## 许可证

[添加您的许可证信息]
# fx-gin

一个基于 Gin 和 Uber FX 构建的企业级 Go Web 项目脚手架，提供完整的项目结构和最佳实践，助力开发者快速构建高质量应用。

## ✨ 特性

- 🏗️ 清晰的分层架构，确保代码可维护性
- 🔄 基于 Uber FX 的依赖注入，实现松耦合设计
- 📚 内置 Swagger 文档支持
- 📝 标准化的日志和配置管理
- 💾 默认使用 SQLite，支持扩展其他数据库
- 🚀 支持 AI 开发工具（Cursor、GitHub Copilot）

## 🚀 快速开始

### 1. 克隆项目
```bash
git clone [项目地址]
cd fx-gin
```

### 2. 配置环境
```bash
cp .env.example .env
# 编辑 .env 文件配置必要的环境变量
```

### 3. 运行项目
```bash
make run
# 或
go run cmd/server/main.go
```

## 📁 项目结构

```
├── cmd/                    # 命令行入口
│   ├── server/            # HTTP 服务器入口
│   └── swagger/           # Swagger 文档生成
├── internal/              # 内部应用代码
│   ├── config/           # 应用配置
│   ├── domain/           # 领域模型和业务规则
│   ├── infra/            # 基础设施代码
│   │   └── db/           # 数据库连接和管理
│   ├── repo/             # 数据访问层
│   ├── service/          # 业务逻辑服务
│   └── transport/        # 传输层
│       └── http/         # HTTP 相关代码
│           ├── handler/  # HTTP 处理器
│           └── middleware/ # HTTP 中间件
├── pkg/                   # 公共包
│   └── logger/           # 日志包
├── scripts/              # 脚本文件
├── test/                 # 测试相关文件
├── docs/                 # 文档
└── Makefile              # 构建脚本
```

## 🛠️ 核心功能

### 1. 项目架构
- 采用领域驱动设计（DDD）原则
- 清晰的分层架构设计
- 模块化的代码组织
- 易于扩展的项目结构

### 2. 数据库支持
- 默认使用 SQLite，快速开始
- 支持扩展其他数据库（MySQL、PostgreSQL 等）
- 自动数据库迁移
- 连接池管理

### 3. API 开发
- 遵循 RESTful API 规范
- 内置请求参数验证(TODO)
- 统一的错误处理机制(TODO)
- 自动生成 Swagger 文档

### 4. 开发体验
- 完整的开发环境配置
- 支持主流 AI 开发工具
- 详细的开发文档
- 丰富的示例代码

## 📚 最佳实践

### 1. 代码规范
- 遵循 Go 语言最佳实践
- 使用依赖注入管理组件
- 保持模块间的低耦合
- 编写完整的单元测试

### 2. 性能优化
- 数据库连接池管理
- 多级缓存机制
- 查询性能优化
- 中间件性能调优

### 3. 安全防护
- 请求参数验证和清理
- SQL 注入防护
- XSS 攻击防护
- CSRF 防护
- 敏感数据加密

### 4. 部署方案
- Docker 容器化支持
- CI/CD 流程配置
- 健康检查机制
- 监控和告警系统

## ❓ 常见问题

### 1. 如何添加新的数据库支持？
- 实现 `internal/infra/db/driver.go` 中的接口
- 添加相应的数据库配置
- 更新数据库迁移脚本

### 2. 如何添加新的中间件？
- 在 `internal/transport/http/middleware` 中创建新文件
- 实现中间件逻辑
- 在路由配置中注册中间件

### 3. 如何实现用户认证？
- 在 `internal/service` 中实现认证服务
- 添加认证中间件
- 实现用户会话管理

## 🤝 贡献指南

欢迎提交 Issue 和 Pull Request 来改进项目。在提交代码前，请确保：

1. 代码符合项目规范
2. 添加必要的测试用例
3. 更新相关文档
4. 提交信息清晰明确

## �� 许可证

MIT License 
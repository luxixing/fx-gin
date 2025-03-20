# API 测试指南

本文档提供了使用curl命令测试主要API接口的示例。

## 用户相关接口

### 注册用户

```bash
curl -X POST "http://localhost:38080/api/v1/users/register" \
     -H "Content-Type: application/json" \
     -d '{
        "username": "testuser",
        "email": "test@example.com",
        "password": "password123"
     }'
```

### 用户登录

```bash
curl -X POST "http://localhost:38080/api/v1/users/login" \
     -H "Content-Type: application/json" \
     -d '{
        "username": "testuser",
        "password": "password123"
     }'
```

### 获取用户信息

```bash
curl -X GET "http://localhost:38080/api/v1/users/1" \
     -H "Content-Type: application/json"
```

### 获取用户档案

```bash
curl -X GET "http://localhost:38080/api/v1/users/1/profile" \
     -H "Content-Type: application/json"
```

### 获取用户角色

```bash
curl -X GET "http://localhost:38080/api/v1/users/1/roles" \
     -H "Content-Type: application/json"
```

### 更新用户信息

```bash
curl -X PUT "http://localhost:38080/api/v1/users/1" \
     -H "Content-Type: application/json" \
     -d '{
        "username": "updateduser",
        "email": "updated@example.com"
     }'
```

### 获取用户列表

```bash
curl -X GET "http://localhost:38080/api/v1/users?page=1&size=10" \
     -H "Content-Type: application/json"
```

## 配置相关接口

### 创建配置

```bash
curl -X POST "http://localhost:38080/api/v1/config" \
     -H "Content-Type: application/json" \
     -d '{
        "key": "app_theme",
        "value": "dark"
     }'
```

### 获取配置

```bash
curl -X GET "http://localhost:38080/api/v1/config/app_theme" \
     -H "Content-Type: application/json"
```

## API流程示例

下面是一个完整的API使用流程示例，从用户注册到数据操作：

1. **注册新用户**
```bash
curl -X POST "http://localhost:38080/api/v1/users/register" \
     -H "Content-Type: application/json" \
     -d '{
        "username": "johnsmith",
        "email": "john@example.com",
        "password": "securepass123"
     }'
```

2. **用户登录获取令牌**
```bash
curl -X POST "http://localhost:38080/api/v1/users/login" \
     -H "Content-Type: application/json" \
     -d '{
        "username": "johnsmith",
        "password": "securepass123"
     }'
```

3. **查看个人资料**
```bash
# 假设用户ID为1
curl -X GET "http://localhost:38080/api/v1/users/1/profile" \
     -H "Content-Type: application/json"
```

4. **更新用户信息**
```bash
curl -X PUT "http://localhost:38080/api/v1/users/1" \
     -H "Content-Type: application/json" \
     -d '{
        "username": "johnsmith",
        "email": "john.updated@example.com"
     }'
``` 
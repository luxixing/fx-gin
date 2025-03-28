# API Testing Guide

This document provides examples of testing main API endpoints using curl commands.

## User Related Endpoints

### Register User

```bash
curl -X POST "http://localhost:38080/api/v1/users/register" \
     -H "Content-Type: application/json" \
     -d '{
        "username": "testuser22",
        "email": "test@example.com",
        "password": "password123"
     }'
```

### User Login

```bash
curl -X POST "http://localhost:38080/api/v1/users/login" \
     -H "Content-Type: application/json" \
     -d '{
        "username": "testuser",
        "password": "password123"
     }'
```

### Get User Information

```bash
curl -X GET "http://localhost:38080/api/v1/users/1" \
     -H "Content-Type: application/json"
```

### Get User Profile

```bash
curl -X GET "http://localhost:38080/api/v1/users/1/profile" \
     -H "Content-Type: application/json"
```

### Get User Roles

```bash
curl -X GET "http://localhost:38080/api/v1/users/1/roles" \
     -H "Content-Type: application/json"
```

### Update User Information

```bash
curl -X PUT "http://localhost:38080/api/v1/users/1" \
     -H "Content-Type: application/json" \
     -d '{
        "username": "updateduser",
        "email": "updated@example.com"
     }'
```

### Get User List

```bash
curl -X GET "http://localhost:38080/api/v1/users?page=1&size=10" \
     -H "Content-Type: application/json"
```
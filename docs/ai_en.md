## AI Development Assistant Guide

This project is an enterprise-level Go Web project scaffold built with Gin and Uber FX. This document provides prompt templates for common development scenarios to help developers use AI-assisted development more efficiently.

### General Development Prompt Templates

#### 1. Adding New Feature Modules

When you need to develop a new feature module, you can use the following prompt:

```
I need to implement a [feature name] module in the fx-gin project. The main functionality of this module is [detailed feature requirements].

Please help me implement this feature following the project's layered architecture (domain, repo, service, handler):

1. Define domain models and interfaces in internal/domain
2. Implement data access layer code in internal/repo
3. Write business logic services in internal/service
4. Create HTTP handlers in internal/transport/http/handler
5. Register related routes in internal/transport/http/router.go

Please ensure:
- Follow DDD design principles
- Use dependency injection for component management
- Add necessary unit tests
- Update Swagger documentation
```

#### 2. Extending Existing Features

When you need to extend existing features:

```
I want to extend the [module name] functionality in the fx-gin project by adding the following capabilities: [detailed description of new features].

Please help me modify the following aspects:
1. Update model and interface definitions in internal/domain
2. Extend existing repository implementations
3. Add new business logic in the service layer
4. Create new API endpoints in the handler
5. Register new routes in router.go

Please ensure:
- Maintain backward compatibility
- Add appropriate parameter validation
- Update relevant documentation
```

#### 3. Fixing Issues or Optimizing Code

When you need to fix issues or optimize code:

```
I found an issue in the [module name] of the fx-gin project: [describe the issue]. The error log or behavior is: [paste error message or describe abnormal behavior].

Please help me:
1. Analyze possible causes
2. Provide a fix
3. Implement code fixes
4. Add necessary test cases
```

### Specific Feature Development Prompts

#### 1. Adding New Data Models and CRUD Operations

```
Please help me implement a new [model name] model and its complete CRUD operations in the fx-gin project.

Requirements:
- Model fields: [list model fields and types]
- Required APIs: Create, Query (single/list), Update, Delete
- Add necessary parameter validation
- Implement relationships: [if any, describe relationships with other models]

Please ensure:
- Define models in internal/domain
- Implement corresponding repository interfaces
- Add database migration scripts
- Update Swagger documentation
```

#### 2. Implementing Authentication and Authorization

```
Please help me implement JWT authentication and role-based access control in the fx-gin project.

Requirements:
1. Implement token generation and validation using JWT standards
2. Create authentication middleware in internal/transport/http/middleware
3. Implement role-based authorization checks
4. Add permission control to existing API endpoints

Please ensure:
- Use secure key management
- Implement token refresh mechanism
- Add necessary logging
```

#### 3. Adding Cache Layer

```
I want to add Redis cache support to the fx-gin project, mainly for caching [specific data type] data.

Please help me:
1. Create cache interface abstraction in internal/infra/cache
2. Implement Redis cache provider
3. Integrate cache logic in [specified service]
4. Add cache invalidation and update strategies

Please ensure:
- Implement graceful degradation
- Add cache warming functionality
- Implement cache monitoring
```

#### 4. Database Migration and Extension

```
I need to migrate the fx-gin project from SQLite to PostgreSQL database.

Please help me:
1. Update database configuration in internal/infra/db
2. Modify existing migration scripts to be compatible with PostgreSQL
3. Adjust query syntax to work with PostgreSQL
4. Provide necessary indexing and optimization suggestions

Please ensure:
- Maintain atomic data migration
- Provide rollback mechanism
- Update connection pool configuration
```

#### 5. Implementing API Documentation Updates

```
Please help me update the Swagger documentation comments for the [module name] in the fx-gin project.

New API endpoints are: [describe new APIs]
Please ensure:
1. Add appropriate comments to generate complete Swagger documentation
2. Include all request parameters, response models, and possible error codes
3. Provide example values
4. Update API descriptions in README.md
```

### Practical Development Tips

- **Provide Context**: Tell the AI which part of the project you're working on and share relevant existing code files.
- **Step-by-Step Requests**: For complex features, request design first, then implementation details.
- **Specify Reference Patterns**: Use phrases like "please follow the implementation pattern of the user module" to maintain consistent code style.
- **Request Code Review**: Ask AI to review your code and provide optimization suggestions.
- **Focus on Performance**: Consider performance impact when implementing new features, add performance tests when necessary.
- **Keep Documentation in Sync**: Ensure documentation is updated when code changes are made.

By using these templates and techniques, you can more effectively leverage the AI assistant to accelerate the development process of your fx-gin project. Remember, the AI assistant is a tool to help improve development efficiency, but final code quality and architectural decisions still require developer expertise. 
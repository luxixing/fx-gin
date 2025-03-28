# fx-gin

A professional Go Web project scaffold based on Gin and Uber FX, providing enterprise-grade project structure and best practices to help developers quickly start high-quality project development.

## Project Overview

fx-gin is a battle-tested Go Web project scaffold, particularly suitable for:

- Developers who need to quickly start new projects
- Teams pursuing code quality and best practices
- Developers using AI development tools (like Cursor, GitHub Copilot)
- Organizations requiring standardized project structure

### Core Advantages

- Enterprise-grade layered architecture ensuring code maintainability
- Uber FX dependency injection for loose coupling design
- Clear project structure and code organization
- Built-in Swagger documentation support
- Standardized logging and configuration management
- SQLite by default with support for other databases

## Project Structure

```
├── cmd/                    # Command line entry points
│   ├── server/            # HTTP server entry
│   └── swagger/           # Swagger documentation generation
├── internal/              # Internal application code
│   ├── config/           # Application configuration
│   ├── domain/           # Domain models and business rules
│   ├── infra/            # Infrastructure code
│   │   └── db/           # Database connection and management
│   ├── repo/             # Data access layer
│   ├── service/          # Business logic services
│   └── transport/        # Transport layer
│       └── http/         # HTTP related code
│           ├── handler/  # HTTP handlers
│           └── middleware/ # HTTP middleware
├── pkg/                   # Public packages
│   └── logger/           # Logging package
├── scripts/              # Script files
├── test/                 # Test related files
├── docs/                 # Documentation
└── Makefile              # Build script
```

## Quick Start

### 1. Clone the Project
```bash
git clone [project-url]
cd fx-gin
```

### 2. Configure Environment
```bash
cp .env.example .env
# Edit .env file with necessary environment variables
```

### 3. Run the Project
```bash
make run
# or
go run cmd/server/main.go
```

## Core Features

### 1. Project Structure
- Clear layered architecture
- Standardized code organization
- Modular design
- Easy to extend structure

### 2. Database Support
- SQLite by default for quick start
- Database interface for extending other databases
- Automatic database migration
- Model relationship management

### 3. API Development Support
- RESTful API specification
- Request parameter validation
- Unified error handling
- Swagger documentation generation

### 4. Development Tool Support
- Support for AI development tools (Cursor, GitHub Copilot)
- Standardized code generation patterns
- Complete testing framework
- Development environment configuration

### 5. AI Development Assistant Guide
The project provides comprehensive AI development assistant guides in both Chinese and English to help developers leverage AI tools more effectively:

- [AI Development Assistant Guide (中文)](docs/ai.md)
- [AI Development Assistant Guide (English)](docs/ai_en.md)

These guides include:
- Templates for common development scenarios
- Best practices for AI-assisted development
- Specific feature implementation guides
- Code review and optimization tips
- Performance considerations
- Documentation maintenance guidelines

Key benefits:
- Standardized prompts for consistent development
- Efficient feature implementation patterns
- Quality assurance guidelines
- Best practices for AI tool usage

## Best Practices

### 1. Code Organization
- Follow Domain-Driven Design principles
- Use dependency injection for component management
- Maintain low coupling between modules
- Write comprehensive unit tests

### 2. Performance Optimization
- Use connection pooling
- Implement caching mechanisms
- Optimize database queries
- Use middleware appropriately

### 3. Security Practices
- Parameter validation and sanitization
- SQL injection prevention
- XSS protection
- CSRF protection

### 4. Deployment Recommendations
- Use Docker for containerization
- Configure CI/CD pipeline
- Implement health checks
- Set up monitoring and alerts

## Common Questions

### 1. How to Add New Database Support?
- Implement the interface in `internal/infra/db/driver.go`
- Add corresponding database configuration
- Update database migration scripts

### 2. How to Add New Middleware?
- Create new file in `internal/transport/http/middleware`
- Implement middleware logic
- Register middleware in route configuration

### 3. How to Implement User Authentication?
- Implement authentication service in `internal/service`
- Add authentication middleware
- Implement user session management

## Contributing

Issues and Pull Requests are welcome to improve the project.

## License

MIT License
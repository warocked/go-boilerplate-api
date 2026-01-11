# Go Boilerplate API

A production-ready Go API boilerplate built with Fiber, GORM, and PostgreSQL. Designed for high performance, scalability, and security.

## Features

- 🚀 **High Performance**: Optimized for handling millions of requests with connection pooling
- 🔐 **Security**: JWT authentication, input validation, security headers, rate limiting
- 🗄️ **Database**: PostgreSQL with GORM ORM, UUID primary keys, connection pooling
- 📦 **Migrations**: Versioned SQL migrations using golang-migrate (industry best practice)
- 🔄 **WebSocket**: Real-time bidirectional communication support
- ⚡ **Redis**: Optional Redis support for caching and sessions
- 🛡️ **Production Ready**: Graceful shutdown, health checks, error handling
- 📝 **Structured Logging**: Zerolog integration with configurable log levels
- 🔧 **Configurable**: Environment-based configuration with sensible defaults

## Tech Stack

- **Framework**: [Fiber v2](https://github.com/gofiber/fiber) - Express-inspired web framework
- **ORM**: [GORM](https://gorm.io/) - Go Object-Relational Mapping
- **Database**: PostgreSQL with connection pooling
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate) - Database migrations
- **WebSocket**: [Fiber WebSocket](https://github.com/gofiber/websocket) - Real-time communication
- **Validation**: [validator](https://github.com/go-playground/validator) - Struct validation
- **Authentication**: JWT (JSON Web Tokens)
- **Logging**: [Zerolog](https://github.com/rs/zerolog) - Structured logging

## Quick Start

### Prerequisites

- Go 1.25.5 or higher
- PostgreSQL 12+ (optional, for database features)
- Redis (optional, for caching)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-boilerplate-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Run migrations** (if using database)
   ```bash
   go run ./cmd/migrate -command up
   ```

5. **Start the server**
   ```bash
   go run ./cmd/api
   ```

The API will start on `http://localhost:8080` (or the port specified in your `.env` file).

## Configuration

Configuration is managed through environment variables. See `.env.example` for all available options.

### Key Configuration Variables

- `PORT` - Server port (default: 8080)
- `IS_PROD` - Production mode (default: false)
- `LOG_LEVEL` - Log level (debug, info, warn, error, fatal)
- `SECRET_KEY` - JWT secret key (required in production, min 32 characters)
- `DATABASE_URL` - PostgreSQL connection string
- `REDIS_URL` - Redis connection string (optional)
- `ALLOWED_ORIGINS` - CORS allowed origins (comma-separated)
- `TIMEZONE` - Timezone for date operations (IANA format, default: Asia/Manila)

## Project Structure

```
go-boilerplate-api/
├── cmd/
│   ├── api/              # Main application entry point
│   └── migrate/          # Database migration CLI
├── internal/
│   └── api/
│       ├── config/       # Configuration management
│       ├── db/           # Database connection and migrations
│       ├── handlers/     # HTTP/WebSocket handlers
│       ├── middlewares/  # Fiber middlewares
│       └── routes/       # Route definitions
├── shared/
│   ├── helpers/          # Shared utility functions
│   └── models/           # Database models
├── docs/                 # Documentation
└── .env.example          # Environment variables template
```

## API Endpoints

### Health Check
- `GET /v1/health` - Health check with database and Redis status

### Authentication
- `POST /v1/login` - User login (placeholder)

### WebSocket
- `WS /ws` - WebSocket endpoint for real-time communication

## Database

### Migrations

This project uses SQL migrations (golang-migrate) which is the industry best practice for production applications.

**Run migrations:**
```bash
go run ./cmd/migrate -command up
```

**Check migration version:**
```bash
go run ./cmd/migrate -command version
```

**Validate migrations:**
```bash
go run ./cmd/migrate -command validate
```

See [Database Migrations Documentation](./docs/database-migrations.md) for detailed information.

### GORM Usage

The project uses GORM for database operations. See [GORM Usage Documentation](./docs/database-gorm-usage.md) for examples and best practices.

**Key Features:**
- UUID primary keys (security best practice)
- Connection pooling and reuse
- Automatic timestamp management
- Transaction support

## WebSocket

WebSocket support is available at `/ws` endpoint. See [WebSocket Documentation](./docs/websocket.md) for usage examples and client code.

## Development

### Build

```bash
# Build the application
go build -o api.exe ./cmd/api

# Build migration tool
go build -o migrate.exe ./cmd/migrate
```

### Run Tests

```bash
go test ./...
```

### Code Quality

```bash
# Format code
go fmt ./...

# Run linter
go vet ./...
```

## Production Deployment

### Checklist

- [ ] Set `IS_PROD=true`
- [ ] Set secure `SECRET_KEY` (min 32 characters)
- [ ] Configure `DATABASE_URL` with SSL
- [ ] Set `ALLOWED_ORIGINS` (explicit origins, not wildcard)
- [ ] Configure `LOG_LEVEL` (info, warn, or error)
- [ ] Set up SSL/TLS certificates
- [ ] Configure firewall and security groups
- [ ] Set up database backups
- [ ] Configure monitoring and logging
- [ ] Test migrations in staging environment

### Environment Variables

Ensure all required environment variables are set in your production environment. Never commit `.env` files to version control.

## Performance

The boilerplate is optimized for high performance:

- **Connection Pooling**: PostgreSQL (100 max connections) and Redis (100 pool size)
- **Prepared Statements**: Enabled for GORM queries
- **Fiber Optimizations**: Optimized timeouts, buffers, and memory usage
- **Graceful Shutdown**: Proper cleanup on application termination

Connection pooling is configured for optimal performance with connection reuse.

## Security

Security features included:

- ✅ JWT authentication
- ✅ Input validation (struct validation)
- ✅ Security headers (Helmet middleware)
- ✅ CORS configuration
- ✅ Rate limiting
- ✅ SQL injection prevention (GORM parameterized queries)
- ✅ UUID primary keys (prevents enumeration attacks)
- ✅ Error handling (no information leakage)

## Documentation

- [Database Migrations](./docs/database-migrations.md) - Migration guidelines
- [GORM Usage](./docs/database-gorm-usage.md) - Database operations guide
- [WebSocket](./docs/websocket.md) - WebSocket usage and examples


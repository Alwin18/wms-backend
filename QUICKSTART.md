# WMS Backend - Quick Start Guide

## Prerequisites

- Go 1.24+ installed
- PostgreSQL database running
- Git

## Installation

1. **Clone the repository**
   ```bash
   cd /home/gobe/Documents/self/wms
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Setup environment variables**
   ```bash
   cp .env.example .env
   ```
   
   Edit `.env` file and configure your database credentials:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=wms_db
   JWT_SECRET=your-super-secret-jwt-key-change-this
   ```

4. **Create database**
   ```bash
   createdb wms_db
   # or using psql:
   psql -U postgres -c "CREATE DATABASE wms_db;"
   ```

5. **Run the application**
   ```bash
   go run cmd/main.go
   ```
   
   Or build and run:
   ```bash
   go build -o bin/wms ./cmd/main.go
   ./bin/wms
   ```

## API Endpoints

### Health Check
```bash
GET http://localhost:8080/health
```

### Authentication

#### Register
```bash
POST http://localhost:8080/api/v1/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe",
  "phone": "08123456789"
}
```

#### Login
```bash
POST http://localhost:8080/api/v1/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "access_token": "eyJhbGc...",
    "refresh_token": "eyJhbGc...",
    "token_type": "Bearer",
    "expires_in": 900,
    "user": {
      "id": 1,
      "email": "user@example.com",
      "name": "John Doe",
      "phone": "08123456789",
      "status": "active"
    }
  }
}
```

#### Refresh Token
```bash
POST http://localhost:8080/api/v1/auth/refresh-token
Content-Type: application/json

{
  "refresh_token": "your_refresh_token"
}
```

#### Change Password (Protected)
```bash
POST http://localhost:8080/api/v1/auth/change-password
Authorization: Bearer your_access_token
Content-Type: application/json

{
  "old_password": "password123",
  "new_password": "newpassword123"
}
```

#### Logout (Protected)
```bash
POST http://localhost:8080/api/v1/auth/logout
Authorization: Bearer your_access_token
```

## Project Structure

```
wms/
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   ├── app.go                  # Application bootstrap
│   ├── config.go               # Configuration loader
│   ├── fiber.go                # Fiber setup
│   ├── gorm.go                 # Database setup
│   ├── logrus.go               # Logger setup
│   └── validator.go            # Validator setup
├── internal/
│   ├── dto/
│   │   └── auth.go             # Auth DTOs
│   ├── handlers/
│   │   └── auth.go             # Auth handlers
│   ├── models/
│   │   ├── audit_log.go        # Audit log model
│   │   ├── permission.go       # Permission model
│   │   ├── role.go             # Role model
│   │   ├── user.go             # User model
│   │   └── warehouse.go        # Warehouse model
│   ├── routes/
│   │   └── auth.go             # Auth routes
│   ├── services/
│   │   └── auth/
│   │       ├── login.go        # Login service
│   │       ├── register.go     # Register service
│   │       ├── refresh_token.go # Refresh token service
│   │       └── password.go     # Password management
│   └── utils/
│       ├── jwt.go              # JWT utilities
│       ├── password.go         # Password utilities
│       └── response.go         # Response utilities
├── pkg/
│   └── middleware/
│       └── middleware.go       # Auth & role middleware
├── .env.example                # Environment template
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Development

### Database Migration

The application automatically runs migrations on startup. Models are defined in `internal/models/`.

### Adding New Routes

1. Create handler in `internal/handlers/`
2. Create service in `internal/services/`
3. Register routes in `internal/routes/`
4. Add route setup in `cmd/main.go`

## Testing with cURL

```bash
# Register
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123","name":"Test User"}'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Access protected endpoint
curl -X POST http://localhost:8080/api/v1/auth/logout \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

## Next Steps

- [ ] Implement User Management endpoints
- [ ] Implement Master Data services (Products, Warehouses, etc.)
- [ ] Implement Inventory service
- [ ] Implement Inbound/Outbound services
- [ ] Add unit tests
- [ ] Add API documentation (Swagger)
- [ ] Implement email service for password reset
- [ ] Add rate limiting
- [ ] Add request logging

## License

MIT

## COVERAGE
```
gcov2lcov -infile test/coverage/coverage.out -outfile test/coverage/coverage.lcov
genhtml test/coverage/coverage.lcov --output-dir coverage-html 

```

# Testing Guide

## Overview

Folder ini berisi test files dan coverage reports untuk WMS Backend.

## Struktur Folder

```
test/
├── coverage/              # Coverage reports (auto-generated, git-ignored)
│   ├── coverage.out      # Raw coverage data
│   └── coverage.html     # HTML coverage report
└── run_coverage.sh       # Script untuk run tests dengan coverage
```

## Menjalankan Tests

### Run Semua Tests
```bash
go test ./...
```

### Run Tests dengan Verbose Output
```bash
go test ./... -v
```

### Run Tests untuk Package Tertentu
```bash
# Test untuk cmd package
go test ./cmd/... -v

# Test untuk internal/services/auth
go test ./internal/services/auth/... -v
```

### Run Tests dengan Coverage
```bash
./test/run_coverage.sh
```

Script ini akan:
1. Run semua tests di project
2. Generate coverage report di `test/coverage/coverage.out`
3. Generate HTML coverage report di `test/coverage/coverage.html`
4. Tampilkan coverage summary di terminal

### View HTML Coverage Report
```bash
# Linux
xdg-open test/coverage/coverage.html

# macOS
open test/coverage/coverage.html

# Windows
start test/coverage/coverage.html
```

## Menulis Tests

### Naming Convention
- Test files harus bernama `*_test.go`
- Test functions harus diawali dengan `Test`
- Letakkan test files di package yang sama dengan code yang di-test

### Contoh Test Structure

```go
package auth_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/Alwin18/wms/internal/services/auth"
)

func TestLogin(t *testing.T) {
    // Arrange
    req := &dto.LoginRequest{
        Email:    "test@example.com",
        Password: "password123",
    }

    // Act
    resp, err := auth.Login(req)

    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, resp)
    assert.NotEmpty(t, resp.AccessToken)
}
```

### Testing dengan Database

Jika test memerlukan database connection:

```go
func TestWithDatabase(t *testing.T) {
    app, err := Initialize()
    if err != nil {
        t.Skipf("Skipping test - database not available: %v", err)
        return
    }
    
    // Your test code here
}
```

## Coverage Goals

- **Overall Coverage**: Target 80%+
- **Critical Paths**: Target 90%+ (auth, inventory, transactions)
- **Utilities**: Target 95%+

## Testing Best Practices

1. **Arrange-Act-Assert Pattern**: Struktur test dengan jelas
2. **Table-Driven Tests**: Gunakan untuk test multiple scenarios
3. **Mock External Dependencies**: Gunakan mock untuk database, API calls, dll
4. **Test Edge Cases**: Jangan hanya test happy path
5. **Clear Test Names**: Nama test harus menjelaskan apa yang di-test

## CI/CD Integration

Coverage report bisa di-integrate dengan CI/CD:

```yaml
# Example GitHub Actions
- name: Run tests with coverage
  run: |
    go test ./... -coverprofile=test/coverage/coverage.out
    go tool cover -func=test/coverage/coverage.out
```

## Troubleshooting

### Tests Skipped karena Database
Jika tests di-skip karena database tidak tersedia:
1. Pastikan PostgreSQL running
2. Setup `.env` file dengan database credentials yang benar
3. Create database: `createdb wms_db`

### Coverage Report Kosong
Jika tidak ada coverage data:
- Pastikan ada test files yang di-run
- Check apakah semua tests di-skip
- Pastikan tests tidak error sebelum coverage di-generate

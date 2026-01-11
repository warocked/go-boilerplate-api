# Database Package - GORM Usage

This package uses GORM (Go Object-Relational Mapping) for database operations.

## Getting Started

### Accessing the Database

```go
import "go-boilerplate-api/internal/api/db"

// Get the GORM DB instance
gormDB := db.GetDB()

// Use GORM methods
var user models.User
gormDB.Where("email = ?", "user@example.com").First(&user)
```

## Models

All models should be defined in `shared/models/` with GORM tags:

```go
type User struct {
    ID        string    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
    Email     string    `gorm:"type:varchar(255);uniqueIndex;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
```

## Common Operations

### Create

```go
user := models.User{
    Email: "user@example.com",
    FirstName: "John",
    LastName: "Doe",
}
result := db.GetDB().Create(&user)
if result.Error != nil {
    // Handle error
}
```

### Read

```go
// Find by ID
var user models.User
db.GetDB().First(&user, userID)

// Find by condition
db.GetDB().Where("email = ?", email).First(&user)

// Find all
var users []models.User
db.GetDB().Find(&users)
```

### Update

```go
// Update single field
db.GetDB().Model(&user).Update("first_name", "Jane")

// Update multiple fields
db.GetDB().Model(&user).Updates(models.User{
    FirstName: "Jane",
    LastName: "Smith",
})
```

### Delete

```go
// Hard delete
db.GetDB().Delete(&user)

// Soft delete (if DeletedAt is defined in model)
db.GetDB().Delete(&user) // Sets DeletedAt timestamp
```

### Transactions

```go
err := db.GetDB().Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user1).Error; err != nil {
        return err
    }
    if err := tx.Create(&user2).Error; err != nil {
        return err
    }
    return nil
})
```

## Connection Pooling & Reuse

**The application is configured for optimal connection reuse and pooling:**

- **Single Global Instance**: All database operations use the same global `DB` instance, ensuring connections are reused across all requests
- **Connection Pool**: Configured with 100 max open connections and 10 idle connections
- **Connection Lifetime**: Connections are recycled after 1 hour to prevent stale connections
- **Idle Timeout**: Unused connections are closed after 30 minutes to free resources
- **Prepared Statements**: Enabled by default for better performance and connection reuse

### Connection Pool Configuration

```go
MaxOpenConns: 100        // Maximum connections in the pool
MaxIdleConns: 10         // Idle connections kept for immediate reuse
ConnMaxLifetime: 1 hour  // Maximum connection age
ConnMaxIdleTime: 30 min  // Maximum idle time before closing
PrepareStmt: true        // Prepared statements for reuse
```

### Monitoring Connection Pool

```go
import "go-boilerplate-api/internal/api/db"

// Get connection pool statistics
stats, err := db.GetPoolStats()
if err == nil {
    fmt.Printf("Open connections: %d/%d\n", stats.OpenConnections, stats.MaxOpenConnections)
    fmt.Printf("In use: %d, Idle: %d\n", stats.InUse, stats.Idle)
}

// Verify connections are being reused
reusing, _ := db.VerifyConnectionReuse()
if reusing {
    fmt.Println("Connections are being reused properly")
}
```

Connection pooling is configured for optimal performance with connection reuse.

## Performance

- **Connection pooling is configured for high performance**
- **Connections are automatically reused** - single global DB instance
- **Prepared statements are enabled by default** - queries are cached and reused
- Use `.Select()` to limit fields when querying
- Use indexes defined in migrations for better query performance

## Security

- GORM automatically uses parameterized queries (SQL injection protection)
- UUID primary keys prevent enumeration attacks
- Always validate input before database operations

## Migrations

Database schema changes are managed through migrations in `internal/api/db/migrations/`.
Migrations run automatically on application startup.

See [Database Migrations](./database-migrations.md) for migration guidelines.

package db

// Security Note: This package uses GORM which automatically uses parameterized queries
// All database queries are automatically protected against SQL injection attacks.
//
// Example safe query with GORM:
//   var user models.User
//   db.DB.Where("id = ?", userID).First(&user)
//
// GORM automatically uses parameterized queries, so you don't need to worry about SQL injection.
// However, when using Raw SQL, always use parameterized queries:
//   db.DB.Raw("SELECT * FROM users WHERE id = ?", userID).Scan(&user)
//
// NEVER use string concatenation with Raw SQL:
//   db.DB.Raw("SELECT * FROM users WHERE id = " + userID) // UNSAFE!

// UUID Primary Keys:
// Always use UUID as primary keys, NEVER use auto-incrementing integers (SERIAL/BIGSERIAL).
// This prevents enumeration attacks and works better in distributed systems.
// Models should use: gorm:"type:uuid;primary_key;default:uuid_generate_v4()"

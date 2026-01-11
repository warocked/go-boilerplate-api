-- Migration: 000001_initial_schema (DOWN)
-- Description: Rollback initial schema
-- WARNING: This will DROP all data in the users table
-- Only use in development or after data backup in production

-- Drop trigger first
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop indexes
DROP INDEX IF EXISTS idx_users_created_at;
DROP INDEX IF EXISTS idx_users_email;

-- Drop table
DROP TABLE IF EXISTS users;

-- Note: We do NOT drop the uuid-ossp extension as it may be used by other objects

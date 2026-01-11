# Database Migrations

This directory contains database migration files that manage schema changes in a versioned, safe manner.

**Note**: This project uses SQL migrations (golang-migrate) which is the industry best practice for production applications, even when using GORM as the ORM. 

## Migration File Naming

Migrations follow this pattern:
- `<version>_<name>.up.sql` - Migration to apply (forward)
- `<version>_<name>.down.sql` - Migration to rollback (reverse)

Example:
- `000001_initial_schema.up.sql`
- `000001_initial_schema.down.sql`

## Safety Guidelines

### ✅ Safe Migration Practices

1. **Always use IF NOT EXISTS / IF EXISTS**
   ```sql
   CREATE TABLE IF NOT EXISTS users (...);
   DROP TABLE IF EXISTS temp_table;
   ```

2. **Add columns with defaults or NULL**
   ```sql
   ALTER TABLE users ADD COLUMN new_field VARCHAR(255) DEFAULT '';
   ```

3. **Use transactions where possible**
   - Migrations run in transactions automatically
   - If migration fails, it's rolled back

4. **Create indexes concurrently in production**
   ```sql
   CREATE INDEX CONCURRENTLY idx_name ON table_name(column_name);
   ```

5. **Test migrations on staging first**
   - Always test migrations on staging/development before production

### ❌ Dangerous Operations (Avoid or Use Carefully)

1. **Dropping columns** - Data loss!
   - Only drop after confirming no application code uses the column
   - Consider deprecation period

2. **Renaming columns/tables** - Break existing queries
   - Add new column, migrate data, then drop old column

3. **Changing column types** - May fail if data incompatible
   - Test thoroughly, use ALTER COLUMN TYPE with USING clause

4. **Dropping tables** - Complete data loss!
   - Only in down migrations or with explicit backup

## Running Migrations

### Automatic (on application start)
Migrations run automatically when the application starts (if DATABASE_URL is set).

### Manual (CLI)
```bash
# Run all pending migrations
go run ./cmd/migrate -command up

# Check current migration version
go run ./cmd/migrate -command version

# Validate migrations
go run ./cmd/migrate -command validate

# Build migration binary
go build -o migrate.exe ./cmd/migrate
./migrate.exe -command version
```

## Creating New Migrations

1. **Determine next version number**
   - Check existing migrations to find the highest version
   - Increment by 1

2. **Create migration files**
   ```bash
   # Option 1: Use migrate CLI tool (recommended)
   migrate create -ext sql -dir internal/api/db/migrations -seq add_user_roles
   
   # Option 2: Create manually
   # Create: 000002_add_user_roles.up.sql
   # Create: 000002_add_user_roles.down.sql
   ```

3. **Write safe migration SQL**
   - Use IF NOT EXISTS / IF EXISTS
   - Add data migration if needed
   - Test thoroughly

4. **Write down migration (rollback)**
   - Reverse all changes from up migration
   - Include data cleanup if needed

## Production Safety Checklist

Before running migrations in production:

- [ ] Test migrations on staging environment
- [ ] Backup database before running migrations
- [ ] Review migration SQL for data safety
- [ ] Ensure rollback (down) migration is tested
- [ ] Run during maintenance window (if large changes)
- [ ] Monitor application logs during migration
- [ ] Verify migration success with `-command version`
- [ ] Check application functionality after migration

## Version Tracking

Migrations are tracked in the `schema_migrations` table:
- `version` - Current migration version
- `dirty` - Whether migration is in dirty state (failed mid-migration)

**Never modify `schema_migrations` table manually!**

## Troubleshooting

### Dirty Migration State
If migrations are in a dirty state:
1. Check database logs for error
2. Manually fix the issue (if safe)
3. Mark migration as clean: `UPDATE schema_migrations SET dirty = false;`
4. Or rollback and re-run: Use down migration, then up again

### Migration Fails
1. Check error message in logs
2. Fix the SQL issue
3. If already applied, create new migration to fix
4. If not applied, fix the migration file and re-run

## Example: Adding a Column

```sql
-- 000002_add_user_roles.up.sql
ALTER TABLE users 
ADD COLUMN IF NOT EXISTS role VARCHAR(50) DEFAULT 'user' NOT NULL;

CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);
```

```sql
-- 000002_add_user_roles.down.sql
DROP INDEX IF EXISTS idx_users_role;
ALTER TABLE users DROP COLUMN IF EXISTS role;
```

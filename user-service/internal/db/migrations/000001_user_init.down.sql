-- Drop trigger
DROP TRIGGER IF EXISTS set_updated_at ON users;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at_column;

-- Drop policy
DROP POLICY IF EXISTS tenant_isolation_policy ON users;

-- Drop table
DROP TABLE IF EXISTS users;


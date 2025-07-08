CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  tenant_id UUID NOT NULL,
  name TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE CHECK (position(' ' in email) = 0),
  password TEXT NOT NULL,
  role TEXT NOT NULL DEFAULT 'user', 
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);



CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;



CREATE TRIGGER set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();


ALTER TABLE users ENABLE ROW LEVEL SECURITY;


CREATE POLICY tenant_isolation_policy ON users
  USING (tenant_id::text = current_setting('app.tenant_id')::text);

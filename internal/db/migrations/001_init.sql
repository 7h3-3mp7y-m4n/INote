CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS notes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug TEXT UNIQUE NOT NULL DEFAULT encode(gen_random_bytes(6), 'base64'),
    content TEXT NOT NULL,
    expires_at TIMESTAMPTZ,
    format TEXT DEFAULT 'plain',
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS groups (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    created_at timestamp WITH TIME ZONE NOT NULL
);
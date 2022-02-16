CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    surname text NOT NULL,
    email text NOT NULL UNIQUE,
    username text NOT NULL UNIQUE,
    password text NULL DEFAULT NULL,
    shell text NOT NULL DEFAULT '/bin/bash',
    public_ssh_key text NOT NULL,
    created_at timestamp WITH TIME ZONE NOT NULL
);
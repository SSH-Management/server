CREATE TABLE  IF NOT EXISTS roles (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    created_at timestamp WITH TIME ZONE NOT NULL
);

ALTER TABLE users ADD COLUMN role_id bigint;

ALTER TABLE users
    ADD CONSTRAINT users_role_id_constraint
    FOREIGN KEY (role_id) REFERENCES roles (id)
    ON DELETE CASCADE;


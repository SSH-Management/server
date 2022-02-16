CREATE TABLE IF NOT EXISTS servers (
    id bigserial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    ip text NOT NULL,
    public_ip text NULL DEFAULT NULL,
    created_at timestamp WITH TIME ZONE NOT NULL,
    group_id bigint NOT NULL
);

ALTER TABLE servers
    ADD CONSTRAINT servers_group_id_constraint
    FOREIGN KEY (group_id) REFERENCES groups (id)
    ON DELETE CASCADE;

CREATE TABLE servers (
    id bigint(20) unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name varchar(60) NOT NULL,
    ip varchar(35) NOT NULL,
    public_ip varchar(35) NULL,
    created_at timestamp NOT NULL,
    group_id bigint(20) unsigned NOT NULL
);

ALTER TABLE servers
    ADD CONSTRAINT servers_group_id_constraint
    FOREIGN KEY (group_id) REFERENCES groups (id)
    ON DELETE CASCADE;

CREATE TABLE roles (
    id bigint(20) unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name varchar(60) NOT NULL,
    created_at timestamp NOT NULL
);

ALTER TABLE users ADD COLUMN role_id bigint(20) unsigned;

ALTER TABLE users
    ADD CONSTRAINT users_role_id_constraint
    FOREIGN KEY (role_id) REFERENCES roles (id)
    ON DELETE CASCADE;


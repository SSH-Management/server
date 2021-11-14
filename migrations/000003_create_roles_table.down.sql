ALTER TABLE users DROP FOREIGN KEY users_role_id_constraint;

ALTER TABLE users DROP COLUMN role_id;

DROP TABLE IF EXISTS roles;
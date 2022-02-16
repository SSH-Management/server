ALTER TABLE user_groups
    ADD CONSTRAINT user_groups_group_id_constraint
    FOREIGN KEY (group_id) REFERENCES groups (id)
    ON DELETE CASCADE;

ALTER TABLE user_groups
    ADD CONSTRAINT user_groups_user_id_constraint
    FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE;
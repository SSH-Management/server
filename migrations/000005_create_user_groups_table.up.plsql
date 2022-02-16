CREATE TABLE IF NOT EXISTS user_groups (
    user_id bigint NOT NULL,
    group_id bigint NOT NULL,
    CONSTRAINT pk_user_groups PRIMARY KEY(user_id, group_id)
);

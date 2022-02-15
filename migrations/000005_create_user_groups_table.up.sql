CREATE TABLE `user_groups` (
    user_id bigint(20) unsigned NOT NULL,
    group_id bigint(20) unsigned NOT NULL,
    CONSTRAINT pk_user_groups PRIMARY KEY(user_id, group_id)
);

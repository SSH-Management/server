CREATE TABLE `users` (
    id bigint(20) unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name varchar(60) NOT NULL,
    surname varchar(60) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    username varchar(60) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    shell varchar(255) NOT NULL,
    public_ssh_key varchar(255) NOT NULL,
    created_at timestamp NOT NULL
);
CREATE TABLE IF NOT EXISTS `user` (
    `id` varchar(40) not null,
    `first_name` varchar(256) default null,
    `last_name` varchar(256) default null,
    `nickname` varchar(256) default null,
    `password` varchar(256) default null,
    `email` varchar(256) default null,
    `country` varchar(256) default null,
    `created_at` timestamp not null default current_timestamp(),
    `updated_at` timestamp  null on update current_timestamp(),
    PRIMARY KEY (`id`)
) ;
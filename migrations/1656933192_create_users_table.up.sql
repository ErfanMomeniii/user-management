CREATE TABLE IF NOT EXISTS `users` (
    `id` varchar(40) not null,
    `first_name` varchar(256) default null,
    `last_name` varchar(256) default null,
    `nickname` varchar(256) default null,
    `password` varchar(256) default null,
    `email` varchar(256) default null,
    `country` varchar(256) default null,
    `created_at` timestamp not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp on update CURRENT_TIMESTAMP not null default CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ;
create table users
(
    id          bigint auto_increment,
    name        varchar(100) null comment 'name',
    status      int(1)       null comment 'status',
    created_at datetime     null comment 'created time',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
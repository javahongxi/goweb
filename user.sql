create table user
(
    id          bigint auto_increment
        primary key,
    username    varchar(20) not null,
    nickname    varchar(20) null,
    gender      tinyint     null,
    age         int         null,
    create_date datetime    null,
    update_date datetime    null
);
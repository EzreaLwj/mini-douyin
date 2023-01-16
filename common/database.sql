create database if not exists mini_douyin charset utf8mb4 collate utf8mb4_general_ci;
use mini_douyin;

create table if not exists tb_user
(
    user_id        bigint primary key auto_increment comment '用户id',
    username       varchar(32) not null comment '用户名称',
    `password`     varchar(32) not null comment '用户密码',
    follow_count   bigint default 0 comment '关注总数',
    follower_count bigint default 0 comment '粉丝总数'
) comment '用户表';

insert into tb_user(username, password, follow_count, follower_count)
values ('ezreal', '123', 1, 1);
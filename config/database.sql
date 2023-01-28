create database if not exists mini_douyin
    character set utf8mb4
    collate utf8mb4_general_ci;

use mini_douyin;

create table if not exists tb_user
(
    `user_id`        bigint primary key auto_increment comment '用户id',
    `username`       varchar(32) not null comment '用户名',
    `password`       varchar(32) not null comment '密码',
    `follow_count`   bigint   default 0 comment '关注总数',
    `follower_count` bigint   default 0 comment '粉丝总数',
    `create_time`    DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
    ) ENGINE = InnoDB,
    DEFAULT CHARSET = utf8mb4, comment '用户表';


create table if not exists tb_video
(
    `video_id`      bigint primary key auto_increment comment '视频id',
    `user_id`       bigint       not null comment '用户id',
    `title`         varchar(255) not null comment '视频标题',
    `play_url`      varchar(512) not null comment '播放地址',
    `cover_url`     varchar(512) not null comment '视频封面地址',
    `comment_count` bigint       not null default 0 comment '视频的评论总数',
    `create_time`   DATETIME              DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`   DATETIME              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
    ) ENGINE = InnoDB,
    DEFAULT CHARSET = utf8mb4, comment '视频表';

create table if not exists tb_like
(
    `like_id`     bigint primary key auto_increment comment '点赞id',
    `user_id`     bigint  not null comment '用户id',
    `video_id`    bigint  not null comment '视频id',


    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB,
    DEFAULT CHARSET = utf8mb4, comment '点赞表';

create table if not exists tb_comment
(
    `comment_id`  bigint primary key auto_increment comment '评论id',
    `user_id`     bigint       not null comment '用户id',
    `video_id`    bigint       not null comment '视频id',
    `content`     varchar(255) not null comment '评论内容',
    `create_date` date         not null comment '评论发布日期',


    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
    ) ENGINE = InnoDB,
    DEFAULT CHARSET = utf8mb4, comment '评论表';

create table if not exists tb_relation
(
    `relation_id` bigint primary key auto_increment comment '关注id',
    `user_id`     bigint not null comment '用户id',
    `to_user_id`  bigint not null comment '被关注者id',


    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE = InnoDB,
    DEFAULT CHARSET = utf8mb4, comment '关注表';


create table if not exists tb_message
(
    `message_id`  bigint primary key auto_increment comment '消息id',
    `user_id`     bigint       not null comment '用户id',
    `to_user_id`  bigint       not null comment '对方用户id',
    `content`     varchar(255) not null comment '消息内容',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
    ) ENGINE = InnoDB,
    DEFAULT CHARSET = utf8mb4, comment '关注表';
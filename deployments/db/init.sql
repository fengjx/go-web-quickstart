create table user
(
    `id` bigint auto_increment primary key,
    `username` varchar(32) not null comment '用户名',
    `pwd` varchar(64) not null comment '密码',
    `salt` varchar(64) not null comment '密码盐',
    `nick` varchar(32) not null comment '昵称',
    `utime` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
    `ctime` timestamp not null default current_timestamp comment '创建时间',
    unique uk_u(username)
)
engine = innodb
default charset = utf8mb4 comment '用户信息表';

create table blog
(
    `id` bigint auto_increment primary key,
    `uid` bigint not null comment '用户ID',
    `title` varchar(128) not null comment '标题',
    `content` text not null comment '内容',
    `create_time` bigint not null comment '创建时间',
    `utime` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
    `ctime` timestamp not null default current_timestamp comment '创建时间',
    index idx_u(uid)
)
engine = innodb
default charset = utf8mb4 comment '博客表';



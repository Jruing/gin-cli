CREATE TABLE Users
(
    id       SERIAL  PRIMARY KEY,
    nickname varchar(32),
    username varchar(32) unique NOT NULL,
    password varchar(32) not null,
    sex      char(1),
    email    varchar(32) not null,
    status   int not null,
    created  timestamp
);
comment on table Users is '用户表';
comment on column Users.sex is '性别';
comment on column Users.nickname is '昵称';
comment on column Users.username is '用户名';
comment on column Users.password is '密码';
comment on column Users.email is '邮箱';
comment on column Users.status is '状态';
comment on column Users.created is '创建时间';

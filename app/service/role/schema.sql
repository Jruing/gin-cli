create TABLE Roles
(
    id       SERIAL  PRIMARY KEY,
    rolename varchar(32) not null,
    status   int not null,
    created  timestamp
);
comment on table Roles is '角色表';
comment on column Roles.rolename is '角色名称';
comment on column Roles.status is '状态';
comment on column Roles.created is '创建时间';
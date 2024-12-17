CREATE TABLE Domains
(
    id      SERIAL  PRIMARY KEY,
    domain  varchar(32) not null,
    status  int not null,
    created timestamp
);
comment on table Domains is '域表';
comment on column Domains.domain is '域名称';
comment on column Domains.status is '状态';
comment on column Domains.created is '创建时间';
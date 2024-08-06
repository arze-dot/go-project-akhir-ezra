-- +migrate Up
-- +migrate StatementBegin

create table category (
    id           BIGINT NOT NULL PRIMARY KEY,
    name         varchar(256),
    description  varchar(1024),
    created_at   TIMESTAMP,
    created_by   BIGINT,
    updated_at   TIMESTAMP,
    updated_by   BIGINT
)

-- +migrate StatementEnd

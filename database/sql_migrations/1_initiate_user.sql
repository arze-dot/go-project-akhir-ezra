-- +migrate Up
-- +migrate StatementBegin

create table users (
    id           BIGINT NOT NULL PRIMARY KEY,
    username     varchar(256),
    password     varchar(256),
    role         varchar(256),
    created_at   TIMESTAMP,
    created_by   BIGINT,
    updated_at   TIMESTAMP,
    updated_by   BIGINT
)

-- +migrate StatementEnd

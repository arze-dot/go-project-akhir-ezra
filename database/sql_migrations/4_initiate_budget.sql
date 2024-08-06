-- +migrate Up
-- +migrate StatementBegin

create table budget (
    id           BIGINT NOT NULL PRIMARY KEY,
    category_id  BIGINT,
    price        integer,
    user_id      BIGINT,
    description  varchar(1024),
    created_at   TIMESTAMP,
    created_by   BIGINT,
    updated_at   TIMESTAMP,
    updated_by   BIGINT,
    FOREIGN KEY (category_id) REFERENCES category(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)

-- +migrate StatementEnd

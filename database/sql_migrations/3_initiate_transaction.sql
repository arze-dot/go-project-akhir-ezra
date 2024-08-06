-- +migrate Up
-- +migrate StatementBegin

create table transaction (
    id           BIGINT NOT NULL PRIMARY KEY,
    income       integer,
    expense      integer,
    category_id  BIGINT,
    user_id      BIGINT,
    title        varchar(256),
    description  varchar(1024),
    created_at   TIMESTAMP,
    created_by   BIGINT,
    updated_at   TIMESTAMP,
    updated_by   BIGINT,
    FOREIGN KEY (category_id) REFERENCES category(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)

-- +migrate StatementEnd

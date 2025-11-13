-- +goose Up
-- +goose StatementBegin
create table if not exists pull_request (
    pull_request_id varchar(8) primary key,
    pull_request_name varchar(255) not null,
    author_id varchar(8) not null references "user"(user_id),
    status varchar(8) not null default 'OPEN',
    merged_at timestamp,
    created_at timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists pull_request;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
create table if not exists user_pull_request (
    user_id varchar(8) not null references "user"(user_id) on delete restrict,
    pull_request_id varchar(8) not null references pull_request(pull_request_id) on delete restrict,
    primary key (user_id, pull_request_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_pull_request;
-- +goose StatementEnd

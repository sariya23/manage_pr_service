-- +goose Up
-- +goose StatementBegin
create table if not exists "user" (
    user_id varchar(8) primary key,
    username varchar(37) not null,
    is_active boolean not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists "user";
-- +goose StatementEnd

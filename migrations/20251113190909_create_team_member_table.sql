-- +goose Up
-- +goose StatementBegin
create table if not exists team_member (
    team_name varchar(255) not null references team(team_name) on delete restrict,
    user_id varchar(8) not null unique references "user"(user_id) on delete restrict
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists team_member;
-- +goose StatementEnd

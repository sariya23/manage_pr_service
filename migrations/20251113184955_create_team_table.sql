-- +goose Up
-- +goose StatementBegin
create table if not exists team (
    team_name varchar(255) primary key
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists team;
-- +goose StatementEnd

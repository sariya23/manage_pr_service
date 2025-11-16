-- +goose Up
-- +goose StatementBegin
alter table team_member
alter column team_name type varchar(32);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

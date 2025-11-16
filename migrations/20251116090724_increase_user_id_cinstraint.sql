-- +goose Up
-- +goose StatementBegin
alter table "user"
alter column user_id type varchar(32);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

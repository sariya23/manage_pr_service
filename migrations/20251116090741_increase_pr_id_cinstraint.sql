-- +goose Up
-- +goose StatementBegin
alter table pull_request
alter column pull_request_id type varchar(32);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
alter table pull_request
alter column assigned_reviewers type varchar(32)[];
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

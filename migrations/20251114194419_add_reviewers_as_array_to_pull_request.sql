-- +goose Up
-- +goose StatementBegin
alter table pull_request
add column assigned_reviewers varchar(8)[] not null default '{}';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table pull_request drop column assigned_reviewers;
-- +goose StatementEnd

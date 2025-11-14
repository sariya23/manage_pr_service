-- +goose Up
-- +goose StatementBegin
drop table if exists reviewer_pull_request;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

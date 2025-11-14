-- +goose Up
-- +goose StatementBegin
alter table user_pull_request rename to reviwer_pull_request;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table reviwer_pull_request rename to user_pull_request;
-- +goose StatementEnd

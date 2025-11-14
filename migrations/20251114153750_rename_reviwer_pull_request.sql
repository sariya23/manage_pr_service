-- +goose Up
-- +goose StatementBegin
alter table reviwer_pull_request rename to reviewer_pull_request;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table reviewer_pull_request rename to reviwer_pull_request;
-- +goose StatementEnd

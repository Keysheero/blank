-- +goose Up
-- +goose StatementBegin
ALTER TABLE student
    ADD COLUMN name VARCHAR(255);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE student
    DROP COLUMN name;

-- +goose StatementEnd

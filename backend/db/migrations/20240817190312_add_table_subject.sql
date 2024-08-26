-- +goose Up
CREATE TABLE  subject(
    id INT UNSIGNED NOT NULL AUTO_INCREMENT ,
    name VARCHAR() NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE subject;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    UserID BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    UserName VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL UNIQUE
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS users;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

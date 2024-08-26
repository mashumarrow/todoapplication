-- +goose Up
CREATE TABLE IF NOT EXISTS classrooms (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    classroom_id BIGINT UNSIGNED NOT NULL UNIQUE,
    classroom_name VARCHAR(255) NOT NULL,
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS classrooms;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

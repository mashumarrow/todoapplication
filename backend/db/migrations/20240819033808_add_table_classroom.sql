-- +goose Up
CREATE TABLE classrooms (
    ClassroomID BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    ClassroomName VARCHAR(255));
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS classrooms;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

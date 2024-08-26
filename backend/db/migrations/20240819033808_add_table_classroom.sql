-- +goose Up
CREATE TABLE Classroom (
    ClassroomID BIGINT UNSIGNED AUTO_INCREMENTPRIMARY KEY,
    ClassroomNAme VARCHAR(255));
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS classrooms;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

-- +goose Up
CREATE TABLE IF NOT EXISTS todos (
    TodoID BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    UserID BIGINT UNSIGNED NOT NULL,
    Title VARCHAR(255) NOT NULL,
    Completed BOOLEAN DEFAULT FALSE,
    Period INT NOT NULL,
    ClassroomName VARCHAR(255),
    SubjectName VARCHAR(255),
    FOREIGN KEY (UserID) REFERENCES users(UserID)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS todos;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

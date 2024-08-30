-- +goose Up
CREATE TABLE IF NOT EXISTS todos (
    UserID BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Completed BOOLEAN DEFAULT FALSE,
    SubjectID BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (SubjectID) REFERENCES subjects(SubjectID),
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

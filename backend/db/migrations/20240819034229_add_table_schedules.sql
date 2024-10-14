-- +goose Up
CREATE TABLE IF NOT EXISTS schedules(
    UserID BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    
    Dayofweek ENUM('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday') NOT NULL,
    Period INT NOT NULL,
    ClassroomName VARCHAR(255),
    SubjectName VARCHAR(255),
    
    FOREIGN KEY (UserID) REFERENCES users(UserID)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS schedules;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

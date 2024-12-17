-- +goose Up
CREATE TABLE IF NOT EXISTS schedules(
    ScheduleID BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    UserID BIGINT UNSIGNED NOT NULL,
    
    DayOfWeek ENUM('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday') NOT NULL,
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

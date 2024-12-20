-- +goose Up
CREATE TABLE IF NOT EXISTS todos (
    ID INT UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    TodoID  VARCHAR(255)   ,
    UserID BIGINT UNSIGNED  NOT NULL,
    Title VARCHAR(255) NOT NULL,
    Completed BOOLEAN DEFAULT FALSE,
    Period INT NOT NULL,
    Classroomname VARCHAR(255),
    Subjectname VARCHAR(255),
    Dayofweek VARCHAR(255),
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

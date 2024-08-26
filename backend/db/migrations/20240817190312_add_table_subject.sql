-- +goose Up
CREATE TABLE  subject(
    SubjectID BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY, 
    SubjectName VARCHAR(255)
   
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE subject;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

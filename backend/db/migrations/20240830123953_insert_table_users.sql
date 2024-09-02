-- +goose Up
INSERT INTO users (UserId, UserName, Email) VALUES (1, 'kkk', 'kkk@example.com');
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DELETE FROM users WHERE userId IN ('1');
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

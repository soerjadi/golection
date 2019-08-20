-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id      BIGSERIAL   NOT NULL PRIMARY KEY,
    username    VARCHAR(20) NOT NULL,
    fullname    VARCHAR(155) NOT NULL,
    passhash    VARCHAR(80) NOT NULL,
    identity_number INTEGER NOT NULL,
    is_voted        BOOLEAN NULL DEFAULT false,
    role            INTEGER NOT NULL,
    loginable       BOOLEAN NULL DEFAULT false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

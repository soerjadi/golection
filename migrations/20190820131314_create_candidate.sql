-- +goose Up
-- +goose StatementBegin
CREATE TABLE candidates (
    id              bigserial   PRIMARY KEY NOT NULL,
    name            varchar(50) not null,
    pic             varchar(155)    null,
    description     text        null,
    vote_count      bigint  null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS candidates;
-- +goose StatementEnd

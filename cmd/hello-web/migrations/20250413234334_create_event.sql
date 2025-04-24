-- +goose Up
-- +goose StatementBegin

CREATE TABLE event (
    id          INTEGER PRIMARY KEY,
    name        VARCHAR(255) NOT NULL UNIQUE DEFAULT '',
    create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE event;

-- +goose StatementEnd

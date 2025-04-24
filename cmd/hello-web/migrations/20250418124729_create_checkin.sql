-- +goose Up
-- +goose StatementBegin
CREATE TABLE checkin (
    id          INTEGER PRIMARY KEY,
    event_id    INTEGER NOT NULL,
    create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_checkin_event_id_event_id FOREIGN KEY (event_id) REFERENCES event (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE checkin;
-- +goose StatementEnd

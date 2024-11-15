-- +migrate Up
CREATE TABLE log_records (
    id INTEGER PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL,
    level TEXT NOT NULL,
    message TEXT NOT NULL
);
-- -migrate Down
DROP TABLE log_records;
		
-- POSTGRES
DROP TABLE IF EXISTS log_records;
CREATE TABLE log_records (
    id INTEGER PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL,
    level TEXT NOT NULL,
    message TEXT NOT NULL
);
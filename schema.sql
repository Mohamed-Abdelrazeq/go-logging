-- SQLite
DROP TABLE IF EXISTS log_records;
CREATE TABLE log_records (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    timestamp DATETIME NOT NULL,
    level TEXT NOT NULL,
    message TEXT NOT NULL
);
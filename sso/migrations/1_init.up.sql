-- ./migrations/1_init.up.sql

CREATE TABLE IF NOT EXISTS users
(
    id        INTEGER PRIMARY KEY,
    email     TEXT    NOT NULL UNIQUE,
    pass_hash BLOB    NOT NULl
);
CREATE INDEX IF NOT EXISTS idx_email ON users (email);

CREATE TABLE IF NOT EXISTS apps
(
    id     INTEGER PRIMARY KEY,
    name   TEXT    NOT NULL UNIQUE,
    secret TEXT    NOT NULL UNIQUE
);
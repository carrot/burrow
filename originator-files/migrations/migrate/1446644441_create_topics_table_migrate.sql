CREATE TABLE topics(
    id                  serial          PRIMARY KEY,
    name                varchar(128)    NOT NULL UNIQUE
);

CREATE TABLE topics(
    id                  serial          PRIMARY KEY,
    name                varchar(128)    NOT NULL UNIQUE,
    created_at          timestamptz     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          timestamptz     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER topics_updated_at
BEFORE UPDATE
ON topics
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

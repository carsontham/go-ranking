-- +goose Up
-- +goose StatementBegin

CREATE TABLE user
(
    id INTEGER NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    score INTEGER NOT NULL
    balance DECIMAL(20,10) NOT NULL CHECK (balance >= 0)
);

INSERT INTO account (name, score)
VALUES ('John', 50),
       ('Daniel', 60),
       ('Poh', 70),

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd

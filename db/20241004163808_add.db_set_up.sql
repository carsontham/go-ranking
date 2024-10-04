-- +goose Up
-- +goose StatementBegin

CREATE TABLE ranked_users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    score INTEGER NOT NULL
);

INSERT INTO ranked_users (name, email, score)
VALUES ('John', 'john@gmail.com', 50),
       ('Daniel', 'daniel@gmail.com', 60),
       ('Poh', 'poh@gmail.com', 70);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd

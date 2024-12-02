CREATE TABLE requests
(
    x         FLOAT        NOT NULL,
    y         FLOAT        NOT NULL,
    radius    FLOAT        NOT NULL,
    result    BOOLEAN NOT NULL,
    send_time TIMESTAMP    NOT NULL,
    author    VARCHAR(255) NOT NULL,
    FOREIGN KEY (author) REFERENCES users (name)
);

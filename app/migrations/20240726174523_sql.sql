-- +goose Up
CREATE TABLE IF NOT EXISTS users (
        id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
        name CHAR(15) DEFAULT NULL,
        password CHAR(30) DEFAULT NULL,
        mobile CHAR(15) DEFAULT NULL,
        crate_at DATETIME DEFAULT NULL
    );

-- +goose Down
DROP TABLE users;

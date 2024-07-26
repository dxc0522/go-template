-- +goose Up

CREATE TABLE IF NOT EXISTS charging_data (
       id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
       waterpower DECIMAL(10, 2) DEFAULT NULL,
       allpower DECIMAL(10, 2) DEFAULT NULL,
       allmoney DECIMAL(10, 2) DEFAULT NULL,
       initpower DECIMAL(10, 2) DEFAULT NULL,
       lastpower DECIMAL(10, 2) DEFAULT NULL,
       difference DECIMAL(10, 2) DEFAULT NULL,
       price BIGINT DEFAULT NULL,
       date DATETIME DEFAULT NULL,
       uuid VARCHAR(255) NOT NULL
);

-- +goose Down

DROP TABLE charging_data;

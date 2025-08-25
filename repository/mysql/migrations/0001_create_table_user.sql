-- +migrate Up
CREATE TABLE users (
                       id int primary key AUTO_INCREMENT,
                       name varchar(255) not null,
                       phone_number varchar(255) unique,
                       password text not null
);

-- +migrate Down
DROP TABLE users;
CREATE TABLE IF NOT EXISTS library.authors
(
    id bigint unsigned auto_increment PRIMARY KEY,
    first_name varchar(255) NOT null,
    last_name varchar(255) NOT NULL,
    patronymic varchar(255) DEFAULT NULL
)
CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
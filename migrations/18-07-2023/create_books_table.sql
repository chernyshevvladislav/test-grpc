CREATE TABLE IF NOT EXISTS library.books
(
    id bigint unsigned auto_increment PRIMARY KEY,
    title varchar(255) NOT NULL
)
CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
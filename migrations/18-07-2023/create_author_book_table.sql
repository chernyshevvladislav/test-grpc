CREATE TABLE IF NOT EXISTS library.author_book
(
    author_id bigint unsigned,
    book_id bigint unsigned,

    foreign key(author_id) references authors(id),
    foreign key(book_id) references books(id)
)
CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
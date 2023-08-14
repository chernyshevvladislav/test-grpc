SET NAMES utf8mb4;
INSERT INTO library.authors (first_name,last_name,patronymic) VALUES("Лев","Толстой","Николаевич"),("Федор","Достоевский","Михайлович"),("Александр","Дюма", null);
INSERT INTO library.books (title) VALUES("Анна Каренина"),("Преступление и наказание"),("Граф Монте-Кристо"), ("Юность");
INSERT INTO library.author_book (author_id, book_id) VALUES(1, 1),(2, 2),(3, 3), (1,4);
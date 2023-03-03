CREATE TABLE authors (
	id int AUTO_INCREMENT,
	name varchar(100) DEFAULT "John Doe" NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE books (
	id int AUTO_INCREMENT,
	name varchar(100) DEFAULT "Lorem Ipsum" NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE library (
    id_author int,
    id_book int,
    FOREIGN KEY (id_author) REFERENCES authors(id),
    FOREIGN KEY (id_book) REFERENCES books(id)
);

INSERT INTO authors(name) VALUES ('J.K. Rowling'), ('Karen C. Timberlake'), ('William Timberlake'), ('Leo Tolstoy');

INSERT INTO books(name) VALUES ("Harry Potter and Philosopher's Stone"),( 'Harry Potter and Chamber of Secrets'),( 'Harry Potter and Prisoner of Azkaban'),('Harry Potter and Goblet of Fire'),( 'Harry Potter and Order of the Phoenix'),('Harry Potter and Half-Blood Prince'),( 'Harry Potter and Deathly Hallows'), ('Chemistry'), ('Organic Chemistry'), ('Advanced Chemistry'), ('Anna Karenina');

INSERT INTO library (id_author, id_book) VALUES (1,1),
(1,2),
(1,3),
(1,4),
(1,5),
(1,6),
(1,7),
(2,8),
(3,8),
(2,9),
(2,10),
(3,10),
(4,11);
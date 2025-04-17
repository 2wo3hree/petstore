CREATE TABLE IF NOT EXISTS books (
                                     id         SERIAL PRIMARY KEY,
                                     title      TEXT   NOT NULL,
                                     author_id  INT    NOT NULL,
                                     CONSTRAINT fk_author FOREIGN KEY (author_id)
    REFERENCES authors(id) ON DELETE RESTRICT
    );
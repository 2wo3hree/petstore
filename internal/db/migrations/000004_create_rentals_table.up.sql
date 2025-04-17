CREATE TABLE IF NOT EXISTS rentals (
                                       id             SERIAL PRIMARY KEY,
                                       user_id        INT       NOT NULL,
                                       book_id        INT       NOT NULL,
                                       date_issued    TIMESTAMP NOT NULL DEFAULT NOW(),
    date_returned  TIMESTAMP NULL,
    CONSTRAINT fk_user   FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT fk_book   FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE RESTRICT,
    CONSTRAINT chk_return_after_issue CHECK (date_returned IS NULL OR date_returned >= date_issued)
    );
-- Уникальные индексы для активных аренд
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_book_active
    ON rentals(book_id)
    WHERE date_returned IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_user_book_active
    ON rentals(user_id, book_id)
    WHERE date_returned IS NULL;
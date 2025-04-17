CREATE TABLE IF NOT EXISTS orders (
                                      id SERIAL PRIMARY KEY,
                                      pet_id INTEGER,
                                      quantity INTEGER,
                                      ship_date TIMESTAMP,
                                      status VARCHAR(20),
    complete BOOLEAN,
    FOREIGN KEY (pet_id) REFERENCES pets (id) ON DELETE SET NULL
    );

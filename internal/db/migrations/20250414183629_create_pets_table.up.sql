CREATE TABLE IF NOT EXISTS pets (
                                    id SERIAL PRIMARY KEY,
                                    name VARCHAR(50) NOT NULL,
    category VARCHAR(50),
    status VARCHAR(20),
    photo_urls TEXT[]
    );

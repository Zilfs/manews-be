CREATE TABLE IF NOT EXISTS "categories" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    slug VARCHAR(200) UNIQUE NOT NULL,
    created_by_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP NULL DEFAULT current_timestamp
);

CREATE INDEX idx_categories_created_by_id ON categories(created_by_id);
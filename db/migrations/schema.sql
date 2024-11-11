DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS contacts;
DROP INDEX IF EXISTS idx_user_email;
DROP INDEX IF EXISTS idx_fk_contact_user;
CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    hash_password TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS contacts(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL phone VARCHAR(100) NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_user_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_fk_contact_user ON contacts(user_id);
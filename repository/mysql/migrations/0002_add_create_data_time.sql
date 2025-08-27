-- +migrate Up
ALTER TABLE users ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- +migrate Down
ALTER TABLE users DROP COLUMN created_at;

-- +migrate Up
ALTER TABLE `users`
    ADD COLUMN `role` ENUM('admin', 'user') NOT NULL DEFAULT 'user ';

-- +migrate Down
ALTER TABLE `users` DROP COLUMN `role`;
-- +migrate Up
CREATE TABLE IF NOT EXISTS `permissions` (
        `id` int PRIMARY KEY AUTO_INCREMENT,
        `title` varchar(191) NOT NULL,
        `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- +migrate Down
DROP TABLE IF EXISTS `permissions`;

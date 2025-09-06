-- +migrate Up
CREATE TABLE IF NOT EXISTS `access_control` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `actor_ic` varchar(191) NOT NULL UNIQUE,
    `actpr_type` varchar(191) NOT NULL,
    `permission_id` INT NOT NULL,
    FOREIGN KEY (`permission_id`) REFERENCES permissions(`id`),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- +migrate Down
DROP TABLE IF EXISTS `permissions`;

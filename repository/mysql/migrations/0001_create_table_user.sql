-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
                       `id` int primary key AUTO_INCREMENT,
                       `name` varchar(191) not null,
                       `phone_number` varchar(191) unique,
                       `password` text not null
);

-- +migrate Down
DROP TABLE `users`;
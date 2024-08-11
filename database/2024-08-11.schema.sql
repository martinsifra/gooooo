SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE TABLE `users` (
    `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
    `name` longtext COLLATE utf8mb4_unicode_ci,
    `email` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
    `birthdate` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

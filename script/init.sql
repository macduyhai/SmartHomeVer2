DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `username` VARCHAR(64) UNIQUE,
    `password` VARCHAR(64),
    `money` INT(11),
    `phone` VARCHAR(255) DEFAULT '',
    `number_video` INT(10)  DEFAULT 0,
    `total_size` INT(10)  DEFAULT 0,
    `max_size` INT(20)  DEFAULT 600000,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX user_name (`username`)
);

DROP TABLE IF EXISTS `devices`;
CREATE TABLE `devices` (
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT(20),
    `device_name` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ,
    `mac` VARCHAR(255) UNIQUE NOT NULL,
    `video_name` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ,
    `video_size` BIGINT(20)  NOT NULL,
    `video_time` BIGINT(20)  NOT NULL,
    `status` TINYINT(3) DEFAULT 0,
    `location` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `map_long` VARCHAR(255)  DEFAULT NULL,
    `map_lat` VARCHAR(255)  DEFAULT NULL,
    `expired` DATETIME     DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX device_mac (`mac`)
);

DROP TABLE IF EXISTS `media`;
CREATE TABLE `media` (
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT(20),
    `video_name` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ,
    `video_size` BIGINT(20)  NOT NULL,
    `video_time` BIGINT(20)  NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`                 DATETIME    DEFAULT NULL,
    INDEX video_name (`video_name`)
);

DROP TABLE IF EXISTS `logs`;
CREATE TABLE `logs` (
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `detail` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `user_id` BIGINT(20) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX log_id (`id`)
);

DROP TABLE IF EXISTS `authenkey`;
CREATE TABLE `authenkey` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `key` VARCHAR(255) NOT NULL,
  `created_at`                 DATETIME    DEFAULT NOW()
);

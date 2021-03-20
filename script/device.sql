DROP TABLE IF EXISTS `devices`;

CREATE TABLE `devices` (
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT(20),
    `device_name` VARCHAR(255) DEFAULT '',
    `mac` VARCHAR(255) UNIQUE NOT NULL,
    `video_name` VARCHAR(255) DEFAULT '',
    `video_size` BIGINT(20)  NOT NULL,
    `video_time` BIGINT(20)  NOT NULL,
    `status` TINYINT(3) DEFAULT 0,
    `location` VARCHAR(255) DEFAULT '',
    `expired` DATETIME     DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX device_mac (`mac`)
);

DROP TABLE IF EXISTS `medias`;

CREATE TABLE `medias` (
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT(20),
    `video_name` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ,
    `video_size` BIGINT(20)  NOT NULL,
    `video_time` BIGINT(20)  NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX video_name (`video_name`)
);
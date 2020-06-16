DROP TABLE IF EXISTS `logs`;

CREATE TABLE `logs` (
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `detail` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
     `money` INT(11) NOT NULL,
    `tag` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `user_id` BIGINT(20) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX log_id (`id`)
);
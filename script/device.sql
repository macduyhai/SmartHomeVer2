DROP TABLE IF EXISTS `devices`;

CREATE TABLE `devices` (
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `mac` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `serial` VARCHAR(64) UNIQUE,
	`id_device` INT(11),
	`name` VARCHAR(64) UNIQUE,
	`type` VARCHAR(64) UNIQUE,
	`laststate` BOOLEAN,
    `newstate` BOOLEAN,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX device_name (`name`)
);
